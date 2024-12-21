package util

import (
	"os"
	"strconv"
)

/*
Wrapper for strconv.ParseInt, with a bit of extra
Handling for if the conversion fails
*/
func ExtractInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		os.Exit(-1)
	}
	return int(i)
}
