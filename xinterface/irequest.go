package xinterface

// IRequest store conn and data
type IRequest interface {
	GetConnection() IConnection
	GetData() []byte
	GetMsgID() uint32
}
