package home

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
	"github.com/jinzhu/gorm"
)

type AddDiaryService struct {
	UserId      uint
	Content     string   `form:"content" json:"content" `
	Address     string   `form:"address" json:"address" `
	Community   string   `form:"community" json:"community"`
	Photos      []string `form:"photos" json:"photos"`
	PhotosThumb []string `form:"photosthumb" json:"photosthumb"`
	CommunityId uint     `form:"communityId" json:"communityId" `  //社区id
	ClassifyId  uint     `form:"classifyId" json:"classifyId" `    //标签ID
	SubTopicId  uint     `form:"sub_topic_id" json:"sub_topic_id"` //标签ID
}

func (diary *AddDiaryService) AddDiary(userId uint) serializer.Response {

	dia := models.Diary{
		UserId:      userId,
		Content:     diary.Content,
		Address:     diary.Address,
		Photos:      diary.Photos,
		PhotosThumb: diary.Photos,
		CommunityId: diary.CommunityId,
		SubTopicId:  diary.SubTopicId,
		ClassifyId:  diary.ClassifyId,
	}

	//创建话题
	models.PG.Create(&dia)

	//更新tag sendNum
	var subTopic models.SubTopic
	subTopic.ID = diary.SubTopicId
	models.PG.Model(&subTopic).UpdateColumn("send_num", gorm.Expr("send_num + ?", 1))

	return serializer.Response{
		Code:  0,
		Data:  nil,
		Msg:   "创建成功",
		Error: "",
	}

}
