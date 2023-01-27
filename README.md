# charsetconv

字符集编解码库，提供以下三种接口。

## 编码（Encode）
将utf-8字符集编码的数据编码成其他字符集形式
```GO
func Encode(src io.Reader, dest io.Writer, destEncoding encoding.Encoding, useStageFile bool) error
func EncodeWith(src io.Reader, dest io.Writer, destCharset Charset, useStageFile bool) error
func EncodeString(src string, destCharset Charset) ([]byte, error)
```


## 解码（Decode）
将非utf-8字符集编码的数据解码成utf-8字符集编码形式
```Go
func Decode(src io.Reader, dest io.Writer, srcEncoding encoding.Encoding, useStageFile bool) error
func DecodeWith(src io.Reader, dest io.Writer, srcCharset Charset, useStageFile bool) error
func DecodeToString(src []byte, srcCharset Charset) (string, error)
```

## 转换（Convert）
在任意两种字符集之间转换编码
```Go
func Convert(src io.Reader, srcEncoding encoding.Encoding, dest io.Writer, destEncoding encoding.Encoding, useStageFile bool) error
func ConvertWith(src io.Reader, srcCharset Charset, dest io.Writer, destCharset Charset, useStageFile bool) error
func DecodeToString(src []byte, srcCharset Charset) (string, error)
```