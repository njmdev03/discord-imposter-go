package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/njmdev03/discord-imposter-go/bot"
	"github.com/njmdev03/discord-imposter-go/formatting"
)

type Command struct {
	Name string
	Description string
	Handler func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

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

	}
	RollHelpText = []string{

	}
	SummaryHelpText = []string{

	}


	InviteLines = []string{
		"Add the bot to your own server and use all of the same features!",
		"",
		"https://discord.com/oauth2/authorize?client_id=1529308930856194048&permissions=2147485696&integration_type=0&scope=bot+applications.commands",
	}
)

func HandleHelp(b *bot.Bot, i *discordgo.InteractionCreate) error {
	b.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: formatting.LinesToString(HelpLines),
		},
	})

	return nil
}

func HandleInvite(b *bot.Bot, i *discordgo.InteractionCreate) error {
	b.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: formatting.LinesToString(HelpLines),
		},
	})

	return nil
}

func HandleRoll(b *bot.Bot, i *discordgo.InteractionCreate) error {
	return nil
}

func HandleSummary(b *bot.Bot, i *discordgo.InteractionCreate) error {
	return nil
}