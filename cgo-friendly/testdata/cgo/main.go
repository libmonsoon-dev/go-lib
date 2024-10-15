package main

import "github.com/libmonsoon-dev/go-lib/cgo-friendly/context"

//export ContextBackground
func ContextBackground() context.Context {
	return context.Background()
}

func main() {

}
