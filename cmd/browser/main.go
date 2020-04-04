// Copyright 2020 The browser Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/abc-inc/browser"
	"log"
	"os"
	"path"
)

func main() {
	log.SetFlags(0)
	if len(os.Args) == 1 {
		usage()
	}

	browser.Open(os.Args[1])
}

func usage() {
	log.Println("Usage:", path.Base(os.Args[0]), " URL")
	os.Exit(1)
}
