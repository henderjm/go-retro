# go-feedback
No one likes opening websites, use this tool to post your retros

* run `go get github.com/henderjm/go-feedback`

## Usage
```
Usage:
  main [OPTIONS] <happy | meh | sad>

Application Options:
  -r, --retro-id= Retro Board Id

Help Options:
  -h, --help      Show this help message

Available commands:
  happy  Express your happiness (aliases: h)
  meh    Raise a potential concern (aliases: m)
  sad    Why so sad?? (aliases: s)
```

## Examples
1. `go-feedback --retro-id 123456790 happy --description "I love this cli 😀"`
