package tagged

import (
	"context"

	"github.com/AikoCute-Offical/Aiko-Core/common/net"
)

type DialFunc func(ctx context.Context, dest net.Destination, tag string) (net.Conn, error)

var Dialer DialFunc
