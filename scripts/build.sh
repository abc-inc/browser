#!/usr/bin/env bash

# Copyright 2020 The browser Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# shellcheck disable=SC2164
cd "$(dirname "$0")/.."
go test browser_test.go
for os in darwin:app linux:bin windows:exe; do
  GOOS="${os%:*}" go build -ldflags "-s -w" -o "browser.${os#*:}" cmd/browser/main.go
done
