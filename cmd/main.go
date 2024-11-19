package main

// библиотеки, необходимые для работы бота
import (
	"fmt"
	"os"

	"TG_simple_bot/message"
	"TG_simple_bot/work"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

// main - функция, в которой происходит инициализация бота и запуск его работы
func main() {

	// получение информации о конфигурации перед запуском бота
	botConfig := work.ConfigReader()

	// токен для авторизации бота в Телеграмм
	// токен получаем из конфигурационного файла
	botToken := botConfig.Token

	// выводит сообщение о работе бота в консоль
	fmt.Println(message.StartMessage(botConfig))

	// запуск и подключение бота по токену
	bot, err := telego.NewBot(botToken)
	if err != nil {
		fmt.Println("Ошибка запуска бота - ", err)
		os.Exit(1)
	}

	// обновление канала связи
	updates, _ := bot.UpdatesViaLongPolling(nil)

	// прекращение обновлений канала при остановке бота
	defer bot.StopLongPolling()

	// выводит в консоль сообщение о том, что бот успешно запущен и работает
	fmt.Println(message.OkMessage())

	firstMessage := false

	// проходим подряд все обновления
	for update := range updates {
		// формируем и отправляем один раз сообщение перед началом диалога с ботом
		// если такое сообщение ещё не было отправлено
		if firstMessage {
			chatID := tu.ID(update.Message.Chat.ID)

			// текст для сообщения возвращает функция message.FirstMessage(botConfig)
			// в которую передаются данные конфигурации
			_, _ = bot.SendMessage(
				tu.Message(chatID, message.FirstMessage(botConfig)))

			// выставляем отметку о том, что первое сообщение уже случилось
			firstMessage = false
		}

		// если в обновлении есть не пустое сообщение, то идём дальше
		if update.Message != nil {
			// получаем ID чата (с каждым пользователем свой чат)
			chatID := tu.ID(update.Message.Chat.ID)

			// myMessage - переменная для хранения текста из сообщения
			myMessage := tu.TextMessage(update.Message.Text)
			// messageString - переменная для ответного сообщения. Сейчас пустая.
			messageString := ""

			switch myMessage.MessageText {
			// если во входящем сообщении "1", то сохраняем в ответное сообщение ссылку на сайт школы
			case "привет покажи ссылку школы":
				messageString = "https://sch2045zg.mskobr.ru/"
			// 	если во входящем сообщении "2", то сохраняем в ответное сообщение текст "Привет Ученик!"
			case "что ты знаешь":
				messageString = "Мы спортсмены а не обарегены аааааа !"
			// во всех остальных случаях возвращаем текст "Для ссылки на сайт школы введи 1.\nДля приветствия 2."
			case "давай поиграем":
				messageString = "@gamee"
			case "напиши : что за команды":
				messageString = "1- привет покажи ссылку школы . 2- что ты знаешь . 3- давай поиграем . 4- кто тебя создал . 5- что они любят .6- мини правила . 7-лучшие друзья . 8- информация о создателях."
			case "кто тебя создал":
				messageString = "Паша капибарка и его помошник Миша"
			case "что они любят":
				messageString = "Паша любит: клинок рассекающий демонов , пабг , майнкрафт и мишу. Миша любит спасть , есть и когда его гладят и любят"
			case "мини правила":
				messageString = "те цифры писать не надо ,  точки и тера писать не надо "
			case "лучшие друзья":
				messageString = "Даша капибарка , янык фуфлик , козин вася"
			case ".":
				messageString = "напиши : что за команды команды"
			case "start":
				messageString = "напиши : что за команды"
			case "информация о создателях":
				messageString = "у Паши капибарки : день рождение 20 ноября родился 2011 ,любит капибарак , имеет рутуб акаунт и тик ток акаунт, скорпион, ходит в медиа . Миша : день рождение 3 ноября , молтеская болонка , скорпион , любимый песель для Паши ."
			case "Start":
				messageString = "что за команды"
			default:

				messageString = "Для ссылки на сайт школы введи 1.\nДля приветствия введи 2."
			}

			// отправляем ответное сообщение в канал
			_, _ = bot.SendMessage(
				tu.Message(chatID, messageString))
		}
	}
}
