// Copyright 2017 Matt Ho
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package swagger

import (
	"fmt"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"
)

func makeRef(name string) string {
	return fmt.Sprintf("#/definitions/%v", name)
}

type reflectType interface {
	PkgPath() string
	Name() string
}

func makeName(t reflectType) string {
	name := filepath.Base(t.PkgPath()) + t.Name()
	return strings.Replace(name, "-", "_", -1)
}

// isExported reports whether the identifier is exported.
func isExported(id string) bool {
	r, _ := utf8.DecodeRuneInString(id)
	return unicode.IsUpper(r)
}
