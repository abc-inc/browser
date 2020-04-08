// Copyright 2020 The browser Authors. All rights reserved.
// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package browser provides utilities for interacting with users' browsers.
package browser

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// Commands returns a list of possible commands to use to open a url.
func Commands() [][]string {
	var cmds [][]string
	if exe := os.Getenv("BROWSER"); exe != "" {
		for _, e := range strings.Split(exe, ":") {
			cmds = append(cmds, strings.Split(e, " "))
		}
	}
	switch runtime.GOOS {
	case "darwin":
		cmds = append(cmds, []string{"/usr/bin/open"})
	case "windows":
		cmds = append(cmds, []string{"cmd.exe", "/c", "start"})
	default:
		if os.Getenv("DISPLAY") != "" {
			// xdg-open is only for use in a desktop environment.
			cmds = append(cmds, []string{"xdg-open"})
		} else if _, err := os.Stat("/proc/version"); err != nil {
			// WSL reports to be linux and if there's no X server available, fallback to Windows.
			if v, err := ioutil.ReadFile("/proc/version"); err != nil && bytes.Contains(bytes.ToLower(v), []byte("microsoft")) {
				cmds = append(cmds, []string{"cmd.exe", "/c", "start"})
			}
		}
	}
	cmds = append(cmds,
		[]string{"chrome"},
		[]string{"google-chrome"},
		[]string{"chromium"},
		[]string{"firefox"},
	)
	return cmds
}

// Open tries to open url in a browser and reports whether it succeeded.
func Open(url string) bool {
	return OpenCmd(url, func(cmd *exec.Cmd) *exec.Cmd { return cmd })
}

// OpenCmd tries to open url in a browser using customized Cmds and reports whether it succeeded.
// If cust returns nil, the Cmd is skipped.
func OpenCmd(url string, cust func(cmd *exec.Cmd) *exec.Cmd) bool {
	for _, args := range Commands() {
		cmd := exec.Command(args[0], append(args[1:], url)...)
		if cmd = cust(cmd); cmd == nil {
			continue
		}
		if cmd.Start() == nil && appearsSuccessful(cmd, 3*time.Second) {
			return true
		}
	}
	return false
}

// appearsSuccessful reports whether the command appears to have run successfully.
// If the command runs longer than the timeout, it's deemed successful.
// If the command runs within the timeout, it's deemed successful if it exited cleanly.
func appearsSuccessful(cmd *exec.Cmd, timeout time.Duration) bool {
	errc := make(chan error, 1)
	go func() {
		errc <- cmd.Wait()
	}()

	select {
	case <-time.After(timeout):
		return true
	case err := <-errc:
		return err == nil
	}
}
