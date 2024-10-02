package ziface

// IServer define a server interface
type IServer interface {
	// Start server
	Start()

	// Stop server
	Stop()

	// Serve Run server
	Serve()

	// AddRouter Add router to server
	AddRouter(mgsId uint32, router IRouter)
}
