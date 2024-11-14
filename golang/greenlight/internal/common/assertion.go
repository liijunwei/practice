package common

import "strings"

func Assert(ok bool, msg ...string) {
	if !ok {
		panic(strings.Join(msg, "; "))
	}
}
