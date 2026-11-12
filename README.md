# hmacgo

[![Go Reference](https://pkg.go.dev/badge/github.com/MinatoAka/hmacgo.svg)](https://pkg.go.dev/github.com/MinatoAka/hmacgo)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Support via Saweria](https://img.shields.io/badge/Support-Saweria-orange)](https://saweria.co/AKaAturDompet)

A lightweight Go library for HMAC-SHA256/512 signatures and sorted URL query strings. Uses only the Go standard library.

## Install

```bash
go get github.com/MinatoAka/hmacgo
```

## Example

```go
params := map[string]string{
    "symbol": "ethidr",
    "side":   "buy",
}

payload := hmacgo.BuildQuery(params) // side=buy&symbol=ethidr
signature := hmacgo.SignSHA256(payload, "your-secret")
```

Use `SignSHA512(payload, secret)` for HMAC-SHA512. `BuildQuery` sorts keys, URL-encodes values, and omits empty values.

## Test

```bash
go test ./...
```

## Support

If you find this library useful, you can support its development via [Saweria](https://saweria.co/AKaAturDompet).

## License

MIT — see [LICENSE](LICENSE).
