package jsondup

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ErrDuplicateKey is returned when duplicate key names are detected
// inside a JSON object
type ErrDuplicateKey struct {
	path []string
	key  string
}

func (e *ErrDuplicateKey) Error() string {
	return fmt.Sprintf(`duplicate key "%s"`, strings.Join(append(e.path, e.key), "."))
}

// ValidateNoDuplicateKeys verifies the provided JSON object contains
// no duplicated keys
//
// The function expects a single JSON object, and will error prior to
// checking for duplicate keys should an invalid input be provided.
func ValidateNoDuplicateKeys(s string) error {
	var out map[string]any
	if err := json.Unmarshal([]byte(s), &out); err != nil {
		return fmt.Errorf("unmarshaling input: %w", err)
	}

	dec := json.NewDecoder(strings.NewReader(s))
	return checkToken(dec, nil)
}

// checkToken walks a JSON object checking for duplicated keys
//
// The function is called recursively on the value of each key
// inside and object, or item inside an array.
//
// Adapted from: https://stackoverflow.com/a/50109335
func checkToken(dec *json.Decoder, path []string) error {
	var dupErrs []error
	t, err := dec.Token()
	if err != nil {
		return err
	}

	delim, ok := t.(json.Delim)
	if !ok {
		// non-delimiter, nothing to do
		return nil
	}

	switch delim {
	case '{':
		keys := make(map[string]bool)
		for dec.More() {
			// Get the field key
			t, err := dec.Token()
			if err != nil {
				return err
			}
			key := t.(string)

			if keys[key] {
				// Duplicate found
				dupErrs = append(dupErrs, &ErrDuplicateKey{path: path, key: key})
			}
			keys[key] = true

			// Check the keys value
			if err := checkToken(dec, append(path, key)); err != nil {
				dupErrs = append(dupErrs, err)
			}
		}

		// consume trailing "}"
		_, err := dec.Token()
		if err != nil {
			return err
		}
	case '[':
		i := 0
		for dec.More() {
			// Check each items value
			if err := checkToken(dec, append(path, strconv.Itoa(i))); err != nil {
				dupErrs = append(dupErrs, err)
			}
			i++
		}

		// consume trailing "]"
		_, err := dec.Token()
		if err != nil {
			return err
		}
	}

	return errors.Join(dupErrs...)
}
