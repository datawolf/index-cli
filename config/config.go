//
// config.go
// Copyright (C) 2016 wanglong <wanglong@laoqinren.net>
//
// Distributed under terms of the MIT license.
//

package config

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	// ConfigFileName is the name of config file
	ConfigFileName = "config.json"
	oldConfigfile  = ".dockercfg"

	defaultIndexserver = "https://index.docker.io/v1/"
)

var (
	configDir = os.Getenv("DOCKER_CONFIG")
)

func init() {
	home := os.Getenv("HOME")
	if home == "" {
		home = "~/"
	}

	if configDir == "" {
		configDir = filepath.Join(home + "/.docker")
	}
}

func ConfigDir() string {
	return configDir
}

// AuthConfig contains authorization infomation for connecting to a Registry
type AuthConfig struct {
	Username      string `json:"uername,omitempty"`
	Password      string `json:"password,omitempty"`
	Auth          string `json:"auth"`
	Email         string `json:"email,omitempty"`
	ServerAddress string `json:"serveraddress:,omitempty"`
	IdentityToken string `json:"identitytoken,omitempty"`
	RegistryToken string `json:"registrytoken,omitempty"`
}

// ConfigFile ~/.docker/config.json file info
type ConfigFile struct {
	AuthConfigs      map[string]AuthConfig `json:"auths"`
	HTTPHeaders      map[string]string     `json:"HttpHeaders,omitempty"`
	PsFormat         string                `json:"psFormat,omitempty"`
	ImagesFormat     string                `json:"imagesFormat,omitempty"`
	DetachKeys       string                `json:"detachKeys,omitempty"`
	CredentialsStore string                `json:"credsStore,omitempty"`
	filename         string                // Note: not serialized - for internal use only
}

// NewConfigFile initilizes am empty configuration file for the given filename 'fn'
func NewConfigFile(fn string) *ConfigFile {
	return &ConfigFile{
		AuthConfigs: make(map[string]AuthConfig),
		HTTPHeaders: make(map[string]string),
		filename:    fn,
	}
}

func (configFile *ConfigFile) LoadFromReader(configData io.Reader) error {
	if err := json.NewDecoder(configData).Decode(&configFile); err != nil {
		return err
	}

	if err := InitAESKey(); err != nil {
		return err
	}

	for addr, ac := range configFile.AuthConfigs {
		data, err := base64.StdEncoding.DecodeString(ac.Auth)
		if err != nil {
			return err
		}

		decAuth, err := AESDecrypt([]byte(data), KeyAES)
		if err != nil {
			return err
		}

		ac.Username, ac.Password, err = DecodeAuth(string(decAuth))
		if err != nil {
			return err
		}

		ac.Auth = ""
		ac.ServerAddress = addr
		configFile.AuthConfigs[addr] = ac
	}

	return nil
}

// Load reads the configuration files in the given directory
func Load(configDir string) (*ConfigFile, error) {
	if configDir == "" {
		configDir = ConfigDir()
	}

	configFile := ConfigFile{
		AuthConfigs: make(map[string]AuthConfig),
		filename:    filepath.Join(configDir, ConfigFileName),
	}

	if _, err := os.Stat(configFile.filename); err == nil {
		file, err := os.Open(configFile.filename)
		if err != nil {
			return &configFile, err
		}
		defer file.Close()
		err = configFile.LoadFromReader(file)
		return &configFile, err
	}

	return &configFile, fmt.Errorf("config file not found")
}

// DecodeAuth decodes a base64 encoded string and returns username and password
func DecodeAuth(authStr string) (string, string, error) {
	decLen := base64.StdEncoding.DecodedLen(len(authStr))
	decoded := make([]byte, decLen)
	authByte := []byte(authStr)

	n, err := base64.StdEncoding.Decode(decoded, authByte)
	if err != nil {
		return "", "", err
	}

	if n > decLen {
		return "", "", fmt.Errorf("something went wrong decoding auth config")
	}

	arr := strings.SplitN(string(decoded), ":", 2)
	if len(arr) != 2 {
		return "", "", fmt.Errorf("Invalid auth configuration file")
	}

	password := strings.Trim(arr[1], "\x00")

	return arr[0], password, nil
}
