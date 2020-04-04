## Introduction
*browser* opens the provided URL (or file) by using an executable (in the following sequence):

1.  command defined by the environment variable `$BROWSER` (if set)
2.  operating system specific launcher
    * macOS: `/usr/bin/open`
	* windows: `cmd /c start`
	* otherwise: `xdg-open` (if environment variable `$DISPLAY` is set)
3.  use one the following binaries from `$PATH`
    * `chrome`
	* `google-chrome`
	* `chromium`
    * `firefox`

## Use *browser* as library
### Installation
```shell script
go get -d -u github.com/abc-inc/browser
```
### Usage
```go
import "github.com/abc-inc/browser"

...

if ok := browser.Open(url); !ok {
    // something bad happened
}
```

## Use *browser* as standalone application
### Installation
```shell script
go get -u github.com/abc-inc/browser
```
### Examples
```shell script
# open a URL in the default browser
browser http://localhost/

# open a local file in the default browser
browser file:///C:/Temp/report.html

# macOS: use Firefox even if another default web browser is set
BROWSER="open -a Firefox" browser http://localhost/file.zip

# download a file with wget or curl - whatever is available
BROWSER="wget:curl -o index.html" browser http://localhost/
```

Since *browser* leverages operating system specific launchers, it can also be used to open other files and directories:
```shell script
# play audio file using the media player associated with the file type
browser file:///home/me/Music/track.mp3

# open a folder in Visual Studio Code or vi (must be installed separately)
BROWSER=code:vi browser file:///Users/me/dev/project
```

## Similar Projects
* [github.com/pkg/browser][gh/pkg/browser]: the de facto standard implementation
* [github.com/hashicorp/terraform/command/webbrowser][gh/hashicorp/terraform]:
encapsulates [pkg/browser][gh/pkg/browser] and provides a mock implementation for testing purposes.

| Feature                   | [abc-inc/browser][gh/abc-inc/browser] | [pkg/browser][gh/pkg/browser] | [hashicorp/terraform][gh/hashicorp/terraform] |
| ------------------------- |:-------------------------------------:|:-----------------------------:|:---------------------------------------------:|
| Windows/Linux/macOS       | :ballot_box_with_check:               |                               |                                               |
| FreeBSD/NetBSD (xdg-open) | :ballot_box_with_check:               |                               |                                               |
| [$BROWSER][man.1] support | :ballot_box_with_check:               |                               |                                               |
| WSL support               | :ballot_box_with_check:               |                               |                                               |
| open from io.Reader       |                                       | :ballot_box_with_check:       |                                               |
| mock support              |                                       |                               | :ballot_box_with_check:                       |

[gh/abc-inc/browser]: https://github.com/abc-inc/browser
[gh/pkg/browser]: https://github.com/pkg/browser
[gh/hashicorp/terraform]: https://github.com/hashicorp/terraform/tree/master/command/webbrowser
[man.1]: http://linuxcommand.org/lc3_man_pages/man1.html
