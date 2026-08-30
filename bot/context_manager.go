package bot

import (
	"errors"
	"time"

	"github.com/bwmarrin/discordgo"
	discordhelpers "github.com/njmdev03/discord-imposter-go/discord-helpers"
	"github.com/njmdev03/discord-imposter-go/game"
)

type ContextManager interface {
	AddGame(g *game.Game, guild_id string, user_id string) error
	RemoveGame(guild_id string, user_id string) error
	CheckGame(guild_id string, user_id string, s *discordgo.Session) (*game.Game, error)
	Vacuum(older_than time.Duration) error
}

var (
    ErrGameAlreadyExists = errors.New("game already exists")
    ErrGameNotFound      = errors.New("game not found")
)

type GuildUser struct {
	GuildID string
	UserID string
}

type GameState struct {
	Imposters []string
	Innocents []string
}

type InternalGameState struct {
	Imposters []string
	Innocents []string
	CreatedAt int64
}

type MemoryContextManager struct {
	Games map[GuildUser] InternalGameState
}

func NewMemoryContextManager() *MemoryContextManager {
	return &MemoryContextManager{
		Games: make(map[GuildUser]InternalGameState),
	}
}

func (cm *MemoryContextManager) AddGame(g *game.Game, guild_id string, user_id string) error {
	gu := GuildUser{
		GuildID: guild_id,
		UserID: user_id,
	}

	_, exists := cm.Games[gu]

	if exists {
		return ErrGameAlreadyExists
	}

	gs := InternalGameState{
		Imposters: discordhelpers.GetIDsFromMembers(g.Imposters),
		Innocents: discordhelpers.GetIDsFromMembers(g.Innocents),
		CreatedAt: time.Now().Unix(),
	}

	cm.Games[gu] = gs

	return nil
}

func (cm *MemoryContextManager) RemoveGame(guild_id string, user_id string) error {
	gu := GuildUser{
		GuildID: guild_id,
		UserID: user_id,
	}

	_, exists := cm.Games[gu]

	if !exists {
		return nil
	}

	delete(cm.Games, gu)

	return nil
}

func (cm *MemoryContextManager) CheckGame(guild_id string, user_id string, s *discordgo.Session) (*game.Game, error) {
	gu := GuildUser{
		GuildID: guild_id,
		UserID: user_id,
	}

	gs, exists := cm.Games[gu]

	if !exists {
		return nil, ErrGameNotFound
	}

	innocents, _ := discordhelpers.GetMembersFromIDs(gs.Innocents, guild_id, s)

	imposters, _ := discordhelpers.GetMembersFromIDs(gs.Imposters, guild_id, s)

	resolved_game := &game.Game{
		Imposters: imposters,
		Innocents: innocents,
	}

	return resolved_game, nil
}

func (cm *MemoryContextManager) Vacuum(older_than time.Duration) error {
	now := time.Now()

	for key, igs := range cm.Games {
		elapsed := now.Sub(time.Unix(igs.CreatedAt, 0))

		if elapsed > older_than {
			delete(cm.Games, key)
		}
	}

	return nil
}