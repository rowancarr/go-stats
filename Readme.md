# go-stats
Learning task to create a small JSON stats endpoint. I used various pages to help and props to:

* [codegangsta.gitbooks.io](https://codegangsta.gitbooks.io/building-web-apps-with-go)
* [golang docs](https://golang.org/doc/effective_go.html#commentary)
* [lunny/diskinfo.go](https://gist.github.com/lunny/9828326)

## Install
```
go Install
```

or download from [releases](https://github.com/rowancarr/go-stats/releases)

## Usage
By default it takes 2 parameters which should be exported before running or in the init.d script

* PORT (default 8080)
* EP_BACKUP_DIR (default /backups)
