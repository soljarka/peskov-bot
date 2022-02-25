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
		"Я сказал вам уже на этот вопрос.",
		"Не смогу ответить сейчас на этот вопрос.",
		"Вопросы к военным.",
		"Тоже к военным.",
		"Следующий вопрос, пожалуйста.",
		"Так ставить вопросы пока преждевременно.",
		"Я оставлю без комментария этот вопрос.",
		"Не понял вопроса.",
		"Это вопрос не к нам.",
		"Мы не имеем информации пока.",
		"Я не могу давать вам какую-то информацию,",
		"Я воздержусь здесь от каких-то других разъяснений.",
		"Мне неизвестно об этом.",
		"Не могу уточнить.",
		"Но мне об этом ничего неизвестно.",
		"Не буду это никак комментировать.",
		"Я не считаю необходимым что-то объяснять.",
		"Дальше. Ребята, на это я отвечал.",
		"Цели были сказаны президентом.",
		"В данном случае пальма первенства и единственным первоисточником здесь должны быть наши военные, наше оборонное ведомство.",
		"Я не думаю, что это может обсуждаться.",
		"Я не стал бы как-то, собственно, комментировать.",
		"Это, собственно, не наша функция.",
		"Это абсолютно не наша функция.",
		"Сейчас как об этом говорить?",
		"Я не располагаю информацией.",
		"Не могу на веру принять ваше утверждение.",
		"Это не та озабоченность, которую мы принимаем к сведению.",
		"Нет, этой темы вообще не затрагивалось.",
		"Президент занимается подготовкой к посланию к Федеральному собранию.",
		"Пока преждевременно говорить об этом.",
		"Будет проводиться анализ.",
		"Стороны пока не планируют перечень тем для обсуждения.",
		"Я ничего не понял из сказанного вами, я буду откровенен.",
		"Такого я просто не могу себе представить.",
		"Вы уверены, что это не фейк?",
		"Не уподобляйтесь уважаемым американским изданиям.",
		"Понятно.",
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
