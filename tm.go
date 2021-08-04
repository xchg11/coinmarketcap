package main

import (
	"log"

	"github.com/Syfaro/telegram-bot-api"
)

func sendMessageBot(msg1 string) error {
	var err error
	// подключаемся к боту с помощью токена
	bot, err := tgbotapi.NewBotAPI("277617333:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	if err != nil {
		log.Println(err)
		return err
	}
	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)
	// инициализируем канал, куда будут прилетать обновления от API
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	msg := tgbotapi.NewMessage(2263580, msg1)
	// и отправляем его
	bot.Send(msg)
	return err
}
