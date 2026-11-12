# hmacgo

[![Go Reference](https://pkg.go.dev/badge/github.com/MinatoAka/hmacgo.svg)](https://pkg.go.dev/github.com/MinatoAka/hmacgo)

`hmacgo` adalah library Go ringan untuk membuat signature HMAC-SHA256 dan HMAC-SHA512, serta menyusun parameter menjadi query string yang terurut.

## Fitur

- Signature HMAC-SHA256 dan HMAC-SHA512 dalam format hexadecimal lowercase.
- Query string dengan key terurut secara alfabetis.
- Nilai parameter kosong diabaikan.
- Nilai query di-encode menggunakan `net/url` dari standard library Go.
- Tanpa dependency eksternal.

## Instalasi

```bash
go get github.com/MinatoAka/hmacgo
```

## Penggunaan

```go
package main

import (
	"fmt"

	"github.com/MinatoAka/hmacgo"
)

func main() {
	params := map[string]string{
		"symbol":    "ethidr",
		"side":      "buy",
		"timestamp": "1700000000000",
	}

	// Key diurutkan dan nilai di-encode menjadi query string.
	payload := hmacgo.BuildQuery(params)
	fmt.Println(payload)
	// side=buy&symbol=ethidr&timestamp=1700000000000

	secret := "your-secret"

	signature256 := hmacgo.SignSHA256(payload, secret)
	signature512 := hmacgo.SignSHA512(payload, secret)

	fmt.Println(signature256)
	fmt.Println(signature512)
}
```

## API

### `BuildQuery(params map[string]string) string`

Menghasilkan query string dari parameter dengan key dalam urutan alfabetis. Parameter dengan nilai kosong tidak disertakan. Key dan value di-encode sesuai aturan `url.Values.Encode`.

### `SignSHA256(payload, secret string) string`

Menghasilkan HMAC-SHA256 dari `payload` menggunakan `secret`, lalu mengembalikan hasilnya sebagai string hexadecimal lowercase (64 karakter).

### `SignSHA512(payload, secret string) string`

Menghasilkan HMAC-SHA512 dari `payload` menggunakan `secret`, lalu mengembalikan hasilnya sebagai string hexadecimal lowercase (128 karakter).

> Pastikan format payload, encoding, dan algoritma yang digunakan sesuai dengan spesifikasi API yang Anda integrasikan. Format signing dapat berbeda antar layanan.

## Menjalankan test

```bash
go test ./...
```

## Lisensi

Didistribusikan di bawah lisensi MIT. Lihat file [LICENSE](LICENSE).
