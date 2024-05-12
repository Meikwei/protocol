package constant

const (

	///ContentType
	//UserRelated.
	// 定义消息类型和通知类型常量

	// 消息内容类型
	ContentTypeBegin             = 100 // 内容类型开始
	Text                         = 101 // 文本消息
	Picture                      = 102 // 图片消息
	Voice                        = 103 // 音频消息
	Video                        = 104 // 视频消息
	File                         = 105 // 文件消息
	AtText                       = 106 // @文本消息
	Merger                       = 107 // 合并消息
	Card                         = 108 // 卡片消息
	Location                     = 109 // 位置消息
	Custom                       = 110 // 自定义消息
	Revoke                       = 111 // 撤销消息
	Typing                       = 113 // 打字消息
	Quote                        = 114 // 引用消息
	AdvancedText                 = 117 // 高级文本消息
	CustomNotTriggerConversation = 119 // 不触发会话的自定义消息
	CustomOnlineOnly             = 120 // 仅在线的自定义消息
	ReactionMessageModifier      = 121 // 反应消息修改器
	ReactionMessageDeleter       = 122 // 反应消息删除器

	// 消息类别
	Common             = 200 // 普通消息
	GroupMsg           = 201 // 群消息
	SignalMsg          = 202 // 信号消息
	CustomNotification = 203 // 自定义通知

	// 系统相关
	NotificationBegin = 1000 // 通知开始

	// 好友通知
	FriendApplicationApprovedNotification = 1201 // 好友申请通过通知
	FriendApplicationRejectedNotification = 1202 // 好友申请拒绝通知
	FriendApplicationNotification         = 1203 // 好友申请通知
	FriendAddedNotification               = 1204 // 好友添加通知
	FriendDeletedNotification             = 1205 // 好友删除通知
	FriendRemarkSetNotification           = 1206 // 好友备注设置通知
	BlackAddedNotification                = 1207 // 加入黑名单通知
	BlackDeletedNotification              = 1208 // 从黑名单移除通知
	FriendInfoUpdatedNotification         = 1209 // 好友信息更新通知
	FriendsInfoUpdateNotification         = 1210 // 好友信息更新通知

	// 会话变更通知
	ConversationChangeNotification = 1300 // 会话选项变更通知

	// 用户通知
	UserNotificationBegin         = 1301 // 用户通知开始
	UserInfoUpdatedNotification   = 1303 // 用户信息更新通知
	UserStatusChangeNotification  = 1304 // 用户状态改变通知
	UserCommandAddNotification    = 1305 // 用户指令添加通知
	UserCommandDeleteNotification = 1306 // 用户指令删除通知
	UserCommandUpdateNotification = 1307 // 用户指令更新通知
	UserNotificationEnd           = 1399 // 用户通知结束

	// 其他通知
	OANotification = 1400 // 其他通知

	// 群组通知
	GroupNotificationBegin = 1500 // 群组通知开始

	// 群组创建与管理
	GroupCreatedNotification                 = 1501 // 群组创建通知
	GroupInfoSetNotification                 = 1502 // 群组信息设置通知
	JoinGroupApplicationNotification         = 1503 // 加入群组申请通知
	MemberQuitNotification                   = 1504 // 成员退出通知
	GroupApplicationAcceptedNotification     = 1505 // 群组申请接受通知
	GroupApplicationRejectedNotification     = 1506 // 群组申请拒绝通知
	GroupOwnerTransferredNotification        = 1507 // 群主转让通知
	MemberKickedNotification                 = 1508 // 成员被踢出通知
	MemberInvitedNotification                = 1509 // 成员被邀请通知
	MemberEnterNotification                  = 1510 // 成员进入通知
	GroupDismissedNotification               = 1511 // 群组解散通知
	GroupMemberMutedNotification             = 1512 // 群成员禁言通知
	GroupMemberCancelMutedNotification       = 1513 // 群成员取消禁言通知
	GroupMutedNotification                   = 1514 // 群禁言通知
	GroupCancelMutedNotification             = 1515 // 群取消禁言通知
	GroupMemberInfoSetNotification           = 1516 // 群成员信息设置通知
	GroupMemberSetToAdminNotification        = 1517 // 群成员设置为管理员通知
	GroupMemberSetToOrdinaryUserNotification = 1518 // 群成员设置为普通成员通知
	GroupInfoSetAnnouncementNotification     = 1519 // 群设置公告通知
	GroupInfoSetNameNotification             = 1520 // 群设置名称通知

	// 超级群组通知
	SuperGroupNotificationBegin  = 1650 // 超级群组通知开始
	SuperGroupUpdateNotification = 1651 // 超级群组更新通知
	MsgDeleteNotification        = 1652 // 消息删除通知
	SuperGroupNotificationEnd    = 1699 // 超级群组通知结束

	// 会话通知
	ConversationPrivateChatNotification = 1701 // 私聊会话通知
	ConversationUnreadNotification      = 1702 // 未读消息通知

	// 消息撤销通知
	MsgRevokeNotification = 2101

	// 业务通知
	BusinessNotificationBegin = 2000 // 业务通知开始
	BusinessNotification      = 2001 // 业务通知
	BusinessNotificationEnd   = 2099 // 业务通知结束

	// 清除会话通知
	ClearConversationNotification = 2101
	DeleteMsgsNotification        = 2102 // 删除消息通知

	// 已读收据
	HasReadReceipt = 2200

	// 通知结束
	NotificationEnd = 5000

	// 消息状态
	MsgNormal  = 1 // 普通消息
	MsgDeleted = 4 // 已删除消息

	// 消息来源
	UserMsgType = 100 // 用户消息
	SysMsgType  = 200 // 系统消息

	// 会话类型
	SingleChatType = 1 // 单聊
	// WriteGroupChatType Not enabled temporarily // 暂未启用的群聊发送类型
	ReadGroupChatType    = 3 // 群聊（只读）
	NotificationChatType = 4 // 通知会话类型

	// Token 类型
	NormalToken  = 0 // 正常令牌
	InValidToken = 1 // 无效令牌
	KickedToken  = 2 // 被踢出令牌
	ExpiredToken = 3 // 过期令牌

	// 多终端登录
	DefalutNotKick          = 0 // 默认不踢出
	AllLoginButSameTermKick = 1 // 全端登录但同端互斥
	SingleTerminalLogin     = 2 // 单一终端登录
	WebAndOther             = 3 // 网页端与其他端
	PcMobileAndWeb          = 4 // PC端与移动端及网页端
	PCAndOther              = 5 // PC端与其他端

	// 在线状态
	OnlineStatus  = "online"       // 在线状态
	OfflineStatus = "offline"      // 离线状态
	Registered    = "registered"   // 已注册
	UnRegistered  = "unregistered" // 未注册

	// 在线状态值
	Online  = 1 // 在线
	Offline = 0 // 离线

	// 消息接收选项
	ReceiveMessage          = 0 // 接收消息
	NotReceiveMessage       = 1 // 不接收消息
	ReceiveNotNotifyMessage = 2 // 接收但不通知消息

	// 定义消息选项和其他相关常量

	// 消息历史记录选项
	IsHistory = "history"
	// 消息是否持久化
	IsPersistent = "persistent"
	// 是否进行离线推送
	IsOfflinePush = "offlinePush"
	// 是否显示未读计数
	IsUnreadCount = "unreadCount"
	// 是否更新会话信息
	IsConversationUpdate = "conversationUpdate"
	// 发送者是否同步
	IsSenderSync = "senderSync"
	// 消息是否不私密
	IsNotPrivate = "notPrivate"
	// 发送者是否更新会话信息
	IsSenderConversationUpdate = "senderConversationUpdate"
	// 发送者是否推送通知
	IsSenderNotificationPush = "senderNotificationPush"
	// 反应是否来自缓存
	IsReactionFromCache = "reactionFromCache"
	// 是否不发送通知
	IsNotNotification = "isNotNotification"
	// 是否发送消息
	IsSendMsg = "isSendMsg"

	// 定义群组状态常量

	GroupOk              = 0 // 群组正常
	GroupBanChat         = 1 // 群组禁言
	GroupStatusDismissed = 2 // 群组解散
	GroupStatusMuted     = 3 // 群组静音

	// 定义群组类型常量

	NormalGroup  = 0 // 普通群
	SuperGroup   = 1 // 超级群
	WorkingGroup = 2 // 工作群

	// 群组禁言和私聊禁止类型

	GroupBaned          = 3
	GroupBanPrivateChat = 4

	// 用户加入群组来源常量

	JoinByAdmin      = 1 // 管理员添加
	JoinByInvitation = 2 // 邀请加入
	JoinBySearch     = 3 // 搜索加入
	JoinByQRCode     = 4 // 通过二维码加入

	// Minio和Aws的过期时间常量（单位：秒）
	MinioDurationTimes = 3600
	AwsDurationTimes   = 3600

	// 各种回调命令

	CallbackBeforeSendSingleMsgCommand                   = "callbackBeforeSendSingleMsgCommand"                   // 发送单聊消息前的回调命令
	CallbackAfterSendSingleMsgCommand                    = "callbackAfterSendSingleMsgCommand"                    // 发送单聊消息后的回调命令
	CallbackBeforeSendGroupMsgCommand                    = "callbackBeforeSendGroupMsgCommand"                    // 发送群聊消息前的回调命令
	CallbackAfterSendGroupMsgCommand                     = "callbackAfterSendGroupMsgCommand"                     // 发送群聊消息后的回调命令
	CallbackMsgModifyCommand                             = "callbackMsgModifyCommand"                             // 消息修改的回调命令
	CallbackUserOnlineCommand                            = "callbackUserOnlineCommand"                            // 用户上线的回调命令
	CallbackUserOfflineCommand                           = "callbackUserOfflineCommand"                           // 用户下线的回调命令
	CallbackUserKickOffCommand                           = "callbackUserKickOffCommand"                           // 用户被踢下的回调命令
	CallbackOfflinePushCommand                           = "callbackOfflinePushCommand"                           // 离线推送的回调命令
	CallbackOnlinePushCommand                            = "callbackOnlinePushCommand"                            // 在线推送的回调命令
	CallbackSuperGroupOnlinePushCommand                  = "callbackSuperGroupOnlinePushCommand"                  // 超级群在线推送的回调命令
	CallbackBeforeAddFriendCommand                       = "callbackBeforeAddFriendCommand"                       // 添加好友前的回调命令
	CallbackBeforeUpdateUserInfoCommand                  = "callbackBeforeUpdateUserInfoCommand"                  // 更新用户信息前的回调命令
	CallbackBeforeCreateGroupCommand                     = "callbackBeforeCreateGroupCommand"                     // 创建群组前的回调命令
	CallbackBeforeMemberJoinGroupCommand                 = "callbackBeforeMemberJoinGroupCommand"                 // 成员加入群组前的回调命令
	CallbackBeforeSetGroupMemberInfoCommand              = "CallbackBeforeSetGroupMemberInfoCommand"              // 设置群成员信息前的回调命令
	CallbackBeforeSetMessageReactionExtensionCommand     = "callbackBeforeSetMessageReactionExtensionCommand"     // 设置消息反应扩展前的回调命令
	CallbackBeforeDeleteMessageReactionExtensionsCommand = "callbackBeforeDeleteMessageReactionExtensionsCommand" // 删除消息反应扩展前的回调命令
	CallbackGetMessageListReactionExtensionsCommand      = "callbackGetMessageListReactionExtensionsCommand"      // 获取消息列表反应扩展的回调命令
	CallbackAddMessageListReactionExtensionsCommand      = "callbackAddMessageListReactionExtensionsCommand"      // 添加消息列表反应扩展的回调命令

	// 回调操作码和处理结果码

	ActionAllow     = 0 // 允许操作
	ActionForbidden = 1 // 禁止操作

	CallbackHandleSuccess = 0 // 回调处理成功
	CallbackHandleFailed  = 1 // 回调处理失败

	// minio上传类型

	OtherType = 1 // 其他类型
	VideoType = 2 // 视频类型
	ImageType = 3 // 图片类型

	// 消息发送状态

	MsgStatusNotExist = 0 // 消息不存在
	MsgIsSending      = 1 // 消息正在发送
	MsgSendSuccessed  = 2 // 消息发送成功
	MsgSendFailed     = 3 // 消息发送失败
)

