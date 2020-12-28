package ziface

/**
  把client连接信息与 请求数据包装到一个Request中
*/

type IRequest interface {
	//得到当前连接
	GetConnection() IConnection
	//得到请求数据
	GetData() []byte
}
