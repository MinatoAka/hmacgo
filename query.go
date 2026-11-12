package hmacgo

import (
	"net/url"
	"sort"
)

// BuildQuery mengurutkan parameter secara abjad dan mengubahnya menjadi URL query string.
// Ini adalah syarat mutlak sebelum melakukan signing API seperti Indodax.
func BuildQuery(params map[string]string) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	values := url.Values{}
	for _, k := range keys {
		if params[k] != "" {
			values.Set(k, params[k])
		}
	}
	return values.Encode()
}
