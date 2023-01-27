package charsetconv

import "testing"

func TestConvertTo(t *testing.T) {
	dest, err := ConvertTo(gbk, GBK, GBK)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if !compareBytes(dest, gbk) {
		t.Fail()
	}
}

func TestConvertTo2(t *testing.T) {
	dest, err := ConvertTo(gbk, GBK, UTF8)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if !compareBytes(dest, utf8) {
		t.Fail()
	}
}

func TestConvertTo3(t *testing.T) {
	dest, err := ConvertTo(utf8, UTF8, EUCJP)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if !compareBytes(dest, eucjp) {
		t.Fail()
	}
}

func TestConvertTo4(t *testing.T) {
	dest, err := ConvertTo(gb18030, GB18030, EUCJP)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if !compareBytes(dest, eucjp) {
		t.Fail()
	}

	dest, err = ConvertTo(eucjp, EUCJP, Big5)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if !compareBytes(dest, big5) {
		t.Fail()
	}
}

func TestConvertTo5(t *testing.T) {
	dest, err := ConvertTo(gb18030, Windows1258, EUCJP)
	if err == nil {
		t.Fail()
	}
	t.Log("err:", err)
	t.Log("dest:", dest)
}
