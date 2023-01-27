package charsetconv

import "testing"

func TestDecodeToString(t *testing.T) {
	dest, err := DecodeToString(gbk, "GBK")
	if err != nil {
		t.Fail()
	}
	if dest != str {
		t.Fail()
	}

	dest, err = DecodeToString(eucjp, "euc-jp")
	if err != nil {
		t.Fail()
	}
	if dest != str {
		t.Fail()
	}

	dest, err = DecodeToString(eucjp, "GBK")
	t.Log(dest)
	if dest == str {
		t.Fail()
	}

	dest, err = DecodeToString(eucjp, "ANSI")
	t.Log(dest)
	t.Log(err)
	if err == nil {
		t.Error(err)
		t.Fail()
	}

}
