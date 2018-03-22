# mapper

[![CircleCI](https://circleci.com/gh/davidsbond/mapper/tree/develop.svg?style=shield)](https://circleci.com/gh/davidsbond/mapper)
[![Coverage Status](https://coveralls.io/repos/github/davidsbond/mapper/badge.svg?branch=develop)](https://coveralls.io/github/davidsbond/mapper?branch=develop)
[![GoDoc](https://godoc.org/github.com/davidsbond/mapper?status.svg)](http://godoc.org/github.com/davidsbond/mapper)
[![Go Report Card](https://goreportcard.com/badge/github.com/davidsbond/mapper)](https://goreportcard.com/report/github.com/davidsbond/mapper)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/davidsbond/mapper/release/LICENSE)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fdavidsbond%2Fmapper.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fdavidsbond%2Fmapper?ref=badge_shield)

A Golang package for mapping fields from a source struct to a target struct using struct tags

## usage

```go
package main

type (
  Source struct {
    Field string `map:"Target:Field"`
  }

  Target struct {
    Field string
  }
)

func main() {
 s := Source{
   Field: "some data",
 }

 t := Target{}

 if err := mapper.Map(s, &t); err != nil {
   // handle error
 }

 // t.Field will now equal s.Field, you can also
 // specify multiple targets separating them with
 // ';' in the struct tag.
}
```
