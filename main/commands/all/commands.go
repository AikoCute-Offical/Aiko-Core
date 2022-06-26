package all

import (
	"github.com/AikoCute-Offical/Aiko-Core/main/commands/all/api"
	"github.com/AikoCute-Offical/Aiko-Core/main/commands/all/tls"
	"github.com/AikoCute-Offical/Aiko-Core/main/commands/base"
)

// go:generate go run github.com/AikoCute-Offical/Aiko-Core/common/errors/errorgen

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		// cmdConvert,
		tls.CmdTLS,
		cmdUUID,
	)
}
