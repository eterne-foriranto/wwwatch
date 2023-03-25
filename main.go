package main

const TeamName = "Охтыжёжик"
const Cutoff = 3
const FileName = "topic_id.dat"

func main() {
	url, status := checkComment(Cutoff, FileName, TeamName)

	if status {
		send(url)
	}
}
