
package aaa

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

var accountingOutput = `
<rpc-reply xmlns:junos="http://xml.juniper.net/junos/22.4R0/junos">
    <aaa-module-statistics>
        <aaa-module-accounting-statistics>
            <requests>69725</requests>
            <accounting-request-failures>0</accounting-request-failures>
            <accounting-request-success>69725</accounting-request-success>
            <acct-acct-on-requests>0</acct-acct-on-requests>
            <acct-start-requests>0</acct-start-requests>
            <acct-interim-requests>0</acct-interim-requests>
            <acct-stop-requests>0</acct-stop-requests>
            <timeouts>16</timeouts>
            <accounting-response-failures>0</accounting-response-failures>
            <accounting-response-success>69709</accounting-response-success>
            <acct-acct-on-responses>0</acct-acct-on-responses>
            <acct-start-responses>0</acct-start-responses>
            <acct-interim-responses>0</acct-interim-responses>
            <acct-stop-responses>0</acct-stop-responses>
            <acct-rollover-requests>4014</acct-rollover-requests>
            <acct-unknown-responses>0</acct-unknown-responses>
            <acct-requests-pending>0</acct-requests-pending>
            <acct-malformed-responses>0</acct-malformed-responses>
            <acct-retransmissions>0</acct-retransmissions>
            <acct-bad-authenticators>0</acct-bad-authenticators>
            <acct-packets-dropped>0</acct-packets-dropped>
            <acct-backup-record-creation-requests>0</acct-backup-record-creation-requests>
            <acct-backup-request-replay-success>0</acct-backup-request-replay-success>
            <acct-backup-request-failures>0</acct-backup-request-failures>
            <acct-backup-request-success>0</acct-backup-request-success>
            <acct-backup-request-timeouts>0</acct-backup-request-timeouts>
            <acct-backup-in-flight-requests>0</acct-backup-in-flight-requests>
            <acct-backup-response-success>0</acct-backup-response-success>
            <acct-backup-radius-requests>0</acct-backup-radius-requests>
            <acct-backup-radius-responses>0</acct-backup-radius-responses>
            <acct-backup-radius-timeouts>0</acct-backup-radius-timeouts>
            <acct-backup-radius-pending-request>0</acct-backup-radius-pending-request>
            <acct-backup-radius-retransmissions>0</acct-backup-radius-retransmissions>
            <acct-backup-malformed-responses>0</acct-backup-malformed-responses>
            <acct-backup-bad-authenticators>0</acct-backup-bad-authenticators>
            <acct-backup-packets-dropped>0</acct-backup-packets-dropped>
            <acct-backup-rollover-requests>0</acct-backup-rollover-requests>
            <acct-backup-unknown-responses>0</acct-backup-unknown-responses>
        </aaa-module-accounting-statistics>
    </aaa-module-statistics>
    <cli>
        <banner>{master:0}</banner>
    </cli>                              
</rpc-reply>`

func TestParseAccounting(t *testing.T) {
	var x accountingRpc
	err := xml.Unmarshal([]byte(accountingOutput), &x)
	assert.NoError(t, err)
	assert.Equal(t, int64(69725), x.Statistics.AccountingStatistics.Requests)
	assert.Equal(t, int64(16), x.Statistics.AccountingStatistics.Timeouts)
}

