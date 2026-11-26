//go:build go1.26

package handshake

import (
	"testing"

	utls "github.com/metacubex/utls"
)

func TestUTLSConnectionStateToStdHelloRetryRequest(t *testing.T) {
	got := connectionStateToStd(utls.ConnectionState{HelloRetryRequest: true})
	if !got.HelloRetryRequest {
		t.Fatal("HelloRetryRequest = false, want true")
	}
}
