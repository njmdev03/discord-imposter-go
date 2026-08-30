package interactions

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/njmdev03/discord-imposter-go/bot"
	discordhelpers "github.com/njmdev03/discord-imposter-go/discord-helpers"
	"github.com/njmdev03/discord-imposter-go/formatting"
	"github.com/njmdev03/discord-imposter-go/game"
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
	if i.Interaction.Context != discordgo.InteractionContextGuild {
		return b.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags: discordgo.MessageFlagsEphemeral,
				Content: "This command can only be used in a server.",
			},
		})
	}

	e := StartGame(b, i)

	if e != nil {
		err := b.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags: discordgo.MessageFlagsEphemeral,
				Content: "There was a problem starting your game.",
			},
		})

		if err != nil {
			return fmt.Errorf("Problem sending interacation response: '%w' after Problem starting game: '%w'", err, e)
		}

		return fmt.Errorf("Problem starting game: %w", e)
	}

	return b.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
			Content: "Roles sent, GLHF",
		},
	})
}

func StartGame(b *bot.Bot, i *discordgo.InteractionCreate) error {
	existing_game, err := b.ContextManager.CheckGame(i.Interaction.GuildID, i.Interaction.Member.User.ID, b.Session)

	if err == bot.ErrGameNotFound {
		// Good
	} else if err != nil {
		return fmt.Errorf("Problem checking for game: %w", err)
	} else if existing_game != nil {
		err := b.ContextManager.RemoveGame(i.Interaction.GuildID, i.Interaction.Member.User.ID)

		if err != nil {
			return fmt.Errorf("Problem removing existing game: %w", err)
		}
	}

	// Find the invoking user's voice state.
	userVoiceState, err := b.Session.State.VoiceState(i.Interaction.GuildID, i.Interaction.Member.User.ID)

	if err != nil || userVoiceState == nil || userVoiceState.ChannelID == "" {
		// User isn't in a voice channel.
		return b.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags: discordgo.MessageFlagsEphemeral,
				Content: "You are not in a voice channel.",
			},
		})
	}

	// Find everyone in the same voice channel.
	vs, err := discordhelpers.GetUsersInVoice(b.Session, i.Interaction.GuildID, userVoiceState.ChannelID)

	if err != nil {
		return fmt.Errorf("Problem getting voice channel users: %w", err)
	}

	members, err := discordhelpers.VoiceStatesToMembers(b.Session, vs)

	// Create Game
	g, err := game.CreateGame(members, int(discordhelpers.GetIntOption(i.Interaction, "imposters", 1)))

	if err != nil {
		return fmt.Errorf("Problem creating game: %w", err)
	}

	// Save game
	err = b.ContextManager.AddGame(g, i.Interaction.GuildID, i.Interaction.Member.User.ID)

	if err != nil {
		return fmt.Errorf("Problem saving game: %w", err)
	}

	// message players
	g.MessagePlayers(b.Session, discordhelpers.GetBoolOption(i.Interaction, "include_allies", false))

	return nil
}

func HandleSummary(b *bot.Bot, i *discordgo.InteractionCreate) error {
	if i.Interaction.Context != discordgo.InteractionContextGuild {
		return b.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags: discordgo.MessageFlagsEphemeral,
				Content: "This command can only be used in a server.",
			},
		})
	}

	g, e := GetExistingGame(b, i)

	if e != nil {
		err := b.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags: discordgo.MessageFlagsEphemeral,
				Content: "There was a problem getting your game summary",
			},
		})

		if err != nil {
			return fmt.Errorf("Problem sending interaction response: '%w' after Problem getting summary: '%w'", err, e)
		}

		return fmt.Errorf("Problem getting summary: %w", e)
	}

	return b.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
			Content: g.ToString(),
		},
	})
}

func GetExistingGame(b *bot.Bot, i *discordgo.InteractionCreate) (*game.Game, error) {
	existing_game, err := b.ContextManager.CheckGame(i.Interaction.GuildID, i.Interaction.Member.User.ID, b.Session)

	if err == bot.ErrGameNotFound {
		return nil, b.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags: discordgo.MessageFlagsEphemeral,
				Content: "It looks like you haven't started a game in this server yet. Try using /roll",
			},
		})
	} else if err != nil {
		return nil, fmt.Errorf("Problem checking for game: %w", err)
	}

	return existing_game, nil
}