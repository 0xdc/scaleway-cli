package vpc

import (
	"github.com/scaleway/scaleway-cli/v2/internal/core"
)

func GetCommands() *core.Commands {
	cmds := GetGeneratedCommands()

	cmds.MustFind("vpc", "private-network", "get").Override(privateNetworkGetBuilder)

	return cmds
}
