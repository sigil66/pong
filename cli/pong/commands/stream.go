package commands

import (

	"github.com/solvent-io/pong"
	"github.com/solvent-io/pong/cli"
	"github.com/spf13/cobra"
	"fmt"

	"errors"
)

type PongStreamCommand struct {
	*cobra.Command
	*cli.Ui
}

func NewPongStreamCommand() *PongStreamCommand {
	cmd := &PongStreamCommand{}
	cmd.Command = &cobra.Command{}
	cmd.Ui = cli.NewUi()
	cmd.Use = "stream"
	cmd.Short = "Stream all events arriving on the eventbus"
	cmd.Long = "Stream all events arriving on the eventbus"
	cmd.PreRunE = cmd.setup
	cmd.RunE = cmd.run

	return cmd
}

func (p *PongStreamCommand) setup(cmd *cobra.Command, args []string) error {
	color, err := cmd.Flags().GetBool("no-color")

	p.NoColor(color)

	return err
}

func (p *PongStreamCommand) run(cmd *cobra.Command, args []string) error {

	eb := pong.NewEventBus("")

	eb.On("error", func(err string) {
		p.Warn(err)
	})

	eb.On("message", func(msg *pong.Message) {
		fmt.Println(msg)
	})


	err := eb.Start()
	if err != nil {
		return err
	}

	select {
		case result := <-eb.Shutdown:
			switch result {
			case 0:
				return nil
			case 1:
				return errors.New("fatal eventbus shutdown")
			}
	}

	return nil
}