var radiusServersOutput = `
<rpc-reply xmlns:junos="http://xml.juniper.net/junos/22.4R0/junos">
    <aaa-module-radius-servers-information>
        <aaa-module-profile-radius-servers>
            <profile-name>dot1x</profile-name>
            <server-address>192.0.2.1</server-address>
            <authentication-port>1812</authentication-port>
            <preauthentication-port>1812</preauthentication-port>
            <accounting-port>1813</accounting-port>
            <available-status>UP</available-status>
            <server-address>192.0.2.2</server-address>
            <authentication-port>1812</authentication-port>
            <preauthentication-port>1812</preauthentication-port>
            <accounting-port>1813</accounting-port>
            <available-status>UP</available-status>
            <server-address>2001:db8::1</server-address>
            <authentication-port>1812</authentication-port>
            <preauthentication-port>1812</preauthentication-port>
            <accounting-port>1813</accounting-port>
            <available-status>UP</available-status>
            <server-address>2001:db8::2</server-address>
            <authentication-port>1812</authentication-port>
            <preauthentication-port>1812</preauthentication-port>
            <accounting-port>1813</accounting-port>
            <available-status>UP</available-status>
        </aaa-module-profile-radius-servers>
        <aaa-module-radius-servers-statistics>
            <server-address>192.0.2.1</server-address>
            <last-rtt>77</last-rtt>
            <authentication-requests>78487</authentication-requests>
            <authentication-rollover-requests>1389</authentication-rollover-requests>
            <authentication-retransmissions>0</authentication-retransmissions>
            <accepts>38613</accepts>
            <rejects>36567</rejects>
            <challenges>3257</challenges>
            <authentication-malformed-responses>0</authentication-malformed-responses>
            <authentication-bad-authenticators>0</authentication-bad-authenticators>
            <authentication-requests-pending>0</authentication-requests-pending>
            <authentication-timeouts>50</authentication-timeouts>
            <authentication-unknown-responses>0</authentication-unknown-responses>
            <authentication-packets-dropped>0</authentication-packets-dropped>
            <preauthentication-requests>0</preauthentication-requests>
            <preauthentication-rollover-requests>0</preauthentication-rollover-requests>
            <preauthentication-retransmissions>0</preauthentication-retransmissions>
            <preauthentication-accepts>0</preauthentication-accepts>
            <preauthentication-rejects>0</preauthentication-rejects>
            <preauthentication-challenges>0</preauthentication-challenges>
            <preauthentication-malformed-responses>0</preauthentication-malformed-responses>
            <preauthentication-bad-authenticators>0</preauthentication-bad-authenticators>
            <preauthentication-requests-pending>0</preauthentication-requests-pending>
            <preauthentication-timeouts>0</preauthentication-timeouts>
            <preauthentication-unknown-responses>0</preauthentication-unknown-responses>
            <preauthentication-packets-dropped>0</preauthentication-packets-dropped>
            <accounting-start-requests>0</accounting-start-requests>
            <accounting-interim-requests>0</accounting-interim-requests>
            <accounting-stop-requests>0</accounting-stop-requests>
            <accounting-rollover-requests>1990</accounting-rollover-requests>
            <accounting-retransmissions>0</accounting-retransmissions>
            <accounting-start-response>0</accounting-start-response>
            <accounting-interim-response>0</accounting-interim-response>
            <accounting-stop-response>0</accounting-stop-response>
            <accounting-malformed-response>0</accounting-malformed-response>
            <accounting-bad-authenticators>0</accounting-bad-authenticators>
            <accounting-requests-pending>0</accounting-requests-pending>
            <accounting-timeouts>34</accounting-timeouts>
            <accounting-unknown-responses>0</accounting-unknown-responses>
            <accounting-packets-dropped>0</accounting-packets-dropped>
            <dynamic-request-coa>0</dynamic-request-coa>
            <dynamic-request-disconnect>0</dynamic-request-disconnect>
            <dynamic-request-unknown>0</dynamic-request-unknown>
            <dynamic-request-malformed>0</dynamic-request-malformed>
            <dynamic-request-bad-authenticators>0</dynamic-request-bad-authenticators>
            <dynamic-request-invalid-length>0</dynamic-request-invalid-length>
            <server-address>192.0.2.2</server-address>
            <last-rtt>1079</last-rtt>
            <authentication-requests>51</authentication-requests>
            <authentication-rollover-requests>50</authentication-rollover-requests>
            <authentication-retransmissions>0</authentication-retransmissions>
            <accepts>15</accepts>
            <rejects>9</rejects>
            <challenges>0</challenges>
            <authentication-malformed-responses>0</authentication-malformed-responses>
            <authentication-bad-authenticators>0</authentication-bad-authenticators>
            <authentication-requests-pending>0</authentication-requests-pending>
            <authentication-timeouts>27</authentication-timeouts>
            <authentication-unknown-responses>0</authentication-unknown-responses>
            <authentication-packets-dropped>0</authentication-packets-dropped>
            <preauthentication-requests>0</preauthentication-requests>
            <preauthentication-rollover-requests>0</preauthentication-rollover-requests>
            <preauthentication-retransmissions>0</preauthentication-retransmissions>
            <preauthentication-accepts>0</preauthentication-accepts>
            <preauthentication-rejects>0</preauthentication-rejects>
            <preauthentication-challenges>0</preauthentication-challenges>
            <preauthentication-malformed-responses>0</preauthentication-malformed-responses>
            <preauthentication-bad-authenticators>0</preauthentication-bad-authenticators>
            <preauthentication-requests-pending>0</preauthentication-requests-pending>
            <preauthentication-timeouts>0</preauthentication-timeouts>
            <preauthentication-unknown-responses>0</preauthentication-unknown-responses>
            <preauthentication-packets-dropped>0</preauthentication-packets-dropped>
            <accounting-start-requests>0</accounting-start-requests>
            <accounting-interim-requests>0</accounting-interim-requests>
            <accounting-stop-requests>0</accounting-stop-requests>
            <accounting-rollover-requests>34</accounting-rollover-requests>
            <accounting-retransmissions>0</accounting-retransmissions>
            <accounting-start-response>0</accounting-start-response>
            <accounting-interim-response>0</accounting-interim-response>
            <accounting-stop-response>0</accounting-stop-response>
            <accounting-malformed-response>0</accounting-malformed-response>
            <accounting-bad-authenticators>0</accounting-bad-authenticators>
            <accounting-requests-pending>0</accounting-requests-pending>
            <accounting-timeouts>16</accounting-timeouts>
            <accounting-unknown-responses>0</accounting-unknown-responses>
            <accounting-packets-dropped>0</accounting-packets-dropped>
            <dynamic-request-coa>0</dynamic-request-coa>
            <dynamic-request-disconnect>0</dynamic-request-disconnect>
            <dynamic-request-unknown>0</dynamic-request-unknown>
            <dynamic-request-malformed>0</dynamic-request-malformed>
            <dynamic-request-bad-authenticators>0</dynamic-request-bad-authenticators>
            <dynamic-request-invalid-length>0</dynamic-request-invalid-length>
            <server-address>2001:db8::1</server-address>
            <last-rtt>0</last-rtt>
            <authentication-requests>1389</authentication-requests>
            <authentication-rollover-requests>0</authentication-rollover-requests>
            <authentication-retransmissions>0</authentication-retransmissions>
            <accepts>0</accepts>
            <rejects>0</rejects>
            <challenges>0</challenges>
            <authentication-malformed-responses>0</authentication-malformed-responses>
            <authentication-bad-authenticators>0</authentication-bad-authenticators>
            <authentication-requests-pending>0</authentication-requests-pending>
            <authentication-timeouts>1389</authentication-timeouts>
            <authentication-unknown-responses>0</authentication-unknown-responses>
            <authentication-packets-dropped>0</authentication-packets-dropped>
            <preauthentication-requests>0</preauthentication-requests>
            <preauthentication-rollover-requests>0</preauthentication-rollover-requests>
            <preauthentication-retransmissions>0</preauthentication-retransmissions>
            <preauthentication-accepts>0</preauthentication-accepts>
            <preauthentication-rejects>0</preauthentication-rejects>
            <preauthentication-challenges>0</preauthentication-challenges>
            <preauthentication-malformed-responses>0</preauthentication-malformed-responses>
            <preauthentication-bad-authenticators>0</preauthentication-bad-authenticators>
            <preauthentication-requests-pending>0</preauthentication-requests-pending>
            <preauthentication-timeouts>0</preauthentication-timeouts>
            <preauthentication-unknown-responses>0</preauthentication-unknown-responses>
            <preauthentication-packets-dropped>0</preauthentication-packets-dropped>
            <accounting-start-requests>0</accounting-start-requests>
            <accounting-interim-requests>0</accounting-interim-requests>
            <accounting-stop-requests>0</accounting-stop-requests>
            <accounting-rollover-requests>0</accounting-rollover-requests>
            <accounting-retransmissions>0</accounting-retransmissions>
            <accounting-start-response>0</accounting-start-response>
            <accounting-interim-response>0</accounting-interim-response>
            <accounting-stop-response>0</accounting-stop-response>
            <accounting-malformed-response>0</accounting-malformed-response>
            <accounting-bad-authenticators>0</accounting-bad-authenticators>
            <accounting-requests-pending>0</accounting-requests-pending>
            <accounting-timeouts>1990</accounting-timeouts>
            <accounting-unknown-responses>0</accounting-unknown-responses>
            <accounting-packets-dropped>0</accounting-packets-dropped>
            <dynamic-request-coa>0</dynamic-request-coa>
            <dynamic-request-disconnect>0</dynamic-request-disconnect>
            <dynamic-request-unknown>0</dynamic-request-unknown>
            <dynamic-request-malformed>0</dynamic-request-malformed>
            <dynamic-request-bad-authenticators>0</dynamic-request-bad-authenticators>
            <dynamic-request-invalid-length>0</dynamic-request-invalid-length>
            <server-address>2001:db8::2</server-address>
            <last-rtt>0</last-rtt>
            <authentication-requests>1389</authentication-requests>
            <authentication-rollover-requests>1389</authentication-rollover-requests>
            <authentication-retransmissions>0</authentication-retransmissions>
            <accepts>0</accepts>
            <rejects>0</rejects>
            <challenges>0</challenges>
            <authentication-malformed-responses>0</authentication-malformed-responses>
            <authentication-bad-authenticators>0</authentication-bad-authenticators>
            <authentication-requests-pending>0</authentication-requests-pending>
            <authentication-timeouts>1389</authentication-timeouts>
            <authentication-unknown-responses>0</authentication-unknown-responses>
            <authentication-packets-dropped>0</authentication-packets-dropped>
            <preauthentication-requests>0</preauthentication-requests>
            <preauthentication-rollover-requests>0</preauthentication-rollover-requests>
            <preauthentication-retransmissions>0</preauthentication-retransmissions>
            <preauthentication-accepts>0</preauthentication-accepts>
            <preauthentication-rejects>0</preauthentication-rejects>
            <preauthentication-challenges>0</preauthentication-challenges>
            <preauthentication-malformed-responses>0</preauthentication-malformed-responses>
            <preauthentication-bad-authenticators>0</preauthentication-bad-authenticators>
            <preauthentication-requests-pending>0</preauthentication-requests-pending>
            <preauthentication-timeouts>0</preauthentication-timeouts>
            <preauthentication-unknown-responses>0</preauthentication-unknown-responses>
            <preauthentication-packets-dropped>0</preauthentication-packets-dropped>
            <accounting-start-requests>0</accounting-start-requests>
            <accounting-interim-requests>0</accounting-interim-requests>
            <accounting-stop-requests>0</accounting-stop-requests>
            <accounting-rollover-requests>1990</accounting-rollover-requests>
            <accounting-retransmissions>0</accounting-retransmissions>
            <accounting-start-response>0</accounting-start-response>
            <accounting-interim-response>0</accounting-interim-response>
            <accounting-stop-response>0</accounting-stop-response>
            <accounting-malformed-response>0</accounting-malformed-response>
            <accounting-bad-authenticators>0</accounting-bad-authenticators>
            <accounting-requests-pending>0</accounting-requests-pending>
            <accounting-timeouts>1990</accounting-timeouts>
            <accounting-unknown-responses>0</accounting-unknown-responses>
            <accounting-packets-dropped>0</accounting-packets-dropped>
            <dynamic-request-coa>0</dynamic-request-coa>
            <dynamic-request-disconnect>0</dynamic-request-disconnect>
            <dynamic-request-unknown>0</dynamic-request-unknown>
            <dynamic-request-malformed>0</dynamic-request-malformed>
            <dynamic-request-bad-authenticators>0</dynamic-request-bad-authenticators>
            <dynamic-request-invalid-length>0</dynamic-request-invalid-length>
        </aaa-module-radius-servers-statistics>
    </aaa-module-radius-servers-information>
</rpc-reply>`

