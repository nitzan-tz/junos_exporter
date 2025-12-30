
package aaa

import (
	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "junos_aaa_"

var (
	accountingRequestsDesc      *prometheus.Desc
	accountingRequestFailuresDesc *prometheus.Desc
	accountingRequestSuccessDesc *prometheus.Desc
	accountingResponseFailuresDesc *prometheus.Desc
	accountingResponseSuccessDesc *prometheus.Desc
	accountingTimeoutsDesc      *prometheus.Desc
	accountingRequestsPendingDesc *prometheus.Desc
	accountingMalformedResponsesDesc *prometheus.Desc
	accountingRetransmissionsDesc *prometheus.Desc
	accountingBadAuthenticatorsDesc *prometheus.Desc
	accountingPacketsDroppedDesc *prometheus.Desc

	authenticationRequestsDesc    *prometheus.Desc
	authenticationAcceptsDesc     *prometheus.Desc
	authenticationRejectsDesc     *prometheus.Desc
	authenticationRadiusFailuresDesc *prometheus.Desc
	authenticationInvalidCredentialsDesc *prometheus.Desc
	authenticationMalformedRequestDesc *prometheus.Desc
	authenticationInternalFailureDesc *prometheus.Desc
	authenticationLocalFailuresDesc *prometheus.Desc
	authenticationLdapFailuresDesc *prometheus.Desc
	authenticationChallengesDesc  *prometheus.Desc
	authenticationTimeoutsDesc    *prometheus.Desc

	radiusServerMaxOutstandingDesc *prometheus.Desc
	radiusServerCurrentOutstandingDesc *prometheus.Desc
	radiusServerPeakOutstandingDesc *prometheus.Desc
	radiusServerFailOutstandingDesc *prometheus.Desc

	radiusServerLastRttDesc *prometheus.Desc
	radiusServerAuthRequestsDesc *prometheus.Desc
	radiusServerAuthRolloverRequestsDesc *prometheus.Desc
	radiusServerAuthRetransmissionsDesc *prometheus.Desc
	radiusServerAcceptsDesc *prometheus.Desc
	radiusServerRejectsDesc *prometheus.Desc
	radiusServerChallengesDesc *prometheus.Desc
	radiusServerAuthMalformedResponsesDesc *prometheus.Desc
	radiusServerAuthBadAuthenticatorsDesc *prometheus.Desc
	radiusServerAuthRequestsPendingDesc *prometheus.Desc
	radiusServerAuthTimeoutsDesc *prometheus.Desc
	radiusServerAuthUnknownResponsesDesc *prometheus.Desc
	radiusServerAuthPacketsDroppedDesc *prometheus.Desc
	radiusServerAcctStartRequestsDesc *prometheus.Desc
	radiusServerAcctInterimRequestsDesc *prometheus.Desc
	radiusServerAcctStopRequestsDesc *prometheus.Desc
	radiusServerAcctRolloverRequestsDesc *prometheus.Desc
	radiusServerAcctRetransmissionsDesc *prometheus.Desc
	radiusServerAcctStartResponseDesc *prometheus.Desc
	radiusServerAcctInterimResponseDesc *prometheus.Desc
	radiusServerAcctStopResponseDesc *prometheus.Desc
	radiusServerAcctMalformedResponseDesc *prometheus.Desc
	radiusServerAcctBadAuthenticatorsDesc *prometheus.Desc
	radiusServerAcctRequestsPendingDesc *prometheus.Desc
	radiusServerAcctTimeoutsDesc *prometheus.Desc
	radiusServerAcctUnknownResponsesDesc *prometheus.Desc
	radiusServerAcctPacketsDroppedDesc *prometheus.Desc

	localCertRequestsDesc *prometheus.Desc
	localCertFailedRequestsDesc *prometheus.Desc
	localCertTotalResponsesDesc *prometheus.Desc
	localCertConfiguredResponsesDesc *prometheus.Desc
)

func init() {
	l := []string{"target"}
	serverLabel := []string{"target", "server_address"}
	// Accounting
	accountingRequestsDesc = prometheus.NewDesc(prefix+"accounting_requests_total", "Total accounting requests", l, nil)
	accountingRequestFailuresDesc = prometheus.NewDesc(prefix+"accounting_request_failures_total", "Total accounting request failures", l, nil)
	accountingRequestSuccessDesc = prometheus.NewDesc(prefix+"accounting_request_success_total", "Total accounting request success", l, nil)
	accountingResponseFailuresDesc = prometheus.NewDesc(prefix+"accounting_response_failures_total", "Total accounting response failures", l, nil)
	accountingResponseSuccessDesc = prometheus.NewDesc(prefix+"accounting_response_success_total", "Total accounting response success", l, nil)
	accountingTimeoutsDesc = prometheus.NewDesc(prefix+"accounting_timeouts_total", "Total accounting timeouts", l, nil)
	accountingRequestsPendingDesc = prometheus.NewDesc(prefix+"accounting_requests_pending", "Pending accounting requests", l, nil)
	accountingMalformedResponsesDesc = prometheus.NewDesc(prefix+"accounting_malformed_responses_total", "Total accounting malformed responses", l, nil)
	accountingRetransmissionsDesc = prometheus.NewDesc(prefix+"accounting_retransmissions_total", "Total accounting retransmissions", l, nil)
	accountingBadAuthenticatorsDesc = prometheus.NewDesc(prefix+"accounting_bad_authenticators_total", "Total accounting bad authenticators", l, nil)
	accountingPacketsDroppedDesc = prometheus.NewDesc(prefix+"accounting_packets_dropped_total", "Total accounting packets dropped", l, nil)

	// Authentication
	authenticationRequestsDesc = prometheus.NewDesc(prefix+"authentication_requests_total", "Total authentication requests", l, nil)
	authenticationAcceptsDesc = prometheus.NewDesc(prefix+"authentication_accepts_total", "Total authentication accepts", l, nil)
	authenticationRejectsDesc = prometheus.NewDesc(prefix+"authentication_rejects_total", "Total authentication rejects", l, nil)
	authenticationRadiusFailuresDesc = prometheus.NewDesc(prefix+"authentication_radius_failures_total", "Total authentication radius failures", l, nil)
	authenticationInvalidCredentialsDesc = prometheus.NewDesc(prefix+"authentication_invalid_credentials_total", "Total authentication invalid credentials", l, nil)
	authenticationMalformedRequestDesc = prometheus.NewDesc(prefix+"authentication_malformed_requests_total", "Total authentication malformed requests", l, nil)
	authenticationInternalFailureDesc = prometheus.NewDesc(prefix+"authentication_internal_failures_total", "Total authentication internal failures", l, nil)
	authenticationLocalFailuresDesc = prometheus.NewDesc(prefix+"authentication_local_failures_total", "Total authentication local failures", l, nil)
	authenticationLdapFailuresDesc = prometheus.NewDesc(prefix+"authentication_ldap_failures_total", "Total authentication ldap failures", l, nil)
	authenticationChallengesDesc = prometheus.NewDesc(prefix+"authentication_challenges_total", "Total authentication challenges", l, nil)
	authenticationTimeoutsDesc = prometheus.NewDesc(prefix+"authentication_timeouts_total", "Total authentication timeouts", l, nil)

	// Radius Server (from show network-access aaa statistics detail radius)
	radiusServerMaxOutstandingDesc = prometheus.NewDesc(prefix+"radius_server_max_outstanding", "Max outstanding requests for radius server", serverLabel, nil)
	radiusServerCurrentOutstandingDesc = prometheus.NewDesc(prefix+"radius_server_current_outstanding", "Current outstanding requests for radius server", serverLabel, nil)
	radiusServerPeakOutstandingDesc = prometheus.NewDesc(prefix+"radius_server_peak_outstanding", "Peak outstanding requests for radius server", serverLabel, nil)
	radiusServerFailOutstandingDesc = prometheus.NewDesc(prefix+"radius_server_fail_outstanding", "Fail outstanding requests for radius server", serverLabel, nil)

	// Radius Server Detailed (from show network-access aaa radius-servers detail)
	radiusServerLastRttDesc = prometheus.NewDesc(prefix+"radius_server_last_rtt", "Last RTT for radius server", serverLabel, nil)
	radiusServerAuthRequestsDesc = prometheus.NewDesc(prefix+"radius_server_auth_requests_total", "Total authentication requests for radius server", serverLabel, nil)
	radiusServerAuthRolloverRequestsDesc = prometheus.NewDesc(prefix+"radius_server_auth_rollover_requests_total", "Total authentication rollover requests for radius server", serverLabel, nil)
	radiusServerAuthRetransmissionsDesc = prometheus.NewDesc(prefix+"radius_server_auth_retransmissions_total", "Total authentication retransmissions for radius server", serverLabel, nil)
	radiusServerAcceptsDesc = prometheus.NewDesc(prefix+"radius_server_accepts_total", "Total accepts for radius server", serverLabel, nil)
	radiusServerRejectsDesc = prometheus.NewDesc(prefix+"radius_server_rejects_total", "Total rejects for radius server", serverLabel, nil)
	radiusServerChallengesDesc = prometheus.NewDesc(prefix+"radius_server_challenges_total", "Total challenges for radius server", serverLabel, nil)
	radiusServerAuthMalformedResponsesDesc = prometheus.NewDesc(prefix+"radius_server_auth_malformed_responses_total", "Total authentication malformed responses for radius server", serverLabel, nil)
	radiusServerAuthBadAuthenticatorsDesc = prometheus.NewDesc(prefix+"radius_server_auth_bad_authenticators_total", "Total authentication bad authenticators for radius server", serverLabel, nil)
	radiusServerAuthRequestsPendingDesc = prometheus.NewDesc(prefix+"radius_server_auth_requests_pending", "Pending authentication requests for radius server", serverLabel, nil)
	radiusServerAuthTimeoutsDesc = prometheus.NewDesc(prefix+"radius_server_auth_timeouts_total", "Total authentication timeouts for radius server", serverLabel, nil)
	radiusServerAuthUnknownResponsesDesc = prometheus.NewDesc(prefix+"radius_server_auth_unknown_responses_total", "Total authentication unknown responses for radius server", serverLabel, nil)
	radiusServerAuthPacketsDroppedDesc = prometheus.NewDesc(prefix+"radius_server_auth_packets_dropped_total", "Total authentication packets dropped for radius server", serverLabel, nil)
	radiusServerAcctStartRequestsDesc = prometheus.NewDesc(prefix+"radius_server_acct_start_requests_total", "Total accounting start requests for radius server", serverLabel, nil)
	radiusServerAcctInterimRequestsDesc = prometheus.NewDesc(prefix+"radius_server_acct_interim_requests_total", "Total accounting interim requests for radius server", serverLabel, nil)
	radiusServerAcctStopRequestsDesc = prometheus.NewDesc(prefix+"radius_server_acct_stop_requests_total", "Total accounting stop requests for radius server", serverLabel, nil)
	radiusServerAcctRolloverRequestsDesc = prometheus.NewDesc(prefix+"radius_server_acct_rollover_requests_total", "Total accounting rollover requests for radius server", serverLabel, nil)
	radiusServerAcctRetransmissionsDesc = prometheus.NewDesc(prefix+"radius_server_acct_retransmissions_total", "Total accounting retransmissions for radius server", serverLabel, nil)
	radiusServerAcctStartResponseDesc = prometheus.NewDesc(prefix+"radius_server_acct_start_response_total", "Total accounting start response for radius server", serverLabel, nil)
	radiusServerAcctInterimResponseDesc = prometheus.NewDesc(prefix+"radius_server_acct_interim_response_total", "Total accounting interim response for radius server", serverLabel, nil)
	radiusServerAcctStopResponseDesc = prometheus.NewDesc(prefix+"radius_server_acct_stop_response_total", "Total accounting stop response for radius server", serverLabel, nil)
	radiusServerAcctMalformedResponseDesc = prometheus.NewDesc(prefix+"radius_server_acct_malformed_response_total", "Total accounting malformed response for radius server", serverLabel, nil)
	radiusServerAcctBadAuthenticatorsDesc = prometheus.NewDesc(prefix+"radius_server_acct_bad_authenticators_total", "Total accounting bad authenticators for radius server", serverLabel, nil)
	radiusServerAcctRequestsPendingDesc = prometheus.NewDesc(prefix+"radius_server_acct_requests_pending", "Pending accounting requests for radius server", serverLabel, nil)
	radiusServerAcctTimeoutsDesc = prometheus.NewDesc(prefix+"radius_server_acct_timeouts_total", "Total accounting timeouts for radius server", serverLabel, nil)
	radiusServerAcctUnknownResponsesDesc = prometheus.NewDesc(prefix+"radius_server_acct_unknown_responses_total", "Total accounting unknown responses for radius server", serverLabel, nil)
	radiusServerAcctPacketsDroppedDesc = prometheus.NewDesc(prefix+"radius_server_acct_packets_dropped_total", "Total accounting packets dropped for radius server", serverLabel, nil)

	// Local Certificate
	localCertRequestsDesc = prometheus.NewDesc(prefix+"local_cert_requests_total", "Total local certificate requests", l, nil)
	localCertFailedRequestsDesc = prometheus.NewDesc(prefix+"local_cert_failed_requests_total", "Total local certificate failed requests", l, nil)
	localCertTotalResponsesDesc = prometheus.NewDesc(prefix+"local_cert_total_responses_total", "Total local certificate total responses", l, nil)
	localCertConfiguredResponsesDesc = prometheus.NewDesc(prefix+"local_cert_configured_responses_total", "Total local certificate configured responses", l, nil)
}

type aaaCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &aaaCollector{}
}

