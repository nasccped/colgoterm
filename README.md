# colgoterm

A Go lang colored terminal printing. This is an simple app that
prints colored flags into the terminal (similar to some `neofetch`
display configs).

## Requirements

To run/install this program, you'll need:
- git
- go tooling

## Setup

1. clone the repository:
```sh
git clone https://github.com/nasccped/colgoterm && cd colgoterm
```

2. install with go:
```sh
go install
```

## Usage

This app provides a small CLI flag parsing:
```txt
$ colgoterm --help
A go lang colored terminal printer.

Usage: colgoterm [OPTIONS]

Options:
  --help                Prints the help panel
  --width  | -w <VALUE> Set the square width (default = 8)
  --height | -h <VALUE> Set the square height (default = 4)
  --gap    | -g <VALUE> Set the gap between squares (default = 4)
```

### Options

- **help:** displays the help panel providing the flags description
  overview. Note that this flag can't be called with other
  flags/alias.

- **others:** all the other flags expects an integer value to set
  measures when printing the content. When value isn't set, the
  default one is used.

## License

This project is under the [MIT license](./LICENSE)!
