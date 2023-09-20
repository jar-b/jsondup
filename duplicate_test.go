package jsondup

import (
	"testing"
)

func TestValidate(t *testing.T) {
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
			name:    "invalid multiple objects",
			s:       `{"a": "foo"}{"b": "bar"}`,
			wantErr: true,
		},
		{
			name:    "invalid multiple arrays",
			s:       `["a"]["b"]`,
			wantErr: true,
		},
		{
			name: "object root duplicate",
			s: `{
  "a": "foo",
  "a": "bar"
}`,
			wantErr: true,
		},
		{
			name: "object nested object duplicate",
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
			name: "object nested array duplicate",
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
			name: "object multiple duplicates",
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
			name: "object valid",
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
		{
			name: "array valid",
			s: `[
  {
    "a": "foo"
  },
  {
    "b": "bar"
  }
]`,
			wantErr: false,
		},
		{
			name: "array duplicate",
			s: `[
  {
    "a": "foo",
    "a": "bar"
  },
  {
    "b": "baz"
  }
]`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Validate(tt.s); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