func TestParseRadiusServers(t *testing.T) {
	var x radiusServersRpc
	err := xml.Unmarshal([]byte(radiusServersOutput), &x)
	assert.NoError(t, err)
	assert.Len(t, x.Information.Statistics.Address, 4)
}

var localCertOutput = `
<rpc-reply xmlns:junos="http://xml.juniper.net/junos/22.4R0/junos">
    <local-cert-statistics-information>
        <local-cert-statistics-table>
            <local-cert-statistics-table>
                <local-cert-statistics-data>
                    <local-cert-counter-name>total-requests</local-cert-counter-name>
                    <local-cert-counter-value>100920</local-cert-counter-value>
                </local-cert-statistics-data>
                <local-cert-statistics-data>
                    <local-cert-counter-name>failed-requests</local-cert-counter-name>
                    <local-cert-counter-value>0</local-cert-counter-value>
                </local-cert-statistics-data>
                <local-cert-statistics-data>
                    <local-cert-counter-name>total-responses</local-cert-counter-name>
                    <local-cert-counter-value>100919</local-cert-counter-value>
                </local-cert-statistics-data>
                <local-cert-statistics-data>
                    <local-cert-counter-name>cofigured-responses</local-cert-counter-name>
                    <local-cert-counter-value>100919</local-cert-counter-value>
                </local-cert-statistics-data>
            </local-cert-statistics-table>
        </local-cert-statistics-table>
    </local-cert-statistics-information>
</rpc-reply>`

func TestParseLocalCert(t *testing.T) {
	var x localCertificateRpc
	err := xml.Unmarshal([]byte(localCertOutput), &x)
	assert.NoError(t, err)
	assert.NotEmpty(t, x.Information.Table.InnerTable.Data)
}
