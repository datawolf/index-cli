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
	"github.com/codegangsta/cli"
	"net/url"
	"os"
	"strings"
	"syscall"

	_ "github.com/datawolf/index-cli/config"
	"github.com/datawolf/index-cli/index"
	"golang.org/x/crypto/ssh/terminal"
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

}
