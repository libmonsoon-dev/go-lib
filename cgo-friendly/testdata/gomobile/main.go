package gomobile

import (
	_ "golang.org/x/mobile/bind"

	"github.com/libmonsoon-dev/go-lib/cgo-friendly/context"
)

func ContextBackground() context.Context {
	return context.Background()
}
