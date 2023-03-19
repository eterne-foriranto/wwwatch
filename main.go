package wwwatch

import (
	"github.com/SevereCloud/vksdk/api"
	"os"
)

func main() {
	token := os.Getenv("USER_TOKEN")
	vk := api.NewVK(token)
}
