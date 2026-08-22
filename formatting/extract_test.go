package formatting

import (
	"testing"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func TestExtractUser(t *testing.T) {
	tests := []struct {
		users []*discordgo.Member
		extract *discordgo.Member
		want []*discordgo.Member
	}{
		{
			[]*discordgo.Member{
				{
					User: &discordgo.User{
						ID: "123",
						Username: "Username",
						Discriminator: "",
						GlobalName: "",
					},
					Nick: "",
				},
			},
			&discordgo.Member{
				User: &discordgo.User{
					ID: "123",
					Username: "Username",
					Discriminator: "",
					GlobalName: "",
				},
				Nick: "",
			},
			[]*discordgo.Member{},
		},
		{
			[]*discordgo.Member{
				{
					User: &discordgo.User{
						ID: "123",
						Username: "Username",
						Discriminator: "",
						GlobalName: "",
					},
					Nick: "",
				},
				{
					User: &discordgo.User{
						ID: "234",
						Username: "Username",
						Discriminator: "",
						GlobalName: "",
					},
					Nick: "",
				},
			},
			&discordgo.Member{
				User: &discordgo.User{
					ID: "123",
					Username: "Username",
					Discriminator: "",
					GlobalName: "",
				},
				Nick: "",
			},
			[]*discordgo.Member{
				{
					User: &discordgo.User{
						ID: "234",
						Username: "Username",
						Discriminator: "",
						GlobalName: "",
					},
					Nick: "",
				},
			},
		},
		{
			[]*discordgo.Member{
				{
					User: &discordgo.User{
						ID: "234",
						Username: "Username",
						Discriminator: "",
						GlobalName: "",
					},
					Nick: "",
				},
			},
			&discordgo.Member{
				User: &discordgo.User{
					ID: "123",
					Username: "Username",
					Discriminator: "",
					GlobalName: "",
				},
				Nick: "",
			},
			[]*discordgo.Member{
				{
					User: &discordgo.User{
						ID: "234",
						Username: "Username",
						Discriminator: "",
						GlobalName: "",
					},
					Nick: "",
				},
			},
		},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("remove %v", test.extract.User.ID)

		t.Run(testname, func(t *testing.T) {
			result := ExtractUser(test.users, test.extract)

			var foundIDs []string

			for _, user := range result {
				if user.User.ID == test.extract.User.ID {
					t.Errorf("Found user that should have been removed %s", user.User.ID)
				} else {
					foundIDs = append(foundIDs, user.User.ID)
				}
			}

			for _, outUser := range test.want {
				found := false

				for _, resUser := range result {
					if outUser.User.ID == resUser.User.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("User missing from output slice %s", outUser.User.ID)
				}
			}
		})
	}
}