package util

import (
	"fmt"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	fmt.Println("Hello, World!")
	fmt.Println("Router initialized:", r)
}

func Test() {
	fmt.Println("abc")
}
