package ziface

/**
消息的抽象层接口
将请求的消息封装到Message中
TLV格式
*/

type IMessage interface {
	//获取Message的id
	GetMsgId() uint32
	//获取Message的长度
	GetMsgLen() uint32
	//获取Message的内容
	GetData() []byte

	SetMsgId(id uint32)

	SetData(data []byte)

	SetDataLen(len uint32)
}
