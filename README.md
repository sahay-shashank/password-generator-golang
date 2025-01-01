# Password Generator

A CLI utility tool that can be used to generate passwords written in GoLang.

## Options

We are providing two options:

- CLI
- WEB

### CLI

This option, will generate and output a password on the STDOUT.

The length flag will also be parsed and used in this case.

### WEB

This option will not parse the length flag. The length should be provided in the url as an argument of 'length'.

example: ``http://localhost:8080/?length=10``

## Defaults and Character Set

The 'length' flag is defaulted to 7 in both the options.

Character set includes:

- abcdefghijklmnopqrstuvwxyz
- ABCDEFGHIJKLMNOPQRSTUVWXYZ
- 0123456789
- !@#$%
