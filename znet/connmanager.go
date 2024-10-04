package znet

import (
	"fmt"
	"sync"
	"zinx/ziface"
)

type ConnManager struct {
	// 连接集合
	connections map[uint32]ziface.IConnection

	// 保护连接集合的锁
	connLock sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}

func (c *ConnManager) Add(conn ziface.IConnection) {
	// 写锁
	c.connLock.Lock()
	defer c.connLock.Unlock()

	// 将连接添加到 ConnManager 中
	c.connections[conn.GetConnID()] = conn

	// 打印日志
	fmt.Println("connection add to ConnManager successfully: conn num = ", c.Len())
}

func (c *ConnManager) Remove(conn ziface.IConnection) {
	// 写锁
	c.connLock.Lock()
	defer c.connLock.Unlock()

	// 删除连接
	delete(c.connections, conn.GetConnID())

	// 打印日志
	fmt.Println("connection remove ConnID = ", conn.GetConnID(), " successfully: conn num = ", c.Len())
}

func (c *ConnManager) Get(connID uint32) (ziface.IConnection, error) {
	// 读锁
	c.connLock.RLock()
	defer c.connLock.RUnlock()

	if conn, ok := c.connections[connID]; ok {
		return conn, nil
	} else {
		return nil, fmt.Errorf("connection not found: %d", connID)
	}
}

func (c *ConnManager) Len() int {
	return len(c.connections)
}

func (c *ConnManager) ClearConn() {
	// 写锁
	c.connLock.Lock()
	defer c.connLock.Unlock()

	// 删除所有连接
	for connID, conn := range c.connections {
		// 停止
		conn.Stop()
		// 删除
		delete(c.connections, connID)
	}

	// 打印日志
	fmt.Println("Clear All Connections successfully: conn num = ", c.Len())
}
