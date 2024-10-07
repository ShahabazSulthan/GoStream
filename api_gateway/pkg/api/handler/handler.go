package handler

import (
	"api-gateway/pkg/client/interfaces"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VideoHandler struct {
	Client interfaces.VideoClient
}

func NewVideoHandler(client interfaces.VideoClient) VideoHandler {
	return VideoHandler{
		Client: client,
	}
}

func (cr *VideoHandler) GetUploadVideo(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", nil)
}

func (cr *VideoHandler) Video(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func (cr *VideoHandler) UploadVideo(c *gin.Context) {
	file, err := c.FormFile("video")
	if err != nil {
		fmt.Println(file, "🔗")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to find the file",
			"error":   err.Error(),
		})
		return
	}
	res, err1 := cr.Client.UploadVideo(c.Request.Context(), file)
	if err1 != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "failed to upload",
			"error":   err1.Error(),
		})
		return
	}
	if res.Status == http.StatusOK {
		c.Redirect(http.StatusSeeOther, "/video")
		return
	}
	c.JSON(int(res.Status), gin.H{
		"Success": res,
	})
}

func (cr *VideoHandler) StreamVideo(c *gin.Context) {
	filename := c.Param("video_id")
	playlist := c.Param("playlist")
	stream, err := cr.Client.StreamVideo(c.Request.Context(), filename, playlist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to stream the video",
			"error":   err.Error(),
		})
		return
	}
	for {
		//reciving the streaming bytes send from the video service
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"message": "error while receiving video chunk:",
				"error":   err.Error(),
			})
			return
		}

		// Process each video chunk received from the server
		c.Header("Content-Type", "application/vnd.apple.mpegurl")
		c.Header("Content-Disposition", "inline")
		c.Writer.Write(resp.VideoChunk)
	}
}

func (cr *VideoHandler) FindAllVideo(c *gin.Context) {
	res, err := cr.Client.FindAllVideo(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "failed to get the videos",
			"error":   err.Error(),
		})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"Video": res.Videos,
	})
}