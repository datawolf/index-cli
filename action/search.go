package action

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/codegangsta/cli"
	"github.com/datawolf/index-cli/index"
)

// Search search docker image in rnd-dockerhub
func Search(c *cli.Context) {
	client := index.NewClient(nil)
	var res *index.RepositoriesSearchResult

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	for _, query := range c.Args() {
		result, _, err := client.Search.Repositories(query, nil)
		if err != nil {
			fmt.Printf("\nerror:%v\n", err)
			return
		}
		res = result
		fmt.Fprintln(w, "NAME\tDESCRIPTION\tSTARS\tOFFICIAL")
		for _, repo := range res.Repositories {
			fmt.Fprintf(w, "rnd-dockerhub.huawei.com/%s\t%s\t%d\t%v\n",
				*repo.Name, *repo.Description, *repo.StarCount, *repo.IsOfficial)
		}
	}
	w.Flush()
}
