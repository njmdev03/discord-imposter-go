package interactions

import (
	"github.com/bwmarrin/discordgo"
	"github.com/njmdev03/discord-imposter-go/bot"
	"github.com/njmdev03/discord-imposter-go/formatting"
)

var (
	HelpLines = []string{
		"Use `/roll` while in a voice a call. The bot will dm all users in that call their rolls.",
		"- Use the `imposters` parameter to set the number of imposters (e.g. `1`, `2`, `3`, defaults to `1`). The maximum number of imposters is the number of players in the call.",
		"- Set `include_allies` to `True` and the imposter's dms will include a list of the other imposters they are allied with.",
		"After running `/roll`, you can use `/summary` to see the rolls everyone had. Starting a new game with `/roll` will overwrite this summary with the new game. Summaries are user-specific, so the player who ran `/roll` must be the one to run `/summary`",
		"Use `/invite` to get a new invite link to add the bot to another server",
		"Use `/help` to view this message`",
	}

	HelpText = []string{
		"Available commands are",
		"- /help",
		"- /invite",
		"- /roll [imposters] [include_allies]",
		"- /summary",
		"",
		"Select a command below to view more help",
	}
	RollHelpText = []string{
		"Usage: `/roll [imposters] [include_allies]`",
		"Use `/roll` while in a voice a call. The bot will select rolls for all users in the voice call and dm them those rolls.",
		"Setting the `imposters` parameter decides how many imposters are selected. (default 1)",
		"Setting `include_allies` to true will tell all of the imposters who the other imposters are. (default false)",
		"Use summary after a roll to view all of the players and their roles",
	}
	SummaryHelpText = []string{
		"Usage: `/summary`",
		"Use `/summary` after a roll to view all of the players and their roles.",
		"Only the player who ran `/roll` can use `/summary`.",
		"Players and rolls are tracked on a per-guild basis.",
		"The role summary will be visible to everyone in the channel it is invoked from.",
	}

	InviteText = []string{
		"Add the bot to your own server!",
		"",
		"https://discord.com/oauth2/authorize?client_id=1529308930856194048&permissions=2147485696&integration_type=0&scope=bot+applications.commands",
	}
)

func HandleHelp(b *bot.Bot, i *discordgo.InteractionCreate) error {
	return b.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
			Content: formatting.LinesToString(HelpText),
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						rollHelpButton,
						summaryHelpButton,
						inviteHelpButton,
					},
				},
			},
		},
	})
}

func HandleRollHelp(b *bot.Bot, i *discordgo.InteractionCreate) error {
	return b.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
			Content: formatting.LinesToString(RollHelpText),
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						summaryHelpButton,
					},
				},
			},
		},
	})
}

func HandleSummaryHelp(b *bot.Bot, i *discordgo.InteractionCreate) error {
	return b.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
			Content: formatting.LinesToString(SummaryHelpText),
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						rollHelpButton,
					},
				},
			},
		},
	})
}

func HandleInvite(b *bot.Bot, i *discordgo.InteractionCreate) error {
	return b.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
			Content: formatting.LinesToString(InviteText),
		},
	})
}

func HandleRoll(b *bot.Bot, i *discordgo.InteractionCreate) error {
	return nil
}

func HandleSummary(b *bot.Bot, i *discordgo.InteractionCreate) error {
	return nil
}