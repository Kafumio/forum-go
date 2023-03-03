package model

import (
	"database/sql"
	"forum_backend/model/constants"
	"time"
)

/**
  @author: KafuMio
  @since: 2023/3/3
  @desc: //TODO User模型及与User相关的信息模型
**/

type User struct {
	Model
	Username         sql.NullString   `gorm:"size:32;unique;" json:"username" form:"username"`                         // 用户名
	Email            sql.NullString   `gorm:"size:128;unique;" json:"email" form:"email"`                              // 邮箱
	EmailVerified    bool             `gorm:"not null;default:false" json:"emailVerified" form:"emailVerified"`        // 邮箱是否验证
	Nickname         string           `gorm:"size:16;" json:"nickname" form:"nickname"`                                // 昵称
	Avatar           string           `gorm:"type:text" json:"avatar" form:"avatar"`                                   // 头像
	Gender           constants.Gender `gorm:"size:16;default:''" json:"gender" form:"gender"`                          // 性别
	Birthday         *time.Time       `json:"birthday" form:"birthday"`                                                // 生日
	BackgroundImage  string           `gorm:"type:text" json:"backgroundImage" form:"backgroundImage"`                 // 个人中心背景图片
	Password         string           `gorm:"size:512" json:"password" form:"password"`                                // 密码
	HomePage         string           `gorm:"size:1024" json:"homePage" form:"homePage"`                               // 个人主页
	Description      string           `gorm:"type:text" json:"description" form:"description"`                         // 个人描述
	Score            int              `gorm:"type:int(11);not null;index:idx_user_score" json:"score" form:"score"`    // 积分
	Status           int              `gorm:"type:int(11);index:idx_user_status;not null" json:"status" form:"status"` // 状态
	TopicCount       int              `gorm:"type:int(11);not null" json:"topicCount" form:"topicCount"`               // 帖子数量
	CommentCount     int              `gorm:"type:int(11);not null" json:"commentCount" form:"commentCount"`           // 跟帖数量
	FollowCount      int              `gorm:"type:int(11);not null" json:"followCount" form:"followCount"`             // 关注数量
	FansCount        int              `gorm:"type:int(11);not null" json:"fansCount" form:"fansCount"`                 // 粉丝数量
	Roles            string           `gorm:"type:text" json:"roles" form:"roles"`                                     // 角色
	ForbiddenEndTime int64            `gorm:"not null;default:0" json:"forbiddenEndTime" form:"forbiddenEndTime"`      // 禁言结束时间
	CreateTime       int64            `json:"createTime" form:"createTime"`                                            // 创建时间
	UpdateTime       int64            `json:"updateTime" form:"updateTime"`                                            // 更新时间
}

type UserToken struct {
	Model
	Token      string `gorm:"size:32;unique;not null" json:"token" form:"token"`
	UserId     int64  `gorm:"not null;index:idx_user_token_user_id;" json:"userId" form:"userId"`
	ExpiredAt  int64  `gorm:"not null" json:"expiredAt" form:"expiredAt"`
	Status     int    `gorm:"type:int(11);not null;index:idx_user_token_status" json:"status" form:"status"`
	CreateTime int64  `gorm:"not null" json:"createTime" form:"createTime"`
}

type ThirdAccount struct {
	Model
	UserId     sql.NullInt64 `gorm:"uniqueIndex:idx_user_id_third_type;" json:"userId" form:"userId"`                                  // 用户编号
	Avatar     string        `gorm:"size:1024" json:"avatar" form:"avatar"`                                                            // 头像
	Nickname   string        `gorm:"size:32" json:"nickname" form:"nickname"`                                                          // 昵称
	ThirdType  string        `gorm:"size:32;not null;uniqueIndex:idx_user_id_third_type,idx_third;" json:"thirdType" form:"thirdType"` // 第三方类型
	ThirdId    string        `gorm:"size:64;not null;uniqueIndex:idx_third;" json:"thirdId" form:"thirdId"`                            // 第三方唯一标识，例如：openId,unionId
	ExtraData  string        `gorm:"type:longtext" json:"extraData" form:"extraData"`                                                  // 扩展数据
	CreateTime int64         `json:"createTime" form:"createTime"`                                                                     // 创建时间
	UpdateTime int64         `json:"updateTime" form:"updateTime"`                                                                     // 更新时间
}

