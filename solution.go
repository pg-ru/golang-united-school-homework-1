package solution

import (
	"github.com/kyokomi/emoji"
)

func GetMessage() string {
	str, _ := emoji.Println("Hello :world_map:!")

	return string(str)
}
