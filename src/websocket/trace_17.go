//go:build !go1.8
// +build !go1.8

package websocket

import (
	"github.com/linuxliu/SunnyNet/src/crypto/tls"
	"github.com/linuxliu/SunnyNet/src/http/httptrace"
)

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	return doHandshake(tlsConn, cfg)
}
