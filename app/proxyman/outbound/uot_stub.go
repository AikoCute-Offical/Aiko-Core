//go:build !go1.18

package outbound

import (
	"context"
	"os"

	"github.com/AikoCute-Offical/Aiko-Core/common/net"
	"github.com/AikoCute-Offical/Aiko-Core/transport/internet/stat"
)

func (h *Handler) getUoTConnection(ctx context.Context, dest net.Destination) (stat.Connection, error) {
	return nil, os.ErrInvalid
}
