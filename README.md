# unique-id
A unique id generator written in Golang.

Possible to generate ID having letters only or numbers only unique id having custom length.
Default length is 16

# Install
```
go get github.com/albinj12/unique-id
```

# Usage
Generate number only ID with default length
```
id, _ := uniqueid.Generateid("i")
```

Generate number only ID with custom length
```
id, _ := uniqueid.Generateid("i", 6)
```

Generate letter only ID with default length
```
id, _ := uniqueid.Generateid("a")

```
Generate letter only ID with default length
```
id, _ := uniqueid.Generateid("a",20)
```

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
