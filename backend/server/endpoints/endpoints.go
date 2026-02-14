package endpoints

import (
	"github.com/gofiber/fiber/v2"

	accountState "server/endpoints/account/state"
	diaryEndpoints "server/endpoints/diary"
	photoEntity "server/endpoints/photo/entity"
	trackLog "server/endpoints/track/log"
	"server/loaders/fiber/middlewares"
)

func Init(router fiber.Router) {
	// * Account
	account := router.Group("/account", middlewares.Session)
	account.Get("/state", accountState.StateHandler)

	// * Diary
	diary := router.Group("/diary", middlewares.Session)
	diary.Get("/graph", diaryEndpoints.ActivityGraphHandler)

	// * Photo
	photo := router.Group("/photo", middlewares.Session)
	photo.Get("/entity/album/list", photoEntity.AlbumListHandler)
	photo.Get("/entity/album/detail", photoEntity.AlbumDetailHandler)
	photo.Get("/entity/section/detail", photoEntity.SectionDetailHandler)
	photo.Get("/entity/photo/list", photoEntity.PhotoListHandler)

	// * Tracking
	tracking := router.Group("/track", middlewares.Session)
	tracking.Put("/log/click", trackLog.SectionDetailHandler)
	tracking.Put("/sig/err", trackLog.SigErrorHandler)
}
