package main

import (
	"fmt"
	"server-go/configs"
)

func main() {
	accessKey := configs.GetConfig().AWS.AccessKey
	fmt.Println(accessKey)
}
