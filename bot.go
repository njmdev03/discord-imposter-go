package discordimpostergo

import (
	"log"

	"github.com/njmdev03/discord-imposter-go/bot"
	"github.com/njmdev03/discord-imposter-go/interactions"
)

func RegisterHandlers(b *bot.Bot) {
	var e error

	for _, command := range interactions.Commands {
		e = b.AddCommand(&command)

		if e != nil {
			log.Fatalf("Error registering command %s: %v", command.ApplicationCommand.Name, e)
		}
	}

	for _, button := range interactions.Buttons {
		e = b.AddButton(button)

		if e != nil {
			log.Fatalf("Error registering button %s: %v", button.Button.CustomID, e)
		}
	}
}
