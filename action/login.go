package action

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"os"
	"os/exec"
	"syscall"
)

func Login(c *cli.Context) {
	if c.NArg() != 0 {
		log.Fatalf("invalid arguments %v", c.Args())
	}

	path, err := exec.LookPath("docker")
	if err != nil {
		log.Fatal("please install \"docker\" first")
	}

	argv := []string{path, "login", "rnd-dockerhub.huawei.com"}
	if err := syscall.Exec(path, argv, os.Environ()); err != nil {
		log.Fatal(err)
	}
}
