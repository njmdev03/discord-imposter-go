package formatting

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func GetDisplayName(m *discordgo.Member) string {
	if m.Nick != "" {
		return m.Nick
	} else if m.User.GlobalName != "" {
		return m.User.GlobalName
	} else {
		return m.User.Username
	}
}

func ExtractUser(users []*discordgo.Member, extract *discordgo.Member) []*discordgo.Member {
	return nil
}

func GetDisplayNamesAppend(users []*discordgo.Member) []string {
	var displayNames []string

	for _, user := range users {
		displayNames = append(displayNames, GetDisplayName(user))
	}

	return displayNames
}

func GetDisplayNamesPreallocated(users []*discordgo.Member) []string {
	displayNames := make([]string, len(users))

	for i, user := range users {
		displayNames[i] = GetDisplayName(user)
	}

	return displayNames
}

func GetDisplayNames(users []*discordgo.Member) []string {
	return GetDisplayNamesPreallocated(users)
}

func GetStringList(strings []string) string {
	var out string

	if len(strings) <= 1 {
		out = strings[0]
	} else if len(strings) <= 2 {
		out = fmt.Sprintf("%s and %s", strings[0], strings[1])
	} else {
		for _, str := range strings[:len(strings)-1] {
			out += fmt.Sprintf("%s, ", str)
		}

		out += fmt.Sprintf("and %s", strings[len(strings)-1])
	}

	return out
}