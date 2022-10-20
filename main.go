package main

import (
	"BMinder/src/services/botservice"
	"BMinder/src/services/notifier"
	"BMinder/src/services/observer"
	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/robfig/cron/v3"
	"log"
	"os"
)

var conn *pgx.Conn

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	bot, err := botservice.NewBot(os.Getenv("BOT_TOKEN"))
	if err != nil {
		panic(err)
	}

	notifyService := notifier.New()
	notifyService.SetObserver(&observer.TGObserver{Bot: bot})

	c := cron.New()
	_, err = c.AddFunc("@every 30m", func() {
		notifyService.Notify(264979026, "test notify")
	})
	if err != nil {
		panic(err)
	}

	c.Start()

	botservice.ListenCommands(bot)
}

//
//func botRun() *tgbotapi.BotAPI {
//
//	bot, err := tgbotapi.NewBotAPI("311927713:AAE8BOsoajS7TTMU87swuEfkPhmIlBV5_Xo")
//	if err != nil {
//		log.Panic(err)
//	}
//
//	bot.Debug = true
//
//	log.Printf("Authorized on account %s", bot.Self.UserName)
//	return bot
//}
//
//func getTest() string {
//	rows, _ := conn.Query("select id, name from test")
//	var res string
//	for rows.Next() {
//		err := rows.Scan(&id, &name)
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Printf("%d. %s\n", id, name)
//		res = res + "\n" + strconv.Itoa(id) + name
//	}
//	return res
//}
