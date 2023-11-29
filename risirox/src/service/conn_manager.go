package service

import (
	"errors"
	log "github.com/sirupsen/logrus"
	iserver "risirox/risirox/src/server"
	"sync"
)

type ConnPoolManager struct {
	//链接池
	connPool map[uint32]iserver.IConnection

	//保护链接池的读写锁
	connLock sync.RWMutex
}

// AddConn 向链接池中添加一个链接
func (connManager *ConnPoolManager) AddConn(conn iserver.IConnection) error {
	connManager.connLock.Lock()
	defer connManager.connLock.Unlock()
	if _, ok := connManager.connPool[conn.GetConnId()]; ok {
		log.Errorf("Connection is already in pool, id:%d", conn.GetConnId())
		return errors.New("connection is already in pool")
	}
	connManager.connPool[conn.GetConnId()] = conn
	return nil
}

// DelConn 删除链接池中的一个链接
func (connManager *ConnPoolManager) DelConn(connId uint32) {
	connManager.connLock.Lock()
	defer connManager.connLock.Unlock()
	delete(connManager.connPool, connId)
}

// ClearConn ClearPool 删除链接池的所有链接
func (connManager *ConnPoolManager) ClearConn() {
	connManager.connLock.RLock()
	defer connManager.connLock.RUnlock()
	for connId, connection := range connManager.connPool {
		connection.Stop()
		delete(connManager.connPool, connId)
	}
}

func (connManager *ConnPoolManager) GetPoolSize() int {
	return len(connManager.connPool)
}

func (connManager *ConnPoolManager) GetConnById(connId uint32) iserver.IConnection {
	return connManager.connPool[connId]
}

func (connManager *ConnPoolManager) GetAllConn() map[uint32]iserver.IConnection {
	return connManager.connPool
}

func NewConnPoolManager() *ConnPoolManager {
	return &ConnPoolManager{
		connPool: make(map[uint32]iserver.IConnection),
	}
}
