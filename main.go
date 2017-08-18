package main
import (
	"log"
	"net/http"
	"os"
	"strings"
	"fmt"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/kecbigmt/go-kecy-linebot/automata/oldLulu_008"
)

func main() {
	bot, err := linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			var text string
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					switch {
					case message.Text == "へい":
						text = "ほー"
					case strings.HasPrefix(message.Text, "L8:"):
						t := strings.Trim(message.Text, "L8:")
						b := make([]byte, len(t))
						for i, l := range t {
							switch l{
							case '0':
								b[i] = uint8(0)
							case '1':
								b[i] = uint8(1)
							default:
								b[i] = uint8(255)
							}
						}
						if err := oldLulu_008.Validate(b); err != nil {
							text = fmt.Sprintf("拒否\n%v", err)
						} else {
							text = "受理"
						}
					default:
						text = message.Text
					}
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(text)).Do(); err != nil {
						log.Print(err)
          }
        }
      }
    }
  })
	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
