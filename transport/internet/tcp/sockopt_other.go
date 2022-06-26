//go:build !linux && !freebsd
// +build !linux,!freebsd

package tcp

import (
	"github.com/AikoCute-Offical/Aiko-Core/common/net"
	"github.com/AikoCute-Offical/Aiko-Core/transport/internet/stat"
)

func GetOriginalDestination(conn stat.Connection) (net.Destination, error) {
	return net.Destination{}, nil
}
