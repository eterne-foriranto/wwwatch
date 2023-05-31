package main

import (
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api"
	object2 "github.com/SevereCloud/vksdk/v2/object"
	"strings"
)

type Announcement struct {
	prefix string
}

func (a Announcement) getUrl(postId int) string {
	return fmt.Sprintf("%v%v%v_%v", Domain, a.prefix, getGroupID(), postId)
}

func (a Announcement) run() {
	timestampFilename := getConfigValue("main", "timestamp_filename")
	processedTimestamp := readIntFromFile(timestampFilename)
	post := getLatestPost()

	if post.Date > processedTimestamp || processedTimestamp == 0 { // 0 means that there was no file
		if keyPresents(post.Text) {
			url := a.getUrl(post.ID)
			send(url)
		}

		writeIntToFile(timestampFilename, post.Date)
	}
}

func getLatestPost() object2.WallWallpost {
	params := api.Params{
		"domain": getConfigValue("vk", "group_code"),
		"filter": "owner",
	}
	vk := getVK()
	posts, err := vk.WallGet(params)

	if err != nil {
		panic(err)
	}

	return posts.Items[0]
}

func keyPresents(text string) bool {
	return strings.Contains(text, getConfigValue("main", "post_key"))
}
