package discordhelpers

import "github.com/bwmarrin/discordgo"

func GetUsersInVoice(s *discordgo.Session, g_id string, c_id string) ([]*discordgo.VoiceState, error) {
	var members []*discordgo.VoiceState

	guild, err := s.State.Guild(g_id)

	if err != nil {
		// handle error
		return nil, err
	}

	for _, vs := range guild.VoiceStates {
		if vs.ChannelID != c_id {
			continue
		}

		members = append(members, vs)
	}

	return members, nil
}

func VoiceStateToMember(s *discordgo.Session, vs *discordgo.VoiceState) (*discordgo.Member, error) {
	member, err := s.GuildMember(vs.GuildID, vs.UserID)

	if err != nil {
		return nil, err
	}

	return member, nil
}

func VoiceStatesToMembers(s *discordgo.Session, vs []*discordgo.VoiceState) ([]*discordgo.Member, error) {
	var members []*discordgo.Member

	for _, member_vs := range vs {
		member, err := VoiceStateToMember(s, member_vs)

		if err != nil {
			return nil, err
		}

		members = append(members, member)
	}

	return members, nil
}

// TODO: Make these calls return safely handle-able errors instead of using discordgo's helpers which can panic

func GetIntOption(i *discordgo.Interaction, option_name string, def int64) int64 {
	imposters_option := i.ApplicationCommandData().GetOption(option_name)

	if imposters_option == nil {
		return def
	} else {
		return imposters_option.IntValue()
	}
}

func GetBoolOption(i *discordgo.Interaction, option_name string, def bool) bool {
	imposters_option := i.ApplicationCommandData().GetOption(option_name)

	if imposters_option == nil {
		return def
	} else {
		return imposters_option.BoolValue()
	}
}

func GetStringOption(i *discordgo.Interaction, option_name string, def string) string {
	imposters_option := i.ApplicationCommandData().GetOption(option_name)

	if imposters_option == nil {
		return def
	} else {
		return imposters_option.StringValue()
	}
}

func GetFloatOption(i *discordgo.Interaction, option_name string, def float64) float64 {
	imposters_option := i.ApplicationCommandData().GetOption(option_name)

	if imposters_option == nil {
		return def
	} else {
		return imposters_option.FloatValue()
	}
}

func DMUser(s *discordgo.Session, user_id string, message string) error {
	// We create the private channel with the user who sent the message.
	channel, err := s.UserChannelCreate(user_id)
	if err != nil {
		// If an error occurred, we failed to create the channel.
		//
		// Some common causes are:
		// 1. We don't share a server with the user (not possible here).
		// 2. We opened enough DM channels quickly enough for Discord to
		//    label us as abusing the endpoint, blocking us from opening
		//    new ones.
		return err
	}

	// Then we send the message through the channel we created.
	_, err = s.ChannelMessageSend(channel.ID, message)
	if err != nil {
		// If an error occurred, we failed to send the message.
		//
		// It may occur either when we do not share a server with the
		// user (highly unlikely as we just received a message) or
		// the user disabled DM in their settings (more likely).
		return err
	}

	return nil
}

func GetIDsFromMembers(members []*discordgo.Member) []string {
	ids := make([]string, len(members))

	for _, member := range members {
		ids = append(ids, member.User.ID)
	}

	return ids
}

func GetMembersFromIDs(ids []string, guild_id string, s *discordgo.Session) ([]*discordgo.Member, error) {
	members := make([]*discordgo.Member, len(ids))

	for _, id := range ids {
		member, err := s.GuildMember(guild_id, id)

		if err != nil {
			return nil, err
		}

		members = append(members, member)
	}

	return members, nil
}