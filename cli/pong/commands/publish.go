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

type PongPublishCommand struct {
	*cobra.Command
	*cli.Ui
}

func NewPongPublishCommand() *PongPublishCommand {
	cmd := &PongPublishCommand{}
	cmd.Command = &cobra.Command{}
	cmd.Ui = cli.NewUi()
	cmd.Use = "publish"
	cmd.Short = "Publish a pong message to an address"
	cmd.Long = "Publish a pong message to an address"
	cmd.PreRunE = cmd.setup
	cmd.RunE = cmd.run

	return cmd
}

func (p *PongPublishCommand) setup(cmd *cobra.Command, args []string) error {
	color, err := cmd.Flags().GetBool("no-color")

	p.NoColor(color)

	return err
}

func (p *PongPublishCommand) run(cmd *cobra.Command, args []string) error {
	var msg string
	var address string

	if cmd.Flags().Arg(0) == "" {
		return errors.New("argument MESSAGE required")
	} else {
		msg = cmd.Flags().Arg(0)
	}

	if cmd.Flags().Arg(1) == "" {
		address = "*"
	}

	eb := pong.NewEventBus("")

	err := eb.Start()
	if err != nil {
		return err
	}

	message := &pong.Message{}
	message.Address = address
	message.Data = make(map[string]interface{})
	message.Data["message"] = msg

	id, err := eb.Publish(message)
	if err != nil {
		return err
	}

	p.Out(fmt.Sprint("Sent message -> ", id))

	return nil
}

