package main
import "fmt"

//  返回一个友好的问候
func Hi(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}