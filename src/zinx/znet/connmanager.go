package znet

import (
	"errors"
	"learning-go/src/zinx/ziface"
	"log"
	"sync"
)

/*
	连接管理模块
*/

type ConnManager struct {
	connections map[uint32]ziface.IConnection
	connLock    sync.RWMutex
}

//创建当前连接管理的方法
func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}

//添加连接
func (connMgr *ConnManager) Add(conn ziface.IConnection) {
	//保护共享资源，加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	//将conn加入到ConnManager中
	connMgr.connections[conn.GetConnID()] = conn
	log.Println("[ConnManager] add connection successfully, connID = ", conn.GetConnID())
}

//删除连接
func (connMgr *ConnManager) Remove(conn ziface.IConnection) {
	//保护共享资源，加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	delete(connMgr.connections, conn.GetConnID())
	log.Println("[ConnManager] remove connection successfully, connID = ", conn.GetConnID())
}

//根据连接id获取连接
func (connMgr *ConnManager) Get(connID uint32) (ziface.IConnection, error) {
	//加读锁
	connMgr.connLock.RLock()
	defer connMgr.connLock.RUnlock()

	if conn, ok := connMgr.connections[connID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("[ConnManager] get connection NOT FOUND")
	}
}

//获取当前连接总数
func (connMgr *ConnManager) Len() int {
	return len(connMgr.connections)
}

//清楚并终止所有的连接
func (connMgr *ConnManager) ClearConn() {
	//保护共享资源，加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	//删除connection，并停止conn工作
	for connID, conn := range connMgr.connections {
		//停止
		conn.Stop()
		//删除
		delete(connMgr.connections, connID)
	}
	log.Println("[ConnManager] Clear All Connections")
}
