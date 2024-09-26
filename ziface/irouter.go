package ziface

// IRouter 路由抽象接口
type IRouter interface {
	// PreHandle 在处理 conn 业务之前的钩子方法 Hook
	PreHandle(request IRequest)

	// Handle 处理 conn 业务的方法 Hook
	Handle(request IRequest)

	// PostHandle 处理 conn 业务之后的钩子方法 Hook
	PostHandle(request IRequest)
}
