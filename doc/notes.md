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
