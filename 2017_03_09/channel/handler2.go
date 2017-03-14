package main

var Queue chan Payload

func init() {
	// buffer chan
	Queue = make(chan Payload, 100)
}

func payloadHandler2(data PayloadCollection) {
	// Go through each payload and queue items individually to be posted to S3
	for _, payload := range data.Payloads {
		Queue <- payload
	}
}

func StartProcessor() {
	for {
		select {
		case job := <-Queue:
			job.UploadToS3() // <-- STILL NOT GOOD
		}
	}
}