// Name returns the name of the collector
func (*aaaCollector) Name() string {
	return "AAA"
}

// Describe describes the metrics
func (*aaaCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- accountingRequestsDesc
	ch <- accountingRequestFailuresDesc
	ch <- accountingRequestSuccessDesc
	ch <- accountingResponseFailuresDesc
	ch <- accountingResponseSuccessDesc
	ch <- accountingTimeoutsDesc
	ch <- accountingRequestsPendingDesc
	ch <- accountingMalformedResponsesDesc
	ch <- accountingRetransmissionsDesc
	ch <- accountingBadAuthenticatorsDesc
	ch <- accountingPacketsDroppedDesc

	ch <- authenticationRequestsDesc
	ch <- authenticationAcceptsDesc
	ch <- authenticationRejectsDesc
	ch <- authenticationRadiusFailuresDesc
	ch <- authenticationInvalidCredentialsDesc
	ch <- authenticationMalformedRequestDesc
	ch <- authenticationInternalFailureDesc
	ch <- authenticationLocalFailuresDesc
	ch <- authenticationLdapFailuresDesc
	ch <- authenticationChallengesDesc
	ch <- authenticationTimeoutsDesc

	ch <- radiusServerMaxOutstandingDesc
	ch <- radiusServerCurrentOutstandingDesc
	ch <- radiusServerPeakOutstandingDesc
	ch <- radiusServerFailOutstandingDesc

	ch <- radiusServerLastRttDesc
	ch <- radiusServerAuthRequestsDesc
	ch <- radiusServerAuthRolloverRequestsDesc
	ch <- radiusServerAuthRetransmissionsDesc
	ch <- radiusServerAcceptsDesc
	ch <- radiusServerRejectsDesc
	ch <- radiusServerChallengesDesc
	ch <- radiusServerAuthMalformedResponsesDesc
	ch <- radiusServerAuthBadAuthenticatorsDesc
	ch <- radiusServerAuthRequestsPendingDesc
	ch <- radiusServerAuthTimeoutsDesc
	ch <- radiusServerAuthUnknownResponsesDesc
	ch <- radiusServerAuthPacketsDroppedDesc
	ch <- radiusServerAcctStartRequestsDesc
	ch <- radiusServerAcctInterimRequestsDesc
	ch <- radiusServerAcctStopRequestsDesc
	ch <- radiusServerAcctRolloverRequestsDesc
	ch <- radiusServerAcctRetransmissionsDesc
	ch <- radiusServerAcctStartResponseDesc
	ch <- radiusServerAcctInterimResponseDesc
	ch <- radiusServerAcctStopResponseDesc
	ch <- radiusServerAcctMalformedResponseDesc
	ch <- radiusServerAcctBadAuthenticatorsDesc
	ch <- radiusServerAcctRequestsPendingDesc
	ch <- radiusServerAcctTimeoutsDesc
	ch <- radiusServerAcctUnknownResponsesDesc
	ch <- radiusServerAcctPacketsDroppedDesc

	ch <- localCertRequestsDesc
	ch <- localCertFailedRequestsDesc
	ch <- localCertTotalResponsesDesc
	ch <- localCertConfiguredResponsesDesc
}

