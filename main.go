package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/datawolf/index-cli/action"
)

var Version = "0.0.0"

func exit(err error) {
	if err != nil {
		logrus.Fatal(err)
	}
}

func beforeApp(c *cli.Context) error {
	if c.Bool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = os.Args[0]
	app.Version = Version
	app.Usage = "The cli tool to access the index and hub."
	app.Before = beforeApp
	app.EnableBashCompletion = true
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "w00291922",
			Email: "long.wanglong@huawei.com",
		},
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "Debug logging",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:     "login",
			Usage:    "login the hub",
			HideHelp: true,
			Action:   action.Login,
		},
		{
			Name:     "logout",
			Usage:    "logout the hub",
			HideHelp: true,
			Action:   action.Logout,
		},
		{
			Name:     "search",
			Usage:    "search rnd-dockerhub for images",
			HideHelp: true,
			Action:   action.Search,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "stars, s",
					Usage: "Only displays with at least x stars",
				},
			},
		},
	}

	exit(app.Run(os.Args))
}
