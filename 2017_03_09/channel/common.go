package main

import "fmt"

type PayloadCollection struct {
	WindowsVersion string    `json:"version"`
	Token          string    `json:"token"`
	Payloads       []Payload `json:"data"`
}

type Payload struct {
	// [redacted]
	storageFolder string
}

func (p *Payload) UploadToS3() error {
	// the storageFolder method ensures that there are no name collision in
	// case we get same timestamp in the key name
	// storage_path := fmt.Sprintf("%v/%v", p.storageFolder, time.Now().UnixNano())

	fmt.Printf("%v", "hello")

	return nil
}
