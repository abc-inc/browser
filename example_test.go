// Copyright 2020 The browser Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package browser_test

import (
	"github.com/abc-inc/browser"
)

func ExampleOpen_file() {
	_ = browser.Open("file:///tmp")
	// Output:
}

func ExampleOpen_url() {
	_ = browser.Open("http://localhost/")
	// Output:
}
