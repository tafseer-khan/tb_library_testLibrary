package lib

import "math/rand"

//export answer
func answer() uint32 {
	return rand.Uint32()
}
