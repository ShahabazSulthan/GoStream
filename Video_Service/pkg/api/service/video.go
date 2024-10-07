package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"video-microservice/pkg/pb"
	"video-microservice/pkg/repository/interfaces"

	"github.com/google/uuid"
)

const storageLocation = "storage"

type VideoServer struct {
	Repo interfaces.VideoRepo
	pb.UnimplementedVideoServiceServer
}

func NewVideoServer(repo interfaces.VideoRepo) pb.VideoServiceServer {
	return &VideoServer{
		Repo: repo,
	}
}

func (c *VideoServer) UploadVideo(stream pb.VideoService_UploadVideoServer) error {
	//making a folder and a file
	fileuid := uuid.New()
	fileName := fileuid.String()
	folderpath := storageLocation + "/" + fileName
	filepath := folderpath + "/" + fileName + ".mp4"
	if err := os.MkdirAll(folderpath, 0755); err != nil {
		return errors.New("failed to create directory")
	}
	newfile, err1 := os.Create(filepath)
	if err1 != nil {
		return errors.New("failed to create file")
	}
	defer newfile.Close()
	//receiving from the streamed bytes from the api_gateway
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if _, err := newfile.Write(chunk.Data); err != nil {
			return err
		}
	}

	chanerr := make(chan error, 2)
	go func() {
		//call to segment the file to hls format using ffmpeg
		err := CreatePlaylistAndSegments(filepath, folderpath)
		chanerr <- err
	}()
	go func() {
		//saving the video id to a database
		err := c.Repo.CreateVideoid(fileName)
		chanerr <- err
	}()
	for i := 1; i <= 2; i++ {
		err := <-chanerr
		if err != nil {
			return err
		}
	}
	// sending a response and closing the sending stream of bytes
	return stream.SendAndClose(&pb.UploadVideoResponse{
		Status:  http.StatusOK,
		Message: "Video successfully uploaded.",
		VideoId: fileName,
	})
}

func (c *VideoServer) StreamVideo(req *pb.StreamVideoRequest, stream pb.VideoService_StreamVideoServer) error {
	chunkSize := 4096 // Set your desired chunk size
	buffer := make([]byte, chunkSize)
	playlistPath := fmt.Sprintf("storage/%s/%s", req.Videoid, req.Playlist)
	plalistfile, _ := os.Open(playlistPath)
	defer plalistfile.Close()
	for {
		n, err := plalistfile.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		// Send the video chunk as a response to the client
		if err := stream.Send(&pb.StreamVideoResponse{
			VideoChunk: buffer[:n],
		}); err != nil {
			return err
		}
	}
	return nil
}

// to find all the video id
func (c *VideoServer) FindAllVideo(ctx context.Context, req *pb.FindAllRequest) (*pb.FindAllResponse, error) {
	res, err := c.Repo.FindAllVideo()
	if err != nil {
		return nil, err
	}
	return &pb.FindAllResponse{
		Status: http.StatusOK,
		Videos: res,
	}, nil
}

// function to segment the video using ffmpeg and storing it as playlist
func CreatePlaylistAndSegments(filePath string, folderPath string) error {
	segmentDuration := 3
	ffmpegCmd := exec.Command(
		"ffmpeg",
		"-i", filePath,
		"-profile:v", "baseline", // baseline profile is compatible with most devices
		"-level", "3.0",
		"-start_number", "0", // start number segments from 0
		"-hls_time", strconv.Itoa(segmentDuration), //duration of each segment in second
		"-hls_list_size", "0", // keep all segments in the playlist
		"-f", "hls",
		fmt.Sprintf("%s/playlist.m3u8", folderPath),
	)
	output, err := ffmpegCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to create HLS: %v \nOutput: %s ", err, string(output))
	}
	return nil
}