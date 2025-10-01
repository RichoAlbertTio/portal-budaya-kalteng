// internal/util/slug.go
package util


import (
"regexp"
"strings"
)


var nonAlphaNum = regexp.MustCompile(`[^a-z0-9]+`)


func Slugify(s string) string {
s = strings.ToLower(strings.TrimSpace(s))
s = nonAlphaNum.ReplaceAllString(s, "-")
return strings.Trim(s, "-")
}