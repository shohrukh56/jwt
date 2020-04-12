package jwt

import "testing"

func TestDecodeBadToken(t *testing.T) {
	payload := struct{
		Exp int64 `json:"exp"`
	}{
		Exp: 0,
	}
	err := Decode("xxx.xxx.xxx.xxx", &payload)
	if err == nil {
		t.Errorf("Decode() error %v", err)
	}
	err = Decode("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiaWF0IjoxNTg0MDExOTAyLCJleHAiOjE1ODQwNDc5MDJ9.BcumALTIODycz_uESHilA5xGbUmEst3T6RAHUCAwcIc", &payload)
	if err != nil {
		t.Errorf("Decode() error %v", err)
	}

}
func TestDecodeCanNotDecode(t *testing.T) {
	payload := struct{
		Exp int64 `json:"exp"`
	}{
		Exp: 0,
	}
	err := Decode("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.8985465.BcumALTIODycz_uESHilA5xGbUmEst3T6RAHUCAwcIc", &payload)
	if err == nil {
		t.Errorf("Decode() error %v", err)
	}
	err = Decode("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiaWF0IjoxNTg0MDExOTAyLCJleHAiOjE1ODQwNDc5MDJ9.BcumALTIODycz_uESHilA5xGbUmEst3T6RAHUCAwcIc", &payload)
	if err != nil {
		t.Errorf("Decode() error %v", err)
	}
}
