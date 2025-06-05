package slice

import (
	"testing"
)

func TestSliceDiff(t *testing.T) {
	tests := []struct {
		name     string
		a        []int
		b        []int
		expected []int
	}{
		{
			name:     "basic difference",
			a:        []int{1, 2, 3, 4, 5},
			b:        []int{3, 4, 5, 6, 7},
			expected: []int{1, 2},
		},
		{
			name:     "empty slice a",
			a:        []int{},
			b:        []int{1, 2, 3},
			expected: []int{},
		},
		{
			name:     "empty slice b",
			a:        []int{1, 2, 3},
			b:        []int{},
			expected: []int{1, 2, 3},
		},
		{
			name:     "no difference",
			a:        []int{1, 2, 3},
			b:        []int{1, 2, 3},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SliceDiff(tt.a, tt.b)
			if len(result) != len(tt.expected) {
				t.Errorf("SliceDiff() length = %v, want %v", len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("SliceDiff()[%d] = %v, want %v", i, result[i], tt.expected[i])
				}
			}
		})
	}
}

func TestSliceIntersection(t *testing.T) {
	tests := []struct {
		name     string
		a        []int
		b        []int
		expected []int
	}{
		{
			name:     "basic intersection",
			a:        []int{1, 2, 3, 4, 5},
			b:        []int{3, 4, 5, 6, 7},
			expected: []int{3, 4, 5},
		},
		{
			name:     "empty intersection",
			a:        []int{1, 2, 3},
			b:        []int{4, 5, 6},
			expected: []int{},
		},
		{
			name:     "empty slice a",
			a:        []int{},
			b:        []int{1, 2, 3},
			expected: []int{},
		},
		{
			name:     "empty slice b",
			a:        []int{1, 2, 3},
			b:        []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SliceIntersection(tt.a, tt.b)
			if len(result) != len(tt.expected) {
				t.Errorf("SliceIntersection() length = %v, want %v", len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("SliceIntersection()[%d] = %v, want %v", i, result[i], tt.expected[i])
				}
			}
		})
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		fn       func(int) string
		expected []string
	}{
		{
			name:  "convert int to string",
			input: []int{1, 2, 3},
			fn: func(i int) string {
				return string(rune(i + '0'))
			},
			expected: []string{"1", "2", "3"},
		},
		{
			name:  "empty slice",
			input: []int{},
			fn: func(i int) string {
				return string(rune(i + '0'))
			},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Map(tt.input, tt.fn)
			if len(result) != len(tt.expected) {
				t.Errorf("Map() length = %v, want %v", len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Map()[%d] = %v, want %v", i, result[i], tt.expected[i])
				}
			}
		})
	}
}

func TestToMap(t *testing.T) {
	type Person struct {
		ID   int
		Name string
	}

	tests := []struct {
		name     string
		input    []Person
		fn       func(Person) (int, string)
		expected map[int]string
	}{
		{
			name: "convert slice to map",
			input: []Person{
				{ID: 1, Name: "Alice"},
				{ID: 2, Name: "Bob"},
				{ID: 3, Name: "Charlie"},
			},
			fn: func(p Person) (int, string) {
				return p.ID, p.Name
			},
			expected: map[int]string{
				1: "Alice",
				2: "Bob",
				3: "Charlie",
			},
		},
		{
			name:  "empty slice",
			input: []Person{},
			fn: func(p Person) (int, string) {
				return p.ID, p.Name
			},
			expected: map[int]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToMap(tt.input, tt.fn)
			if len(result) != len(tt.expected) {
				t.Errorf("ToMap() length = %v, want %v", len(result), len(tt.expected))
				return
			}
			for k, v := range tt.expected {
				if result[k] != v {
					t.Errorf("ToMap()[%v] = %v, want %v", k, result[k], v)
				}
			}
		})
	}
}
