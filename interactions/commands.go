package interactions

import (
	"github.com/bwmarrin/discordgo"
	"github.com/njmdev03/discord-imposter-go/bot"
	"github.com/njmdev03/discord-imposter-go/utils"
)

var (
	Commands = []bot.Command{
		{
			ApplicationCommand: discordgo.ApplicationCommand{
				Type: discordgo.ChatApplicationCommand,
				Name:        "help",
				Description: "Show a help message",
			},
			Handler: HandleHelp,
		},
		{
			ApplicationCommand: discordgo.ApplicationCommand{
				Type: discordgo.ChatApplicationCommand,
				Name:        "invite",
				Description: "Add the bot to another server",
			},
			Handler: HandleInvite,
		},
		{
			ApplicationCommand: discordgo.ApplicationCommand{
				Type: discordgo.ChatApplicationCommand,
				Name:        "roll",
				Description: "Randomly select imposters from a voice channel",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionInteger,
						Name:        "imposters",
						Description: "The number of imposters to select",
						MinValue:    utils.Ptr(1.0),
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        "include_allies",
						Description: "Tell imposters who is on their team",
					},
				},
			},
			Handler: HandleRoll,
		},
		{
			ApplicationCommand: discordgo.ApplicationCommand{
				Type: discordgo.ChatApplicationCommand,
				Name:        "summary",
				Description: "Post the rolls from the last game",
			},
			Handler: HandleSummary,
		},
	}
)