package game

import (
	"fmt"
	"math/rand"

	"github.com/bwmarrin/discordgo"
	discordhelpers "github.com/njmdev03/discord-imposter-go/discord-helpers"
	"github.com/njmdev03/discord-imposter-go/formatting"
)

type Game struct {
	Imposters []*discordgo.Member
	Innocents []*discordgo.Member
}

func (g *Game) ToString() string {
	var imp_tense_pluraltiy string

	if len(g.Imposters) > 1 {
		imp_tense_pluraltiy = "s were"
	} else {
		imp_tense_pluraltiy = " was"
	}

	var ino_tense_pluraltiy string

	if len(g.Innocents) > 1 {
		ino_tense_pluraltiy = "were"
	} else {
		ino_tense_pluraltiy = "was"
	}


	return fmt.Sprintf(
		"The imposter%s %s, and %s %s innocent",
		imp_tense_pluraltiy,
		formatting.GetStringList(formatting.GetDisplayNames(g.Imposters)),
		formatting.GetStringList(formatting.GetDisplayNames(g.Innocents)),
		ino_tense_pluraltiy,
	)
}

func CreateGame(players []*discordgo.Member, imposter_count int) (*Game, error) {
	imps, inno, err := SelectImposters(players, imposter_count)

	if err != nil {
		return nil, err
	}

	return &Game{
		Imposters: imps,
		Innocents: inno,
	}, nil
}

func Cut[T any](s []T, i int) []T {
	return append(s[:i], s[i+1:]...)
}

// func GetIDsFromVoice(members []*discordgo.VoiceState) []string {
// 	var out []string

// 	for _, vs := range members {
// 		out = append(out, vs.Member.User.ID)
// 	}

// 	return out
// }

func SelectImposters[T any](members []T, imposter_count int) ([]T, []T, error) {
	if imposter_count > len(members) {
		return nil, nil, fmt.Errorf("Asked to select to many imposters from members. %v imposters from %v members", imposter_count, len(members))
	}

	var imposters []T

	for i := 0; i < imposter_count; i++ {
		x := rand.Intn(len(members))

		imp := members[x]
		members = Cut(members, x)

		imposters = append(imposters, imp)
	}

	return imposters, members, nil
}

func (g *Game) GetImposterMessageCallbacks(s *discordgo.Session, send_allies bool) []func() error {
	callbacks := make([]func() error, len(g.Imposters) - 1)

	for _, member := range g.Imposters {
		var quant string

		if len(g.Imposters) > 1 {
			quant = "an"
		} else {
			quant = "the"
		}

		var allies string

		if send_allies {
			allies = formatting.GetStringList(formatting.GetDisplayNames(formatting.ExtractUser(g.Imposters, member)))
		}

		msg := fmt.Sprintf("You are %s impasta!%s", quant, allies)

		callbacks = append(callbacks, func() error {
			return discordhelpers.DMUser(s, member.User.ID, msg)
		})
	}

	return callbacks
}

func (g *Game) GetInnocentMessageCallbacks(s *discordgo.Session) []func() error {
	callbacks := make([]func() error, len(g.Innocents) - 1)

	msg := fmt.Sprintf("You are Innocent!")

	for _, member := range g.Innocents {
		callbacks = append(callbacks, func() error {
			return discordhelpers.DMUser(s, member.User.ID, msg)
		})
	}

	return callbacks
}

func (g *Game) MessagePlayers(s *discordgo.Session, send_allies bool) []error {
	imp_callbacks := g.GetImposterMessageCallbacks(s, send_allies)
	inn_callbacks := g.GetInnocentMessageCallbacks(s)

	callbacks := append(imp_callbacks, inn_callbacks...)

	errors := make([]error, 0)

	for _, callback := range callbacks {
		err := callback()

		if err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}