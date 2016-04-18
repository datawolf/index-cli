//
// users.go
// Copyright (C) 2016 wanglong <wanglong@laoqinren.net>
//
// Distributed under terms of the MIT license.
//

package action

import (
	"bufio"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/datawolf/index-cli/config"
	"github.com/datawolf/index-cli/index"
	"golang.org/x/crypto/ssh/terminal"
	"net/url"
	"os"
	"strings"
	"syscall"
)

func CreateUser(c *cli.Context) {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Please input USERNAME you want to create: ")
	username, _ := r.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Please input PASSWORD: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)
	password = strings.TrimSpace(password)

	fmt.Printf("\n")
	fmt.Print("Please re-input PASSWORD: ")
	bytePassword, _ = terminal.ReadPassword(int(syscall.Stdin))
	password2 := string(bytePassword)
	password2 = strings.TrimSpace(password2)

	if password != password2 {
		fmt.Printf("\nSorry, passwords do not match\n")
		os.Exit(1)
	}

	fmt.Printf("\n")
	fmt.Print("Please input EMAIL: ")
	email, _ := r.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Please input PHONE: ")
	phone, _ := r.ReadString('\n')
	phone = strings.TrimSpace(phone)

	user := &index.User{
		Username: &username,
		Password: &password,
		Email:    &email,
		Phone:    &phone,
	}

	client := index.NewClient(nil)
	rel, err := url.Parse(index.EuropaURL)

	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		os.Exit(1)
	}
	client.BaseURL = rel

	result, _, err := client.Users.Create(user)
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s", result)
}

func UpdateUser(c *cli.Context) {
	// Get the username and password
	configFile, err := config.Load("")
	if err != nil {
		log.Fatal("Failed to loading the config file")
	}

	ac := configFile.AuthConfigs["rnd-dockerhub.huawei.com"]
	if ac.Username == "" && ac.Password == "" {
		log.Fatal("Please login in the hub, using command \"index-cli login\"")
	}

	username := strings.TrimSpace(ac.Username)
	password := strings.TrimSpace(ac.Password)

	// Get the new password
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Please input NEW PASSWORD: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	newPassword := string(bytePassword)
	newPassword = strings.TrimSpace(newPassword)

	fmt.Printf("\n")
	fmt.Print("Please re-input NEW PASSWORD: ")
	bytePassword, _ = terminal.ReadPassword(int(syscall.Stdin))
	newPassword2 := string(bytePassword)
	newPassword2 = strings.TrimSpace(newPassword2)

	if newPassword != newPassword2 {
		fmt.Printf("\nSorry, passwords do not match\n")
		os.Exit(1)
	}

	// Get the new Email and Phone
	fmt.Printf("\n")
	fmt.Print("Please input EMAIL: ")
	email, _ := r.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Please input PHONE: ")
	phone, _ := r.ReadString('\n')
	phone = strings.TrimSpace(phone)

	user := &index.User{
		Username:    &username,
		Password:    &password,
		NewPassword: &newPassword,
		Email:       &email,
		Phone:       &phone,
	}

	client := index.NewClient(nil)
	rel, err := url.Parse(index.EuropaURL)

	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		os.Exit(1)
	}
	client.BaseURL = rel

	result, _, err := client.Users.Update(user)
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s", result)
}
