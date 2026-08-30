package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	Session *discordgo.Session

	InteractionManager *InteractionManager
	CommandManager *CommandManager
	ContextManager ContextManager
}

func (b *Bot) AddCommand(c *Command) error {
	e := b.CommandManager.AddCommand(c)

	if e != nil {
		return e
	}

	e = b.InteractionManager.AddCommand(c)

	if e != nil {
		return e
	}

	return nil
}

func (b *Bot) AddButton(c *Button) error {
	e := b.InteractionManager.AddButton(c)

	if e != nil {
		return e
	}

	return nil
}

func (b *Bot) AddSelectMenu(c *SelectMenu) error {
	e := b.InteractionManager.AddSelectMenu(c)

	if e != nil {
		return e
	}

	return nil
}

func (b *Bot) AddModal(m *Modal) error {
	e := b.InteractionManager.AddModal(m)

	if e != nil {
		return e
	}

	return nil
}

func (b *Bot) HandleInteraction(i *discordgo.InteractionCreate) error {
	var name string

	switch i.Interaction.Type {
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

	if e != nil || callback == nil {
		return fmt.Errorf("Error handling interaction %w", e)
	}

	e = callback(b, i)

	if e != nil {
		return fmt.Errorf("Problem in interaction handler for %v with name %s: %w", i.Type, name, e)
	}

	return nil
}

func NewBot(s *discordgo.Session, c ContextManager) (*Bot, error) {
	b := &Bot{
		Session: s,
		InteractionManager: NewInteractionManager(),
		CommandManager: NewCommandManager(),
		ContextManager: c,
	}

	return b, nil
}