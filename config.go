package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var (
	config = &tConfig{}
)

type tConfig struct {
	AppName       string
	ListenAddress string
	ListenPort    int
	DBPath        string
	Path          string
	Debug         bool
}

func (c *tConfig) Read(path string) error {
	byts, err := ioutil.ReadFile(path)
	if err != nil {
		c = &tConfig{}
		return err
	}
	err = json.Unmarshal(byts, c)
	if err != nil {
		c = &tConfig{}
		return err
	}
	return nil
}

func (c *tConfig) Save() error {
	byts, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(c.Path, byts, os.ModePerm)
}
