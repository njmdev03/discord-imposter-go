package formatting

import (
	"fmt"
	"testing"

	"github.com/bwmarrin/discordgo"
)

func TestGetStringList(t *testing.T) {
	tests := []struct {
		input []string
		want string
	}{
		{
			[]string{"string",},
			"string",
		},
		{
			[]string{"string1", "string2"},
			"string1 and string2",
		},
		{
			[]string{"string1", "string2", "string3"},
			"string1, string2, and string3",
		},
		{
			[]string{"string1", "string2", "string3", "string4"},
			"string1, string2, string3, and string4",
		},
		{
			[]string{"string with spaces",},
			"string with spaces",
		},
		{
			[]string{"string 1", "string 2"},
			"string 1 and string 2",
		},
		{
			[]string{"string 1", "string 2", "string 3"},
			"string 1, string 2, and string 3",
		},
		{
			[]string{"string 1", "string 2", "string 3", "string 4"},
			"string 1, string 2, string 3, and string 4",
		},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%v", test.input)

		t.Run(testname, func(t *testing.T) {
			result := GetStringList(test.input)
			if result != test.want {
				t.Errorf("got %s, want %s", result, test.want)
			}
		})
	}
}

func TestGetDisplayName(t *testing.T) {
	tests := []struct {
		member discordgo.Member
		want string
	}{
		{
			discordgo.Member{
				User: &discordgo.User{
					Username: "Username",
					Discriminator: "",
					GlobalName: "",
				},
				Nick: "",
			},
			"Username",
		},
		{
			discordgo.Member{
				User: &discordgo.User{
					Username: "User name",
					Discriminator: "",
					GlobalName: "",
				},
				Nick: "",
			},
			"User name",
		},
		{
			discordgo.Member{
				User: &discordgo.User{
					Username: "Username",
					Discriminator: "",
					GlobalName: "GlobalName",
				},
				Nick: "",
			},
			"GlobalName",
		},
		{
			discordgo.Member{
				User: &discordgo.User{
					Username: "Username",
					Discriminator: "",
					GlobalName: "Global Name",
				},
				Nick: "",
			},
			"Global Name",
		},
		{
			discordgo.Member{
				User: &discordgo.User{
					Username: "Username",
					Discriminator: "",
					GlobalName: "",
				},
				Nick: "Nick",
			},
			"Nick",
		},
		{
			discordgo.Member{
				User: &discordgo.User{
					Username: "Username",
					Discriminator: "",
					GlobalName: "GlobalName",
				},
				Nick: "Nick",
			},
			"Nick",
		},
		{
			discordgo.Member{
				User: &discordgo.User{
					Username: "Username",
					Discriminator: "",
					GlobalName: "GlobalName",
				},
				Nick: "Nick name",
			},
			"Nick name",
		},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%v", test.member)

		t.Run(testname, func(t *testing.T) {
			result := GetDisplayName(&test.member)
			if result != test.want {
				t.Errorf("got %s, want %s", result, test.want)
			}
		})
	}
}

func TestGetDisplayNames(t *testing.T) {
	tests := []struct {
		members []*discordgo.Member
		want []string
	}{
		{
			[]*discordgo.Member{
				{
					User: &discordgo.User{
						Username: "Username",
						Discriminator: "",
						GlobalName: "",
					},
					Nick: "",
				},
			},
			[]string{"Username"},
		},
		{
			[]*discordgo.Member{
				{
					User: &discordgo.User{
						Username: "Username",
						Discriminator: "",
						GlobalName: "",
					},
					Nick: "",
				},
				{
					User: &discordgo.User{
						Username: "Username",
						Discriminator: "",
						GlobalName: "",
					},
					Nick: "",
				},
			},
			[]string{"Username", "Username"},
		},
		{
			[]*discordgo.Member{},
			[]string{},
		},
		{
			[]*discordgo.Member{
				{
					User: &discordgo.User{
						Username: "Username",
						Discriminator: "",
						GlobalName: "",
					},
					Nick: "",
				},
				{
					User: &discordgo.User{
						Username: "Username",
						Discriminator: "",
						GlobalName: "GlobalName",
					},
					Nick: "",
				},
			},
			[]string{"Username", "GlobalName"},
		},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%v", test.members)

		t.Run(testname, func(t *testing.T) {
			results := GetDisplayNames(test.members)
			for i, result := range results {
				if result != test.want[i] {
					t.Errorf("got %s, want %s", result, test.want[i])
				}
			}
		})
	}
}

func TestLineToString(t *testing.T) {
	tests := []struct {
		lines []string
		want string
	}{
		{
			[]string{},
			"",
		},
		{
			[]string{
				"Line1",
			},
			"Line1",
		},
		{
			[]string{
				"Line 1",
			},
			"Line 1",
		},
		{
			[]string{
				"Line 1",
				"Line 2",
			},
			"Line 1\nLine 2",
		},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%v", test.lines)

		t.Run(testname, func(t *testing.T) {
			result := LinesToString(test.lines)
			if result != test.want {
				t.Errorf("got %s, want %s", result, test.want)
			}
		})
	}
}