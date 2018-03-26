package commands

import (
	"fmt"
	"strings"

	"github.com/nlopes/slack"
	"github.com/spf13/viper"
)

/*
Remove from Slack message BoarD and space
*/
func PrepareSlackMessage(message string) string {
	BotID := strings.ToLower(viper.GetString("BotID"))
	RemoveIDSpace := strings.NewReplacer("<@"+BotID+">", "", " ", "")
	return RemoveIDSpace.Replace(strings.ToLower(message))
}

func CheckMessageTalks(message string) bool {
	answer := viper.GetStringMap("BotQA")
	println(answer)
	if val, ok := answer[message]; ok {
		fmt.Println(val)
		return true
	}
	return false
}
func CheckMessageJenkinsBuild(message string) bool {
	message = strings.ToLower(message)
	answer := viper.GetStringMap("JenkinsJob")
	if val, ok := answer[message]; ok {
		fmt.Println(val)
		return true
	}
	return false
}
func CheckMessageEnvGit(message string) bool {
	message = strings.ToLower(message)
	answer := viper.GetStringMap("EnvGit")
	if val, ok := answer[message]; ok {
		fmt.Println(val)
		return true
	}
	return false
}

func CheckMessageInstAWS(message string) bool {
	message = strings.ToLower(message)
	answer := viper.GetStringMap("EnvAWS")
	if val, ok := answer[message]; ok {
		fmt.Println(val)
		return true
	}
	return false
}

/*
Func catch messege for bot and try undesten what it must do
*/
func CatchMessForBot(SlackEvent *slack.MessageEvent) {
	SlackMessage := PrepareSlackMessage(SlackEvent.Msg.Text)
	fmt.Println(SlackMessage)
	switch {

	case CheckMessageTalks(SlackMessage):
		BotMsg(SlackEvent.Msg.User, SlackMessage, SlackEvent.Msg.Channel)
		// SendMassage(SlackEvent.Msg.User, SlackMessage)
	// case strings.Contains(strings.ToLower(SlackEvent.Msg.Text), "latest qa"):
	case CheckMessageEnvGit(SlackMessage):
		GetEnvGit(SlackEvent.Msg.User, SlackMessage, SlackEvent.Msg.Channel)
	
	case CheckMessageInstAWS(SlackMessage):
		AWSIntsansesFilter(SlackEvent.Msg.User, SlackMessage, SlackEvent.Msg.Channel)


	default:
		SendMassageHelp(SlackEvent.Msg.User, SlackEvent.Msg.Channel)
	}
}
