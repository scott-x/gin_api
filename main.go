package main

import (
	"fmt"
	"github.com/scott-x/gutils/cmd"
	"github.com/scott-x/gutils/fs"
	"path"
	"strings"
)

var (
	HOME         = fs.HOME
	TEMPLATE_DIR = HOME + "/go/src/github.com/scott-x/templates/gin_scaffold"
)

func main() {
	//check template dir is exists or not
	if !fs.DirExists(TEMPLATE_DIR) {
		cmd.Warning(TEMPLATE_DIR + " is not exist")
		return
	}
	//check the directory is correct or not
	if !fs.DirExists("./.git") {
		cmd.Warning("failed to find .git in current directory, please make sure the directory is correct")
		return
	}
	cmd.AddQuestion("github", "what's your github username:", "please input the correct username:", "^[a-zA-Z].*")
	res := cmd.Exec()
	github := res["github"]
	current := path.Base(fs.Dir())
	err := fs.CopyFolder(TEMPLATE_DIR, "./")
	if err != nil {
		cmd.Warning("ops, something is wrong while copying")
		return
	}
	filesShouldModify := []string{
		"controller/api.go",
		"controller/demo.go",
		"dao/demo.go",
		"dao/user.go",
		"dto/api.go",
		"dto/demo.go",
		"middleware/recovery.go",
		"middleware/request_log.go",
		"middleware/translation.go",
		"router/route.go",
		"go.mod",
		"main.go",
		"conf/dev/base.toml",
	}
	m := make(map[string]string)
	m["scott-x"] = strings.TrimSpace(github)
	m["gin_scaffold"] = current
	for _, v := range filesShouldModify {
		err = fs.ReadAndReplace("./"+v, m)
		if err != nil {
			fmt.Print(err)
			cmd.Warning("replace " + v + " error")
		}
	}
	cmd.Info("Congraduations, gin_scaffold structure was build successfully....")
}
