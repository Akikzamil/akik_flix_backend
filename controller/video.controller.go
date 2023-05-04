package controller

import (
	"akikflix/database"
	"akikflix/model"
	"akikflix/util"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func UploadVideo(c *fiber.Ctx) error {
	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(503).SendString(err.Error())
	}
	filePath, _ := util.GetSixDigitRandomNumber()
	filePath += file.Filename
	videoModel := model.Video{Name: file.Filename, Path: filePath}
	c.SaveFile(file, fmt.Sprintf("./files/%s", filePath))
	errr := processVideo(filePath)
	if errr != nil {
		os.Remove(filePath)
		return c.Status(504).SendString(errr.Error())
	}
	database.CreateVideoModel(&videoModel)
	return c.JSON(videoModel)
}

func processVideo(filePath string) error {
	inputFilePath := "files/"
	inputFilePath += filePath

	// path to the output M3U8 file
	outputFilePath := "files/"
	extension := filepath.Ext(filePath)
	name := filePath[0 : len(filePath)-len(extension)]
	outputFilePath += name + ".m3u8"
	// chunkFilePath := "files/" + name + "%v%03d.ts"

	// ffmpeg -i big.mp4 -b:v 1M -g 60 -hls_time 2 -hls_list_size 0 -hls_segment_size 500000 output.m3u8

	// command to execute ffmpeg
	cmd := exec.Command("ffmpeg", "-i", inputFilePath,"-b:v","1M","-g","60", "-hls_time", "2", "-hls_list_size", "0", "-hls_segment_size", "500000", outputFilePath)

	// run the command and check for errors
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func GetVideo(c *fiber.Ctx) error {
	id := c.Params("id")
	video, isVideoExist := database.GetVideo(id)
	if !isVideoExist {
		return c.Status(404).SendString("Doesn't exist")
	}

	videoPath := getOgVideoPath(video.Path)
	return c.SendString(videoPath)
}

func getOgVideoPath(videoPath string) string {
	extension := filepath.Ext(videoPath)
	name := videoPath[0 : len(videoPath)-len(extension)]
	name += ".m3u8"
	// name = "files/" + name
	return name
}

func GetVideos(c *fiber.Ctx) error{
	videos := database.GetVideos()
	return c.JSON(videos)
}
