package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"log"
	"log/slog"
	"os"
	"tests/internal/config"
	"tests/internal/models"
	"tests/pkg/database"
	"time"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config reading error: %s", err)
	}

	//init database connection
	db, err := database.New(cfg.Storage.DBName, cfg.Storage.Username, cfg.Storage.Password)
	if err != nil {
		slog.Error("database conn failed", "error", err)
		return
	}
	start := time.Now()
	info := getInfoAboutAllEntTests(db)
	end := time.Now()

	wasted := end.Sub(start)
	fmt.Println(wasted)
	injson, err := json.Marshal(info)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.OpenFile("qqq", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()
	_, err = file.Write(injson)
	if err != nil {
		log.Fatal(err)
	}
}

type EntTestInfo struct {
	SubsInfo []SubInfo
}

type SubInfo struct {
	Title      string
	TopicsInfo []TopicInfo
}

type TopicInfo struct {
	Title             string
	TotalCount        int64
	SingleAnswerCount int64
	MultiAnswerCount  int64
}

func getInfoAboutAllEntTests(db *gorm.DB) *EntTestInfo {
	subjectsIDs := []int{188, 189, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199, 200, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217}

	var entTestInfo EntTestInfo
	for _, s := range subjectsIDs {
		var subInfo SubInfo
		var sub models.Subject
		res := db.Where("id = ?", s).First(&sub)
		if res.Error != nil {
			log.Fatal(res.Error)
		}

		var count int64
		c := db.Model(&models.Question{}).Where("subject_id=?", s).Count(&count)
		if c.Error != nil {
			log.Fatal(c.Error)
		}
		subInfo.Title = fmt.Sprintf("%d - %s (%d)", sub.Id, sub.Title, count)

		var tops []models.Topic
		res = db.Where("subject_id=?", s).Find(&tops)
		if res.Error != nil {
			log.Fatal(res.Error)
		}

		for _, t := range tops {
			var topInfo TopicInfo

			topInfo.Title = fmt.Sprintf("%s", t.Title)

			res := db.Model(&models.Question{}).Where("topic_id=?", t.ID).Count(&topInfo.TotalCount)
			if res.Error != nil {
				log.Fatal(res.Error)
			}
			res = db.Model(&models.Question{}).Where("topic_id=? and count_variants=?", t.ID, 4).Count(&topInfo.SingleAnswerCount)
			if res.Error != nil {
				log.Fatal(res.Error)
			}
			res = db.Model(&models.Question{}).Where("topic_id=? and count_variants>?", t.ID, 4).Count(&topInfo.MultiAnswerCount)
			if res.Error != nil {
				log.Fatal(res.Error)
			}
			subInfo.TopicsInfo = append(subInfo.TopicsInfo, topInfo)
		}
		entTestInfo.SubsInfo = append(entTestInfo.SubsInfo, subInfo)
	}
	return &entTestInfo
}
