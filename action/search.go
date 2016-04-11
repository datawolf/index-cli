package action

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"strings"

	"github.com/datawolf/index-cli/config"
	"github.com/datawolf/index-cli/index"
)

func Search(c *cli.Context) {
	// Note: Search does not need to auth.
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
	var res *index.RepositoriesSearchResult
	for _, query := range c.Args() {
		result, _, err := client.Search.Repositories(query, nil)
		if err != nil {
			fmt.Printf("\nerror:%v\n", err)
			return
		}
		//fmt.Printf("\n%v\n", index.Stringify(result))
		res = result
		//fmt.Printf("\nFound %v results about %s\n", *res.NumberResults, *res.QueryString)
		fmt.Printf("NAME\t\t\t\tDESCRIPTION\t\t\t\t\t\tSTARS\tOFFICIAL\n")
		for _, repo := range res.Repositories {
			fmt.Printf("rnd-dockerhub.huawei.com/%s\t\t%s\t%d\t%v\n",
				*repo.Name, *repo.Description, *repo.StarCount, *repo.IsOfficial)
		}

	}
}
