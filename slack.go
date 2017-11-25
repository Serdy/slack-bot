package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nlopes/slack"
	"github.com/serdy/slack/commands"
	"github.com/spf13/viper"
)

func LoadGlobalConfig(relativeSourcePath, configFilename string) error {
	if relativeSourcePath == "" {
		relativeSourcePath = "."
	}

	viper.SetConfigName(configFilename)

	viper.AddConfigPath(relativeSourcePath)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config file not found...")
		return err
	}

	return nil
}

func main() {
	LoadGlobalConfig("", "config")

	BotID := viper.GetString("BotID")
	SlackToken := viper.GetString("SlackToken")
	// //
	// time.Sleep(30 * time.Second)
	api := slack.New(SlackToken)
	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	api.SetDebug(true)
	slack.SetLogger(logger)

	rtm := api.NewRTM()
	go rtm.ManageConnection()
	// #######################################
Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:

			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				// Ignore hello

			case *slack.ConnectedEvent:
				fmt.Println("Infos:", ev.Info)
				fmt.Println("Connection counter:", ev.ConnectionCount)
				// Replace #general with your Channel ID
				rtm.SendMessage(rtm.NewOutgoingMessage("SSHello world", "#general"))
			case *slack.MessageEvent:
				callerID := ev.Msg.User
				if ev.Msg.Type == "message" && callerID != BotID && ev.Msg.SubType != "message_deleted" &&
					(strings.Contains(ev.Msg.Text, "<@"+BotID+">") || strings.HasPrefix(ev.Msg.Channel, "D")) {
					commands.CatchMessForBot(ev)
				}
				// ############################################
			case *slack.PresenceChangeEvent:
				fmt.Printf("Presence Change: %v\n", ev)

			case *slack.LatencyReport:
				fmt.Printf("Current latency: %v\n", ev.Value)

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop

			default:
				// Ignore other events..
				// fmt.Printf("Unexpected: %v\n", msg.Data)
			}
		}
	}

}
