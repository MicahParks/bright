[![Go Report Card](https://goreportcard.com/badge/github.com/MicahParks/bright)](https://goreportcard.com/report/github.com/MicahParks/bright)
# bright
This is a simple program to set monitor brightness using [`xrandr`](https://wiki.archlinux.org/title/xrandr).

I made this because Ubuntu does not come with a built-in way to change monitor brightness.

```bash
go install github.com/MicahParks/bright@latest
```

## Usage
Set brightness to 80%
```
$ bright 80
Brightness set.
```

Only set between 20% and 80%
```
$ bright 101
Must be between 20 and 100.
```

Set brightness to 100%
```
$ bright
Brightness set.
```
