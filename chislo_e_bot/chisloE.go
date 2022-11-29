package main

import (
	"fmt"
	"math"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func сколькоРазФакториал(ввод int) (вывод int, e float64) {
	степень := 1.0
	for i := 0; i < ввод; i++ {
		степень = степень * (1.0 / 10.0)
	}
	сумма, i, факториал, флаг := 1.0, 1, 1, false
	for !флаг {
		факториал = факториал * i
		сумма = сумма + 1/float64(факториал)
		i++
		if math.E-сумма < float64(степень) {
			флаг = true
			e = сумма
		}
	}
	return i - 1, e
}
func доКакогоЗнака(ввод int) (вывод int, e float64) {
	степень := 1.0
	сумма, флаг, факториал := 1.0, false, 1
	fmt.Println("начинаю счёт")
	for i := 1; i < ввод; i++ {
		fmt.Println("новый цикл")
		факториал = факториал * i
		сумма = сумма + 1/float64(факториал)
		fmt.Println(сумма)
	}
	fmt.Println("заканчиваю счёт")
	for i := 0; !флаг; i++ {
		if math.E-сумма > float64(степень) {
			флаг = true
			e = сумма
			вывод = i
		}
		степень = степень * (1.0 / 10.0)
	}
	fmt.Println("заканчиваю функцию")
	return вывод, e
}

func main() {
	userMap := make(map[int64]bool)
	bot, err := tgbotapi.NewBotAPI(BOT_ID)
	if err != nil {
		panic(err)
	}
	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		var messageContent string
		if update.Message == nil {
			continue
		}
		integer, err := strconv.Atoi(update.Message.Text)
		if err != nil {
			if update.Message.Text == "/mode" {
				fmt.Println("смена режима")
				_, ok := userMap[update.Message.From.ID]
				if !userMap[update.Message.From.ID] || !ok {
					userMap[update.Message.From.ID] = true
					messageContent = "переключаю на второй режим"
				} else {
					userMap[update.Message.From.ID] = false
					messageContent = "переключаю на первый режим"
				}
			} else {
				messageContent = "введите натуральное число или /mode для смены режима"
			}
		} else {
			if !userMap[update.Message.From.ID] {
				fmt.Println("активирован первый режим")
				if integer < 16 {
					сколькоРаз, числоЕ := сколькоРазФакториал(integer)
					messageContent = fmt.Sprintf("число циклов: %v, высчитанное значение константы Эйлера = %v", сколькоРаз, числоЕ)
				} else {
					messageContent = "слишком большое число"
				}
			} else {
				if integer < 18 {
					fmt.Println("активирован второй режим")
					доКакогоЗнака, числоЕ := доКакогоЗнака(integer)
					messageContent = fmt.Sprintf("константа высчитана до %v знака после запятой, высчитанное значение константы Эйлера = %v", доКакогоЗнака, числоЕ)
				} else {
					messageContent = "слишком большое число"
				}
			}
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, messageContent)
		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}
}
