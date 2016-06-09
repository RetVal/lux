package tornago

import (
	"fmt"
	"io"
	"os"
)

var printer io.Writer

func init() {
	printer = os.Stdout
}

func d(a ...interface{}) {
	fmt.Println(a...)
}

func df(s string, a ...interface{}) {
	fmt.Printf(s, a...)
}
