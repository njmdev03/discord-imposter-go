package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type CommandManager struct {
	Commands map[string]discordgo.ApplicationCommand
}

func NewCommandManager() *CommandManager {
	return &CommandManager{
		Commands: make(map[string]discordgo.ApplicationCommand, 0),
	}
}

func (m *CommandManager) RegisterCommands(s *discordgo.Session) []error {
	var errors []error

	for _, command := range m.Commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", &command)
		if err != nil {
			errors = append(errors, fmt.Errorf("Cannot create '%v' command: %v", command.Name, err))
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (m *CommandManager) AddCommand(c *Command) error {
	_, exists := m.Commands[c.ApplicationCommand.Name]

	if exists {
		return fmt.Errorf("Command already added: %s", c.ApplicationCommand.Name)
	}

	m.Commands[c.ApplicationCommand.Name] = c.ApplicationCommand

	return nil
}