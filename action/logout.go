package action

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func Logout(c *cli.Context) {
	log.Info("logout successfully")
}
