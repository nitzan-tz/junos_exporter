// SPDX-License-Identifier: MIT

package system

import (
	"encoding/xml"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_system_"

var (
	mbufsCurrentDesc *prometheus.Desc
	mbufsCacheDesc   *prometheus.Desc
	mbufsTotalDesc   *prometheus.Desc
	mbufsDeniedDesc  *prometheus.Desc

	mbufClustersCurrentDesc *prometheus.Desc
	mbufClustersCacheDesc   *prometheus.Desc
	mbufClustersTotalDesc   *prometheus.Desc
	mbufClustersMaxDesc     *prometheus.Desc
	mbufClustersDeniedDesc  *prometheus.Desc

	mbufClustersFromPacketZoneCurrentDesc *prometheus.Desc
	mbufClustersFromPacketZoneCacheDesc   *prometheus.Desc

	jumboClustersCurrentDesc *prometheus.Desc
	jumboClustersCacheDesc   *prometheus.Desc
	jumboClustersTotalDesc   *prometheus.Desc
	jumboClustersMaxDesc     *prometheus.Desc
	jumboClustersDeniedDesc  *prometheus.Desc

	networkAllocCurrentDesc *prometheus.Desc
	networkAllocCacheDesc   *prometheus.Desc
	networkAllocTotalDesc   *prometheus.Desc

	sfbufsDeniedDesc  *prometheus.Desc
	sfbufsDelayedDesc *prometheus.Desc

	mbufAndClustersDeniedDesc *prometheus.Desc
	ioInitDesc                *prometheus.Desc

	hardwareInfoDesc *prometheus.Desc

	licenseUsedDesc      *prometheus.Desc
	licenseInstalledDesc *prometheus.Desc
	licenseNeededDesc    *prometheus.Desc
	licenseExpiryDesc    *prometheus.Desc

	// regex
	regex1Ints        *regexp.Regexp = regexp.MustCompile(`^(\d+).*`)
	regex2Ints        *regexp.Regexp = regexp.MustCompile(`^(\d+)\/(\d+).*`)
	regex3Ints        *regexp.Regexp = regexp.MustCompile(`^(\d+)\/(\d+)\/(\d+).*`)
	regex4Ints        *regexp.Regexp = regexp.MustCompile(`^(\d+)\/(\d+)\/(\d+)\/(\d+).*`)
	regexNetworkAlloc *regexp.Regexp = regexp.MustCompile(`^(\d+)K\/(\d+)K\/(\d+)K.*`)
)

type systemCollector struct {
}

func init() {
	var l []string

	l = []string{"target"}
	mbufsCurrentDesc = prometheus.NewDesc(prefix+"mbufs_bytes_current", "Current number of bytes in mbufs", l, nil)
	mbufsCacheDesc = prometheus.NewDesc(prefix+"mbufs_bytes_cache", "Cached number of bytes in mbufs", l, nil)
	mbufsTotalDesc = prometheus.NewDesc(prefix+"mbufs_bytes_total", "Total nuumber of bytes in mbufs", l, nil)
	mbufsDeniedDesc = prometheus.NewDesc(prefix+"mbufs_denied_count", "Number of mbuf requests denied", l, nil)

	mbufClustersCurrentDesc = prometheus.NewDesc(prefix+"mbuf_cluster_bytes_current", "Current number of bytes in mbuf clusters", l, nil)
	mbufClustersCacheDesc = prometheus.NewDesc(prefix+"mbuf_cluster_bytes_cache", "Cached number of bytes in mbuf clusters", l, nil)
	mbufClustersTotalDesc = prometheus.NewDesc(prefix+"mbuf_cluster_bytes_total", "Total number of bytes in mbuf clusters", l, nil)
	mbufClustersMaxDesc = prometheus.NewDesc(prefix+"mbuf_cluster_bytes_max", "Max number of bytes in mbuf clusters", l, nil)
	mbufClustersDeniedDesc = prometheus.NewDesc(prefix+"mbufs_and_clusters_denied_count", "", l, nil)

	mbufClustersFromPacketZoneCurrentDesc = prometheus.NewDesc(prefix+"mbuf_and_clusters_from_packet_zone_bytes_current", "Current number of bytes used for mbuf+clusters in packet zone", l, nil)
	mbufClustersFromPacketZoneCacheDesc = prometheus.NewDesc(prefix+"mbuf_and_clusters_from_packet_zone_bytes_cache", "Cached number of bytes used for mbuf+clusters in packet zone", l, nil)

	l = append(l, "page_size")
	jumboClustersCurrentDesc = prometheus.NewDesc(prefix+"jumbo_clusters_current", "Current jumbo clusters in use.", l, nil)
	jumboClustersCacheDesc = prometheus.NewDesc(prefix+"jumbo_clusters_cache", "Cached jumbo clusters in use", l, nil)
	jumboClustersTotalDesc = prometheus.NewDesc(prefix+"jumbo_clusters_total", "Total jumbo clusters in use", l, nil)
	jumboClustersMaxDesc = prometheus.NewDesc(prefix+"jumbo_clusters_max", "Max jumbo clusters in use", l, nil)
	jumboClustersDeniedDesc = prometheus.NewDesc(prefix+"jumbo_clusters_denied_count", "Number of jumbo cluster requests denied", l, nil)

	l = []string{"target"}
	networkAllocCurrentDesc = prometheus.NewDesc(prefix+"network_allocated_bytes_current", "Current number of bytes allocated for network", l, nil)
	networkAllocCacheDesc = prometheus.NewDesc(prefix+"network_allocated_bytes_cache", "Cached number of bytes allocated for network", l, nil)
	networkAllocTotalDesc = prometheus.NewDesc(prefix+"network_allocated_bytes_total", "Total number of bytes allocated for network", l, nil)

	sfbufsDeniedDesc = prometheus.NewDesc(prefix+"sfbufs_denied_count", "Number of sfbuf requests denied", l, nil)
	sfbufsDelayedDesc = prometheus.NewDesc(prefix+"sfbufs_delayed_count", "Number of sfbuf requests delayed", l, nil)

	ioInitDesc = prometheus.NewDesc(prefix+"io_requests_count", "Number of I/O requests initiated", l, nil)
	mbufAndClustersDeniedDesc = prometheus.NewDesc(prefix+"mbuf_and_clusters_denied_count", "Number of mbuf+cluster requests denied", l, nil)

	l = append(l, "model", "os", "os_version", "serial", "hostname", "alias", "slot_id", "state")
	hardwareInfoDesc = prometheus.NewDesc(prefix+"hardware_info", "Hardware information about this system", l, nil)

	l = []string{"target"}
	l = append(l, "feature_name", "feature_description")
	licenseUsedDesc = prometheus.NewDesc(prefix+"license_used", "Amount of license used", l, nil)
	licenseInstalledDesc = prometheus.NewDesc(prefix+"license_installed", "Amount of license installed", l, nil)
	licenseNeededDesc = prometheus.NewDesc(prefix+"license_needed", "Amount of license needed", l, nil)
	licenseExpiryDesc = prometheus.NewDesc(prefix+"license_expiry", "Days until expiry, if applicable; -1 = expired; +Inf = permanent; -Inf = invalid", l, nil)
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &systemCollector{}
}

// Name returns the name of the collector
func (*systemCollector) Name() string {
	return "System"
}

// Describe describes the metrics
func (*systemCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- mbufsCurrentDesc
	ch <- mbufsCacheDesc
	ch <- mbufsTotalDesc
	ch <- mbufsDeniedDesc
	ch <- mbufClustersCurrentDesc
	ch <- mbufClustersCacheDesc
	ch <- mbufClustersTotalDesc
	ch <- mbufClustersMaxDesc
	ch <- mbufClustersDeniedDesc
	ch <- mbufClustersFromPacketZoneCurrentDesc
	ch <- mbufClustersFromPacketZoneCacheDesc
	ch <- jumboClustersCurrentDesc
	ch <- jumboClustersCacheDesc
	ch <- jumboClustersTotalDesc
	ch <- jumboClustersMaxDesc
	ch <- jumboClustersDeniedDesc
	ch <- networkAllocCurrentDesc
	ch <- networkAllocCacheDesc
	ch <- networkAllocTotalDesc
	ch <- sfbufsDeniedDesc
	ch <- sfbufsDelayedDesc
	ch <- ioInitDesc
	ch <- hardwareInfoDesc
	ch <- licenseUsedDesc
	ch <- licenseInstalledDesc
	ch <- licenseNeededDesc
	ch <- licenseExpiryDesc
}

// Collect collects metrics from JunOS
func (c *systemCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	err := c.CollectSystem(client, ch, labelValues)
	if err != nil {
		return err
	}

	return nil
}

func (c *systemCollector) CollectSystem(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	err := c.collectBuffers(client, ch, labelValues)
	if err != nil {
		return fmt.Errorf("could not get system buffers: %w", err)
	}

	err = c.collectSystemInformation(client, ch, labelValues)
	if err != nil {
		return fmt.Errorf("could not get system information: %w", err)
	}

	if client.IsSatelliteEnabled() {
		c.collectSatelites(client, ch, labelValues)
	}

	if client.IsScrapingLicenseEnabled() {
		c.collectLicense(client, ch, labelValues)
	}

	return nil
}

func (c *systemCollector) collectBuffers(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	r := &buffers{}

	err := client.RunCommandAndParseWithParser("show system buffers", func(b []byte) error {
		if string(b[:]) == "\nerror: syntax error, expecting <command>: buffers\n" || strings.Contains(string(b[:]), "error: command is not valid on the") {
			log.Infof("system doesn't support show system buffers command")
			return nil
		}
		err := xml.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		ch <- prometheus.MustNewConstMetric(mbufsCurrentDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufsCurrent), labelValues...)
		ch <- prometheus.MustNewConstMetric(mbufsCacheDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufsCache), labelValues...)
		ch <- prometheus.MustNewConstMetric(mbufsTotalDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufsTotal), labelValues...)
		ch <- prometheus.MustNewConstMetric(mbufsDeniedDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufsDenied), labelValues...)

		ch <- prometheus.MustNewConstMetric(mbufClustersCurrentDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufClustersCurrent), labelValues...)
		ch <- prometheus.MustNewConstMetric(mbufClustersCacheDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufClustersCache), labelValues...)
		ch <- prometheus.MustNewConstMetric(mbufClustersTotalDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufClustersTotal), labelValues...)
		ch <- prometheus.MustNewConstMetric(mbufClustersMaxDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufClustersMax), labelValues...)
		ch <- prometheus.MustNewConstMetric(mbufClustersDeniedDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufClustersDenied), labelValues...)

		ch <- prometheus.MustNewConstMetric(mbufClustersFromPacketZoneCurrentDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufClustersFromPacketZoneCurrent), labelValues...)
		ch <- prometheus.MustNewConstMetric(mbufClustersFromPacketZoneCacheDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufClustersFromPacketZoneCache), labelValues...)

		l := append(labelValues, "4k")
		ch <- prometheus.MustNewConstMetric(jumboClustersCurrentDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersCurrent4K), l...)
		ch <- prometheus.MustNewConstMetric(jumboClustersCacheDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersCache4K), l...)
		ch <- prometheus.MustNewConstMetric(jumboClustersTotalDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersTotal4K), l...)
		ch <- prometheus.MustNewConstMetric(jumboClustersMaxDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersMax4K), l...)
		ch <- prometheus.MustNewConstMetric(jumboClustersDeniedDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersDenied4K), l...)

		l = append(labelValues, "9k")
		ch <- prometheus.MustNewConstMetric(jumboClustersCurrentDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersCurrent9K), l...)
		ch <- prometheus.MustNewConstMetric(jumboClustersCacheDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersCache9K), l...)
		ch <- prometheus.MustNewConstMetric(jumboClustersTotalDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersTotal9K), l...)
		ch <- prometheus.MustNewConstMetric(jumboClustersMaxDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersMax9K), l...)
		ch <- prometheus.MustNewConstMetric(jumboClustersDeniedDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersDenied9K), l...)

		l = append(labelValues, "16k")
		ch <- prometheus.MustNewConstMetric(jumboClustersCurrentDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersCurrent16K), l...)
		ch <- prometheus.MustNewConstMetric(jumboClustersCacheDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersCache16K), l...)
		ch <- prometheus.MustNewConstMetric(jumboClustersTotalDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersTotal16K), l...)
		ch <- prometheus.MustNewConstMetric(jumboClustersMaxDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersMax16K), l...)
		ch <- prometheus.MustNewConstMetric(jumboClustersDeniedDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersDenied16K), l...)

		ch <- prometheus.MustNewConstMetric(sfbufsDeniedDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.SfbufsDenied), labelValues...)
		ch <- prometheus.MustNewConstMetric(sfbufsDelayedDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.SfbufsDelayed), labelValues...)

		ch <- prometheus.MustNewConstMetric(mbufAndClustersDeniedDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufAndClustersDenied), labelValues...)
		ch <- prometheus.MustNewConstMetric(ioInitDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.IoInit), labelValues...)

		// network alloc values seem to be Kb
		ch <- prometheus.MustNewConstMetric(networkAllocCurrentDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.NetworkAllocCurrent*1024), labelValues...)
		ch <- prometheus.MustNewConstMetric(networkAllocCacheDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.NetworkAllocCache*1024), labelValues...)
		ch <- prometheus.MustNewConstMetric(networkAllocTotalDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.NetworkAllocTotal*1024), labelValues...)
		return nil
	})

	if err != nil {
		return err
	}

	if r.Output == "" {
		return nil
	}

	// this data still needs to be parsed
	lines := strings.Split(r.Output, "\n")

	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}

	// trim away empty lines
	lines = lines[1 : len(lines)-1]

	// NOTE: matches[0][0] is always the whole line

	// "3216/15519/18735 mbufs in use (current/cache/total)"
	matches := regex3Ints.FindAllStringSubmatch(lines[0], 1)
	if len(matches) >= 1 && len(matches[0]) >= 4 {
		r.MemoryStatistics.MbufsCurrent, _ = strconv.Atoi(matches[0][1])
		r.MemoryStatistics.MbufsCache, _ = strconv.Atoi(matches[0][2])
		r.MemoryStatistics.MbufsTotal, _ = strconv.Atoi(matches[0][3])
	}

	// "3074/14458/17532/2039110 mbuf clusters in use (current/cache/total/max)"
	matches = regex4Ints.FindAllStringSubmatch(lines[1], 1)
	if len(matches) >= 1 && len(matches[0]) >= 5 {
		r.MemoryStatistics.MbufClustersCurrent, _ = strconv.Atoi(matches[0][1])
		r.MemoryStatistics.MbufClustersCache, _ = strconv.Atoi(matches[0][2])
		r.MemoryStatistics.MbufClustersTotal, _ = strconv.Atoi(matches[0][3])
		r.MemoryStatistics.MbufClustersMax, _ = strconv.Atoi(matches[0][4])
	}

	// "3069/7557 mbuf+clusters out of packet secondary zone in use (current/cache)"
	matches = regex2Ints.FindAllStringSubmatch(lines[2], 1)
	if len(matches) >= 1 && len(matches[0]) >= 3 {
		r.MemoryStatistics.MbufClustersFromPacketZoneCurrent, _ = strconv.Atoi(matches[0][1])
		r.MemoryStatistics.MbufClustersFromPacketZoneCache, _ = strconv.Atoi(matches[0][2])
	}

	// "0/1101/1101/1019555 4k (page size) jumbo clusters in use (current/cache/total/max)"
	matches = regex4Ints.FindAllStringSubmatch(lines[3], 1)
	if len(matches) >= 1 && len(matches[0]) >= 5 {
		r.MemoryStatistics.JumboClustersCurrent4K, _ = strconv.Atoi(matches[0][1])
		r.MemoryStatistics.JumboClustersCache4K, _ = strconv.Atoi(matches[0][2])
		r.MemoryStatistics.JumboClustersTotal4K, _ = strconv.Atoi(matches[0][3])
		r.MemoryStatistics.JumboClustersMax4K, _ = strconv.Atoi(matches[0][4])
	}

	// "0/1101/1101/1019555 9k (page size) jumbo clusters in use (current/cache/total/max)"
	matches = regex4Ints.FindAllStringSubmatch(lines[4], 1)
	if len(matches) >= 1 && len(matches[0]) >= 5 {
		r.MemoryStatistics.JumboClustersCurrent9K, _ = strconv.Atoi(matches[0][1])
		r.MemoryStatistics.JumboClustersCache9K, _ = strconv.Atoi(matches[0][2])
		r.MemoryStatistics.JumboClustersTotal9K, _ = strconv.Atoi(matches[0][3])
		r.MemoryStatistics.JumboClustersMax9K, _ = strconv.Atoi(matches[0][4])
	}

	// "0/1101/1101/1019555 16k (page size) jumbo clusters in use (current/cache/total/max)"
	matches = regex4Ints.FindAllStringSubmatch(lines[5], 1)
	if len(matches) >= 1 && len(matches[0]) >= 5 {
		r.MemoryStatistics.JumboClustersCurrent16K, _ = strconv.Atoi(matches[0][1])
		r.MemoryStatistics.JumboClustersCache16K, _ = strconv.Atoi(matches[0][2])
		r.MemoryStatistics.JumboClustersTotal16K, _ = strconv.Atoi(matches[0][3])
		r.MemoryStatistics.JumboClustersMax16K, _ = strconv.Atoi(matches[0][4])
	}

	// "6952K/37199K/44152K bytes allocated to network (current/cache/total)"
	matches = regexNetworkAlloc.FindAllStringSubmatch(lines[6], 1)
	if len(matches) >= 1 && len(matches[0]) >= 4 {
		r.MemoryStatistics.NetworkAllocCurrent, _ = strconv.Atoi(matches[0][1])
		r.MemoryStatistics.NetworkAllocCache, _ = strconv.Atoi(matches[0][2])
		r.MemoryStatistics.NetworkAllocTotal, _ = strconv.Atoi(matches[0][3])
	}

	// "0/0/0 requests for mbufs denied (mbufs/clusters/mbuf+clusters)"
	matches = regex3Ints.FindAllStringSubmatch(lines[7], 1)
	if len(matches) >= 1 && len(matches[0]) >= 4 {
		r.MemoryStatistics.MbufsDenied, _ = strconv.Atoi(matches[0][1])
		r.MemoryStatistics.MbufClustersDenied, _ = strconv.Atoi(matches[0][2])
		r.MemoryStatistics.MbufAndClustersDenied, _ = strconv.Atoi(matches[0][2])
	}

	// "0/0/0 requests for jumbo clusters denied (4k/9k/16k)"
	matches = regex3Ints.FindAllStringSubmatch(lines[8], 1)
	if len(matches) >= 1 && len(matches[0]) >= 4 {
		r.MemoryStatistics.JumboClustersDenied4K, _ = strconv.Atoi(matches[0][1])
		r.MemoryStatistics.JumboClustersDenied9K, _ = strconv.Atoi(matches[0][2])
		r.MemoryStatistics.JumboClustersDenied16K, _ = strconv.Atoi(matches[0][3])
	}

	// "0 requests for sfbufs denied"
	matches = regex1Ints.FindAllStringSubmatch(lines[9], 1)
	if len(matches) >= 1 && len(matches[0]) >= 2 {
		r.MemoryStatistics.SfbufsDenied, _ = strconv.Atoi(matches[0][1])
	}

	// "0 requests for sfbufs delayed"
	matches = regex1Ints.FindAllStringSubmatch(lines[10], 1)
	if len(matches) >= 1 {
		r.MemoryStatistics.SfbufsDelayed, _ = strconv.Atoi(matches[0][1])
	}

	// "0 requests for I/O initiated by sendfile"
	matches = regex1Ints.FindAllStringSubmatch(lines[11], 1)
	if len(matches) >= 1 {
		r.MemoryStatistics.IoInit, _ = strconv.Atoi(matches[0][1])
	}

	ch <- prometheus.MustNewConstMetric(mbufsCurrentDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufsCurrent), labelValues...)
	ch <- prometheus.MustNewConstMetric(mbufsCacheDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufsCache), labelValues...)
	ch <- prometheus.MustNewConstMetric(mbufsTotalDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufsTotal), labelValues...)
	ch <- prometheus.MustNewConstMetric(mbufsDeniedDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufsDenied), labelValues...)

	ch <- prometheus.MustNewConstMetric(mbufClustersCurrentDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufClustersCurrent), labelValues...)
	ch <- prometheus.MustNewConstMetric(mbufClustersCacheDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufClustersCache), labelValues...)
	ch <- prometheus.MustNewConstMetric(mbufClustersTotalDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufClustersTotal), labelValues...)
	ch <- prometheus.MustNewConstMetric(mbufClustersMaxDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufClustersMax), labelValues...)
	ch <- prometheus.MustNewConstMetric(mbufClustersDeniedDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufClustersDenied), labelValues...)

	ch <- prometheus.MustNewConstMetric(mbufClustersFromPacketZoneCurrentDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufClustersFromPacketZoneCurrent), labelValues...)
	ch <- prometheus.MustNewConstMetric(mbufClustersFromPacketZoneCacheDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufClustersFromPacketZoneCache), labelValues...)

	l := append(labelValues, "4k")
	ch <- prometheus.MustNewConstMetric(jumboClustersCurrentDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersCurrent4K), l...)
	ch <- prometheus.MustNewConstMetric(jumboClustersCacheDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersCache4K), l...)
	ch <- prometheus.MustNewConstMetric(jumboClustersTotalDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersTotal4K), l...)
	ch <- prometheus.MustNewConstMetric(jumboClustersMaxDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersMax4K), l...)
	ch <- prometheus.MustNewConstMetric(jumboClustersDeniedDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersDenied4K), l...)

	l = append(labelValues, "9k")
	ch <- prometheus.MustNewConstMetric(jumboClustersCurrentDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersCurrent9K), l...)
	ch <- prometheus.MustNewConstMetric(jumboClustersCacheDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersCache9K), l...)
	ch <- prometheus.MustNewConstMetric(jumboClustersTotalDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersTotal9K), l...)
	ch <- prometheus.MustNewConstMetric(jumboClustersMaxDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersMax9K), l...)
	ch <- prometheus.MustNewConstMetric(jumboClustersDeniedDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersDenied9K), l...)

	l = append(labelValues, "16k")
	ch <- prometheus.MustNewConstMetric(jumboClustersCurrentDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersCurrent16K), l...)
	ch <- prometheus.MustNewConstMetric(jumboClustersCacheDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersCache16K), l...)
	ch <- prometheus.MustNewConstMetric(jumboClustersTotalDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersTotal16K), l...)
	ch <- prometheus.MustNewConstMetric(jumboClustersMaxDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersMax16K), l...)
	ch <- prometheus.MustNewConstMetric(jumboClustersDeniedDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersDenied16K), l...)

	ch <- prometheus.MustNewConstMetric(sfbufsDeniedDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.SfbufsDenied), labelValues...)
	ch <- prometheus.MustNewConstMetric(sfbufsDelayedDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.SfbufsDelayed), labelValues...)

	ch <- prometheus.MustNewConstMetric(mbufAndClustersDeniedDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufAndClustersDenied), labelValues...)
	ch <- prometheus.MustNewConstMetric(ioInitDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.IoInit), labelValues...)

	// network alloc values seem to be Kb
	ch <- prometheus.MustNewConstMetric(networkAllocCurrentDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.NetworkAllocCurrent*1024), labelValues...)
	ch <- prometheus.MustNewConstMetric(networkAllocCacheDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.NetworkAllocCache*1024), labelValues...)
	ch <- prometheus.MustNewConstMetric(networkAllocTotalDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.NetworkAllocTotal*1024), labelValues...)

	return nil
}

