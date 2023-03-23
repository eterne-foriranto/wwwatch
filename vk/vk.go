package vk

import (
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/agnivade/levenshtein"
	"os"
	"strconv"
	"strings"
)

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

func processTopic(vk *api.VK, params api.Params, topicId int,
	lowerTeamName string, cutoff int, fileName string) (string, bool) {

	params["topic_id"] = topicId
	comments, err := vk.BoardGetComments(params)

	if err != nil {
		fmt.Println(err)
	}

	for _, comment := range comments.Items {
		lowerCommentText := strings.ToLower(comment.Text)
		distance := levenshtein.ComputeDistance(lowerCommentText, lowerTeamName)

		if distance <= cutoff {
			writeTopicId(fileName, topicId)
			return comment.Text, true
		}
	}
	return "", false
}

func CheckComment(cutoff int, fileName string, teamName string) (string, bool) {
	lowerTeamName := strings.ToLower(teamName)
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

	topicId := topics.Items[0].ID
	processedTopicId := readTopicId(fileName)

	if topicId != processedTopicId {
		return processTopic(vk, params, topicId, lowerTeamName, cutoff,
			fileName)
	}

	return "", false
}
