package main

import (
	"fmt"
	"github.com/SevereCloud/vksdk/api"
	"github.com/agnivade/levenshtein"
	"os"
	"strings"
)

const TeamName = "Охтыжёжик"
const Cutoff = 3

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

	topicId := topics.Items[0].ID
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
}
