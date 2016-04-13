package action

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/datawolf/index-cli/index"
)

func Status(c *cli.Context) {
	client := index.NewClient(nil)

	result, _, err := client.Status.Get()
	if err != nil {
		fmt.Printf("\nerror:%v\n", err)
		return
	}

	fmt.Printf("%s", result)
}
