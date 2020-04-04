// Copyright 2020 The browser Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package browser_test

import (
	"github.com/abc-inc/browser"
)

func ExampleOpen_File() {
	_ = browser.Open("file:///tmp")
	// Output:
}

func ExampleOpen_URL() {
	_ = browser.Open("http://localhost/")
	// Output:
}
