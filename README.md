# CLI Dictionary

[![GoDoc](https://godoc.org/github.com/Rompei/dict?status.png)](https://godoc.org/github.com/Rompei/dict)
[![Build Status](https://drone.io/github.com/Rompei/dict/status.png)](https://drone.io/github.com/Rompei/dict/latest)


Dictionary on command line interface implemented with Golang.

![Sample](https://gyazo.com/f6a5ea8720e45bcc53be9d31155d2643.gif "sample")

##How to use it

```
go get github.com/Rompei/dict
cd /path/to/github.com/Rompei/dict
go build
```

Or download executable file at [drone.io](https://drone.io/github.com/Rompei/dict/files)


On shell configure file

```
alias dict="./path/to/binary -t <Dest-language> -f <Src-language> <Src>"
```

Put command, for example

```
$ ./dict -f en -t kor -t jpn -t rus -t ind -t swe -t tha cat
$ ko:고양이,고양잇과
$ ru:кошка,кот,кат,пизда,ко́шка,блевать,изрыгать,кошачий,парень
$ th:แมว,วิฬาร์,maeaew,wílaa,อาเจียน,อ้วก
$ sv:katt,kattdjur,fitta,mutta,slida,snippa,vagina,kisse,katta,kille,spy,tamkatt,typ
$ id:kucing,cat,meong
$ ja:猫,ネコ,ねこ,にゃあにゃあ,neko,にゃにゃ,にゃんにゃん,ぬこ,カト,キャット
```

Language code is available right here [ISO 639-3](https://en.wikipedia.org/wiki/List_of_ISO_639-3_codes)

##LICENSE

[BSD 3-Clause license](http://opensource.org/licenses/BSD-3-Clause)
