package routes

import (
	"net/http"

	"github.com/StackStoreHQ/go-fm/modules"
	"github.com/StackStoreHQ/go-fm/tools"
	"github.com/labstack/echo/v4"
)

func GetMusicCover(ctx echo.Context) error {
	musicInfo := modules.MusicReader.GetMusicInfoStoreData()
	err := ctx.Blob(http.StatusOK, musicInfo.Cover.MimeType, musicInfo.Cover.Data)
	if err != nil {
		modules.Logger.Error(err)
		return err
	}
	return err
}

func GetMusicInfo(ctx echo.Context) error {

	musicInfo := modules.MusicReader.GetMusicInfo()
	err := ctx.JSON(http.StatusOK, tools.Response.GetResponseBody(musicInfo))
	if err != nil {
		modules.Logger.Error(err)
		return err
	}
	return nil
}
