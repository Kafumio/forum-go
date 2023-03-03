package model

import "forum_backend/model/constants"

/**
  @author: KafuMio
  @since: 2023/3/3
  @desc: //TODO 话题模型
**/

// TopicNode 话题节点
type TopicNode struct {
	Model
	Name        string `gorm:"size:32;unique" json:"name" form:"name"`                     // 名称
	Description string `gorm:"size:1024" json:"description" form:"description"`            // 描述
	Logo        string `gorm:"size:1024" json:"logo" form:"logo"`                          // 图标
	SortNo      int    `gorm:"type:int(11);index:idx_sort_no" json:"sortNo" form:"sortNo"` // 排序编号
	Status      int    `gorm:"type:int(11);not null" json:"status" form:"status"`          // 状态
	CreateTime  int64  `json:"createTime" form:"createTime"`                               // 创建时间
}

// 话题
type Topic struct {
	Model
	Type              constants.TopicType `gorm:"type:int(11);not null:default:0" json:"type" form:"type"`                         // 类型
	NodeId            int64               `gorm:"not null;index:idx_node_id;" json:"nodeId" form:"nodeId"`                         // 节点编号
	UserId            int64               `gorm:"not null;index:idx_topic_user_id;" json:"userId" form:"userId"`                   // 用户
	Title             string              `gorm:"size:128" json:"title" form:"title"`                                              // 标题
	Content           string              `gorm:"type:longtext" json:"content" form:"content"`                                     // 内容
	ImageList         string              `gorm:"type:longtext" json:"imageList" form:"imageList"`                                 // 图片
	HideContent       string              `gorm:"type:longtext" json:"hideContent" form:"hideContent"`                             // 回复可见内容
	Recommend         bool                `gorm:"not null;index:idx_recommend" json:"recommend" form:"recommend"`                  // 是否推荐
	RecommendTime     int64               `gorm:"not null" json:"recommendTime" form:"recommendTime"`                              // 推荐时间
	Sticky            bool                `gorm:"not null;index:idx_sticky_sticky_time" json:"sticky" form:"sticky"`               // 置顶
	StickyTime        int64               `gorm:"not null;index:idx_sticky_sticky_time" json:"stickyTime" form:"stickyTime"`       // 置顶时间
	ViewCount         int64               `gorm:"not null" json:"viewCount" form:"viewCount"`                                      // 查看数量
	CommentCount      int64               `gorm:"not null" json:"commentCount" form:"commentCount"`                                // 跟帖数量
	LikeCount         int64               `gorm:"not null" json:"likeCount" form:"likeCount"`                                      // 点赞数量
	Status            int                 `gorm:"type:int(11);index:idx_topic_status;" json:"status" form:"status"`                // 状态：0：正常、1：删除
	LastCommentTime   int64               `gorm:"index:idx_topic_last_comment_time" json:"lastCommentTime" form:"lastCommentTime"` // 最后回复时间
	LastCommentUserId int64               `json:"lastCommentUserId" form:"lastCommentUserId"`                                      // 最后回复用户
	UserAgent         string              `gorm:"size:1024" json:"userAgent" form:"userAgent"`                                     // UserAgent
	Ip                string              `gorm:"size:128" json:"ip" form:"ip"`                                                    // IP
	CreateTime        int64               `gorm:"index:idx_topic_create_time" json:"createTime" form:"createTime"`                 // 创建时间
	ExtraData         string              `gorm:"type:text" json:"extraData" form:"extraData"`                                     // 扩展数据
}

// 话题标签
type TopicTag struct {
	Model
	TopicId           int64 `gorm:"not null;index:idx_topic_tag_topic_id;" json:"topicId" form:"topicId"`                // 主题编号
	TagId             int64 `gorm:"not null;index:idx_topic_tag_tag_id;" json:"tagId" form:"tagId"`                      // 标签编号
	Status            int64 `gorm:"not null;index:idx_topic_tag_status" json:"status" form:"status"`                     // 状态：正常、删除
	LastCommentTime   int64 `gorm:"index:idx_topic_tag_last_comment_time" json:"lastCommentTime" form:"lastCommentTime"` // 最后回复时间
	LastCommentUserId int64 `json:"lastCommentUserId" form:"lastCommentUserId"`                                          // 最后回复用户
	CreateTime        int64 `json:"createTime" form:"createTime"`                                                        // 创建时间
}