// Collect collects metrics from JunOS
func (c *aaaCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	err := c.collectAccounting(client, ch, labelValues)
	if err != nil {
		return err
	}
	err = c.collectAuthentication(client, ch, labelValues)
	if err != nil {
		return err
	}
	err = c.collectRadius(client, ch, labelValues)
	if err != nil {
		return err
	}
	err = c.collectRadiusServers(client, ch, labelValues)
	if err != nil {
		return err
	}
	err = c.collectLocalCert(client, ch, labelValues)
	if err != nil {
		return err
	}
	return nil
}

func (c *aaaCollector) collectAccounting(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = accountingRpc{}
	err := client.RunCommandAndParse("show network-access aaa statistics accounting detail", &x)
	if err != nil {
		return err
	}

	stats := x.Statistics.AccountingStatistics
	ch <- prometheus.MustNewConstMetric(accountingRequestsDesc, prometheus.CounterValue, float64(stats.Requests), labelValues...)
	ch <- prometheus.MustNewConstMetric(accountingRequestFailuresDesc, prometheus.CounterValue, float64(stats.RequestFailures), labelValues...)
	ch <- prometheus.MustNewConstMetric(accountingRequestSuccessDesc, prometheus.CounterValue, float64(stats.RequestSuccess), labelValues...)
	ch <- prometheus.MustNewConstMetric(accountingResponseFailuresDesc, prometheus.CounterValue, float64(stats.ResponseFailures), labelValues...)
	ch <- prometheus.MustNewConstMetric(accountingResponseSuccessDesc, prometheus.CounterValue, float64(stats.ResponseSuccess), labelValues...)
	ch <- prometheus.MustNewConstMetric(accountingTimeoutsDesc, prometheus.CounterValue, float64(stats.Timeouts), labelValues...)
	ch <- prometheus.MustNewConstMetric(accountingRequestsPendingDesc, prometheus.GaugeValue, float64(stats.RequestsPending), labelValues...)
	ch <- prometheus.MustNewConstMetric(accountingMalformedResponsesDesc, prometheus.CounterValue, float64(stats.MalformedResponses), labelValues...)
	ch <- prometheus.MustNewConstMetric(accountingRetransmissionsDesc, prometheus.CounterValue, float64(stats.Retransmissions), labelValues...)
	ch <- prometheus.MustNewConstMetric(accountingBadAuthenticatorsDesc, prometheus.CounterValue, float64(stats.BadAuthenticators), labelValues...)
	ch <- prometheus.MustNewConstMetric(accountingPacketsDroppedDesc, prometheus.CounterValue, float64(stats.PacketsDropped), labelValues...)

	return nil
}

