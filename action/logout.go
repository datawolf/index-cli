package action

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"

	"os"
	"os/exec"
	"syscall"
)

func Logout(c *cli.Context) {
	path, err := exec.LookPath("docker")
	if err != nil {
		log.Fatal("please install \"docker\" first")
	}

	os.Args[0] = path
	if err := syscall.Exec(os.Args[0], os.Args, os.Environ()); err != nil {
		log.Fatal(err)
	}
}
