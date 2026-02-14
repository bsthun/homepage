package diary

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"server/loaders/config"
	loggerUtils "server/utils/logger"
)

func AddDiary() error {
	strapiDialector := mysql.New(
		mysql.Config{
			DSN:               config.C.MySqlStrapiDsn,
			DefaultStringSize: 255,
		},
	)

	// * Open strapi database connection
	db, err := gorm.Open(strapiDialector, &gorm.Config{
		Logger: nil,
	})
	if err != nil {
		loggerUtils.Log(logrus.Fatal, "UNABLE TO LOAD STRAPI GORM MYSQL DATABASE")
	}

	// * Add diary
	for iter := time.Date(2001, 9, 7, 0, 0, 0, 0, time.UTC); iter.Before(time.Now()); iter = iter.AddDate(0, 0, 1) {
		row := map[string]any{
			"date":               iter.Format("2006-01-02"),
			"summary":            nil,
			"story":              nil,
			"graph_gain":         0,
			"graph_emotional":    0,
			"graph_productivity": 0,
			"task":               "{}",
			"created_at":         time.Now(),
			"updated_at":         time.Now(),
			"created_by_id":      1,
			"updated_by_id":      1,
		}
		if result := db.Table("diaries").Create(row); result.Error != nil {
			logrus.Error(result.Error)
			loggerUtils.Log(logrus.Fatal, "UNABLE TO ADD DIARY")
		}
	}

	return nil
}
