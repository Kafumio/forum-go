package model

/**
  @author: KafuMio
  @since: 2023/3/3
  @desc: //TODO 文章及标签模型
**/

// 文章
type Article struct {
	Model
	CategoryId   int64  `gorm:"default:0;index:idx_category_id" json:"categoryId" form:"categoryId"` // 分类ID
	UserId       int64  `gorm:"index:idx_article_user_id" json:"userId" form:"userId"`               // 所属用户编号
	Title        string `gorm:"size:128;not null;" json:"title" form:"title"`                        // 标题
	Summary      string `gorm:"type:text" json:"summary" form:"summary"`                             // 摘要
	Content      string `gorm:"type:longtext;not null;" json:"content" form:"content"`               // 内容
	ContentType  string `gorm:"type:varchar(32);not null" json:"contentType" form:"contentType"`     // 内容类型：markdown、html
	Cover        string `gorm:"type:text;" json:"cover" form:"cover"`                                // 封面图
	Status       int    `gorm:"type:int(11);index:idx_article_status" json:"status" form:"status"`   // 状态
	SourceUrl    string `gorm:"type:text" json:"sourceUrl" form:"sourceUrl"`                         // 原文链接
	ViewCount    int64  `gorm:"not null;" json:"viewCount" form:"viewCount"`                         // 查看数量
	CommentCount int64  `gorm:"default:0" json:"commentCount" form:"commentCount"`                   // 评论数量
	LikeCount    int64  `gorm:"default:0" json:"likeCount" form:"likeCount"`                         // 点赞数量
	CreateTime   int64  `json:"createTime" form:"createTime"`                                        // 创建时间
	UpdateTime   int64  `json:"updateTime" form:"updateTime"`                                        // 更新时间
}

// 标签
type Tag struct {
	Model
	Name        string `gorm:"size:32;unique;not null" json:"name" form:"name"`
	Description string `gorm:"size:1024" json:"description" form:"description"`
	Status      int    `gorm:"type:int(11);index:idx_tag_status;not null" json:"status" form:"status"`
	CreateTime  int64  `json:"createTime" form:"createTime"`
	UpdateTime  int64  `json:"updateTime" form:"updateTime"`
}

// 文章标签
type ArticleTag struct {
	Model
	ArticleId  int64 `gorm:"not null;index:idx_article_id;" json:"articleId" form:"articleId"`  // 文章编号
	TagId      int64 `gorm:"not null;index:idx_article_tag_tag_id;" json:"tagId" form:"tagId"`  // 标签编号
	Status     int64 `gorm:"not null;index:idx_article_tag_status" json:"status" form:"status"` // 状态：正常、删除
	CreateTime int64 `json:"createTime" form:"createTime"`                                      // 创建时间
}
