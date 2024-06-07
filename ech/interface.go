package quic

import (
	"github.com/sagernet/quic-go"
	"github.com/sagernet/quic-go/internal/handshake_ech"
	"github.com/sagernet/quic-go/internal/protocol"
)

// The StreamID is the ID of a QUIC stream.
type StreamID = protocol.StreamID

// A Version is a QUIC version number.
type Version = protocol.Version

// A VersionNumber is a QUIC version number.
// Deprecated: VersionNumber was renamed to Version.
type VersionNumber = Version

const (
	// Version1 is RFC 9000
	Version1 = protocol.Version1
	// Version2 is RFC 9369
	Version2 = protocol.Version2
)

// A ClientToken is a token received by the client.
// It can be used to skip address validation on future connection attempts.
type ClientToken = quic.ClientToken

type TokenStore = quic.TokenStore

// Err0RTTRejected is the returned from:
// * Open{Uni}Stream{Sync}
// * Accept{Uni}Stream
// * Stream.Read and Stream.Write
// when the server rejects a 0-RTT connection attempt.
var Err0RTTRejected = quic.Err0RTTRejected

// ConnectionTracingKey can be used to associate a ConnectionTracer with a Connection.
// It is set on the Connection.Context() context,
// as well as on the context passed to logging.Tracer.NewConnectionTracer.
// Deprecated: Applications can set their own tracing key using Transport.ConnContext.
var ConnectionTracingKey = quic.ConnectionTracingKey

// ConnectionTracingID is the type of the context value saved under the ConnectionTracingKey.
// Deprecated: Applications can set their own tracing key using Transport.ConnContext.
type ConnectionTracingID = quic.ConnectionTracingID

// QUICVersionContextKey can be used to find out the QUIC version of a TLS handshake from the
// context returned by tls.Config.ClientHelloInfo.Context.
var QUICVersionContextKey = handshake.QUICVersionContextKey

// Stream is the interface implemented by QUIC streams
// In addition to the errors listed on the Connection,
// calls to stream functions can return a StreamError if the stream is canceled.
type Stream = quic.Stream

// A ReceiveStream is a unidirectional Receive Stream.
type ReceiveStream = quic.ReceiveStream

// A SendStream is a unidirectional Send Stream.
type SendStream = quic.SendStream

// A Connection is a QUIC connection between two peers.
// Calls to the connection (and to streams) can return the following types of errors:
// * ApplicationError: for errors triggered by the application running on top of QUIC
// * TransportError: for errors triggered by the QUIC transport (in many cases a misbehaving peer)
// * IdleTimeoutError: when the peer goes away unexpectedly (this is a net.Error timeout error)
// * HandshakeTimeoutError: when the cryptographic handshake takes too long (this is a net.Error timeout error)
// * StatelessResetError: when we receive a stateless reset (this is a net.Error temporary error)
// * VersionNegotiationError: returned by the client, when there's no version overlap between the peers
type Connection = quic.Connection

// An EarlyConnection is a connection that is handshaking.
// Data sent during the handshake is encrypted using the forward secure keys.
// When using client certificates, the client's identity is only verified
// after completion of the handshake.
type EarlyConnection = quic.EarlyConnection

// StatelessResetKey is a key used to derive stateless reset tokens.
type StatelessResetKey = quic.StatelessResetKey

// TokenGeneratorKey is a key used to encrypt session resumption tokens.
type TokenGeneratorKey = handshake.TokenProtectorKey

// A ConnectionID is a QUIC Connection ID, as defined in RFC 9000.
// It is not able to handle QUIC Connection IDs longer than 20 bytes,
// as they are allowed by RFC 8999.
type ConnectionID = quic.ConnectionID

// ConnectionIDFromBytes interprets b as a Connection ID. It panics if b is
// longer than 20 bytes.
func ConnectionIDFromBytes(b []byte) ConnectionID {
	return protocol.ParseConnectionID(b)
}

// A ConnectionIDGenerator is an interface that allows clients to implement their own format
// for the Connection IDs that servers/clients use as SrcConnectionID in QUIC packets.
//
// Connection IDs generated by an implementation should always produce IDs of constant size.
type ConnectionIDGenerator = quic.ConnectionIDGenerator

// Config contains all configuration data needed for a QUIC server or client.
type Config = quic.Config

type ClientHelloInfo = quic.ClientHelloInfo

// ConnectionState records basic details about a QUIC connection
type ConnectionState = quic.ConnectionState