func (c *systemCollector) collectSystemInformation(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	r := &systemInformation{}
	err := client.RunCommandAndParse("show system information", r)
	if err != nil {
		return err
	}

	hardwareLabels := append(labelValues,
		r.SysInfo.Model,
		r.SysInfo.OS,
		r.SysInfo.OSVersion,
		r.SysInfo.Serial,
		r.SysInfo.Hostname,
		"", "", "")

	ch <- prometheus.MustNewConstMetric(hardwareInfoDesc, prometheus.GaugeValue, float64(1), hardwareLabels...)

	return nil
}

func (c *systemCollector) collectSatelites(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) {
	r := &satelliteChassis{}
	err := client.RunCommandAndParse("show chassis satellite detail", r)
	if err != nil {
		// there are various error messages when satellite is not enabled; thus here we just ignore the error and continue
		return
	}

	for i := range r.SatelliteInfo.Satellite {
		l := append(labelValues,
			strings.ToLower(r.SatelliteInfo.Satellite[i].Model),
			"satellite",
			r.SatelliteInfo.Satellite[i].Version,
			r.SatelliteInfo.Satellite[i].Serial,
			"",
			r.SatelliteInfo.Satellite[i].Alias,
			strconv.Itoa(r.SatelliteInfo.Satellite[i].SlotID),
			strings.ToLower(r.SatelliteInfo.Satellite[i].State))

		ch <- prometheus.MustNewConstMetric(hardwareInfoDesc, prometheus.GaugeValue, float64(1), l...)
	}
}