// 数据扩散类型
const (
	WriteDiffusion = 0 // 写扩散
	ReadDiffusion  = 1 // 读扩散
)

// 通知可靠性类型
const (
	UnreliableNotification    = 1 // 不可靠的通知
	ReliableNotificationNoMsg = 2 // 可靠的通知，没有消息
	ReliableNotificationMsg   = 3 // 可靠的通知，包含消息
)

// @符号在群组消息中的类型
const (
	AtAllString       = "AtAllTag" // @所有人字符串
	AtNormal          = 0          // 普通@某人
	AtMe              = 1          // @自己
	AtAll             = 2          // @所有人
	AtAllAtMe         = 3          // @所有人且@自己
	GroupNotification = 4          // 群通知
)

// 各种消息内容类型到推送内容的映射
var ContentType2PushContent = map[int64]string{
	Picture:   "[PICTURE]",      // 图片
	Voice:     "[VOICE]",        // 音频
	Video:     "[VIDEO]",        // 视频
	File:      "[File]",         // 文件
	Text:      "[TEXT]",         // 文本
	AtText:    "[@TEXT]",        // @文本
	GroupMsg:  "[GROUPMSG]]",    // 群消息
	Common:    "[NEWMSG]",       // 普通消息
	SignalMsg: "[SIGNALINVITE]", // 信号消息
}

