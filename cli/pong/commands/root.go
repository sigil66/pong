package commands

import (
	"github.com/spf13/cobra"
	"github.com/solvent-io/pong/cli"
)


type PongRootCommand struct {
	*cobra.Command
	*cli.Ui
}

func NewPongRootCommand() *PongRootCommand {
	cmd := &PongRootCommand{}
	cmd.Command = &cobra.Command{}
	cmd.Ui = cli.NewUi()
	cmd.Use = "pong"
	cmd.Short = "pong, test command for pong eventbus"
	cmd.Long = "pong, test command for pong eventbus"
	cmd.PreRunE = cmd.setup
	cmd.RunE = cmd.run

	cmd.PersistentFlags().Bool("no-color", false, "Disable color")

	cmd.AddCommand(NewPongStreamCommand().Command)

	return cmd
}

func (p *PongRootCommand) setup(cmd *cobra.Command, args []string) error {
	color, err := cmd.Flags().GetBool("no-color")

	p.NoColor(color)

	return err
}

func (p *PongRootCommand) run(cmd *cobra.Command, args []string) error {
	p.Help()
	return nil
}