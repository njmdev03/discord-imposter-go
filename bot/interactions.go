package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type InteractionCallback func(*Bot, *discordgo.InteractionCreate) error

type data struct {
	x int
	y int
}

type inter interface {
	data
}

func test(i *inter) {
	
}

type InteractionHandler interface {
	data
	GetType() discordgo.InteractionType
	GetID() string
	GetCallback() InteractionCallback
}

type Command struct {
	ApplicationCommand discordgo.ApplicationCommand
	Callback InteractionCallback
}

func (m *Command) GetID() string {
	return m.ApplicationCommand.Name
}

func (m *Command) GetType() discordgo.InteractionType {
	return discordgo.InteractionApplicationCommand
}

func (m *Command) GetCallback() InteractionCallback {
	return m.Callback
}

type Component struct {
	MessageComponent discordgo.MessageComponent
	Callback InteractionCallback
}

func (m *Component) GetID() string {
	switch t := m.MessageComponent.(type) {
	case *discordgo.Button:
		return t.CustomID
	case *discordgo.SelectMenu:
		return t.CustomID
	}

	return ""
}

func (m *Component) GetType() discordgo.InteractionType {
	return discordgo.InteractionMessageComponent
}

func (m *Component) GetCallback() InteractionCallback {
	return m.Callback
}

type Modal struct {
	CustomID string
	MessageComponent discordgo.MessageComponent
	Callback InteractionCallback
}

func (m *Modal) GetID() string {
	return m.CustomID
}

func (m *Modal) GetType() discordgo.InteractionType {
	return discordgo.InteractionModalSubmit
}

func (m *Modal) GetCallback() InteractionCallback {
	return m.Callback
}

func (m *Modal) GetInteractionResponseData() *discordgo.InteractionResponseData {
	return nil
}

type InteractionManager struct {
	handlers map[discordgo.InteractionType] map[string] InteractionHandler
}

func NewInteractionManager() *InteractionManager {
	return &InteractionManager{}
}

func (m *InteractionManager) AddHandler(h InteractionHandler) error {
	_, exists := m.handlers[h.GetType()][h.GetID()]

	if exists {
		return fmt.Errorf("Handler for %v already exists with ID %s", h.GetType(), h.GetID())
	}

	m.handlers[h.GetType()][h.GetID()] = h

	return nil
}

func (m *InteractionManager) GetHandler(t discordgo.InteractionType, id string) (InteractionHandler, error) {
	handler, exists := m.handlers[t][id]

	if !exists {
		return nil, fmt.Errorf("Could not find handler of type %v with id %s", t, id)
	}

	return handler, nil
}

func (m *InteractionManager) GetCallback(t discordgo.InteractionType, id string) (InteractionCallback, error) {
	h, e := m.GetHandler(t, id)

	if e != nil {
		return nil, fmt.Errorf("%w", e)
	}

	return h.GetCallback(), nil
}

