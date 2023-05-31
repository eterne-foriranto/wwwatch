package main

import (
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api"
)

const Domain = "https://vk.com/"

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
