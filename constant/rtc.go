package constant

// 以下常量定义了不同类型的信号通知。
// 这些常量用于标识与会议相关的不同事件或状态改变。
const (
	SignalingNotificationBegin               = 1600 // 信号通知开始的标识
	SignalingNotification                    = 1601 // 通用信号通知
	RoomParticipantsConnectedNotification    = 1602 // 房间参与者已连接通知
	RoomParticipantsDisconnectedNotification = 1603 // 房间参与者已断开连接通知
	StreamChangedNotification                = 1604 // 流变更通知
	CustomSignalNotification                 = 1605 // 自定义信号通知

	SignalingNotificationEnd = 1699 // 信号通知结束的标识
)

// 以下常量定义了会议的结束状态。
const (
	MeetingNotEnd = 0 // 会议未结束
	MeetingEnd    = 1 // 会议已结束
)
