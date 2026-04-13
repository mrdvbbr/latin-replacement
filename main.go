package main

import (
	"fmt"
	"log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"latin-replacement/transliterate"
)

func main() {
	token := "8786610478:AAFQ-c8lL9uR_NCA510Euze0tDFSokUT8MI"

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal("failed to create bot:", err)
	}

	log.Printf("Authorized as @%s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := update.Message
		if !transliterate.HasCyrillic(msg.Text) {
			continue
		}

		latin := transliterate.Do(msg.Text)
		username := senderName(msg.From)

		// Try to delete the original message (requires Delete Messages admin right).
		del := tgbotapi.NewDeleteMessage(msg.Chat.ID, msg.MessageID)
		if _, err := bot.Request(del); err != nil {
			log.Printf("could not delete message %d: %v", msg.MessageID, err)
		}

		reply := tgbotapi.NewMessage(
			msg.Chat.ID,
			fmt.Sprintf("%s aytmoqchi bo'ldiki: %s\n%s lotincha yoz olipta !", username, latin, username),
		)
		if msg.ReplyToMessage != nil {
			reply.ReplyToMessageID = msg.ReplyToMessage.MessageID
		}
		if _, err := bot.Send(reply); err != nil {
			log.Printf("could not send replacement message: %v", err)
		}
	}
}

// senderName returns "@username" when available, otherwise the display name.
func senderName(u *tgbotapi.User) string {
	if u == nil {
		return "Someone"
	}
	if u.UserName != "" {
		return "@" + u.UserName
	}
	name := u.FirstName
	if u.LastName != "" {
		name += " " + u.LastName
	}
	return name
}
