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
	app.Usage = "The cli tool that access rnd-dockerhub.huawei.com."
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
			Usage:    "Login the hub",
			HideHelp: true,
			Action:   action.Login,
		},
		{
			Name:     "logout",
			Usage:    "Logout the hub",
			HideHelp: true,
			Action:   action.Logout,
		},
		{
			Name:     "status",
			Usage:    "Get the status of the rnd-dockerhub",
			HideHelp: true,
			Action:   action.Status,
		},
		{
			Name:     "userinfo",
			Usage:    "Get the user info of yourself",
			HideHelp: true,
			Action:   action.UserInfo,
		},
		{
			Name:        "user",
			ShortName:   "u",
			Usage:       "Control user",
			HideHelp:    true,
			Subcommands: userSubcommand(),
		},
		{
			Name:        "repo",
			ShortName:   "r",
			Usage:       "Control the Repository",
			HideHelp:    true,
			Subcommands: repoSubcommand(),
		},
		{
			Name:     "search",
			Usage:    "Search rnd-dockerhub for images",
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

func repoSubcommand() []cli.Command {
	return []cli.Command{
		{
			Name:   "get",
			Usage:  "Get Specific Repository's Property",
			Action: action.RepoGetProperty,
		},
		{
			Name:   "set",
			Usage:  "Set Specific Repository's Property",
			Action: action.RepoSetProperty,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "access, a",
					Usage: "Set access level for an images(private,public,protect)",
				},
				cli.StringFlag{
					Name:  "description, d",
					Usage: "Set description an images",
				},
			},
		},
		{
			Name:   "del",
			Usage:  "Delete Specific Repository",
			Action: action.RepoDelete,
		},
		{
			Name:   "rmi",
			Usage:  "Delete Repository With Specific tag",
			Action: action.RepoDeleteTag,
		},
	}
}

func userSubcommand() []cli.Command {
	return []cli.Command{
		{
			Name:   "create",
			Usage:  "create an user",
			Action: action.CreateUser,
		},
		{
			Name:   "update",
			Usage:  "update an user",
			Action: action.UpdateUser,
		},
	}
}
