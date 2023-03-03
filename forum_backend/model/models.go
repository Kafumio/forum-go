package model

/**
  @author: KafuMio
  @since: 2023/3/3
  @desc: //TODO 用models集成各种模型
**/

var Models = []interface{}{
	&User{}, &UserToken{}, &Tag{}, &Article{}, &ArticleTag{}, &Comment{}, &Favorite{}, &Topic{}, &TopicNode{},
	&TopicTag{}, &UserLike{}, &Message{}, &SysConfig{}, &Link{}, &ThirdAccount{},
	&UserScoreLog{}, &OperateLog{}, &EmailCode{}, &CheckIn{}, &UserFollow{}, &UserFeed{}, &UserReport{},
	&ForbiddenWord{},
}

type Model struct {
	Id int64 `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
}
