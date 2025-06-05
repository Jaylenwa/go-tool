package compress

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"io"
)

// CompressStruct compresses any struct into a base64-encoded string.
// The process involves:
// 1. Converting the struct to JSON
// 2. Compressing the JSON using gzip
// 3. Encoding the compressed data to base64 string
//
// Returns:
// - string: The base64-encoded compressed string
// - error: Any error that occurred during the process
//
// Example:
//
//	type User struct { Name string `json:"name"` }
//	user := User{Name: "John"}
//	compressed, err := CompressStruct(user)
func CompressStruct(data any) (string, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	gzipWriter := gzip.NewWriter(&buf)
	_, err = gzipWriter.Write(jsonBytes)
	if err != nil {
		return "", err
	}
	gzipWriter.Close()

	encodedStr := base64.StdEncoding.EncodeToString(buf.Bytes())
	return encodedStr, nil
}

// DecompressStruct decompresses a base64-encoded string back into a struct.
// The process involves:
// 1. Decoding the base64 string
// 2. Decompressing the data using gzip
// 3. Unmarshaling the JSON into the provided struct
//
// Parameters:
// - str: The base64-encoded compressed string
// - out: A pointer to the struct that will receive the decompressed data
//
// Returns:
// - error: Any error that occurred during the process
//
// Example:
//
//	var user User
//	err := DecompressStruct(compressedString, &user)
func DecompressStruct(str string, out any) error {
	compressedData, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return err
	}

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

	return json.Unmarshal(jsonBuf.Bytes(), out)
}
