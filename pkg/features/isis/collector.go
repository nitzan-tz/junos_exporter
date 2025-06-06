// SPDX-License-Identifier: MIT

package isis

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/czerwonk/junos_exporter/pkg/collector"

	log "github.com/sirupsen/logrus"
)

const prefix string = "junos_isis_"

var (
	upCountDesc       *prometheus.Desc
	totalCountDesc    *prometheus.Desc
	adjStateDesc      *prometheus.Desc
	adjCountDesc      *prometheus.Desc
	adjPriorityDesc   *prometheus.Desc
	adjMetricDesc     *prometheus.Desc
	adjHelloTimerDesc *prometheus.Desc
	adjHoldTimerDesc  *prometheus.Desc
	lspIntervalDesc   *prometheus.Desc
	csnpIntervalDesc  *prometheus.Desc
	helloPaddingDesc  *prometheus.Desc
	maxHelloSizeDesc  *prometheus.Desc
	nodeCoverageDesc  *prometheus.Desc
	backupPathDesc    *prometheus.Desc
)

func init() {
	l := []string{"target"}
	upCountDesc = prometheus.NewDesc(prefix+"up_count", "Number of ISIS Adjacencies in state up", l, nil)
	totalCountDesc = prometheus.NewDesc(prefix+"total_count", "Number of ISIS Adjacencies", l, nil)
	l = append(l, "interface_name")
	lspIntervalDesc = prometheus.NewDesc(prefix+"lsp_interval_ms", "The ISIS LSP interval", l, nil)
	csnpIntervalDesc = prometheus.NewDesc(prefix+"csnp_interval_seconds", "The ISIS CSNP interval", l, nil)
	helloPaddingDesc = prometheus.NewDesc(prefix+"hello_padding", "The ISIS hello padding (0 = UNKNOWN, 1 = ADAPTIVE, 2 = DISABLE, 3 = LOOSE, 4 = STRICT)", l, nil)
	maxHelloSizeDesc = prometheus.NewDesc(prefix+"max_hello_size_bytes", "The ISIS max hello size", l, nil)
	l = append(l, "system_name", "level")
	adjStateDesc = prometheus.NewDesc(prefix+"adjacency_state", "The ISIS Adjacency state (0 = DOWN, 1 = UP, 2 = NEW, 3 = ONE-WAY, 4 =INITIALIZING , 5 = REJECTED)", l, nil)
	interfaceMetricsLabels := []string{"target", "interface_name", "level"}
	adjCountDesc = prometheus.NewDesc(prefix+"adjacency_count", "The number of ISIS adjacencies for an interface", interfaceMetricsLabels, nil)
	adjPriorityDesc = prometheus.NewDesc(prefix+"adjacency_priority", "The ISIS adjacency priority", interfaceMetricsLabels, nil)
	adjMetricDesc = prometheus.NewDesc(prefix+"adjacency_metric", "The ISIS adjacency metric", interfaceMetricsLabels, nil)
	adjHelloTimerDesc = prometheus.NewDesc(prefix+"adjacency_hello_timer_seconds", "The ISIS adjacency hello timer", interfaceMetricsLabels, nil)
	adjHoldTimerDesc = prometheus.NewDesc(prefix+"adjacency_hold_timer_seconds", "The ISIS adjacency hold timer", interfaceMetricsLabels, nil)
	coverageLabels := []string{"target", "topology", "level", "node_coverage", "ipv4_route_coverage", "ipv6_route_coverage", "clns_route_coverage", "ipv4_mpls_route_coverage", "ipv6_mpls_route_coverage", "ipv4_mpls_sspf_route_coverage", "ipv6_mpls_sspf_route_coverage"}
	nodeCoverageDesc = prometheus.NewDesc(prefix+"backup_node_coverage", "The ISIS backup node coverage in percents", coverageLabels, nil)
	backupPathLabels := []string{"target", "node_name", "backup_path_via", "backup_path_via_interface"}
	backupPathDesc = prometheus.NewDesc(prefix+"backup_path", "An ISIS backup path", backupPathLabels, nil)
}

type isisCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &isisCollector{}
}

// Name returns the name of the collector
func (*isisCollector) Name() string {
	return "ISIS"
}

// Describe describes the metrics
func (*isisCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- upCountDesc
	ch <- totalCountDesc
	ch <- adjCountDesc
	ch <- adjPriorityDesc
	ch <- adjMetricDesc
	ch <- adjHelloTimerDesc
	ch <- adjHoldTimerDesc
	ch <- lspIntervalDesc
	ch <- csnpIntervalDesc
	ch <- helloPaddingDesc
	ch <- maxHelloSizeDesc
	ch <- nodeCoverageDesc
	ch <- backupPathDesc
}

// Collect collects metrics from JunOS
func (c *isisCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	adjancies, err := c.isisAdjancies(client)
	if err != nil {
		return err
	}

	ch <- prometheus.MustNewConstMetric(upCountDesc, prometheus.GaugeValue, adjancies.Up, labelValues...)
	ch <- prometheus.MustNewConstMetric(totalCountDesc, prometheus.GaugeValue, adjancies.Total, labelValues...)

	if adjancies.Adjacencies != nil {
		for _, adj := range adjancies.Adjacencies {
			localLabelvalues := append(labelValues, adj.InterfaceName, adj.SystemName, strconv.Itoa(int(adj.Level)))
			state := 0.0
			switch adj.AdjacencyState {
			case "Down":
				state = 0.0
			case "Up":
				state = 1.0
			case "New":
				state = 2.0
			case "One-way":
				state = 3.0
			case "Initializing":
				state = 4.0
			case "Rejected":
				state = 5.0
			}

			ch <- prometheus.MustNewConstMetric(adjStateDesc, prometheus.GaugeValue, state, localLabelvalues...)
		}
	}

	var ifas interfaces
	err = client.RunCommandAndParse("show isis interface extensive", &ifas)
	if err != nil {
		return errors.Wrap(err, "failed to run command 'show isis interface extensive'")
	}
	c.isisInterfaces(ifas, ch, labelValues)

	var coverage backupCoverage
	err = client.RunCommandAndParse("show isis backup coverage", &coverage)
	if err != nil {
		return errors.Wrap(err, "failed to run command 'show isis backup coverage'")
	}
	c.isisBackupCoverage(coverage, ch, labelValues)

	var backupPath backupSPF
	err = client.RunCommandAndParse("show isis backup spf results", &backupPath)
	if err != nil {
		return errors.Wrap(err, "failed to run command 'show isis backup spf results'")
	}
	c.isisBackupPath(backupPath, ch, labelValues)
	return nil
}

func (c *isisCollector) isisAdjancies(client collector.Client) (*adjacencies, error) {
	up := 0
	total := 0

	var x = result{}
	err := client.RunCommandAndParse("show isis adjacency", &x)
	if err != nil {
		return nil, err
	}

	for _, adjacency := range x.Information.Adjacencies {
		if adjacency.AdjacencyState == "Up" {
			up++
		}
		total++
	}

	return &adjacencies{Up: float64(up), Total: float64(total), Adjacencies: x.Information.Adjacencies}, nil
}

