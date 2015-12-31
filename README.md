# CLI Dictionary

[![GoDoc](https://godoc.org/github.com/Rompei/dict?status.png)](https://godoc.org/github.com/Rompei/dict)
[![Build Status](https://drone.io/github.com/Rompei/dict/status.png)](https://drone.io/github.com/Rompei/dict/latest)


Dictionary on command line interface implemented with Golang.

##How to use it

```
go get github.com/Rompei/dict
cd /path/to/github.com/Rompei/dict
go build
```

Or download executable file at [drone.io](https://drone.io/github.com/Rompei/dict/files)


On shell congure file

```
alias dict="./path/to/binary -t <Dest-language> -f <Src-language> <Src>"
```

Put command, for example

```
dict -t ja -f en cat
```

Language code is available right here [ISO 639-3](https://en.wikipedia.org/wiki/List_of_ISO_639-3_codes)

##LICENSE

[BSD 3-Clause license](http://opensource.org/licenses/BSD-3-Clause)
