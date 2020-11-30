package main

import (
	"fmt"

	"github.com/raifpy/Go/errHandler"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	const apikey = "1395033625:AAEybfqbSNQb0f6S7l0owkivk5DQolsfYiM"
	bot, err := tgbotapi.NewBotAPI(apikey)
	errHandler.HandlerExit(err)
	//fmt.Println(bot, err)

	bot.Debug = true
	fmt.Println("Bot ; ")
	fmt.Println("ID     : ", bot.Self.ID)
	fmt.Println("FName  : ", bot.Self.FirstName)
	fmt.Println("isBot  : ", bot.Self.IsBot)
	fmt.Println("LanCde : ", bot.Self.LanguageCode)
	fmt.Println("LName  : ", bot.Self.LastName)
	fmt.Println("UName  : ", bot.Self.UserName)
	//fmt.Println(bot.Self.String())

	upd := tgbotapi.NewUpdate(0)

	eleman, err := bot.GetUpdatesChan(upd)
	errHandler.HandlerExit(err)

	for el := range eleman {
		fmt.Println(el.Message.From)
		fmt.Println(el.Message.Chat)
		fmt.Println(el.Message.Command())
		msg := tgbotapi.NewMessage(el.Message.Chat.ID, el.Message.From.UserName+" mal .")
		msg.ParseMode = "markdown"
		bot.Send(msg)

	}

}
