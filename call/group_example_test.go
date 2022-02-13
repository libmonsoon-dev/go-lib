package call_test

import (
	"fmt"
	"time"

	"github.com/libmonsoon-dev/go-lib/call"
)

func ExampleGroup() {
	before := time.Now()
	group := call.NewGroup(call.GroupConfig{})
	group.DoFunc("sleep 1", sleep(1)).
		DoFunc("sleep 2", sleep(2)).
		DoFunc("sleep 3", sleep(3))

	args, err := group.GetResult()
	// Output: Arguments: nil
	// Error: <nil>
	// Execution took 3 seconds
	fmt.Println("Arguments:", args)
	fmt.Printf("Error: %v\n", err)
	fmt.Printf("Execution took %d seconds", int(time.Now().Sub(before).Seconds()))
}

func sleep(seconds int) call.DoFunc {
	return func() error {
		time.Sleep(time.Second * time.Duration(seconds))
		return nil
	}
}
