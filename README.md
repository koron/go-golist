# go list wrapper

Wrapper for `go list` command to get package information.

```console
$ go get github.com/koron/go-golist
```

Example to obtain information of package in `/path/for/package` directory:

```go
pkg, err := golist.GetPackage("/path/for/package")
```

Example to obtain info of package in current directory:

```go
pkg, err := golist.GetPackage(".")

// or

pkg, err := golist.GetPackage("")
```

## LICENSE

MIT license. See LICENSE.