func (c *aaaCollector) collectAuthentication(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = authenticationRpc{}
	err := client.RunCommandAndParse("show network-access aaa statistics authentication detail", &x)
	if err != nil {
		return err
	}

	stats := x.Statistics.AuthenticationStatistics
	ch <- prometheus.MustNewConstMetric(authenticationRequestsDesc, prometheus.CounterValue, float64(stats.Requests), labelValues...)
	ch <- prometheus.MustNewConstMetric(authenticationAcceptsDesc, prometheus.CounterValue, float64(stats.Accepts), labelValues...)
	ch <- prometheus.MustNewConstMetric(authenticationRejectsDesc, prometheus.CounterValue, float64(stats.Rejects), labelValues...)
	ch <- prometheus.MustNewConstMetric(authenticationRadiusFailuresDesc, prometheus.CounterValue, float64(stats.RadiusFailures), labelValues...)
	ch <- prometheus.MustNewConstMetric(authenticationInvalidCredentialsDesc, prometheus.CounterValue, float64(stats.InvalidCredentials), labelValues...)
	ch <- prometheus.MustNewConstMetric(authenticationMalformedRequestDesc, prometheus.CounterValue, float64(stats.MalformedRequest), labelValues...)
	ch <- prometheus.MustNewConstMetric(authenticationInternalFailureDesc, prometheus.CounterValue, float64(stats.InternalFailure), labelValues...)
	ch <- prometheus.MustNewConstMetric(authenticationLocalFailuresDesc, prometheus.CounterValue, float64(stats.LocalFailures), labelValues...)
	ch <- prometheus.MustNewConstMetric(authenticationLdapFailuresDesc, prometheus.CounterValue, float64(stats.LdapFailures), labelValues...)
	ch <- prometheus.MustNewConstMetric(authenticationChallengesDesc, prometheus.CounterValue, float64(stats.Challenges), labelValues...)
	ch <- prometheus.MustNewConstMetric(authenticationTimeoutsDesc, prometheus.CounterValue, float64(stats.Timeouts), labelValues...)

	return nil
}

