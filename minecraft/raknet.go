package minecraft

import (
	"context"
	"net"

	"github.com/sandertv/go-raknet"
	"github.com/sirupsen/logrus"
)

// RakNet is an implementation of a RakNet v10 Network.
type RakNet struct {
	l *logrus.Logger
}

// DialContext ...
func (r RakNet) DialContext(ctx context.Context, address string) (net.Conn, error) {
	return raknet.Dialer{ErrorLog: r.l.WithField("net origin", "raknet").Logger}.DialContext(ctx, address)
}

// PingContext ...
func (r RakNet) PingContext(ctx context.Context, address string) (response []byte, err error) {
	return raknet.Dialer{ErrorLog: r.l.WithField("net origin", "raknet").Logger}.PingContext(ctx, address)
}

// Listen ...
func (r RakNet) Listen(address string) (NetworkListener, error) {
	return raknet.ListenConfig{ErrorLog: r.l.WithField("net origin", "raknet").Logger}.Listen(address)
}

// init registers the RakNet network.
func init() {
	RegisterNetwork("raknet", func(l *logrus.Logger) Network { return RakNet{l: l} })
}
