package main

import (
	"github.com/astaxie/beego/config"
	"strconv"
)

func getConfigValue(sectionName string, key string) string {
	cnf, err := config.NewConfig("ini", "config.ini")

	if err != nil {
		panic(err)
	}

	section, err := cnf.GetSection(sectionName)

	if err != nil {
		panic(err)
	}
	return section[key]
}

func main() {
	processPost()
	teamName := getConfigValue("main", "team_name")
	cutoff, err := strconv.Atoi(getConfigValue("main", "cutoff"))
	topicIdFilename := getConfigValue("main", "topic_id_filename")

	if err != nil {
		panic(err)
	}

	url, status := checkComment(cutoff, topicIdFilename, teamName)

	if status {
		send(url)
	}
}
