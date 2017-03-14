package happenbefore

import "time"

var A int

func init() {
	time.Sleep(time.Second)
	A = 1
}
