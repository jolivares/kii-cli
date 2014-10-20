package kiicli

import (
	"fmt"

	"github.com/tmtk75/cli"
)

func Profile() *GlobalConfig {
	return globalConfig
}

var ProfileCommands = []cli.Command{
	cli.Command{
		Name: "ls",
		Action: func(c *cli.Context) {
			file, _ := loadIniFile()
			for _, s := range *file {
				if v, has := s["profile"]; has {
					fmt.Printf("default-profile: %v\n", v)
				}
			}
			for k, s := range *file {
				if _, has := s["app_id"]; !has {
					continue
				}
				fmt.Printf("%v\t%v\n", s["app_id"], k)
			}
		},
	},
}
