package msgcodes

// MessageCode 消息类型
type MessageCode int32

const (
	// Fail 失败/拒绝
	Fail MessageCode = 0
	// Success 成功
	Success MessageCode = 1
	// Loading 加载中/等待中
	Loading MessageCode = 2
	// Loaded 加载完成
	Loaded MessageCode = 3
	// Null 空
	Null MessageCode = 3
	// VertifyError 验证失败
	VertifyError MessageCode = 3
)
