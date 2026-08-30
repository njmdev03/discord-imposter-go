package bot

import (
	"context"
	// "errors"
	"fmt"
	// "os"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/jackc/pgx/v5"
	discordhelpers "github.com/njmdev03/discord-imposter-go/discord-helpers"
	"github.com/njmdev03/discord-imposter-go/game"
)

var (
	insertGameQuery = `
	INSERT INTO Games (guild_id, creator_id, created_at, imposters, innocents)
	VALUES ($1, $2, $3, $4, $5)
	`

	getGameQuery = `
	SELECT imposters, innocents
	FROM Games
	WHERE guild_id = $1 AND creator_id = $2
	`

	deleteGameQuery = `
	DELETE FROM Games
	WHERE guild_id = $1 AND creator_id = $2
	`

	vacuumQuery = `
	DELETE FROM Games
	WHERE created_at < $1
	`
)

type PostgresContextManager struct {
	Conn *pgx.Conn
}

func NewPostgresContextManager(conn *pgx.Conn) *PostgresContextManager {
	return &PostgresContextManager{
		Conn: conn,
	}
}

func (cm *PostgresContextManager) AddGame(g *game.Game, guild_id string, user_id string) error {

	result, err := cm.Conn.Exec(
		context.Background(),
		insertGameQuery, guild_id, user_id, time.Now().Unix(),
		discordhelpers.GetIDsFromMembers(g.Imposters),
		discordhelpers.GetIDsFromMembers(g.Innocents),
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrGameAlreadyExists
	}

	return nil
}

func (cm *PostgresContextManager) RemoveGame(guild_id string, user_id string) error {

	_, err := cm.Conn.Exec(context.Background(), deleteGameQuery, guild_id, user_id)

	if err != nil {
		return fmt.Errorf("Problem Removing game from database: %w", err)
	}

	return nil
}

func (cm *PostgresContextManager) CheckGame(guild_id string, user_id string, s *discordgo.Session) (*game.Game, error) {
	igs := struct {
		Imps []string
		Inno []string
	} {}

	err := cm.Conn.QueryRow(context.Background(), getGameQuery, guild_id, user_id).Scan(&igs.Imps, &igs.Inno)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrGameNotFound
		} else {
			return nil, fmt.Errorf("Problem querying database: %w", err)
		}
	}

	innocents, _ := discordhelpers.GetMembersFromIDs(igs.Inno, guild_id, s)

	imposters, _ := discordhelpers.GetMembersFromIDs(igs.Imps, guild_id, s)

	gs := &game.Game{
		Imposters: imposters,
		Innocents: innocents,
	}

	return gs, nil
}

func (cm *PostgresContextManager) Vacuum(older_than time.Duration) error {
	reference := time.Now().Add(-older_than)

	_, err := cm.Conn.Exec(context.Background(), vacuumQuery, reference.Unix())

	if err != nil {
		return fmt.Errorf("Error vacuuming database: %w", err)
	}

	return nil
}