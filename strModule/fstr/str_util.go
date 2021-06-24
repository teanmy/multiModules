package fstr

import (
	"fmt"
	"strings"
)

func Trim(str string) string {
	fmt.Println("Trim been called")
	return strings.TrimSpace(str)
}
