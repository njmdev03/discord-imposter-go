package bot

import "github.com/bwmarrin/discordgo"

type InteractionCallback func(*Bot, *discordgo.InteractionCreate) error

type Command struct {
	ApplicationCommand discordgo.ApplicationCommand
	Handler InteractionCallback
}

type Button struct {
	Button discordgo.Button
	Handler InteractionCallback
}

type SelectMenu struct {
	SelectMenu discordgo.SelectMenu
	Handler InteractionCallback
}

type Modal struct {
	CustomID string
	Title string
	Components []discordgo.MessageComponent
	Handler InteractionCallback
}

func (m *Modal) BuildModalResponse() discordgo.InteractionResponseData {
	return discordgo.InteractionResponseData{}
}