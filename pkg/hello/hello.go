package hello

import "fmt"

func SayHello(name string) string {
	return fmt.Sprintf("Hello %s!\n", name)
}

func PrintSayHello(name string) {
	fmt.Print(SayHello(name))
}
