# base32check1-go

Go implementation of the Base32Check1 algorithm.

This implementation is a port from [base32check-python](https://github.com/bitmarck-service/base32check-python) and [base32check-java](https://github.com/bitmarck-service/base32check-java) developed by [BITMARCK Service GmbH](https://github.com/bitmarck-service) to support [DiGA](https://www.bfarm.de/EN/MedicalDevices/DiGA/_node.html) submissions.

## Installation

Use go get

```sh
go get github.com/kirinus/base32check1-go/v1
```

Then import the base32check1 package into your own code

```sh
import "github.com/kirinus/base32check1-go/v1"
```

## Usage

To compute and validate a checksum

```go
base32check1.Compute("CONSECRATIO") // "X"
base32check1.Validate("CONSECRATIO") // false

base32check1.Compute("CAFEDEAD") // "A"
base32check1.Validate("CAFEDEAD") // true
```

## Changelog

See the GitHub [release history](https://github.com/kirinus/base32check1-go/releases).
