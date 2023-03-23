package main

import (
	"fmt"
	"github.com/SevereCloud/vksdk/api"
	"github.com/agnivade/levenshtein"
	"os"
	"strconv"
	"strings"
)

const TeamName = "Охтыжёжик"
const Cutoff = 3
const FileName = "topic_id.dat"

func readTopicId(fileName string) int {
	fi, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}

	defer fi.Close()
	b := make([]byte, 8)
	_, err = fi.Read(b)

	if err != nil {
		fmt.Println(err)
	}

	topicId, err := strconv.Atoi(string(b))

	if err != nil {
		fmt.Println(err)
	}

	return topicId
}

func writeTopicId(fileName string, topicId int) {
	f, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()
	_, err = f.WriteString(strconv.Itoa(topicId))

	if err != nil {
		fmt.Println(err)
	}
}

func processTopic(vk *api.VK, params api.Params, topicId int, lowerTeamName string) {
	params["topic_id"] = topicId
	comments, err := vk.BoardGetComments(params)

	if err != nil {
		fmt.Println(err)
	}

	for _, comment := range comments.Items {
		lowerCommentText := strings.ToLower(comment.Text)
		distance := levenshtein.ComputeDistance(lowerCommentText, lowerTeamName)

		if distance <= Cutoff {
			fmt.Println(comment.Text)
		}
	}
	writeTopicId(FileName, topicId)
}

func main() {
	lowerTeamName := strings.ToLower(TeamName)
	token := os.Getenv("USER_TOKEN")
	vk := api.NewVK(token)
	params := api.Params{"group_id": "chgk_ivanovo"}
	groups, err := vk.GroupsGetByID(params)

	if err != nil {
		fmt.Println(err)
	}

	groupId := groups[0].ID
	params["group_id"] = groupId
	topics, err := vk.BoardGetTopics(params)

	if err != nil {
		fmt.Println(err)
	}

	topic := topics.Items[0]
	topicId := topic.ID
	processedTopicId := readTopicId(FileName)

	if topicId != processedTopicId {
		processTopic(vk, params, topicId, lowerTeamName)
	}
}
