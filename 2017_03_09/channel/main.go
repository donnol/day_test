package main

func main() {
	data := PayloadCollection{}

	// 1
	payloadHandler(data)

	// 2
	payloadHandler2(data)

	// 3
	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()

	payloadHandler3(data)
}
