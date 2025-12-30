
package aaa

type accountingRpc struct {
	Statistics struct {
		AccountingStatistics struct {
			Requests           int64 `xml:"requests"`
			RequestFailures    int64 `xml:"accounting-request-failures"`
			RequestSuccess     int64 `xml:"accounting-request-success"`
			ResponseFailures   int64 `xml:"accounting-response-failures"`
			ResponseSuccess    int64 `xml:"accounting-response-success"`
			Timeouts           int64 `xml:"timeouts"`
			RequestsPending    int64 `xml:"acct-requests-pending"`
			MalformedResponses int64 `xml:"acct-malformed-responses"`
			Retransmissions    int64 `xml:"acct-retransmissions"`
			BadAuthenticators  int64 `xml:"acct-bad-authenticators"`
			PacketsDropped     int64 `xml:"acct-packets-dropped"`
		} `xml:"aaa-module-accounting-statistics"`
	} `xml:"aaa-module-statistics"`
}

type authenticationRpc struct {
	Statistics struct {
		AuthenticationStatistics struct {
			Requests              int64 `xml:"requests"`
			Accepts               int64 `xml:"accepts"`
			Rejects               int64 `xml:"rejects"`
			RadiusFailures        int64 `xml:"radius-failures"`
			InvalidCredentials    int64 `xml:"rejects-invalid-credentials"`
			MalformedRequest      int64 `xml:"rejects-malformed-request"`
			InternalFailure       int64 `xml:"rejects-internal-failure"`
			LocalFailures         int64 `xml:"local-failures"`
			LdapFailures          int64 `xml:"ldap-failures"`
			Challenges            int64 `xml:"challenges"`
			Timeouts              int64 `xml:"timeouts"`
		} `xml:"aaa-module-authentication-statistics"`
	} `xml:"aaa-module-statistics"`
}

type radiusRpc struct {
    Statistics struct {
        RadiusStatistics struct {
            Servers []struct {
                Address           string `xml:"server-address"`
                Profile           string `xml:"profile"`
                MaxOutstanding    int64  `xml:"max-outstanding"`
                CurrentOutstanding int64  `xml:"current-outstanding"`
                PeakOutstanding    int64  `xml:"peak-outstanding"`
                FailOutstanding    int64  `xml:"fail-outstanding"`
            } `xml:"radius-server"`
        } `xml:"aaa-module-radius-statistics"`
    } `xml:"aaa-module-statistics"`
}

type radiusServersRpc struct {
	Information struct {
		Statistics struct {
			Address                         []string `xml:"server-address"`
			LastRtt                         []int64  `xml:"last-rtt"`
			AuthenticationRequests          []int64  `xml:"authentication-requests"`
			AuthenticationRolloverRequests  []int64  `xml:"authentication-rollover-requests"`
			AuthenticationRetransmissions   []int64  `xml:"authentication-retransmissions"`
			Accepts                         []int64  `xml:"accepts"`
			Rejects                         []int64  `xml:"rejects"`
			Challenges                      []int64  `xml:"challenges"`
			AuthenticationMalformedResponses []int64 `xml:"authentication-malformed-responses"`
			AuthenticationBadAuthenticators []int64  `xml:"authentication-bad-authenticators"`
			AuthenticationRequestsPending   []int64  `xml:"authentication-requests-pending"`
			AuthenticationTimeouts          []int64  `xml:"authentication-timeouts"`
			AuthenticationUnknownResponses  []int64  `xml:"authentication-unknown-responses"`
			AuthenticationPacketsDropped    []int64  `xml:"authentication-packets-dropped"`
			AccountingStartRequests         []int64  `xml:"accounting-start-requests"`
			AccountingInterimRequests       []int64  `xml:"accounting-interim-requests"`
			AccountingStopRequests          []int64  `xml:"accounting-stop-requests"`
			AccountingRolloverRequests      []int64  `xml:"accounting-rollover-requests"`
			AccountingRetransmissions       []int64  `xml:"accounting-retransmissions"`
			AccountingStartResponse         []int64  `xml:"accounting-start-response"`
			AccountingInterimResponse       []int64  `xml:"accounting-interim-response"`
			AccountingStopResponse          []int64  `xml:"accounting-stop-response"`
			AccountingMalformedResponse     []int64  `xml:"accounting-malformed-response"`
			AccountingBadAuthenticators     []int64  `xml:"accounting-bad-authenticators"`
			AccountingRequestsPending       []int64  `xml:"accounting-requests-pending"`
			AccountingTimeouts              []int64  `xml:"accounting-timeouts"`
			AccountingUnknownResponses      []int64  `xml:"accounting-unknown-responses"`
			AccountingPacketsDropped        []int64  `xml:"accounting-packets-dropped"`
			DynamicRequestCoa               []int64  `xml:"dynamic-request-coa"`
			DynamicRequestDisconnect        []int64  `xml:"dynamic-request-disconnect"`
			DynamicRequestUnknown           []int64  `xml:"dynamic-request-unknown"`
			DynamicRequestMalformed         []int64  `xml:"dynamic-request-malformed"`
			DynamicRequestBadAuthenticators []int64  `xml:"dynamic-request-bad-authenticators"`
			DynamicRequestInvalidLength     []int64  `xml:"dynamic-request-invalid-length"`
		} `xml:"aaa-module-radius-servers-statistics"`
	} `xml:"aaa-module-radius-servers-information"`
}

type localCertificateRpc struct {
	Information struct {
		Table struct {
			InnerTable struct {
				Data []struct {
					Name  string `xml:"local-cert-counter-name"`
					Value int64  `xml:"local-cert-counter-value"`
				} `xml:"local-cert-statistics-data"`
			} `xml:"local-cert-statistics-table"`
		} `xml:"local-cert-statistics-table"`
	} `xml:"local-cert-statistics-information"`
}