func (c *systemCollector) collectLicense(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) {
	r := &licenseInformation{}
	err := client.RunCommandAndParse("show system license usage", r)

	if err != nil {
		return
	}
	for _, lic := range r.LicenseInfo.License {
		licenseLabels := append(labelValues,
			strings.ToLower(lic.Name),
			strings.ToLower(lic.Description))

		ch <- prometheus.MustNewConstMetric(licenseUsedDesc, prometheus.GaugeValue, float64(lic.Used), licenseLabels...)
		ch <- prometheus.MustNewConstMetric(licenseInstalledDesc, prometheus.GaugeValue, float64(lic.Installed), licenseLabels...)
		ch <- prometheus.MustNewConstMetric(licenseNeededDesc, prometheus.GaugeValue, float64(lic.Needed), licenseLabels...)

		if strings.Compare(lic.ValidityType, "expired") == 0 {
			ch <- prometheus.MustNewConstMetric(licenseExpiryDesc, prometheus.GaugeValue, float64(-1), licenseLabels...)
		} else if strings.Compare(lic.ValidityType, "permanent") == 0 {
			ch <- prometheus.MustNewConstMetric(licenseExpiryDesc, prometheus.GaugeValue, float64(math.Inf(1)), licenseLabels...)
		} else {
			expiry_str := strings.ToLower(lic.EndDate)
			expiry, err := time.Parse("2006-01-02", expiry_str)
			if err != nil {
				ch <- prometheus.MustNewConstMetric(licenseExpiryDesc, prometheus.GaugeValue, float64(math.Inf(-1)), licenseLabels...)
			} else {
				license_ttl_days := time.Until(expiry).Hours() / 24.0
				ch <- prometheus.MustNewConstMetric(licenseExpiryDesc, prometheus.GaugeValue, float64(license_ttl_days), licenseLabels...)
			}
		}
	}
}
