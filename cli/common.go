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

package cli

import (
	"log"
	"os"

	"github.com/mitchellh/colorstring"
)

type Ui struct {
	color bool

	out *log.Logger

	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
}

func NewUi() *Ui {
	ui := &Ui{}
	ui.color = true

	ui.out = log.New(os.Stdout, "", 0)

	ui.debug = log.New(os.Stdout, "", 0)
	ui.info = log.New(os.Stdout, "", 0)
	ui.warn = log.New(os.Stdout, "", 0)
	ui.error = log.New(os.Stderr, "", 0)

	return ui
}

func (u *Ui) NoColor(disable bool) *Ui {
	if disable {
		u.color = false
	} else {
		u.color = true
	}

	return u
}

func (u *Ui) Out(content string) {
	u.out.Print(content)
}

func (u *Ui) Debug(content string) {
	if u.color {
		u.info.Println(colorstring.Color("[magenta]" + content))
	} else {
		u.info.Println(content)
	}
}

func (u *Ui) Info(content string) {
	if u.color {
		u.info.Println(colorstring.Color("[green]" + content))
	} else {
		u.info.Println(content)
	}
}

func (u *Ui) Warn(content string) {
	if u.color {
		u.warn.Println(colorstring.Color("[yellow]" + content))
	} else {
		u.warn.Println(content)
	}
}

func (u *Ui) Error(content string) {
	if u.color {
		u.error.Println(colorstring.Color("[red]" + content))
	} else {
		u.error.Println(content)
	}
}

func (u *Ui) Fatal(content string) {
	u.Error(content)
	os.Exit(1)
}

