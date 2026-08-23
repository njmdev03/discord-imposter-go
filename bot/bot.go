package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	Session *discordgo.Session

	InteractionManager *InteractionManager
}

func (b *Bot) HandleInteraction(i *discordgo.InteractionCreate) error {
	var name string

	switch i.Type {
	case discordgo.InteractionModalSubmit:
		name = i.ModalSubmitData().CustomID
	case discordgo.InteractionMessageComponent:
		name = i.MessageComponentData().CustomID
	case discordgo.InteractionApplicationCommand:
		name = i.ApplicationCommandData().Name
	case discordgo.InteractionApplicationCommandAutocomplete:
		name = i.ApplicationCommandData().Name
	case discordgo.InteractionPing:
		return fmt.Errorf("Ping interactions not handled")
	}

	callback, e := b.InteractionManager.GetCallback(i.Type, name)

	if e != nil {
		return fmt.Errorf("Error handling interaction %w", e)
	}

	e = callback(b, i)

	if e != nil {
		return fmt.Errorf("Problem in interaction handler for %v with name %s: %w", i.Type, name, e)
	}

	return nil
}

func NewBot(s *discordgo.Session, im *InteractionManager) *Bot {
	return &Bot{
		Session: s,
		InteractionManager: im,
	}
}