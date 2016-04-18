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
	"github.com/docker/go-units"
	"os"
	"strings"
	"text/tabwriter"
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
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
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
		fmt.Fprintln(w, "NUM\tNAME:TAG\tSIZE")
		for i, tag := range res.ImageList {
			fmt.Fprintf(w, "%d\trnd-dockerhub.huawei.com/%s:%s\t%s\n", i+1, *res.RepoName, *tag.Tag, units.HumanSize(float64(*tag.Size)))
		}

		w.Flush()
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

func RepoDeleteTag(c *cli.Context) {
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

	for _, repoWithTag := range c.Args() {
		arr := strings.SplitN(repoWithTag, ":", 2)
		if len(arr) != 2 {
			fmt.Printf("\nError: The args should be like \"images:tag\"\n")
			continue
		}

		repo, tag := arr[0], arr[1]
		_, resp, err := client.Repositories.DeleteTag(repo, tag)
		if err != nil {
			fmt.Printf("\nError: %v\n", err)
			os.Exit(1)
		}

		if resp.StatusCode == 401 {
			log.Errorf("Unauthorized")
			continue
		}

		if resp.StatusCode == 404 {
			log.Errorf("The image(%s) with tag(%s) does not exist\n", repo, tag)
			continue
		}

		if resp.StatusCode == 406 {
			log.Errorf("StatusNotAcceptable")
			continue
		}

		fmt.Printf("Delete image(%s) with tha tag(%s) success.\n", repo, tag)
	}
}

func RepoDelete(c *cli.Context) {
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

	for _, repo := range c.Args() {
		_, resp, err := client.Repositories.DeleteRepo(repo)
		if err != nil {
			fmt.Printf("\nError: %v\n", err)
			os.Exit(1)
		}

		if resp.StatusCode == 401 {
			log.Errorf("Unauthorized")
			continue
		}

		if resp.StatusCode == 406 {
			log.Errorf("StatusNotAcceptable")
			continue
		}

		fmt.Printf("Delete repo(%s) success.\n", repo)
	}
}
