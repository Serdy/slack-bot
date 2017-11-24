package commands

import (
	"fmt"

	"github.com/bndr/gojenkins"
	"github.com/spf13/viper"
)

type JenkinsAnswerMap struct {
	TimeBuild string
	GitTag    string
	GitCommit string
	JobURL    string
	JobName   string
}

func jenkins_conn() (*gojenkins.Jenkins, error) {

	JenkinsURL := viper.GetString("JenkinsURL")
	JenkinsUser := viper.GetString("JenkinsUser")
	JenkinsPass := viper.GetString("JenkinsPass")
	jenkins, err := gojenkins.CreateJenkins(JenkinsURL, JenkinsUser, JenkinsPass).Init()

	if err != nil {
		return nil, err
	}
	return jenkins, nil

}

func JenkinsJobInfo(NameUser string, JenkinsJobName string) {
	// JenkinsInfoMap := make(map[string]int)

	JenkinsInfo := new(JenkinsAnswerMap)

	answer := viper.GetStringMap("JenkinsJob")
	jobName, _ := answer[JenkinsJobName].(string)
	JenkinsConn, _ := jenkins_conn()
	job, err := JenkinsConn.GetJob(jobName)

	if err != nil {
		fmt.Println(err)
		return
	}
	secess, err := job.GetLastSuccessfulBuild()
	if err != nil {
		panic(err)
	}
	// #################################################
	// #################################################
	builds, err := JenkinsConn.GetJob(jobName)

	if err != nil {
		panic(err)
	}
	println(builds)

	// ######################################
	build, _ := job.GetBuild(secess.GetBuildNumber())
	fmt.Println(build.GetParameters())
	var GitTag string
	println(build)
	for _, param := range build.GetParameters() {
		if string(param.Name) == "git_tag" {
			GitTag = param.Value
			break
		}
	}
	println(GitTag)

	SendMassageJenkinsStatus(NameUser, JenkinsInfo)
}
