// Code generated by sdkcodegen; DO NOT EDIT.

package workwx

// ChatInfo 群聊信息
type ChatInfo struct {
	// ChatID 群聊唯一标志
	ChatID string `json:"chatid"`
	// Name 群聊名
	Name string `json:"name"`
	// OwnerUserID 群主id
	OwnerUserID string `json:"owner"`
	// MemberUserIDs 群成员id列表
	MemberUserIDs []string `json:"userlist"`
}

// ReqChatListOwnerFilter 群主过滤
type ReqChatListOwnerFilter struct {
	// UserIDList 用户ID列表。最多100个
	UserIDList []string `json:"userid_list"`
}

// ReqChatList 获取客户群列表参数
type ReqChatList struct {
	// StatusFilter 客户群跟进状态过滤
	StatusFilter int64 `json:"status_filter"`
	// OwnerFilter 群主过滤
	OwnerFilter ReqChatListOwnerFilter `json:"owner_filter"`
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用不填
	Cursor string `json:"cursor"`
	// Limit 分页，预期请求的数据量，取值范围 1 ~ 1000
	Limit int64 `json:"limit"`
}

// RespGroupChatList 客户群列表数据
type RespGroupChatList struct {
	// ChatID 客户群ID
	ChatID string `json:"chat_id"`
	// Status 客户群跟进状态 0 - 跟进人正常 1 - 跟进人离职 2 - 离职继承中 3 - 离职继承完成
	Status int64 `json:"status"`
}

// RespAppchatList 客户群列表结果
type RespAppchatList struct {
	// GroupChatList 客户群列表
	GroupChatList []RespGroupChatList `json:"group_chat_list"`
	// NextCursor 分页游标
	NextCursor string `json:"next_cursor"`
}

// ChatMemberList 客户群成员列表
type ChatMemberList struct {
	// UserID 群成员ID
	UserID string `json:"userid"`
	// Type 群成员类型 1 - 企业成员  2 - 外部联系人
	Type int64 `json:"type"`
	// UnionID 微信unionid
	UnionID string `json:"unionid"`
	// JoinTime 入群时间
	JoinTime int64 `json:"join_time"`
	// JoinScene 入群方式。1 - 由群成员邀请入群（直接邀请入群）2 - 由群成员邀请入群（通过邀请链接入群）3 - 通过扫描群二维码入群
	JoinScene int64 `json:"join_scene"`
	// Invitor 邀请者。目前仅当是由本企业内部成员邀请入群时会返回该值
	Invitor ChatMemberListInvitor `json:"invitor"`
	// GroupNickname 在群里的昵称
	GroupNickname string `json:"group_nickname"`
	// Name 在群里名字
	Name string `json:"name"`
}

// ChatMemberListInvitor 入群邀请者
type ChatMemberListInvitor struct {
	// UserID 邀请者ID
	UserID string `json:"userid"`
}

// ChatAdminList 客户群管理员列表
type ChatAdminList struct {
	// UserID 管理员ID
	UserID string `json:"userid"`
}

// RespAppChatInfo 客户群详情
type RespAppChatInfo struct {
	// ChatID 客户群ID
	ChatID string `json:"chat_id"`
	// Name 客户群名称
	Name string `json:"name"`
	// Owner 群主ID
	Owner string `json:"owner"`
	// CreateTime 群创建时间
	CreateTime int64 `json:"create_time"`
	// Notice 群公告
	Notice string `json:"notice"`
	// MemberList 群成员列表
	MemberList []*ChatMemberList `json:"member_list"`
	// AdminList 群管理员列表
	AdminList []*ChatAdminList `json:"admin_list"`
}

// ChatNeedName 是否需要返回群成员的名字 0-不返回；1-返回。默认不返回
const ChatNeedName int64 = 1
