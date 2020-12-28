package ziface

/**
定义一个Message的粘包与拆包接口
直接面向TCP数据流
*/

type IDataPack interface {
	GetHeadLen() uint32
	Pack(msg IMessage) ([]byte, error)
	Unpack([]byte) (IMessage, error)
}
