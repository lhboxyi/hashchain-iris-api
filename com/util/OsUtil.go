package util

import (
	"fmt"
	"log"
	"os/user"
)

/**
 * 获取当前用户信息
 */
func GetCurrentUser() string {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("groupId:%s,homeDir:%s,userName:%s", currentUser.Gid, currentUser.HomeDir, currentUser.Username)
}
