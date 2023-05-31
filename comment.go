package main

import (
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/agnivade/levenshtein"
	"strconv"
	"strings"
)

type Comment struct {
	prefix string
}

func (c Comment) getUrl(groupId int, topicId int, commentId int) string {
	return fmt.Sprintf("%v%v%v_%v?post=%v", Domain, c.prefix, groupId, topicId,
		commentId)
}

func (c Comment) run() {
	teamName := getConfigValue("main", "team_name")
	cutoff, err := strconv.Atoi(getConfigValue("main", "cutoff"))
	topicIdFilename := getConfigValue("main", "topic_id_filename")

	if err != nil {
		panic(err)
	}

	url, status := c.check(cutoff, topicIdFilename, teamName)

	if status {
		send(url)
	}
}

func (c Comment) check(cutoff int, topicIdFileName string, teamName string) (string, bool) {
	lowerTeamName := strings.ToLower(teamName)
	vk := getVK()
	groupId := getGroupID()
	params := api.Params{"group_id": groupId}
	topics, err := vk.BoardGetTopics(params)

	if err != nil {
		fmt.Println(err)
	}

	topicId := topics.Items[0].ID
	processedTopicId := readIntFromFile(topicIdFileName)

	if topicId != processedTopicId {
		return c.processTopic(vk, params, topicId, lowerTeamName, cutoff,
			topicIdFileName, groupId)
	}

	return "", false
}
func (c Comment) processTopic(vk *api.VK, params api.Params, topicId int,
	lowerTeamName string, cutoff int, topicIdFileName string, groupId int) (string,
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
			writeIntToFile(topicIdFileName, topicId)
			return c.getUrl(groupId, topicId, comment.ID), true
		}
	}
	return "", false
}
