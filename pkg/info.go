package pkg

import (
	"fmt"
	"os"
)

func Prompt(mode int, info string, args ...any) {
	if mode == 1 {
		fmt.Fprintf(os.Stdout, info, args...)
	} else if mode == 2 {
		//ToDo 实现gui的信息提示
	}
}
