//go:build go1.26

package handshake

import (
	"crypto/tls"

	utls "github.com/metacubex/utls"
)

func copyUTLSHelloRetryRequest(dst *tls.ConnectionState, src utls.ConnectionState) {
	dst.HelloRetryRequest = src.HelloRetryRequest
}
