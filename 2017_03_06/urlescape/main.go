package main

import (
	"fmt"
	"net/url"
)

func main() {
	keywords := []string{
		"\xe4\xb8",  // unicode 16进制字符串
		"%\xe4\xb8", // panic: invalid URL escape "%\xe4\xb8"
	}

	for _, single := range keywords {
		singleKeyword, err := url.QueryUnescape(single)
		if err != nil {
			panic(err)
		}
		fmt.Println(singleKeyword)
	}
}
