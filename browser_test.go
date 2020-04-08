// Copyright 2020 The browser Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package browser_test

import (
	"os"
	"os/exec"
	"strconv"
	"testing"

	"github.com/abc-inc/browser"
)

func TestCommands(t *testing.T) {
	_ = os.Setenv("BROWSER", "echo")
	commands := browser.Commands()
	if len(commands) < 5 {
		t.Error("want >=5 commands, got " + strconv.Itoa(len(commands)))
	}
}

func TestOpenCmdNil(t *testing.T) {
	if browser.OpenCmd("/", func(cmd *exec.Cmd) *exec.Cmd { return nil }) {
		t.Error("want false, got true")
	}
}
