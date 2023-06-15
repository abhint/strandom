# Strandom

The Strandom package provides a utility for generating random strings with customizable length. It combines numbers, lowercase letters, and uppercase letters in a shuffled manner to create unique and random strings.

## Installation

To use the Strandom package, you need to have Go installed and set up on your machine. Then, run the following command to install the package:

```shell
go get github.com/abhint/strandom
```
## Usage
Import the package in your Go code:

```go
import "github.com/abhint/strandom"
```
Generate a random string of a specified length:

```go
randomString := strandom.RandomString(10) // Generate a random string of length 10
fmt.Println(randomString) // xqha5apkJG
```

If the length is not provided or set to 0, it defaults to 8.

Feel free to explore and customize the package as per your needs.

## Contribution
Contributions are welcome! If you have any suggestions, bug reports, or feature requests, please open an issue or submit a pull request on the GitHub repository.

## License
This package is released under the MIT License. See the [LICENSE](LICENSE) file for more details.
