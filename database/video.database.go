package database

import "akikflix/model"

func CreateVideoModel(video *model.Video) (model.Video, error) {
	result := DB.Create(&video)
	if result.RowsAffected == 0 {
		return *video, result.Error
	}
	return *video, nil
}

func GetVideo(id string) (model.Video, bool) {
	var video model.Video
	result := DB.Find(&video, id)

	if result.RowsAffected == 0 {
		return model.Video{}, false
	}

	return video, true
}

func GetVideos() ([]model.Video) {
	var videos []model.Video
	DB.Find(&videos)
	return videos
}