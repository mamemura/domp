# htmlparser



## Description

## Usage

domc [option] [selector]

```bash
$ curl www.example.com | domc '#id > .class a'
```

### Ootion

#### --ouput, -o
output format.
* csv  => csv format.
* json => json format.
* text => space seperate text.

default "text"

```bash
$ curl www.example.com | domc -o csv '#id > .class a'
```



#### --query, -q
query is output data.
* text
* html
* outerhtml
* attrs
* attr@attribute

default "text"

And can use multiple query. ( seperator is "|" )
```bash
$ curl www.example.com | domc -q 'attr@alt|attr@src' -o csv 'img'
# image1,/img/1.jpg
# image2,/img/2.jpg
# ...
```


## Install

To install, use `go get`:

```bash
$ go get -d github.com/mamemura/domcutter
```

## Contribution

1. Fork ([https://github.com/mamemura/domcutter/fork](https://github.com/mamemura/domcutter/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[mamemura](https://github.com/mamemura)

## TODO
* Write test.
