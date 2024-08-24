// Copyright 2022 Juan Pablo Tosso and the OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

package transformations

import (
	"path/filepath"
	"strings"
)

func normalisePath(data string) (string, bool, error) {
	leng := len(data)
	if leng < 1 {
		return data, false, nil
	}
	clean := filepath.Clean(data)
	if filepath.Separator == '\\' {
		clean = strings.ReplaceAll(clean, "\\", "/")
	}

	if clean == "." {
		return "", true, nil
	}

	if data[len(data)-1] == '/' || data[len(data)-1] == '\\' {
		return clean + "/", true, nil
	}
	return clean, data != clean, nil
}
