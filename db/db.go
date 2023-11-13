package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"dingchatgpt/config"
	"time"
)

type GptInfo struct {
	Timestamp      int    `gorm:"column:timestamp;"`
	AskUser        string `gorm:"column:ask_user;"`
	AskBotQuestion string `gorm:"column:ask_bot_question;"`
	BotAnswer      string `gorm:"column:bot_answer;"`
	UseToken       int    `gorm:"column:use_token;"`
}

func (c *GptInfo) TableName() string {
	return "gpt_ask"
}

type AskGptDB struct {
}

func GptDBModel() *AskGptDB {
	return &AskGptDB{}
}

var db *gorm.DB

func Connect() {
	var err error
	//config.InitFlat()
	//flag.Parse()
	appConfig := config.GetConfig()


	//dsn := "root:123456docker/@tcp(192.168.1.88:13806)/chatgpt?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", appConfig.DBUser, appConfig.DBPasswd, appConfig.DBIp, appConfig.DBPort, appConfig.DBName)
	fmt.Println(dsn)
	//dsn := "root:My:S3cr3t/@tcp(127.0.0.1:30306)/chatgpt?charset=utf8mb4&parseTime=True&loc=Local"



	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
}

// 插入提问人，提问问题，机器人回答，提问时间戳，使用tokens
func (a *AskGptDB) InsertMessage(askUser, askQuestion, botAnswer string, timestamp, useTokens int) error {
	err := db.Select("timestamp", "ask_user", "ask_bot_question", "bot_answer", "use_token").Create(&GptInfo{
		Timestamp:      timestamp,
		AskUser:        askUser,
		AskBotQuestion: askQuestion,
		BotAnswer:      botAnswer,
		UseToken:       useTokens,
	}).Error
	return err
}

// 查询半个小时内提问人的最后问的5问题，最后按照时间戳升序排列
func (a *AskGptDB) SelectLast5Ask(userName string) ([]GptInfo, error) {

	var gptMessage []GptInfo

	selectTime := int(time.Now().Unix()) - 30*60
	//subQuery := db.Model(&GptInfo{}).Where("timestamp >= ? and ask_user = ?", selectTime, userName).Order("timestamp desc").Limit(5)
	subQuery := db.Where("timestamp >= ? and ask_user = ?", selectTime, userName).Order("timestamp desc").Limit(5).Find(&GptInfo{})
	err := db.Table("(?) as t", subQuery).Order("timestamp ASC").Find(&gptMessage).Error
	//db.Table("(?) as t", subQuery).Order("timestamp ASC").Find(&gptMessage).Error.Error()
	//fmt.Println(gptMessage)
	return gptMessage, err
}
