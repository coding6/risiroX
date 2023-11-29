package iserver

type IConnPoolManager interface {
	// AddConn 向链接池中添加一个链接
	AddConn(conn IConnection) error

	// DelConn 删除链接池中的一个链接
	DelConn(connId uint32)

	// ClearConn ClearPool 删除所有链接
	ClearConn()

	// GetConnById 通过id获取获取指定链接
	GetConnById(connId uint32) IConnection

	// GetPoolSize 获取目前链接池的大小
	GetPoolSize() int

	// GetAllConn 获取当前连接池的所有连接
	GetAllConn() map[uint32]IConnection
}
