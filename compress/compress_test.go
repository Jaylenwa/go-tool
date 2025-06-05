package compress

import (
	"testing"
)

type TestStruct struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Tags    []string `json:"tags"`
	Details struct {
		Address string `json:"address"`
		Phone   string `json:"phone"`
	} `json:"details"`
}

func TestCompressAndDecompressStruct(t *testing.T) {
	tests := []struct {
		name     string
		input    TestStruct
		expected TestStruct
	}{
		{
			name: "basic struct",
			input: TestStruct{
				Name: "John Doe",
				Age:  30,
				Tags: []string{"test", "compress"},
				Details: struct {
					Address string `json:"address"`
					Phone   string `json:"phone"`
				}{
					Address: "123 Main St",
					Phone:   "123-456-7890",
				},
			},
			expected: TestStruct{
				Name: "John Doe",
				Age:  30,
				Tags: []string{"test", "compress"},
				Details: struct {
					Address string `json:"address"`
					Phone   string `json:"phone"`
				}{
					Address: "123 Main St",
					Phone:   "123-456-7890",
				},
			},
		},
		{
			name: "empty struct",
			input: TestStruct{
				Name: "",
				Age:  0,
				Tags: []string{},
				Details: struct {
					Address string `json:"address"`
					Phone   string `json:"phone"`
				}{
					Address: "",
					Phone:   "",
				},
			},
			expected: TestStruct{
				Name: "",
				Age:  0,
				Tags: []string{},
				Details: struct {
					Address string `json:"address"`
					Phone   string `json:"phone"`
				}{
					Address: "",
					Phone:   "",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 测试压缩
			compressed, err := CompressStruct(tt.input)
			if err != nil {
				t.Errorf("CompressStruct() error = %v", err)
				return
			}
			if compressed == "" {
				t.Error("CompressStruct() returned empty string")
				return
			}

			// 测试解压缩
			var result TestStruct
			err = DecompressStruct(compressed, &result)
			if err != nil {
				t.Errorf("DecompressStruct() error = %v", err)
				return
			}

			// 验证结果
			if result.Name != tt.expected.Name {
				t.Errorf("Name = %v, want %v", result.Name, tt.expected.Name)
			}
			if result.Age != tt.expected.Age {
				t.Errorf("Age = %v, want %v", result.Age, tt.expected.Age)
			}
			if len(result.Tags) != len(tt.expected.Tags) {
				t.Errorf("Tags length = %v, want %v", len(result.Tags), len(tt.expected.Tags))
				return
			}
			for i := range result.Tags {
				if result.Tags[i] != tt.expected.Tags[i] {
					t.Errorf("Tags[%d] = %v, want %v", i, result.Tags[i], tt.expected.Tags[i])
				}
			}
			if result.Details.Address != tt.expected.Details.Address {
				t.Errorf("Details.Address = %v, want %v", result.Details.Address, tt.expected.Details.Address)
			}
			if result.Details.Phone != tt.expected.Details.Phone {
				t.Errorf("Details.Phone = %v, want %v", result.Details.Phone, tt.expected.Details.Phone)
			}
		})
	}
}

func TestCompressStructError(t *testing.T) {
	// 测试无法序列化的类型
	invalidData := make(chan int)
	_, err := CompressStruct(invalidData)
	if err == nil {
		t.Error("CompressStruct() expected error for invalid data type")
	}
}

func TestDecompressStructError(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedErr bool
	}{
		{
			name:        "invalid base64",
			input:       "invalid-base64-string",
			expectedErr: true,
		},
		{
			name:        "invalid gzip data",
			input:       "SGVsbG8gV29ybGQ=", // "Hello World" in base64
			expectedErr: true,
		},
		{
			name:        "invalid json data",
			input:       "H4sIAAAAAAAA/6pWyk9JVbJSMjIwMLBQqgUAAP//AQAA//8=", // compressed invalid json
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result TestStruct
			err := DecompressStruct(tt.input, &result)
			if (err != nil) != tt.expectedErr {
				t.Errorf("DecompressStruct() error = %v, wantErr %v", err, tt.expectedErr)
			}
		})
	}
}
