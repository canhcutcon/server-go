package main

import (
	"fmt"
	"server-go/internal/configs"
)

func main() {
	accessKey := configs.GetConfig().AWS.AccessKey
	fmt.Println(accessKey)

}
