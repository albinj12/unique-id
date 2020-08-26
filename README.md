# unique-id

[![Build Status](https://github.com/albinj12/unique-id/workflows/Tests/badge.svg)](https://github.com/albinj12/unique-id/actions)
[![GoDoc](https://godoc.org/github.com/albinj12/unique-id?status.svg)](https://godoc.org/github.com/albinj12/unique-id)
[![Go Report Card](https://goreportcard.com/badge/github.com/albinj12/unique-id)](https://goreportcard.com/report/github.com/albinj12/unique-id)
[![License](https://img.shields.io/badge/license-MIT%20License-blue.svg)](LICENSE)


A unique id generator written in Golang.

Possible to generate ID having letters only or numbers only or alphanumeric unique id having custom length.
Default length is 16

# Install
```
go get github.com/albinj12/unique-id
```

# Usage
Generate number only ID with default length
```
id, err := uniqueid.Generateid("n")
```

Generate number only ID with custom length
```
id, err := uniqueid.Generateid("n", 6)
```

Generate letter only ID with default length
```
id, err := uniqueid.Generateid("l")

```
Generate letter only ID with custom length
```
id, err := uniqueid.Generateid("l",20)
```

Generate alphanumeric ID with default length 
```
id, err := uniqueid.Generateid("a")
```

Generate alphanumeric ID with custom length 
```
id, err := uniqueid.Generateid("a", 18)
```

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
