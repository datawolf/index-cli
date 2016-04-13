package action

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"strings"

	"github.com/datawolf/index-cli/config"
	"github.com/datawolf/index-cli/index"
)

func UserInfo(c *cli.Context) {
	configFile, err := config.Load("")
	if err != nil {
		log.Fatal("Failed to loading the config file")
	}

	ac := configFile.AuthConfigs["rnd-dockerhub.huawei.com"]
	if ac.Username == "" && ac.Password == "" {
		log.Fatal("Please login in the hub, using command \"index-cli login\"")
	}

	tp := index.BasicAuthTransport{
		Username: strings.TrimSpace(ac.Username),
		Password: strings.TrimSpace(ac.Password),
	}

	client := index.NewClient(tp.Client())
	var res *index.UserInfo
	result, _, err := client.UserInfo.Get()
	if err != nil {
		fmt.Printf("\nerror:%v\n", err)
		return
	}
	res = result
	fmt.Printf("User Name : %v\n", *res.Username)
	fmt.Printf("Namespace : %v\n", *res.Namespace)
	fmt.Printf("Product   : %v\n", *res.Product)
	fmt.Printf("Quote     : %v\n", *res.Quote)
	fmt.Printf("Used Space: %v\n", *res.UsedSpace)
	fmt.Printf("Number Of Image         : %v\n", *res.NumberImage)
	fmt.Printf("Number of Image(private): %v\n", *res.NumberImagePrivate)
	fmt.Printf("Number of Image(protect): %v\n", *res.NumberImageProtect)
	fmt.Printf("Number of Image(public) : %v\n", *res.NumberImagePublic)
}
