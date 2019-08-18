package xnet

import (
	"errors"
	"fmt"
	"goserv/xinterface"
	"sync"
)

// ConnMgr ...
type ConnMgr struct {
	connections map[uint32]xinterface.IConnection
	connLock    sync.RWMutex
}

// NewConnMgr ...
func NewConnMgr() *ConnMgr {
	return &ConnMgr{
		connections: make(map[uint32]xinterface.IConnection),
	}
}

// Add ..
func (cm *ConnMgr) Add(conn xinterface.IConnection) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	cm.connections[conn.GetConnID()] = conn

	fmt.Println("connections", conn.GetConnID(), "add to ConnMgr successfully, num:", cm.Len())
}

// Remove ..
func (cm *ConnMgr) Remove(conn xinterface.IConnection) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	delete(cm.connections, conn.GetConnID())

	fmt.Println("connections", conn.GetConnID(), "del from ConnMgr successfully, num:", cm.Len())
}

// Get ..
func (cm *ConnMgr) Get(connID uint32) (xinterface.IConnection, error) {
	cm.connLock.RLock()
	defer cm.connLock.RUnlock()

	if connection, ok := cm.connections[connID]; ok {
		return connection, nil
	} else {
		return nil, errors.New("key not found!")
	}
}

// Len ..
func (cm *ConnMgr) Len() int {
	return len(cm.connections)
}

// ClearConn ...
func (cm *ConnMgr) ClearConn() {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	for connID, conn := range cm.connections {
		conn.Stop()

		delete(cm.connections, connID)
	}

	fmt.Println("All connections clear, num=", cm.Len())
}