// 用户积分流水
type UserScoreLog struct {
	Model
	UserId      int64  `gorm:"not null;index:idx_user_score_log_user_id" json:"userId" form:"userId"`   // 用户编号
	SourceType  string `gorm:"not null;index:idx_user_score_score" json:"sourceType" form:"sourceType"` // 积分来源类型
	SourceId    string `gorm:"not null;index:idx_user_score_score" json:"sourceId" form:"sourceId"`     // 积分来源编号
	Description string `json:"description" form:"description"`                                          // 描述
	Type        int    `gorm:"type:int(11)" json:"type" form:"type"`                                    // 类型(增加、减少)
	Score       int    `gorm:"type:int(11)" json:"score" form:"score"`                                  // 积分
	CreateTime  int64  `json:"createTime" form:"createTime"`                                            // 创建时间
}

// UserFeed 用户信息流
type UserFeed struct {
	Model
	UserId     int64  `gorm:"not null;uniqueIndex:idx_data;index:idx_user_id;index:idx_search" json:"userId"`                   // 用户编号
	DataId     int64  `gorm:"not null;uniqueIndex:idx_data;index:idx_data_id" json:"dataId" form:"dataId"`                      // 数据ID
	DataType   string `gorm:"not null;uniqueIndex:idx_data;index:idx_data_id;index:idx_search" json:"dataType" form:"dataType"` // 数据类型
	AuthorId   int64  `gorm:"not null;index:idx_user_id" json:"authorId" form:"authorId"`                                       // 作者编号
	CreateTime int64  `gorm:"type:bigint;not null;index:idx_search" json:"createTime" form:"createTime"`                        // 数据的创建时间
}

// 消息
type Message struct {
	Model
	FromId       int64  `gorm:"not null" json:"fromId" form:"fromId"`                            // 消息发送人
	UserId       int64  `gorm:"not null;index:idx_message_user_id;" json:"userId" form:"userId"` // 用户编号(消息接收人)
	Title        string `gorm:"size:1024" json:"title" form:"title"`                             // 消息标题
	Content      string `gorm:"type:text;not null" json:"content" form:"content"`                // 消息内容
	QuoteContent string `gorm:"type:text" json:"quoteContent" form:"quoteContent"`               // 引用内容
	Type         int    `gorm:"type:int(11);not null" json:"type" form:"type"`                   // 消息类型
	ExtraData    string `gorm:"type:text" json:"extraData" form:"extraData"`                     // 扩展数据
	Status       int    `gorm:"type:int(11);not null" json:"status" form:"status"`               // 状态：0：未读、1：已读
	CreateTime   int64  `json:"createTime" form:"createTime"`                                    // 创建时间
}

// 操作日志
type OperateLog struct {
	Model
	UserId      int64  `gorm:"not null;index:idx_operate_log_user_id" json:"userId" form:"userId"`  // 用户编号
	OpType      string `gorm:"not null;index:idx_op_type;size:32" json:"opType" form:"opType"`      // 操作类型
	DataType    string `gorm:"not null;index:idx_operate_log_data" json:"dataType" form:"dataType"` // 数据类型
	DataId      int64  `gorm:"not null;index:idx_operate_log_data" json:"dataId" form:"dataId" `    // 数据编号
	Description string `gorm:"not null;size:1024" json:"description" form:"description"`            // 描述
	Ip          string `gorm:"size:128" json:"ip" form:"ip"`                                        // ip地址
	UserAgent   string `gorm:"type:text" json:"userAgent" form:"userAgent"`                         // UserAgent
	Referer     string `gorm:"type:text" json:"referer" form:"referer"`                             // Referer
	CreateTime  int64  `json:"createTime" form:"createTime"`                                        // 创建时间
}

// 邮箱验证码
type EmailCode struct {
	Model
	UserId     int64  `gorm:"not null;index:idx_user_score_log_user_id" json:"userId" form:"userId"` // 用户编号
	Email      string `gorm:"not null;size:128" json:"email" form:"email"`                           // 邮箱
	Code       string `gorm:"not null;size:8" json:"code" form:"code"`                               // 验证码
	Token      string `gorm:"not null;size:32;unique" json:"token" form:"token"`                     // 验证码token
	Title      string `gorm:"size:1024" json:"title" form:"title"`                                   // 标题
	Content    string `gorm:"type:text" json:"content" form:"content"`                               // 内容
	Used       bool   `gorm:"not null" json:"used" form:"used"`                                      // 是否使用
	CreateTime int64  `json:"createTime" form:"createTime"`                                          // 创建时间
}
