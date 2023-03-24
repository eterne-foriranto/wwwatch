package main

import (
	"eterne-foriranto/wwwatch/vk"
	"fmt"
)

const TeamName = "Охтыжёжик"
const Cutoff = 3
const FileName = "topic_id.dat"

func main() {
	comment, status := vk.CheckComment(Cutoff, FileName, TeamName)
	fmt.Println(comment)
	fmt.Println(status)
}
