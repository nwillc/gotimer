[![license](https://img.shields.io/github/license/nwillc/gotimer.svg)](https://tldrlegal.com/license/-isc-license)
[![CI](https://github.com/nwillc/gotimer/workflows/CI/badge.svg)](https://github.com/nwillc/gotimer/actions?query=workflow%3CI)
[![Go Report Card](https://goreportcard.com/badge/github.com/nwillc/gotimer)](https://goreportcard.com/report/github.com/nwillc/gotimer)
---
# Go Timer

A simple terminal based digital timer written in Go to use as a Pomodoro timer.

![gotimer](gotimer.png)

## Installation

Install with `go install`:

```bash
  go install github.com/nwillc/gotimer@latest
```

## Usage

```bash
$ ./gotimer -h
A simple terminal based digital count down timer, may be used as a Pomodoro timer.

Usage:
  gotimer [flags]

Flags:
  -c, --color string   Color of timer. (default "orangered")
  -f, --font string    Font size to use. (default "7")
  -h, --help           help for gotimer
  -t, --time string    Time to count down. (default "25m")
  -v, --version        Display version.
```

 - The time duration is given in hours, minutes and seconds: `#h#m#s`. 
 - Exit with Ctrl-C or ESC
 - Pause with SPACE

## Tech Stack

**Language:** Go 1.18+

**Packages:** genfuncs, tcell