// 消息属性字段
const (
	FieldRecvMsgOpt    = 1  // 接收消息选项
	FieldIsPinned      = 2  // 消息是否被固定
	FieldAttachedInfo  = 3  // 附加信息
	FieldIsPrivateChat = 4  // 是否为私聊消息
	FieldGroupAtType   = 5  // 群@类型
	FieldEx            = 7  // 扩展字段
	FieldUnread        = 8  // 未读数
	FieldBurnDuration  = 9  // 焚烧时间
	FieldHasReadSeq    = 10 // 已读序列
)

// 用户和群组角色类型
const (
	IMOrdinaryUser       = 0 // 普通用户
	AppOrdinaryUsers     = 1 // 应用普通用户
	AppAdmin             = 2 // 应用管理员
	AppNotificationAdmin = 3 // 应用通知管理员

	GroupOwner         = 100 // 群组所有者
	GroupAdmin         = 60  // 群组管理员
	GroupOrdinaryUsers = 20  // 群组普通成员

	GroupResponseAgree  = 1  // 群组操作同意
	GroupResponseRefuse = -1 // 群组操作拒绝

	FriendResponseNotHandle = 0  // 好友请求未处理
	FriendResponseAgree     = 1  // 好友请求同意
	FriendResponseRefuse    = -1 // 好友请求拒绝

	Male   = 1 // 男性
	Female = 2 // 女性
)

