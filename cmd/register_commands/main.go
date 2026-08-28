package main

import (
	"flag"
	"log"
	"os"

	// "os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	discordimpostergo "github.com/njmdev03/discord-imposter-go"
	"github.com/njmdev03/discord-imposter-go/bot"
)

// Bot parameters
var (
	GuildID        string
	BotToken       string

	RemoveCommands = flag.Bool("rmcmd", false, "Remove all registered commands")
	ListCommands   = flag.Bool("list", false, "List all commands currently registered with Discord")
)

var (
	b *bot.Bot
	s *discordgo.Session
)

func init() {
	useEnvFlag := flag.Bool("env", true, "Use a .env file to load secrets instead of parameters")

	guildIDFlag        := flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	botTokenFlag       := flag.String("token", "", "Bot access token")

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

	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	defer s.Close()

	if *RemoveCommands {
		commands := GetCommands()

		log.Printf("Found %v commands to remove", len(commands))

		for i, command := range commands {
			e := s.ApplicationCommandDelete(s.State.User.ID, "", command.ID)

			if e != nil {
				log.Fatalf("[%v/%v] Failed to remove command %s", i + 1, len(commands), command.Name)
			} else {
				log.Printf("[%v/%v] Removed command %s", i + 1, len(commands), command.Name)
			}
		}
	} else if *ListCommands {
		commands := GetCommands()

		log.Printf("Found %v commands", len(commands))

		for _, command := range commands {
			log.Printf("  /%s: %s", command.Name, command.Description)
		}

		log.Println()
	} else {
		RegisterCommands()
	}

	// stop := make(chan os.Signal, 1)
	// signal.Notify(stop, os.Interrupt)
	// log.Println("Press Ctrl+C to exit")
	// <-stop

	// log.Println("Gracefully shutting down.")
}

func RegisterCommands() {
	e := b.CommandManager.RegisterCommands(s)

	if e != nil {
		for _, e := range e {
			log.Printf("Error registering command: %v", e)
		}

		log.Println()
	}

	log.Printf("%v/%v commands registered", len(b.CommandManager.Commands) - len(e), len(b.CommandManager.Commands))
}

func GetCommands() []*discordgo.ApplicationCommand {
	commands, e := s.ApplicationCommands(s.State.User.ID, "")

	if e != nil {
		log.Fatalf("Problem getting application commands: %v", e)
	}

	return commands
}