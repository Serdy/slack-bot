package commands

import (
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/viper"
)

type AWSInstancesMap struct{
	Instances AWSInstancesData
} 
type AWSInstancesData struct {
		TagName string
		DNSName string
}
type AWSInstancesDataSlice struct {
	Instances []struct {
		TagName string 
		DNSName string 
	} 
}

func createinstancesmap(resp *ec2.DescribeInstancesOutput) *AWSInstancesDataSlice{
	NewAWSInstancesMap := new(AWSInstancesDataSlice)
	AWSInstances := new(AWSInstancesData)
	for key, instance := range resp.Reservations {
		AWSInstances.DNSName = (*instance.Instances[0].PublicDnsName)
		for _, tag := range instance.Instances[0].Tags {
			if *tag.Key == "Name"{
				fmt.Println(*tag.Value, key)
				AWSInstances.TagName = *tag.Value
				break
			}
		}
		NewAWSInstancesMap.Instances = append(NewAWSInstancesMap.Instances, *AWSInstances)
	}	
	fmt.Println(NewAWSInstancesMap.Instances[0].DNSName)
	return NewAWSInstancesMap
}

// AWSIntsansesFilter ... Get aws instances from AWS
func AWSIntsansesFilter(NameUser string, envName string, Channel string) {
	EnvName := viper.GetString("EnvAWS." + envName)
	AWSTag := viper.GetString("AWSTag")
	AWSRegion := viper.GetString("AWSRegion")

	sess := session.Must(session.NewSession())

	nameFilter := AWSTag + EnvName + "-"
	
	svc := ec2.New(sess, &aws.Config{Region: aws.String(AWSRegion)})
	fmt.Printf("listing instances with tag %v in: %v\n", nameFilter, AWSRegion)
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:Name"),
				Values: []*string{
					aws.String(strings.Join([]string{"*", nameFilter, "*"}, "")),
				},
			},
		},
	}

	resp, err := svc.DescribeInstances(params)
	if err != nil {
		fmt.Println("there was an error listing instances in", AWSRegion, err.Error())
		log.Fatal(err.Error())
	}
	InstansecMap := createinstancesmap(resp)
	SendMassageAWSEnv(NameUser, InstansecMap, Channel)
}