const (
	OperationID     = "operationID"
	OpUserID        = "opUserID"
	ConnID          = "connID"
	OpUserPlatform  = "platform"
	Token           = "token"
	RpcCustomHeader = "customHeader" // rpc中间件自定义ctx参数
	CheckKey        = "CheckKey"
	TriggerID       = "triggerID"
	RemoteAddr      = "remoteAddr"
)

const (
	BecomeFriendByImport = 1 // 管理员导入
	BecomeFriendByApply  = 2 // 申请添加
)

const (
	ApplyNeedVerificationInviteDirectly = 0 // 申请需要同意 邀请直接进
	AllNeedVerification                 = 1 // 所有人进群需要验证，除了群主管理员邀请进群
	Directly                            = 2 // 直接进群
)

const (
	GroupRPCRecvSize = 30
	GroupRPCSendSize = 30
)

const FriendAcceptTip = "You have successfully become friends, so start chatting"

func GroupIsBanChat(status int32) bool {
	return status == GroupStatusMuted
}

func GroupIsBanPrivateChat(status int32) bool {
	return status == GroupBanPrivateChat
}

const LogFileName = "OpenIM.log"

const LocalHost = "0.0.0.0"

// flag parse.
const (
	FlagPort                  = "port"
	FlagWsPort                = "ws_port"
	FlagTransferProgressIndex = "transferProgressIndex"
	FlagPrometheusPort        = "prometheus_port"
	FlagConf                  = "config_folder_path"
)

const OpenIMCommonConfigKey = "OpenIMServerConfig"

const CallbackCommand = "command"

const BatchNum = 100

// Subscribe to user constants
const (
	SubscriberUser = 1
	Unsubscribe    = 2
)
