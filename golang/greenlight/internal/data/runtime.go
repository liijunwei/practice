package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	val := fmt.Sprintf("%d mins", r)
	quotedVal := strconv.Quote(val)

	return []byte(quotedVal), nil
}
