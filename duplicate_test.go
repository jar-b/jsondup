package jsondup

import (
	"testing"
)

func TestValidateNoDuplicateKeys(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		wantErr bool
	}{
		{
			name:    "invalid",
			s:       "{{{",
			wantErr: true,
		},
		{
			name: "root",
			s: `{
  "a": "foo",
  "a": "bar"
}`,
			wantErr: true,
		},
		{
			name: "nested object",
			s: `{
  "a": "foo",
  "b": {
    "c": "bar",
    "c": "baz"
  }
}`,
			wantErr: true,
		},
		{
			name: "nested array",
			s: `{
  "a": "foo",
  "b": {
    "c": "bar",
    "d": [
      {
        "e": "foo",
        "e": "bar"
      },
      {
        "f": "baz",
        "g": "qux"
      }
    ]
  }
}`,
			wantErr: true,
		},
		{
			name: "multiple",
			s: `{
  "a": "foo",
  "a": "bar",
  "b": {
    "c": "baz",
    "c": "qux",
    "d": [
      {
        "e": "foo"
      },
      {
        "f": "bar",
        "f": "baz",
        "g": "qux"
      }
    ]
  }
}`,
			wantErr: true,
		},
		{
			name: "valid",
			s: `{
  "a": "foo",
  "b": {
    "c": "bar",
    "d": [
      {
        "e": "baz"
      },
      {
        "f": "qux",
        "g": "foo"
      }
    ]
  }
}`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateNoDuplicateKeys(tt.s); (err != nil) != tt.wantErr {
				t.Errorf("ValidateNoDuplicateKeys() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