func (c *isisCollector) isisInterfaces(interfaces interfaces, ch chan<- prometheus.Metric, labelValues []string) {
	for _, i := range interfaces.IsisInterfaceInformation.IsisInterface {
		if strings.ToLower(i.InterfaceLevelData.Passive) == "passive" {
			continue
		}
		labels := append(labelValues,
			i.InterfaceName,
			i.InterfaceLevelData.Level)
		ch <- prometheus.MustNewConstMetric(adjCountDesc, prometheus.CounterValue, i.InterfaceLevelData.AdjacencyCount, labels...)
		ch <- prometheus.MustNewConstMetric(adjPriorityDesc, prometheus.GaugeValue, i.InterfaceLevelData.InterfacePriority, labels...)
		ch <- prometheus.MustNewConstMetric(adjMetricDesc, prometheus.GaugeValue, i.InterfaceLevelData.Metric, labels...)
		ch <- prometheus.MustNewConstMetric(adjHelloTimerDesc, prometheus.GaugeValue, i.InterfaceLevelData.HelloTime, labels...)
		ch <- prometheus.MustNewConstMetric(adjHoldTimerDesc, prometheus.GaugeValue, i.InterfaceLevelData.HoldTime, labels...)
		additionaLabels := append(labelValues, i.InterfaceName)
		helloPadding := getHelloPadding(i.HelloPadding)
		ch <- prometheus.MustNewConstMetric(lspIntervalDesc, prometheus.GaugeValue, i.LSPInterval, additionaLabels...)
		ch <- prometheus.MustNewConstMetric(csnpIntervalDesc, prometheus.GaugeValue, i.CSNPInterval, additionaLabels...)
		ch <- prometheus.MustNewConstMetric(helloPaddingDesc, prometheus.GaugeValue, helloPadding, additionaLabels...)
		ch <- prometheus.MustNewConstMetric(maxHelloSizeDesc, prometheus.GaugeValue, i.MaxHelloSize, additionaLabels...)
	}
}

func (c *isisCollector) isisBackupCoverage(coverage backupCoverage, ch chan<- prometheus.Metric, labelValues []string) {
	compactCoverage := coverage.IsisBackupCoverageInformation.IsisBackupCoverage
	labels := append(labelValues, compactCoverage.IsisTopologyID, compactCoverage.Level, compactCoverage.IsisNodeCoverage,
		compactCoverage.IsisRouteCoverageIpv4, compactCoverage.IsisRouteCoverageIpv6,
		compactCoverage.IsisRouteCoverageClns, compactCoverage.IsisRouteCoverageIpv4Mpls,
		compactCoverage.IsisRouteCoverageIpv6Mpls, compactCoverage.IsisRouteCoverageIpv4MplsSspf,
		compactCoverage.IsisRouteCoverageIpv6MplsSspf)
	ch <- prometheus.MustNewConstMetric(nodeCoverageDesc, prometheus.GaugeValue, percentageToFloat64(compactCoverage.IsisNodeCoverage), labels...)
}

func (c *isisCollector) isisBackupPath(backupPath backupSPF, ch chan<- prometheus.Metric, labelValues []string) {
	for _, node := range backupPath.IsisSpfInformation.IsisSpf {
		for _, bpSFPResult := range node.IsisBackupSpfResult {
			for _, _ = range bpSFPResult.NoCoverageReasonElement {
				labelValues := append(labelValues, strings.TrimSuffix(bpSFPResult.NodeID, ".00"), "", "")
				ch <- prometheus.MustNewConstMetric(backupPathDesc, prometheus.GaugeValue, 0.0, labelValues...)
			}
			labelValues := append(labelValues, strings.TrimSuffix(bpSFPResult.NodeID, ".00"), bpSFPResult.BackupNextHopElement.IsisNextHop, bpSFPResult.BackupNextHopElement.InterfaceName)
			ch <- prometheus.MustNewConstMetric(backupPathDesc, prometheus.GaugeValue, 1.0, labelValues...)
		}
	}
}

func getHelloPadding(h string) float64 {
	switch strings.ToLower(h) {
	case "adaptive":
		return 1.0
	case "disable":
		return 2.0
	case "loose":
		return 3.0
	case "strict":
		return 4.0
	default:
		return 0.0
	}
}

func percentageToFloat64(percentageStr string) float64 {
	trimmed := strings.TrimSuffix(percentageStr, "%")
	value, err := strconv.ParseFloat(trimmed, 64)
	if err != nil {
		log.Errorf("failed to turn percentage value into float64: %v", err)
		return 0
	}
	return value
}
