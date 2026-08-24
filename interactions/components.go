package interactions

import (
	"github.com/bwmarrin/discordgo"
	"github.com/njmdev03/discord-imposter-go/bot"
)

var (
	rollHelpButton = discordgo.Button{
		Label: "/roll help",
		Style: discordgo.PrimaryButton,
		Emoji: &discordgo.ComponentEmoji{
			Name: "🎲",
		},
		CustomID: "roll_help_button",
	}
	summaryHelpButton = discordgo.Button{
		Label: "/summary help",
		Style: discordgo.PrimaryButton,
		Emoji: &discordgo.ComponentEmoji{
			Name: "📒",
		},
		CustomID: "summary_help_button",
	}
	inviteHelpButton = discordgo.Button{
		Label: "/invite help",
		Style: discordgo.PrimaryButton,
		Emoji: &discordgo.ComponentEmoji{
			Name: "➕",
		},
		CustomID: "invite_help_button",
	}

	Buttons = []*bot.Button{
		{
			Button: rollHelpButton,
			Handler: HandleRollHelp,
		},
		{
			Button: summaryHelpButton,
			Handler: HandleSummaryHelp,
		},
		{
			Button: inviteHelpButton,
			Handler: HandleInvite,
		},
	}
)


