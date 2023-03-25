package main

import "eterne-foriranto/wwwatch/vk"

const TeamName = "Охтыжёжик"
const Cutoff = 3
const FileName = "topic_id.dat"

func main() {
	url, status := vk.CheckComment(Cutoff, FileName, TeamName)

	if status {
		send(url)
	}
}
