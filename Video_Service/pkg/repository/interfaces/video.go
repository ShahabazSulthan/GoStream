package interfaces

import "video-microservice/pkg/pb"

type VideoRepo interface {
	CreateVideoid(string) error
	FindAllVideo() ([]*pb.VideoID, error)
}
