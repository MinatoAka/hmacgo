package hmacgo

import "testing"

func TestSignSHA256(t *testing.T) {
	secret := "rahasia_jenderal"
	payload := "symbol=ethidr&side=buy&timestamp=1700000000000"
	
	sign := SignSHA256(payload, secret)
	// SHA256 menghasilkan 64 karakter hex
	if len(sign) != 64 {
		t.Errorf("Panjang hash harus 64 karakter hex, dapat %d", len(sign))
	}
}

func TestBuildQuery(t *testing.T) {
	params := map[string]string{
		"symbol": "ethidr",
		"side":   "buy",
		"empty":  "", // Harus diabaikan
	}
	expected := "side=buy&symbol=ethidr"
	got := BuildQuery(params)
	
	if got != expected {
		t.Errorf("BuildQuery() = %q, mau %q", got, expected)
	}
}
