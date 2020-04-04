// Copyright 2020 The browser Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package browser_test

import (
	"github.com/abc-inc/browser"
	"os"
	"strconv"
	"testing"
)

func TestCommands(t *testing.T) {
	_ = os.Setenv("BROWSER", "echo")
	commands := browser.Commands()
	if len(commands) < 5 {
		t.Error("want >=5 commands, got " + strconv.Itoa(len(commands)))
	}
}
