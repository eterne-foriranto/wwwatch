package main

import (
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/agnivade/levenshtein"
	"os"
	"strconv"
	"strings"
)

const Domain = "https://vk.com/"
const Prefix = "topic-"

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
	lowerTeamName string, cutoff int, fileName string, groupId int) (string,
	bool) {

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
			return getUrl(groupId, topicId, comment.ID), true
		}
	}
	return "", false
}

func getUrl(groupId int, topicId int, commentId int) string {
	return fmt.Sprintf("%v%v%v_%v?post=%v", Domain, Prefix, groupId, topicId,
		commentId)
}

func getVK() *api.VK {
	token := getConfigValue("vk", "token")
	return api.NewVK(token)
}

func getGroupID() int {
	groupCode := getConfigValue("vk", "group_code")
	params := api.Params{"group_id": groupCode}
	vk := getVK()
	groups, err := vk.GroupsGetByID(params)

	if err != nil {
		fmt.Println(err)
	}

	return groups[0].ID
}

func checkComment(cutoff int, fileName string, teamName string) (string, bool) {
	lowerTeamName := strings.ToLower(teamName)
	vk := getVK()
	groupId := getGroupID()
	params := api.Params{"group_id": groupId}
	topics, err := vk.BoardGetTopics(params)

	if err != nil {
		fmt.Println(err)
	}

	topicId := topics.Items[0].ID
	processedTopicId := readTopicId(fileName)

	if topicId != processedTopicId {
		return processTopic(vk, params, topicId, lowerTeamName, cutoff,
			fileName, groupId)
	}

	return "", false
}
