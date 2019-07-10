package xinterface

// IDataPack ..
type IDataPack interface {
	GetHeadLen() uint32

	Pack(msg IMessage) ([]byte, error)

	UnPack(data []byte) (IMessage, error)
}
