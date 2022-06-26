package command_test

import (
	"context"
	"testing"

	"github.com/AikoCute-Offical/Aiko-Core/app/dispatcher"
	"github.com/AikoCute-Offical/Aiko-Core/app/log"
	. "github.com/AikoCute-Offical/Aiko-Core/app/log/command"
	"github.com/AikoCute-Offical/Aiko-Core/app/proxyman"
	_ "github.com/AikoCute-Offical/Aiko-Core/app/proxyman/inbound"
	_ "github.com/AikoCute-Offical/Aiko-Core/app/proxyman/outbound"
	"github.com/AikoCute-Offical/Aiko-Core/common"
	"github.com/AikoCute-Offical/Aiko-Core/common/serial"
	"github.com/AikoCute-Offical/Aiko-Core/core"
)

func TestLoggerRestart(t *testing.T) {
	v, err := core.New(&core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{}),
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	})
	common.Must(err)
	common.Must(v.Start())

	server := &LoggerServer{
		V: v,
	}
	common.Must2(server.RestartLogger(context.Background(), &RestartLoggerRequest{}))
}
