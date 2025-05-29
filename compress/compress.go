package compress

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"io"
)

// CompressStruct 将任意结构体压缩为base64字符串
func CompressStruct(data any) (string, error) {
	// 1. 结构体转 JSON
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// 2. Gzip 压缩
	var buf bytes.Buffer
	gzipWriter := gzip.NewWriter(&buf)
	_, err = gzipWriter.Write(jsonBytes)
	if err != nil {
		return "", err
	}
	gzipWriter.Close()

	// 3. 转 base64 字符串
	encodedStr := base64.StdEncoding.EncodeToString(buf.Bytes())
	return encodedStr, nil
}

// DecompressStruct 将base64字符串解压为结构体
func DecompressStruct(str string, out any) error {
	// 1. base64 解码
	compressedData, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return err
	}

	// 2. Gzip 解压
	buf := bytes.NewReader(compressedData)
	gzipReader, err := gzip.NewReader(buf)
	if err != nil {
		return err
	}
	defer gzipReader.Close()
	var jsonBuf bytes.Buffer
	_, err = io.Copy(&jsonBuf, gzipReader)
	if err != nil {
		return err
	}

	// 3. 解析 JSON 到结构体
	return json.Unmarshal(jsonBuf.Bytes(), out)
}
