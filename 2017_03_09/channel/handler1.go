package main

func payloadHandler(data PayloadCollection) {
	for _, payload := range data.Payloads {
		go payload.UploadToS3() // <----- DON'T DO THIS
	}
}
