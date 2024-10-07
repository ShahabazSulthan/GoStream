package repository

import (
	"video-microservice/pkg/domain"
	"video-microservice/pkg/pb"
	"video-microservice/pkg/repository/interfaces"

	"gorm.io/gorm"
)

type videoRepo struct {
	DB *gorm.DB
}

func NewVideoRepo(db *gorm.DB) interfaces.VideoRepo {
	return &videoRepo{
		DB: db,
	}
}

func (c *videoRepo) CreateVideoid(videoid string) error {
	if err := c.DB.Create(&domain.Video{VideoId: videoid}).Error; err != nil {
		return err
	}
	return nil
}

func (c *videoRepo) FindAllVideo() ([]*pb.VideoID, error) {
	var videoid []*pb.VideoID
	if err := c.DB.Model(&domain.Video{}).Find(&videoid).Error; err != nil {
		return nil, err
	}
	return videoid, nil
}