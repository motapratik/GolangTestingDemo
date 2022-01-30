package function

import "fmt"

func Greeting(msg string) string {
	return fmt.Sprintf("Hello %v!", msg)
}

func Addition(a int, b int) int {
	return a + b
}

func DisplayMsg(msg string) string {
	return fmt.Sprintf("Msg: %v", msg)
}
