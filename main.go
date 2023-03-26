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
	teamName := getConfigValue("main", "team_name")
	cutoff, err := strconv.Atoi(getConfigValue("main", "cutoff"))
	filename := getConfigValue("main", "filename")

	if err != nil {
		panic(err)
	}

	url, status := checkComment(cutoff, filename, teamName)

	if status {
		send(url)
	}
}
