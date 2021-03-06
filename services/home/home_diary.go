package home

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
)

type ListDiaryService struct {
	Page        int `form:"page" json:"page" `
	Size        int `form:"size" json:"size" `
	CommunityId int `form:"communityId" json:"communityId" `
	ClassifyId  int `form:"classifyId" json:"classifyId"`
	SubTopicId  int `form:"sub_topic_id" json:"sub_topic_id"`
	UserId      int `form:"user_id" json:"user_id"` //传递userid 进入用户胡中心
}

func (service *ListDiaryService) GetDiarys(userId int64) serializer.Response {

	var diarys []*models.Diary

	total := 0

	if service.Size == 0 {
		service.Size = 10
	}
	if service.Page == 0 {
		service.Page = 1
	}

	start := (service.Page - 1) * service.Size

	//根据前端是否传递ClassifyId 返回对应的数据
	if service.ClassifyId > 0 {
		if err := models.PG.Where("classify_id=? AND community_id=?", service.ClassifyId, service.CommunityId).Model(models.Diary{}).Count(&total).Error; err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库连接错误",
				Error: err.Error(),
			}
		}

		if err := models.PG.
			Preload("UserInfo").
			Preload("CommunityInfo").
			Preload("SubTopicInfo").
			Where("classify_id=? AND community_id=?", service.ClassifyId, service.CommunityId).Order("id desc").Limit(service.Size).Offset(start).Find(&diarys).Error; err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库连接错误",
				Error: err.Error(),
			}
		}

	} else if service.SubTopicId > 0 {
		if err := models.PG.Where(" community_id=? AND sub_topic_id=?", service.CommunityId, service.SubTopicId).Model(models.Diary{}).Count(&total).Error; err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库连接错误",
				Error: err.Error(),
			}
		}

		if err := models.PG.
			Preload("UserInfo").
			Preload("CommunityInfo").
			Preload("SubTopicInfo").
			Where(" community_id=? AND sub_topic_id=?", service.CommunityId, service.SubTopicId).Order("id desc").Limit(service.Size).Offset(start).Find(&diarys).Error; err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库连接错误",
				Error: err.Error(),
			}
		}

	} else if service.UserId > 0 {
		if err := models.PG.Where(" user_id=? ", service.UserId).Model(models.Diary{}).Count(&total).Error; err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库连接错误",
				Error: err.Error(),
			}
		}

		if err := models.PG.
			Preload("UserInfo").
			Preload("CommunityInfo").
			Preload("SubTopicInfo").
			Where(" user_id=? ", service.UserId).Order("id desc").Limit(service.Size).Offset(start).Find(&diarys).Error; err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库连接错误",
				Error: err.Error(),
			}
		}





	} else {
		if err := models.PG.Where("community_id=?", service.CommunityId).Model(models.Diary{}).Count(&total).Error; err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库连接错误",
				Error: err.Error(),
			}
		}

		if err := models.PG.
			Preload("UserInfo").
			Preload("CommunityInfo").
			Preload("SubTopicInfo").
			Where("community_id=?", service.CommunityId).Order("id desc").Limit(service.Size).Offset(start).Find(&diarys).Error; err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库连接错误",
				Error: err.Error(),
			}
		}
	}

	return serializer.BuildListResponse(serializer.BuildDiarys(diarys, userId), uint(total))

}
