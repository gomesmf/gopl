# Learning Go

[The Go Programming Language](http://www.gopl.io/)

[Go Class](https://www.youtube.com/playlist?list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6)

# Notes

Go get private repo

```
git config --global url.git@github.com:.insteadOf https://github.com/

export GOPRIVATE=github.com/<username>/<repo>

go get github.com/<username>/<repo>/<module>
```

Install binaries in bin/

```
$ source .env
$ cd ch02-progstruct/cfremote
$ go install
$ cfremote 23
23.0°F = -5.0°C, 23.0°C = 73.4°F
```
