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

	cmd.AddCommand(NewPongPublishCommand().Command)
	cmd.AddCommand(NewPongStreamCommand().Command)
	cmd.AddCommand(NewPongSubscribeCommand().Command)

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