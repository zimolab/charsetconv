// Package charsetconv
// @Description:
// 这个包实现了文本编码转换功能，用于将基于字符集的文本数据转化为另一字符集。
// 比如将UTF8文本转化为GBK文本，反之亦然。由于go语言原生基于UTF8编码，因此，在go语言语境中， 文本编解码的含义是特定的：
// 编码（encode）是指，将UTF8文本转换为以另一字符集编码的文本数据
// 解码（decode）是指，将以其他字符集编码的文本数据转换为UTF8编码的文本
//
// 本包提供了Encode、Decode两种接口，同时还提供了Convert接口，即在任意两种字符集之间进行转换的接口，提供了诸如直接从GBK编码到Big5编码转换的能力。
//
// 需要注意的是，编码转换并不一定总是能够成功的，结果取决于进行转换的两种字符集之间的兼容性，而且还和转换的方向有很大关系。
// 例如：将一段Windows-1252字符集编码的文本转换为GBK形式，大概率是会成功的，而反过来则很可能会失败， 究其原因，在Windows-1252字符集中并不包含汉字的编码。
//
// 本包基于golang.org/x/text/encoding封装，目的在于提供一个简单、易用、直观的接口。
package charsetconv

import (
	"golang.org/x/text/encoding"
	"io"
)

// Encode
//
//	@Description: 编码函数，将utf-8数据转换为指定编码形式
//	@param src 待编码的数据源，注意需确保该数据源以utf-8格式编码
//	@param dest 编码后的数据输出位置
//	@param destEncoding 目标编码格式
//	@param useStageFile 是否使用中转文件，如果设置了使用中转区文件，则先将编码后的数据缓存到中转区文件，编码成功后再从中转区文件拷贝到dest中
//	@return error
func Encode(src io.Reader, dest io.Writer, destEncoding encoding.Encoding, useStageFile bool) error {
	return doTransform(makeTransformReader(src, destEncoding, true), dest, useStageFile)
}

// EncodeWith
//
//	@Description: 编码函数，将utf-8数据转换为指定编码形式
//	@param src 待编码的数据源，注意需确保该数据源以utf-8格式编码
//	@param dest 编码后的数据输出位置
//	@param destCharset 目标字符集名称，例如：gbk, Big5等
//	@param useStageFile 是否使用中转文件，如果设置了使用中转区文件，则先将编码后的数据缓存到中转区文件，编码成功后再从中转区文件拷贝到dest中
//	@return error
func EncodeWith(src io.Reader, dest io.Writer, destCharset Charset, useStageFile bool) error {
	en, err := EncodingOf(string(destCharset))
	if err != nil {
		return err
	}
	return Encode(src, dest, en, useStageFile)
}

// EncodeString
//
//	@Description: 编码函数，将字符串编码为指定字符集数据
//	@param src 待编码的字符串
//	@param destCharset 目标字符集
//	@return []byte
//	@return error
func EncodeString(src string, destCharset Charset) ([]byte, error) {
	srcReader := MakeStringReader(src)
	destBuff := MakeByteBuffer(0)
	err := EncodeWith(srcReader, destBuff, destCharset, false)
	if err != nil {
		return nil, err
	}
	return destBuff.Bytes(), nil
}
