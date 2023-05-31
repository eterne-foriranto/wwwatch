package main

import (
	"github.com/astaxie/beego/config"
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
	announcement := Announcement{"wall-"}
	announcement.run()
	comment := Comment{"topic-"}
	comment.run()
}
