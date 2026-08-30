package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	discordimpostergo "github.com/njmdev03/discord-imposter-go"
	"github.com/njmdev03/discord-imposter-go/bot"
)

// Bot parameters
var (
	GuildID string
	BotToken string
	DatabaseURL string

	UsePostgresBackend = false
)

var (
	b *bot.Bot
	s *discordgo.Session
	// p *pgx.Conn
	postgres_ctx *bot.PostgresContextManager
	memory_ctx *bot.MemoryContextManager
)

func init() {
	useEnvFlag := flag.Bool("env", true, "Use a .env file to load secrets instead of parameters")
	usePostgresFlag := flag.Bool("pg", false, "Use a postgres db instead of in-memory map for data storage")
	botTokenFlag := flag.String("token", "", "Bot access token")
	guildIDFlag := flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	databaseURLFlag := flag.String("database", "", "The url to the postgres db for data storage")

	flag.Parse()

	UsePostgresBackend = *usePostgresFlag

	if !*useEnvFlag {
		BotToken = *botTokenFlag
		GuildID = *guildIDFlag
		DatabaseURL = *databaseURLFlag
	} else {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		BotToken = os.Getenv("BOT_TOKEN")
		GuildID = os.Getenv("GUILD_ID")
		DatabaseURL = os.Getenv("DATABASE_URL")
	}
}

func init() {
	if UsePostgresBackend {
		conn, err := pgx.Connect(context.Background(), DatabaseURL)

		if err != nil {
			log.Fatalf("Problem connecting to database: %v", err)
		}

		postgres_ctx = bot.NewPostgresContextManager(conn)
	} else {
		memory_ctx = bot.NewMemoryContextManager()
	}
}

func init() {
	var err error
	s, err = discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	s.Identify.Intents = discordgo.IntentGuildVoiceStates | discordgo.IntentsGuilds
}

func init() {
	var e error

	if UsePostgresBackend {
		b, e = bot.NewBot(s, postgres_ctx)
	} else {
		b, e = bot.NewBot(s, memory_ctx)
	}

	if e != nil {
		log.Fatalf("Problem creating bot: %v", e)
	}
}

func main() {
	discordimpostergo.RegisterHandlers(b)

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// log.Printf("Interaction received %v %v", i.Interaction.ID, i.Interaction.Type)

		e := b.HandleInteraction(i)

		if e != nil {
			log.Fatalf("Problem handling interaction: %v", e)
		}
	})

	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	defer signal.Stop(stop)

	log.Println("Press Ctrl+C to exit")

	frequency := 2 * time.Minute
	ticker := time.NewTicker(frequency)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Println("Vacuuming storage")

			if UsePostgresBackend {
				postgres_ctx.Vacuum(frequency)
			} else {
				memory_ctx.Vacuum(frequency)
			}

		case <-stop:
			log.Println("Gracefully shutting down.")
			return
		}
	}
}