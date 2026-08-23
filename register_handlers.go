package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/njmdev03/discord-imposter-go/bot"
	"github.com/njmdev03/discord-imposter-go/commands"
)

var (
	commandHandlers = []bot.Command{
		{
			ApplicationCommand: discordgo.ApplicationCommand{
				Name: "help",
				Description: "Show a help message",
			},
			Callback: commands.HandleHelp,
		},
		{
			ApplicationCommand: discordgo.ApplicationCommand{
				Name: "invite",
				Description: "Add the bot to another server",
			},
			Callback: commands.HandleInvite,
		},
		{
			ApplicationCommand: discordgo.ApplicationCommand{
				Name: "roll",
				Description: "Randomly select imposters from a voice channel",
				Options: &[]discordgo.ApplicationCommandOption{
					{
						Type: discordgo.ApplicationCommandOptionInteger,
						Name: "imposters",
						Description: "The number of imposters to select",
						MinValue: 1,
					},
					{
						Type: discordgo.ApplicationCommandOptionBoolean,
						Name: "include allies",
						Description: "Tell imposters who is on their team",
					},
				},
			},
			Callback: commands.HandleRoll,
		},
		{
			ApplicationCommand: discordgo.ApplicationCommand{
				Name: "summary",
				Description: "Post the rolls from the last game",
			},
			Callback: commands.HandleSummary,
		},
	}
)

var (
	helpRollComponentHandler = bot.Component{
		
	}
	helpSummaryComponentHandler = bot.Component{

	}
	helpInviteComponentHandler = bot.Component{

	}
)

func RegisterHandlers(b *bot.Bot) {
	b.InteractionManager.AddHandler()
}