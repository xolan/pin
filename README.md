# Examples

`pin add -i laa --cmd="ls -la"`

`pin list`

`pin do -i laa`

# Development

## Dependencies

* Go
  * Arch Linux: `pacman -S go`

I also recommend the Atom editor with the __go-plus__ package installed.

## Building

Assumes you have golang installed...

`git clone git@github.com:xolan/pin.git`

Change directory to repo root, for example `cd ~/dev/pin`

`source setpath.sh`

`cd src/github.com/xolan/pin`

`go get`

`go build`

# Usage

`pin help`:

```
Pin is a command pinner, similar to aliasing

Usage:
  pin [flags]
  pin [command]

Available Commands:
  do          Execute command with the given identifier
  list        Display a list of pinned commands
  add         Add (pin) a command
  remove      Remove (unpin) a command
  gendocs     Generate documentation for this program
  help        Help about any command

Flags:
  -h, --help=false: help for pin
  -v, --verbose=false: verbose output


Use "pin help [command]" for more information about a command.
```
