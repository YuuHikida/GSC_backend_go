package service

import "fmt"

func CheackGitHubAccout(userName string) {
	url := fmt.Sprintf("https://github.com/users/%s/contributions", userName)

}
