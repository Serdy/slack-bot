package commands

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/nlopes/slack"
	"github.com/spf13/viper"
)

func BotMsg(NameUser string, MsgQuestion string, Channel string) {
	BotMassege := viper.GetString("BotQA." + MsgQuestion)
	SendMassage(NameUser, BotMassege, Channel)
}
func SendMassage(name string, MsgQuestion string, Channel string) {
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
	api.PostMessage(Channel, "<@"+name+">: "+MsgQuestion, params)

}
func SendMassageEnvGit(name string, MsgQuestion *EnvGitMap, Channel string) {
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
	api.PostMessage(Channel, "<@"+name+">: ", params)
}
func SendMassageAWSEnv(name string, AWSInst *AWSInstancesDataSlice, Channel string) {
	SlackToken := viper.GetString("SlackToken")
	BotName := viper.GetString("BotName")
	BotIconURL := viper.GetString("BotIconURL")
	params := slack.PostMessageParameters{}
	api := slack.New(SlackToken)
	fmt.Println(reflect.TypeOf(api))
	params.Username = BotName
	params.IconURL = BotIconURL
	fmt.Println(AWSInst)
	for _, instance := range AWSInst.Instances {
		attachment := slack.Attachment{
			Text:      "Name_Instance: " + instance.TagName + "\n DNS_Instance: " + instance.DNSName,
		}
		params.Attachments = []slack.Attachment{attachment}
		api.PostMessage(Channel, "<@"+name+">: ", params)
}
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

func SendMassageHelp(name string, Channel string) {
	SlackToken := viper.GetString("SlackToken")
	BotName := viper.GetString("BotName")
	BotIconURL := viper.GetString("BotIconURL")
	api := slack.New(SlackToken)
	fmt.Println(reflect.TypeOf(api))
	params := slack.PostMessageParameters{}
	params.Username = BotName
	params.IconURL = BotIconURL
	api.PostMessage(Channel, "I'm sorry <@"+name+">, I don't understand that. Try asking: <@"+strings.ToLower(BotName)+"> help", params)

}
