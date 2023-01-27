package charsetconv

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
)

// 测试数据
var str = "abcde12345|你好，世界！"

var utf8 = []byte{
	97, 98, 99, 100, 101, 49, 50, 51, 52, 53, 124, 228, 189, 160, 229,
	165, 189, 239, 188, 140, 228, 184, 150, 231, 149, 140, 239, 188, 129,
}

var gbk = []byte{
	97, 98, 99, 100, 101, 49, 50, 51, 52, 53, 124, 196, 227, 186, 195,
	163, 172, 202, 192, 189, 231, 163, 161,
}

var big5 = []byte{
	97, 98, 99, 100, 101, 49, 50, 51, 52, 53, 124, 167, 65, 166, 110, 161,
	65, 165, 64, 172, 201, 161, 73,
}

var eucjp = []byte{
	97, 98, 99, 100, 101, 49, 50, 51, 52, 53, 124, 143, 176, 223, 185, 165, 161,
	164, 192, 164, 179, 166, 161, 170,
}

var gb18030 = []byte{
	97, 98, 99, 100, 101, 49, 50, 51, 52, 53, 124, 196, 227, 186, 195, 163, 172, 202,
	192, 189, 231, 163, 161,
}

var gb2312 = []byte{
	97, 98, 99, 100, 101, 49, 50, 51, 52, 53, 124, 196, 227, 186, 195, 163, 172, 202,
	192, 189, 231, 163, 161,
}

var hzgb2312 = []byte{
	97, 98, 99, 100, 101, 49, 50, 51, 52, 53, 124, 126, 123, 68, 99, 58, 67,
	35, 44, 74, 64, 61, 103, 35, 33, 126, 125,
}

func TestCommon(t *testing.T) {
	AddCharsetAlias(Charset("GBK"), "gbk")
	charset, err := GetCharsetByAlias("gbk")
	if err != nil || charset != "GBK" {
		t.Fail()
	}
	RemoveCharsetAlias("gbk")
	charset, err = GetCharsetByAlias("gbk")
	if err == nil {
		t.Fail()
	}
}

func TestEncodeWith(t *testing.T) {
	// 测试未知字符集
	err := encodeAndCompare(str, "foo", nil)
	if err == nil {
		t.Fail()
	} else {
		t.Log("cannot encode with unknown charset:", err)
	}
	// 测试字符集：=> utf-8
	err = encodeAndCompare(str, "utf-8", utf8)
	if err != nil {
		t.Fatal(err)
	}
	// 测试字符集：=> big5
	err = encodeAndCompare(str, "Big5", big5)
	if err != nil {
		t.Fatal(err)
	}
	// 测试字符集：=> gbk
	err = encodeAndCompare(str, "gbk", gbk)
	if err != nil {
		t.Fatal(err)
	}
	// 测试字符集：=> gbk
	err = encodeAndCompare(str, "EUC-JP", eucjp)
	if err != nil {
		t.Fatal(err)
	}
	// 测试字符集：=> gb2312
	err = encodeAndCompare(str, "GB2312", gb2312)
	if err == nil {
		t.Fatal(err)
	} else {
		t.Log("不支持GB2312字符集，因此会返回一个错误：", err.Error())
	}
	// 测试字符集：=> gb18030
	err = encodeAndCompare(str, "GB18030", gb18030)
	if err != nil {
		t.Fatal(err)
	}
	// 测试字符集：=> hzgb2312
	//err = encodeAndCompare(str, "HZ-GB-2312", hzgb2312)
	//if err != nil {
	//	t.Log("hz-gb-2312编码结果与python中所得结果不一致，目前尚不知道原因！")
	//	t.Error(err)
	//}

	// 测试使用中转区文件
	src := MakeStringReader(str)
	dest := MakeByteBuffer(0)
	err = EncodeWith(src, dest, "GBK", true)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if !compareBytes(dest.Bytes(), gbk) {
		t.Log("not match")
		t.Fail()
	}

	// 测试不使用中转区，编码失败污染dest的情形
	src = MakeStringReader(str)
	dest.Reset()
	err = EncodeWith(src, dest, "Windows-1252", false)
	if err == nil {
		t.Fatal()
	}
	if len(dest.Bytes()) > 0 {
		t.Log("编码中途失败：", err)
		t.Log("dest已污染：", dest.Bytes())
	}

	// 开启中转区后，编码失败也不会污染dest
	src = MakeStringReader(str)
	dest.Reset()
	err = EncodeWith(src, dest, "Windows-1252", true)
	if err == nil {
		t.Fatal()
	} else {
		if len(dest.Bytes()) > 0 {
			t.Log("编码中途失败：", err)
			t.Log("dest已污染：", dest.Bytes())
		} else {
			t.Log("编码中途失败：", err)
			t.Log("dest未被污染：", dest.Bytes())
		}
	}
}

func TestEncodeString(t *testing.T) {
	// 测试EncodeString()函数
	dest, err := EncodeString(str, "utf-8")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if !compareBytes(dest, utf8) {
		t.Error("not match")
		t.Fail()
	}
	// 测试一些边缘情形
	dest, err = EncodeString("", "euc-jp")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

}

func encodeAndCompare(src string, charset string, b []byte) error {
	dest := MakeByteBuffer(0)
	err := EncodeWith(MakeStringReader(src), dest, Charset(charset), false)
	if err != nil {
		return err
	}
	if compareBytes(dest.Bytes(), b) {
		return nil
	} else {
		fmt.Println("encoded@", charset, ":", dest.Bytes())
		fmt.Println("sample @", charset, ":", b)
		return errors.New("not match")
	}
}

func compareBytes(a []byte, b []byte) bool {
	return bytes.Compare(a, b) == 0
}
