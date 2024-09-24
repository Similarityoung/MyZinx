package ziface

// IServer define a server interface
type IServer interface {
	// Start server
	Start()

	// Stop server
	Stop()

	// Serve Run server
	Serve()
}
