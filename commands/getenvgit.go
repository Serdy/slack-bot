package commands

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type EnvGitMap struct {
	Environment    string
	EnvironmentURL string
	GitTag         string
	GitCommit      string
}

type HybrisVersionJSON struct {
	Git struct {
		Branch   string `json:"branch"`
		Describe string `json:"describe"`
		Commit   string `json:"commit"`
	} `json:"git"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

// Get json from url
func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

// Get git version from env
func GetEnvGit(NameUser string, JenkinsJobName string) {
	EnvName := viper.GetString("EnvGit." + JenkinsJobName)
	EnvVersionURL := viper.GetString(EnvName + ".VersionURL")
	EnvURLs := viper.GetString(EnvName + ".EnvURLs")
	ResponseEnvGitMap := new(EnvGitMap)
	HybrisVersionResponseJSON := new(HybrisVersionJSON)
	getJson(EnvVersionURL, HybrisVersionResponseJSON)

	ResponseEnvGitMap.Environment = strings.ToUpper(EnvName)
	ResponseEnvGitMap.EnvironmentURL = EnvURLs
	ResponseEnvGitMap.GitTag = HybrisVersionResponseJSON.Git.Describe
	ResponseEnvGitMap.GitCommit = HybrisVersionResponseJSON.Git.Commit

	SendMassageEnvGit(NameUser, ResponseEnvGitMap)
}
