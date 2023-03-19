package main

import (
	"fmt"
	"github.com/SevereCloud/vksdk/api"
	"os"
)

func main() {
	token := os.Getenv("USER_TOKEN")
	vk := api.NewVK(token)
	params := api.Params{"group_id": "chgk_ivanovo"}
	group, err := vk.GroupsGetByID(params)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(group)
}
