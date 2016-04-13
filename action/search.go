package action

import (
	"fmt"
	"github.com/codegangsta/cli"

	"github.com/datawolf/index-cli/index"
)

// Search search docker image in rnd-dockerhub
func Search(c *cli.Context) {
	client := index.NewClient(nil)
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
