package hello_go

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("hello, %v, weclome!", name)
	return message
}