func (c *aaaCollector) collectRadius(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = radiusRpc{}
	err := client.RunCommandAndParse("show network-access aaa statistics detail radius", &x)
	if err != nil {
		return err
	}

	for _, s := range x.Statistics.RadiusStatistics.Servers {
		l := append(labelValues, s.Address)
		ch <- prometheus.MustNewConstMetric(radiusServerMaxOutstandingDesc, prometheus.GaugeValue, float64(s.MaxOutstanding), l...)
		ch <- prometheus.MustNewConstMetric(radiusServerCurrentOutstandingDesc, prometheus.GaugeValue, float64(s.CurrentOutstanding), l...)
		ch <- prometheus.MustNewConstMetric(radiusServerPeakOutstandingDesc, prometheus.GaugeValue, float64(s.PeakOutstanding), l...)
		ch <- prometheus.MustNewConstMetric(radiusServerFailOutstandingDesc, prometheus.GaugeValue, float64(s.FailOutstanding), l...)
	}

	return nil
}

func (c *aaaCollector) collectRadiusServers(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = radiusServersRpc{}
	err := client.RunCommandAndParse("show network-access aaa radius-servers detail", &x)
	if err != nil {
		return err
	}

	stats := x.Information.Statistics
	for i, address := range stats.Address {
		l := append(labelValues, address)

		// Helper to safely get value from slice or return 0
		safeGet := func(slice []int64, idx int) float64 {
			if idx < len(slice) {
				return float64(slice[idx])
			}
			return 0
		}

		ch <- prometheus.MustNewConstMetric(radiusServerLastRttDesc, prometheus.GaugeValue, safeGet(stats.LastRtt, i), l...)
		ch <- prometheus.MustNewConstMetric(radiusServerAuthRequestsDesc, prometheus.CounterValue, safeGet(stats.AuthenticationRequests, i), l...)
		ch <- prometheus.MustNewConstMetric(radiusServerAuthRolloverRequestsDesc, prometheus.CounterValue, safeGet(stats.AuthenticationRolloverRequests, i), l...)
		ch <- prometheus.MustNewConstMetric(radiusServerAuthRetransmissionsDesc, prometheus.CounterValue, safeGet(stats.AuthenticationRetransmissions, i), l...)
		ch <- prometheus.MustNewConstMetric(radiusServerAcceptsDesc, prometheus.CounterValue, safeGet(stats.Accepts, i), l...)
		ch <- prometheus.MustNewConstMetric(radiusServerRejectsDesc, prometheus.CounterValue, safeGet(stats.Rejects, i), l...)
		ch <- prometheus.MustNewConstMetric(radiusServerChallengesDesc, prometheus.CounterValue, safeGet(stats.Challenges, i), l...)
		ch <- prometheus.MustNewConstMetric(radiusServerAuthMalformedResponsesDesc, prometheus.CounterValue, safeGet(stats.AuthenticationMalformedResponses, i), l...)
		ch <- prometheus.MustNewConstMetric(radiusServerAuthBadAuthenticatorsDesc, prometheus.CounterValue, safeGet(stats.AuthenticationBadAuthenticators, i), l...)
		ch <- prometheus.MustNewConstMetric(radiusServerAuthRequestsPendingDesc, prometheus.GaugeValue, safeGet(stats.AuthenticationRequestsPending, i), l...)
		ch <- prometheus.MustNewConstMetric(radiusServerAuthTimeoutsDesc, prometheus.CounterValue, safeGet(stats.AuthenticationTimeouts, i), l...)
		ch <- prometheus.MustNewConstMetric(radiusServerAuthUnknownResponsesDesc, prometheus.CounterValue, safeGet(stats.AuthenticationUnknownResponses, i), l...)
		ch <- prometheus.MustNewConstMetric(radiusServerAuthPacketsDroppedDesc, prometheus.CounterValue, safeGet(stats.AuthenticationPacketsDropped, i), l...)
		ch <- prometheus.MustNewConstMetric(radiusServerAcctStartRequestsDesc, prometheus.CounterValue, safeGet(stats.AccountingStartRequests, i), l...)
		ch <- prometheus.MustNewConstMetric(radiusServerAcctInterimRequestsDesc, prometheus.CounterValue, safeGet(stats.AccountingInterimRequests, i), l...)
		ch <- prometheus.MustNewConstMetric(radiusServerAcctStopRequestsDesc, prometheus.CounterValue, safeGet(stats.AccountingStopRequests, i), l...)
		ch <- prometheus.MustNewConstMetric(radiusServerAcctRolloverRequestsDesc, prometheus.CounterValue, safeGet(stats.AccountingRolloverRequests, i), l...)
		ch <- prometheus.MustNewConstMetric(radiusServerAcctRetransmissionsDesc, prometheus.CounterValue, safeGet(stats.AccountingRetransmissions, i), l...)
		ch <- prometheus.MustNewConstMetric(radiusServerAcctStartResponseDesc, prometheus.CounterValue, safeGet(stats.AccountingStartResponse, i), l...)
		ch <- prometheus.MustNewConstMetric(radiusServerAcctInterimResponseDesc, prometheus.CounterValue, safeGet(stats.AccountingInterimResponse, i), l...)
		ch <- prometheus.MustNewConstMetric(radiusServerAcctStopResponseDesc, prometheus.CounterValue, safeGet(stats.AccountingStopResponse, i), l...)
		ch <- prometheus.MustNewConstMetric(radiusServerAcctMalformedResponseDesc, prometheus.CounterValue, safeGet(stats.AccountingMalformedResponse, i), l...)
		ch <- prometheus.MustNewConstMetric(radiusServerAcctBadAuthenticatorsDesc, prometheus.CounterValue, safeGet(stats.AccountingBadAuthenticators, i), l...)
		ch <- prometheus.MustNewConstMetric(radiusServerAcctRequestsPendingDesc, prometheus.GaugeValue, safeGet(stats.AccountingRequestsPending, i), l...)
		ch <- prometheus.MustNewConstMetric(radiusServerAcctTimeoutsDesc, prometheus.CounterValue, safeGet(stats.AccountingTimeouts, i), l...)
		ch <- prometheus.MustNewConstMetric(radiusServerAcctUnknownResponsesDesc, prometheus.CounterValue, safeGet(stats.AccountingUnknownResponses, i), l...)
		ch <- prometheus.MustNewConstMetric(radiusServerAcctPacketsDroppedDesc, prometheus.CounterValue, safeGet(stats.AccountingPacketsDropped, i), l...)
	}

	return nil
}

func (c *aaaCollector) collectLocalCert(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = localCertificateRpc{}
	err := client.RunCommandAndParse("show network-access local-certificate statistics extensive", &x)
	if err != nil {
		return err
	}

	for _, data := range x.Information.Table.InnerTable.Data {
		switch data.Name {
		case "total-requests":
			ch <- prometheus.MustNewConstMetric(localCertRequestsDesc, prometheus.CounterValue, float64(data.Value), labelValues...)
		case "failed-requests":
			ch <- prometheus.MustNewConstMetric(localCertFailedRequestsDesc, prometheus.CounterValue, float64(data.Value), labelValues...)
		case "total-responses":
			ch <- prometheus.MustNewConstMetric(localCertTotalResponsesDesc, prometheus.CounterValue, float64(data.Value), labelValues...)
		case "cofigured-responses":
			ch <- prometheus.MustNewConstMetric(localCertConfiguredResponsesDesc, prometheus.CounterValue, float64(data.Value), labelValues...)
		}
	}

	return nil
}
