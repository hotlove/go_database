package websocket

// 请求类型
var (
	MESSAGE_TYPE_LOING  = 1 // 登陆请求
	MESSAGE_TYPE_LOGOUT = 2 // 断开请求
	MESSAGE_TYPE_COMMON = 3 // 聊天
	MESSAGE_TYPE_HEART  = 4 // 心跳检测
)

type Packet struct {
	Token         string // 请求token
	Type          int    // 请求类型
	SendProfileId int    // 发送人id
	RecvProfileId int    // 接收人id
	Message       string // 请求类容
}
