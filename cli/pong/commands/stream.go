/*
 * Copyright 2017 Zachary Schneider
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at

 * http://www.apache.org/licenses/LICENSE-2.0

 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
