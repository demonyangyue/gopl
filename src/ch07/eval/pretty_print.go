
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package eval

import (
	"fmt"
	"strings"
)

//!+String

func (v Var) String() string {
	return string(v)
}

func (l literal ) String() string {
	return fmt.Sprintf("%g", l)
}

func (u unary) String() string {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Sprintf("unexpected unary op %q", u.op)
	}
	return fmt.Sprintf("%c%s", u.op, u.x)
}

func (b binary) String() string {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Sprintf("unexpected binary op %q", b.op)
	}
	return fmt.Sprintf("( %s %c %s)", b.x, b.op, b.y)
}

func (c call) String() string {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Sprintf("unknown function %q", c.fn)
	}
	if len(c.args) != arity {
		return fmt.Sprintf("call to %s has %d args, want %d",
			c.fn, len(c.args), arity)
	}
	return fmt.Sprintf("%s(%s)", c.fn, strings.Join(Map(c.args, func(arg Expr) string {
		return arg.String()
	}), ", "))
}

func Map(vs []Expr, f func(Expr) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

//!-String
