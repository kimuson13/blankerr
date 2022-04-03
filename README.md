# blankerr [![Go Reference](https://pkg.go.dev/badge/github.com/kimuson13/blankerr.svg)](https://pkg.go.dev/github.com/kimuson13/blankerr)
blankerr is check blank err return values and check call function without error handling
## Install
```
$ go install github.com/kimuson13/blankerr/cmd/blankerr@latest
```
## How to use
If you want to check error with a package, give th e package path of interst as the first argument:
```
$ blankerr github.com/kimuson13/blankerr
```
To check all package beneath the current directory:
```
$ blankerr ./...
```
## Demo
If you create such a Go file bellow
```
$ ls
a.go
$ cat a.go
package a

import "errors"

func hoge() (string, error) {
    return "hoge", errors.New("hoge")
}

func fuga() {
    h, _ := hoge()
    print(h)
}
```
`func fuga()` ignores error in return values of `hoge()`.
We need to handle this error. So use `blankerr` like that.
```
$ blankerr a
a/a.go:16:9 hoge has error type in return values
```
`blankerr` help error handling
## Feature Outlook
### not check some standard libraly
Some standard library that have an error return type but are documented to never return an error(eg: `fmt.Print()`)  
So, I want to ignore these functions.
