package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type key struct {
	Type discordgo.InteractionType
	ID string
}

type InteractionManager struct {
	handlers map[key] InteractionCallback
}

func NewInteractionManager() *InteractionManager {
	return &InteractionManager{
		handlers: make(map[key] InteractionCallback, 0),
	}
}

func (m *InteractionManager) AddHandler(t discordgo.InteractionType, id string, c InteractionCallback) error {
	k := key{
		ID: id,
		Type: t,
	}

	_, exists := m.handlers[k]

	if exists {
		return fmt.Errorf("Handler for %v already exists with ID %s", t, id)
	}

	m.handlers[k] = c

	return nil
}

func (m *InteractionManager) AddCommand(c *Command) error {
	return m.AddHandler(discordgo.InteractionApplicationCommand, c.ApplicationCommand.Name, c.Handler)
}

func (m *InteractionManager) AddButton(b *Button) error {
	return m.AddHandler(discordgo.InteractionMessageComponent, b.Button.CustomID, b.Handler)
}

func (m *InteractionManager) AddSelectMenu(s *SelectMenu) error {
	return m.AddHandler(discordgo.InteractionMessageComponent, s.SelectMenu.CustomID, s.Handler)
}

func (m *InteractionManager) AddModal(c *Modal) error {
	return m.AddHandler(discordgo.InteractionModalSubmit, c.CustomID, c.Handler)
}

// TODO: Switch parameters to more useful artifact for bot
func (m *InteractionManager) GetCallback(t discordgo.InteractionType, id string) (InteractionCallback, error) {
	handler, exists := m.handlers[key{ Type: t, ID: id, }]

	if !exists {
		return nil, fmt.Errorf("Could not find handler of type %v with id %s", t, id)
	}

	return handler, nil
}