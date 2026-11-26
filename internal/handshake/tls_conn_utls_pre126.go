//go:build !go1.26

package handshake

import (
	"crypto/tls"

	utls "github.com/metacubex/utls"
)

func copyUTLSHelloRetryRequest(*tls.ConnectionState, utls.ConnectionState) {}
