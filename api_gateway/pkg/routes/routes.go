package routes

import (
	"api-gateway/pkg/api/handler"

	"github.com/gin-gonic/gin"
)

func RegisterVideoRoutes(api *gin.RouterGroup, videoHandler handler.VideoHandler) {
	api.GET("/upload", videoHandler.GetUploadVideo)
	api.GET("/video", videoHandler.Video)
	api.POST("/upload", videoHandler.UploadVideo)
	api.GET("/stream/:video_id/:playlist", videoHandler.StreamVideo)
	api.GET("/video/all", videoHandler.FindAllVideo)
}