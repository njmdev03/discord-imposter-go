package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	discordimpostergo "github.com/njmdev03/discord-imposter-go"
	"github.com/njmdev03/discord-imposter-go/bot"
)

// Bot parameters
var (
	GuildID string
	BotToken string
)

var (
	b *bot.Bot
	s *discordgo.Session
)

func init() {
	useEnvFlag := flag.Bool("env", true, "Use a .env file to load secrets instead of parameters")
	botTokenFlag := flag.String("token", "", "Bot access token")
	guildIDFlag := flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")

	flag.Parse()

	if !*useEnvFlag {
		BotToken = *botTokenFlag
		GuildID = *guildIDFlag
	} else {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		BotToken = os.Getenv("BOT_TOKEN")
		GuildID = os.Getenv("GUILD_ID")
	}
}

func init() {
	var err error
	s, err = discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

func init() {
	var e error

	b, e = bot.NewBot(s)

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
		log.Printf("Interaction received %v", i.Interaction)

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
	log.Println("Press Ctrl+C to exit")
	<-stop

	log.Println("Gracefully shutting down.")
}