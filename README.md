# DOMparser

## Description
Can type for JQuery like because this tool is wrapping the "goquery".

## Usage

domp [option] [selector]

```bash
$ curl www.example.com | domp '#id > .class a'
```

### Ootion

#### --ouput, -o
output format.
* csv  => csv format.
* json => json format.
* text => space seperate text.

default "text"

```bash
$ curl www.example.com | domp -o csv '#id > .class a'
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
$ curl www.example.com | domp -q 'attr@alt|attr@src' -o csv 'img'
# image1,/img/1.jpg
# image2,/img/2.jpg
# ...
```


## Install

To install, use `go get`:

```bash
$ go get -d github.com/muramako/domp
```

## Contribution

1. Fork ([https://github.com/muramako/domp/fork](https://github.com/muramako/domp/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[muramako](https://github.com/muramako)

## TODO
* Write test.
