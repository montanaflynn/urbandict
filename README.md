# Urban Dictionary CLI

Get definitions and synonyms from [Urban Dictionary](http://urbandictionary.com) via the command line.

### Install

```
go get github.com/montanaflynn/urbandict
```

### Usage

```
usage: urbandict [<flags>] <word>

Flags:
  --help           Show help (also see --help-long and --help-man).
  -d, --debug      Enable debug mode.
  -s, --synonyms   Retrieve synonyms instead of definitions.
  --version        Show application version.

Args:
  <word>  Word to define.
