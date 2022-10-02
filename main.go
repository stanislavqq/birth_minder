package main

import (
	"BMinder/src/services/botservice"
	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
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

	bot.SendTextToChat(264979026, "Test!")

	botservice.ListenCommands(bot)
	//db, err := sql.Open("sqlite3", "./database.sqlite")

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
