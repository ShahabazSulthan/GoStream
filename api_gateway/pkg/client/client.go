package client

import (
	"api-gateway/pkg/client/interfaces"
	"api-gateway/pkg/config"
	"api-gateway/pkg/pb"
	"context"
	"fmt"
	"io"
	"mime/multipart"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type videoClient struct {
	Server pb.VideoServiceClient
}

func InitClient(c *config.Config) (pb.VideoServiceClient, error) {
	cc, err := grpc.Dial(c.VideoService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return pb.NewVideoServiceClient(cc), nil
}

func NewVideoClient(server pb.VideoServiceClient) interfaces.VideoClient {
	return &videoClient{
		Server: server,
	}
}

func (c *videoClient) UploadVideo(ctx context.Context, file *multipart.FileHeader) (*pb.UploadVideoResponse, error) {
	upLoadfile, err := file.Open()
	if err != nil {
		fmt.Println("err-1")
		return nil, err
	}
	defer upLoadfile.Close()
	stream, err := c.Server.UploadVideo(ctx)
	if err != nil {
		fmt.Println("err-2")
		return nil, err
	}
	chunkSize := 4096
	buffer := make([]byte, chunkSize)
	for {
		n, err := upLoadfile.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if err := stream.Send(&pb.UploadVideoRequest{
			Filename: file.Filename,
			Data:     buffer[:n],
		}); err != nil {
			return nil, err
		}
	}
	//the final response is recieved and the streaming is closed
	response, err := stream.CloseAndRecv()
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *videoClient) StreamVideo(ctx context.Context, filename, playlist string) (pb.VideoService_StreamVideoClient, error) {
	res, err := c.Server.StreamVideo(ctx, &pb.StreamVideoRequest{
		Videoid:  filename,
		Playlist: playlist,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *videoClient) FindAllVideo(ctx context.Context) (*pb.FindAllResponse, error) {
	res, err := c.Server.FindAllVideo(ctx, &pb.FindAllRequest{})
	if err != nil {
		return nil, err
	}
	return res, nil
}
