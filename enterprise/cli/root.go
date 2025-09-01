package cli

import (
	"github.com/DanielRondonGarcia/coder/v2/cli"
	"github.com/DanielRondonGarcia/serpent"
)

type RootCmd struct {
	cli.RootCmd
}

func (r *RootCmd) enterpriseOnly() []*serpent.Command {
	return []*serpent.Command{
		r.Server(nil),
		r.workspaceProxy(),
		r.features(),
		r.licenses(),
		r.groups(),
		r.prebuilds(),
		r.provisionerDaemons(),
		r.provisionerd(),
		r.externalWorkspaces(),
	}
}

func (r *RootCmd) EnterpriseSubcommands() []*serpent.Command {
	all := append(r.CoreSubcommands(), r.enterpriseOnly()...)
	return all
}
