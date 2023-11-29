package iserver

type IPackUtil interface {

	// Pack message装包
	Pack(msg IMessage) ([]byte, error)

	// UnPack message解包
	UnPack(data []byte) (IMessage, error)
}
