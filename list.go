package lib

import (
	"math/rand"

	"github.com/taubyte/go-sdk/event"
)

//export answer
func answer(e event.Event) uint32 {
	return rand.Uint32()
}
