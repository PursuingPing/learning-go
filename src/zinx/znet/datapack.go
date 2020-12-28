package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"learning-go/src/zinx/utils"
	"learning-go/src/zinx/ziface"
	"log"
)

/**
实现粘包、拆包处理
*/

type DataPack struct{}

//拆包、封装包的实例的初始化方法
func NewDataPack() *DataPack {
	return &DataPack{}
}

func (dp *DataPack) GetHeadLen() uint32 {
	//DataLen + Id  4+4字节
	return 8
}

//封装包
func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	//创建一个buf
	dataBuff := bytes.NewBuffer([]byte{})
	//将dataLen写进
	err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLen())
	if err != nil {
		log.Println("[DataPack]When Pack writing dataLen error", err)
		return nil, err
	}
	//将MsgId写进
	err = binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId())
	if err != nil {
		log.Println("[DataPack]When Pack writing MsgId error", err)
		return nil, err
	}
	//将data写进
	err = binary.Write(dataBuff, binary.LittleEndian, msg.GetData())
	if err != nil {
		log.Println("[DataPack]When Pack writing Data error", err)
		return nil, err
	}

	return dataBuff.Bytes(), nil

}

/**
拆包，将包的Header信息读取出，再读出内容
*/
func (dp *DataPack) Unpack(binaryData []byte) (ziface.IMessage, error) {
	//创建一个从输入二进制数据的ioReader
	reader := bytes.NewReader(binaryData)

	//只取出Head信息得到dataLen和MsgId
	msg := &Message{}
	if err := binary.Read(reader, binary.LittleEndian, &msg.DataLen); err != nil {
		log.Println("[DataPack]When UnPack reading DataLen error", err)
		return nil, err
	}

	if err := binary.Read(reader, binary.LittleEndian, &msg.Id); err != nil {
		log.Println("[DataPack]When UnPack reading MsgId error", err)
		return nil, err
	}

	//是否超出框架定义的包的最大长度
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("Too Large massage data row ")
	}

	return msg, nil
}
