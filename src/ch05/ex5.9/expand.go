package expand

import (
	"regexp"
)

var pattern = regexp.MustCompile(`\${(\w+)}|\$(\w+)`)

func expand(s string, f func(string) string) string {


	wrapper := func(s string) string {

		groups := pattern.FindStringSubmatch(s)
		if len(groups[1]) > 0 {
			s = groups[1]
		} else {
			s = groups[2]
		}
		return f(s)
	}
	return pattern.ReplaceAllStringFunc(s, wrapper)
}
