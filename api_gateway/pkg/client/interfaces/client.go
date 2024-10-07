package interfaces

import (
	"api-gateway/pkg/pb"
	"context"
	"mime/multipart"
)

type VideoClient interface {
	UploadVideo(context.Context, *multipart.FileHeader) (*pb.UploadVideoResponse, error)
	StreamVideo(context.Context, string, string) (pb.VideoService_StreamVideoClient, error)
	FindAllVideo(context.Context) (*pb.FindAllResponse, error)
}