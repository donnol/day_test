package main

import (
	"fmt"

	gaad "github.com/Comcast/gaad"
)

func main() {
	var buf []byte
	buf = []byte("<ADTS+AAC data> \r\n[hello]")

	// Parsing the buffer
	adts, err := gaad.ParseADTS(buf)
	if err != nil {
		panic(err)
	}

	// Looping through top level elements and accessing sub-elements
	var sbr bool
	if adts.Fill_elements != nil {
		for _, e := range adts.Fill_elements {
			if e.Extension_payload != nil &&
				e.Extension_payload.Extension_type == gaad.EXT_SBR_DATA {
				sbr = true
			}
		}
	}
	fmt.Println(sbr)
}
