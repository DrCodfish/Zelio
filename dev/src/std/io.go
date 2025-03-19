package std

import "fmt"

func Println(a ...interface{}) (n int, err error) {
    return fmt.Println(a...)
}