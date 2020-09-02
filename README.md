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
### Number only ID
Generate number only ID with default length
```
id, err := uniqueid.Generateid("n")
```

Generate number only ID with custom length
```
id, err := uniqueid.Generateid("n", 6)
```

Generate number only ID with default length and prefix
```
id, err := uniqueid.Generateid("n", nil, "OD")
```

### Letter only ID
Generate letter only ID with default length
```
id, err := uniqueid.Generateid("l")

```
Generate letter only ID with custom length
```
id, err := uniqueid.Generateid("l",20)
```

Generate letter only ID with default length and prefix
```
id, err := uniqueid.Generateid("l", nil, "USER")
```

### Alphanumerical ID
Generate alphanumeric ID with default length 
```
id, err := uniqueid.Generateid("a")
```

Generate alphanumeric ID with custom length 
```
id, err := uniqueid.Generateid("a", 18)
```

Generate alphnumerical ID with default length and prefix
```
id, err := uniqueid.Generateid("a", nil, "uid")
```
<br />
To generate ID with **custom length** and **prefix**, provide length insted of nil as the second argument.While adding prefix the total length will be given length plus length of the prefix string provided.


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
