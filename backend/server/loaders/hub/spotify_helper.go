package hub

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"share/types/model"
	"share/utils/text"
	"share/utils/value"

	"server/loaders/mysql"

	"share/types/enum"
)

func spotifyRecord(entity enum.SpotifyEntity, val string) {
	date := text.GetDateString()
	if result := mysql.DB.Clauses(clause.OnConflict{
		DoUpdates: clause.Assignments(
			map[string]any{
				"count": gorm.Expr("count + 1"),
			},
		),
	}).Create(&model.DiarySpotifyAnalytics{
		Date:   &date,
		Entity: &entity,
		Value:  &val,
		Count:  value.Ptr[uint64](1),
	}); result.Error != nil {
		logrus.Warnf("Unable to update spotify analytics (%s): %v", entity, result.Error)
	}
}
