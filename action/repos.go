//
// repos.go
// Copyright (C) 2016 wanglong <wanglong@laoqinren.net>
//
// Distributed under terms of the MIT license.
//

package action

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/datawolf/index-cli/config"
	"github.com/datawolf/index-cli/index"
	"os"
	"strings"

	"github.com/docker/go-units"
)

func RepoGetProperty(c *cli.Context) {
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

	var res *index.Property
	for _, repo := range c.Args() {
		result, resp, err := client.Repositories.Get(repo)
		if err != nil {
			log.Error(err)
		}

		if resp.StatusCode == 401 {
			log.Errorf("Unauthorized(Maybe not found \"%s\") in rnd-dockerhub", repo)
			continue
		}

		if resp.StatusCode == 406 {
			log.Errorf("StatusNotAcceptable")
			continue
		}
		res = result

		fmt.Printf("Image Name         : %s\n", *res.RepoName)
		fmt.Printf("Image Size         : %s\n", units.HumanSize(float64(*res.Size)))
		fmt.Printf("Number of Images   : %d\n", *res.NumberImage)
		fmt.Printf("Access Level       : %s\n", *res.Property)
		if res.NumberDL != nil {
			fmt.Printf("Number of Download : %d\n", *res.NumberDL)
		} else {
			fmt.Printf("Number of Download : 0\n")
		}
		fmt.Printf("No.\tIMAGE with TAG\t\t\t\t\tSIZE\n")
		for i, tag := range res.ImageList {
			fmt.Printf("%d\trnd-dockerhub.huawei.com/%s:%s		\t\t%s\n", i+1, *res.RepoName, *tag.Tag, units.HumanSize(float64(*tag.Size)))
		}
	}
}

func RepoSetProperty(c *cli.Context) {
	access := c.String("access")
	if access == "" {
		log.Fatal("Can not proceed without -a <access level> specified")
	}

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
	property := &index.Property{
		Property: &access,
	}
	for _, repo := range c.Args() {
		result, resp, err := client.Repositories.Set(repo, property)
		if err != nil {
			fmt.Printf("\nerror: %v\n", err)
			os.Exit(1)
		}

		if resp.StatusCode == 401 {
			log.Errorf("Unauthorized(Maybe not found \"%s\") in rnd-dockerhub", repo)
			continue
		}

		if resp.StatusCode == 406 {
			log.Errorf("StatusNotAcceptable")
			continue
		}
		fmt.Printf("Set %s Access Level to %s: %s\n", repo, access, result)
	}
}
