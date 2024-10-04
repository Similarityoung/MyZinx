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

	// GetConnManager Get connection manager
	GetConnManager() IConnManager

	// SetOnConnStart Set callback function when connection start
	SetOnConnStart(hookFunc func(conn IConnection))

	// SetOnConnStop Set callback function when connection stop
	SetOnConnStop(hookFunc func(conn IConnection))

	// CallOnConnStart Call callback function when connection start
	CallOnConnStart(conn IConnection)

	// CallOnConnStop Call callback function when connection stop
	CallOnConnStop(conn IConnection)
}
