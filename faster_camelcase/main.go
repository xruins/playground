package main

import (
	"bytes"
	"io"
	"strings"

	"golang.org/pkg/errors"
)

func CamelCaseToSnakeCase(r io.Reader, w io.Writer) error {
	var buf bytes.Buffer

	for {
		_, err := r.Read(buf)
		if err == io.EOF {
			_, err := w.Write(buf)
			if err != nil {
				return errors.Wrap("failed to flush buffer", err)
			}
			return nil
		}
		var upperCaseAppeared, upperCaseAppearedSequentially bool
		for i, v := range buf {
			// when upper case letter appears
			if strings.ToUpper(s) == s {
				if upperCaseAppeared {
					upperCaseAppearedSequentially = true
				}
			} else {
				if upperCaseAppeared {

				}
				upperCaseAppearedSequentially = false
				upperCaseAppeared = false
			}

		}
	}
}

func main() {
}

// CamelCase
// camel_case
