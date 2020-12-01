package amf

import (
	"fmt"
	"strconv"
)

func StringToInt64(v string) (ret int64, err error) {
	return strconv.ParseInt(v, 10, 64)
}

func StringToUint64(v string) (ret uint64, err error) {
	var out int64
	out, err = strconv.ParseInt(v, 10, 64)
	return uint64(out), err
}

func UInt64ToString(v uint64) string {
	return fmt.Sprintf("%v", v)
}

func Int64ToString(v int64) string {
	return fmt.Sprintf("%v", v)
}
