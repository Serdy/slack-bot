package commands

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/nlopes/slack"
	"github.com/spf13/viper"
)

func BotMsg(NameUser string, MsgQuestion string) {
	BotMassege := viper.GetString("BotQA." + MsgQuestion)
	SendMassage(NameUser, BotMassege)
}
func SendMassage(name string, MsgQuestion string) {
	SlackToken := viper.GetString("SlackToken")
	BotName := viper.GetString("BotName")
	BotIconURL := viper.GetString("BotIconURL")
	params := slack.PostMessageParameters{}
	api := slack.New(SlackToken)
	fmt.Println(reflect.TypeOf(api))
	params.Username = BotName
	params.IconURL = BotIconURL
	fmt.Println(MsgQuestion)
	// time.Sleep(30 * time.Second)
	api.PostMessage("epam", "<@"+name+">: "+MsgQuestion, params)

}
func SendMassageEnvGit(name string, MsgQuestion *EnvGitMap) {
	SlackToken := viper.GetString("SlackToken")
	BotName := viper.GetString("BotName")
	BotIconURL := viper.GetString("BotIconURL")
	params := slack.PostMessageParameters{}
	api := slack.New(SlackToken)
	fmt.Println(reflect.TypeOf(api))
	params.Username = BotName
	params.IconURL = BotIconURL

	attachment := slack.Attachment{
		Pretext:   MsgQuestion.Environment,
		Text:      "GitTag: " + MsgQuestion.GitTag + "\nGit " + MsgQuestion.GitCommit,
		Title:     "Environment URL",
		TitleLink: MsgQuestion.EnvironmentURL,
	}

	params.Attachments = []slack.Attachment{attachment}
	api.PostMessage("epam", "<@"+name+">: ", params)
}
func SendMassageJenkinsStatus(name string, MsgQuestion *JenkinsAnswerMap) {
	SlackToken := viper.GetString("SlackToken")
	BotName := viper.GetString("BotName")
	BotIconURL := viper.GetString("BotIconURL")
	params := slack.PostMessageParameters{}
	api := slack.New(SlackToken)
	fmt.Println(reflect.TypeOf(api))
	params.Username = BotName
	params.IconURL = BotIconURL

	test := slack.AttachmentAction{
		Name:  "first",
		Text:  "TextSS",
		Type:  "button",
		Value: "Value",
	}

	// testcall {
	// 	Actions:    []slack.AttachmentAction{test},
	// 	CallbackID: "wopr_game",
	// 	Team       Team               `json:"team"`
	// 	Channel    Channel            `json:"channel"`
	// 	User       User               `json:"user"`
	//
	// 	OriginalMessage Message `json:"original_message"`
	//
	// 	ActionTs     string `json:"action_ts"`
	// 	MessageTs    string `json:"message_ts"`
	// 	AttachmentID string `json:"attachment_id"`
	// 	Token        string `json:"token"`
	// 	ResponseURL  string `json:"response_url"`
	// }

	attachment := slack.Attachment{
		Pretext:    MsgQuestion.JobName,
		Text:       "GitTag: " + MsgQuestion.GitTag + "\nGit " + MsgQuestion.GitCommit + "\nTime build: " + MsgQuestion.TimeBuild,
		Title:      "Jenkins Job",
		TitleLink:  MsgQuestion.JobURL,
		Color:      "36a64f",
		CallbackID: "TTTTTTTTTTTTTTTTTTTT",
		Actions:    []slack.AttachmentAction{test},
	}

	params.Attachments = []slack.Attachment{attachment}
	// params.Attachments = []slack.AttachmentAction{test}
	api.PostMessage("epam", "<@"+name+">: ", params)
}

func SendMassageHelp(name string) {
	SlackToken := viper.GetString("SlackToken")
	BotName := viper.GetString("BotName")
	BotIconURL := viper.GetString("BotIconURL")

	api := slack.New(SlackToken)
	fmt.Println(reflect.TypeOf(api))
	params := slack.PostMessageParameters{}
	params.Username = BotName
	params.IconURL = BotIconURL
	api.PostMessage("epam", "I'm sorry <@"+name+">, I don't understand that. Try asking: <@"+strings.ToLower(BotName)+"> help", params)

}
