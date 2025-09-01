package cli

import (
	"fmt"
	"strconv"

	"github.com/coder/coder/v2/cli/cliui"
	"github.com/coder/coder/v2/cli/cliutil"
	"github.com/coder/coder/v2/codersdk"
	"github.com/coder/serpent"
)

func (r *RootCmd) licenses() *serpent.Command {
	cmd := &serpent.Command{
		Short:   "Add, delete, and list licenses",
		Use:     "licenses",
		Aliases: []string{"license"},
		Handler: func(inv *serpent.Invocation) error {
			return inv.Command.HelpHandler(inv)
		},
		Children: []*serpent.Command{
			r.licenseAdd(),
			r.licensesList(),
			r.licenseDelete(),
		},
	}
	return cmd
}

func (r *RootCmd) licenseAdd() *serpent.Command {
	client := new(codersdk.Client)
	cmd := &serpent.Command{
		Use:   "add [-f file | -l license]",
		Short: "Add license to Coder deployment",
		Middleware: serpent.Chain(
			serpent.RequireNArgs(0),
			r.InitClient(client),
		),
		Handler: func(inv *serpent.Invocation) error {
			_, _ = fmt.Fprintf(inv.Stdout, "License functionality is always enabled. No license required.\n")
			return nil
		},
	}
	return cmd
}

func (r *RootCmd) licensesList() *serpent.Command {
	formatter := cliutil.NewLicenseFormatter()
	client := new(codersdk.Client)
	cmd := &serpent.Command{
		Use:     "list",
		Short:   "List licenses (including expired)",
		Aliases: []string{"ls"},
		Middleware: serpent.Chain(
			serpent.RequireNArgs(0),
			r.InitClient(client),
		),
		Handler: func(inv *serpent.Invocation) error {
			_, _ = fmt.Fprintf(inv.Stdout, "License functionality is always enabled. No licenses to display.\n")
			return nil
		},
	}
	formatter.AttachOptions(&cmd.Options)
	return cmd
}

func (r *RootCmd) licenseDelete() *serpent.Command {
	client := new(codersdk.Client)
	cmd := &serpent.Command{
		Use:     "delete <id>",
		Short:   "Delete license by ID",
		Aliases: []string{"del"},
		Middleware: serpent.Chain(
			serpent.RequireNArgs(1),
			r.InitClient(client),
		),
		Handler: func(inv *serpent.Invocation) error {
			_, _ = fmt.Fprintf(inv.Stdout, "License functionality is always enabled. No license deletion required.\n")
			return nil
		},
	}
	return cmd
}
