package model

/**
  @author: KafuMio
  @since: 2023/3/3
  @desc: //TODO User各种行为的模型
**/

// 用户点赞
type UserLike struct {
	Model
	UserId     int64  `gorm:"not null;uniqueIndex:idx_user_like_unique;" json:"userId" form:"userId"`                                            // 用户
	EntityId   int64  `gorm:"not null;uniqueIndex:idx_user_like_unique;index:idx_user_like_entity;" json:"topicId" form:"topicId"`               // 实体编号
	EntityType string `gorm:"not null;size:32;uniqueIndex:idx_user_like_unique;index:idx_user_like_entity;" json:"entityType" form:"entityType"` // 实体类型
	CreateTime int64  `json:"createTime" form:"createTime"`                                                                                      // 创建时间
}

// UserFollow 粉丝关注
type UserFollow struct {
	Model
	UserId     int64 `gorm:"not null;uniqueIndex:idx_user_id" json:"userId"`           // 用户编号
	OtherId    int64 `gorm:"not null;uniqueIndex:idx_user_id" json:"otherId"`          // 对方的ID（被关注用户编号）
	Status     int   `gorm:"type:int(11);not null" json:"status"`                      // 关注状态
	CreateTime int64 `gorm:"type:bigint;not null" json:"createTime" form:"createTime"` // 创建时间
}

// UserReport 用户举报
type UserReport struct {
	Model
	DataId      int64  `json:"dataId" form:"dataId"`           // 举报数据ID
	DataType    string `json:"dataType" form:"dataType"`       // 举报数据类型
	UserId      int64  `json:"userId" form:"userId"`           // 举报人ID
	Reason      string `json:"reason" form:"reason"`           // 举报原因
	AuditStatus int64  `json:"auditStatus" form:"auditStatus"` // 审核状态
	AuditTime   int64  `json:"auditTime" form:"auditTime"`     // 审核时间
	AuditUserId int64  `json:"auditUserId" form:"auditUserId"` // 审核人ID
	CreateTime  int64  `json:"createTime" form:"createTime"`   // 举报时间
}

// 收藏
type Favorite struct {
	Model
	UserId     int64  `gorm:"index:idx_favorite_user_id;not null" json:"userId" form:"userId"`                     // 用户编号
	EntityType string `gorm:"index:idx_favorite_entity_type;size:32;not null" json:"entityType" form:"entityType"` // 收藏实体类型
	EntityId   int64  `gorm:"index:idx_favorite_entity_id;not null" json:"entityId" form:"entityId"`               // 收藏实体编号
	CreateTime int64  `json:"createTime" form:"createTime"`                                                        // 创建时间
}

// 评论
type Comment struct {
	Model
	UserId       int64  `gorm:"index:idx_comment_user_id;not null" json:"userId" form:"userId"`             // 用户编号
	EntityType   string `gorm:"index:idx_comment_entity_type;not null" json:"entityType" form:"entityType"` // 被评论实体类型
	EntityId     int64  `gorm:"index:idx_comment_entity_id;not null" json:"entityId" form:"entityId"`       // 被评论实体编号
	Content      string `gorm:"type:text;not null" json:"content" form:"content"`                           // 内容
	ImageList    string `gorm:"type:longtext" json:"imageList" form:"imageList"`                            // 图片
	ContentType  string `gorm:"type:varchar(32);not null" json:"contentType" form:"contentType"`            // 内容类型：markdown、html
	QuoteId      int64  `gorm:"not null"  json:"quoteId" form:"quoteId"`                                    // 引用的评论编号
	LikeCount    int64  `gorm:"not null;default:0" json:"likeCount" form:"likeCount"`                       // 点赞数量
	CommentCount int64  `gorm:"not null;default:0" json:"commentCount" form:"commentCount"`                 // 评论数量
	UserAgent    string `gorm:"size:1024" json:"userAgent" form:"userAgent"`                                // UserAgent
	Ip           string `gorm:"size:128" json:"ip" form:"ip"`                                               // IP
	Status       int    `gorm:"type:int(11);index:idx_comment_status" json:"status" form:"status"`          // 状态：0：待审核、1：审核通过、2：审核失败、3：已发布
	CreateTime   int64  `json:"createTime" form:"createTime"`                                               // 创建时间
}

// 签到
type CheckIn struct {
	Model
	UserId          int64 `gorm:"not null;uniqueIndex:idx_user_id" json:"userId" form:"userId"`         // 用户编号
	LatestDayName   int   `gorm:"type:int(11);not null;index:idx_latest" json:"dayName" form:"dayName"` // 最后一次签到
	ConsecutiveDays int   `gorm:"type:int(11);not null;" json:"consecutiveDays" form:"consecutiveDays"` // 连续签到天数
	CreateTime      int64 `json:"createTime" form:"createTime"`                                         // 创建时间
	UpdateTime      int64 `gorm:"index:idx_latest" json:"updateTime" form:"updateTime"`                 // 更新时间
}
