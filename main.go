package main

import (
	"BMinder/Bot"
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strconv"
)

var bot *tgbotapi.BotAPI

var (
	id   int
	name string
)

var conn *pgx.Conn

type BDUser struct {
	ID        int32          `db:"id"`
	Name      string         `db:"name"`
	Day       int32          `db:"day"`
	Month     int32          `db:"month"`
	Year      sql.NullInt32  `db:"year"`
	Comment   sql.NullString `db:"comment"`
	AccountID int32          `db:"account_id"`
}

type Account struct {
	ID int32 `db:"id"`
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot := Bot.Create(os.Getenv("BOT_TOKEN"))

	db, err := sql.Open("sqlite3", "./database.sqlite")

	rows, err := db.Query("SELECT * FROM 'birth_users'")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	//var users []BDUser

	for rows.Next() {
		user := BDUser{}
		err := rows.Scan(&user.ID, &user.Name, &user.Day, &user.Month, &user.Year, &user.AccountID, &user.Comment)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(user)
		fmt.Printf("%d. %s\n", user.ID, user.Name)
	}

	rows.Close()
	defer db.Close()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() {
			fmt.Println(update.Message.IsCommand(), update.Message.Text)
			if update.Message.Text == "/start" {
				stmt, err := db.Prepare("INSERT INTO accounts(id) VALUES(?)")
				checkErr(err)
				fmt.Println(update.Message.From.ID)
				res, err := stmt.Exec(update.Message.From.ID)
				checkErr(err)

				id, err := res.LastInsertId()
				checkErr(err)

				fmt.Println(id)
				message := tgbotapi.NewMessage(update.Message.Chat.ID, "Success register!")
				bot.Send(message)
			}
		}
	}

}

func botRun() *tgbotapi.BotAPI {

	bot, err := tgbotapi.NewBotAPI("311927713:AAE8BOsoajS7TTMU87swuEfkPhmIlBV5_Xo")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	return bot
}

func getTest() string {
	rows, _ := conn.Query("select id, name from test")
	var res string
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d. %s\n", id, name)
		res = res + "\n" + strconv.Itoa(id) + name
	}
	return res
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
