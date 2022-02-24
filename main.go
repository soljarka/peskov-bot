package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var c Config
	err := envconfig.Process("peskovbot", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	Start(c)
}

func Start(c Config) {

	phrases := []string{
		"Я не могу давать вам какую-то информацию.",
		"Я воздержусь здесь от каких-то других разъяснений.",
		"Мне неизвестно об этом.",
		"Я сказал вам уже на этот вопрос.",
		"Не могу уточнить.",
		"Но мне об этом ничего неизвестно.",
		"Не буду это никак комментировать.",
		"Не смогу ответить сейчас на этот вопрос.",
		"Вопросы к военным.",
		"Тоже к военным.",
		"Следующий вопрос, пожалуйста.",
		"Это специальная операция, и я не считаю необходимым что-то объяснять.",
		"Дальше. Ребята, на это я отвечал.",
	}

	b, err := tb.NewBot(tb.Settings{
		URL:    c.BotAPIURL,
		Token:  c.BotToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Println("Failed to start the bot.")
	}

	botCommands := []tb.Command{
		{
			Text:        "/peskov",
			Description: "Сказать что-нибудь по-песковски.",
		},
	}

	b.SetCommands(botCommands)

	b.Handle("/peskov", func(m *tb.Message) {
		logMessage(m)

		b.Send(m.Chat, phrases[rand.Intn(len(phrases))])
	})

	b.Start()
}

func logMessage(m *tb.Message) {
	log.Printf("%v sent %s", m.Sender.ID, m.Text)
}

type Config struct {
	BotAPIURL string `default:"https://api.telegram.org"`
	BotToken  string `required:"true"`
}

func init() {
	log.Println("Loading .env")
	err := godotenv.Load()
	if err != nil {
		log.Print("Couldn't find .env file")
	}
}
