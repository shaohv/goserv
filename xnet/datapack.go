package xnet

import (
	"bytes"
	"encoding/binary"
	"goserv/xinterface"
)

// DataPack ..
type DataPack struct{}

// NewDataPack ..
func NewDataPack() *DataPack {
	return &DataPack{}
}

// GetHeadLen ..
func (dp *DataPack) GetHeadLen() uint32 {
	return 8
}

// Pack ..
func (dp *DataPack) Pack(msg xinterface.IMessage) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})

	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

// UnPack ..
func (dp *DataPack) UnPack(binData []byte) (xinterface.IMessage, error) {
	dataBuff := bytes.NewReader(binData)

	msg := &Message{}

	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	return msg, nil
}
