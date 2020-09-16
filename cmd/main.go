package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/lucasvmiguel/task/internal/command"
	"github.com/lucasvmiguel/task/internal/factory"
	"github.com/lucasvmiguel/task/internal/versioncontrol/git"

	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

// CLI Configuration
// All config fields that can be passed to the CLI
type configuration struct {
	IssueTracker struct {
		Provider string `yaml:"provider"`
		Host     string `yaml:"host"`
		Username string `yaml:"username"`
		Key      string `yaml:"key"`
	} `yaml:"issue-tracker"`
	GitRepo struct {
		Provider    string `yaml:"provider"`
		Host        string `yaml:"host"`
		Token       string `yaml:"token"`
		Org         string `yaml:"org"`
		Repository  string `yaml:"repository"`
		PullRequest struct {
			New struct {
				Title       string `yaml:"title"`
				Description string `yaml:"description"`
			} `yaml:"new"`
		} `yaml:"pull-request"`
	} `yaml:"git-repo"`
}

// CLI Commands
// All commands that can be executed by the CLI
var (
	startCMD = &cli.Command{
		Name:  "start",
		Usage: "start a task",
		Action: func(c *cli.Context) error {
			cfg := config(c.String("config-path"))

			gitRepo, err := factory.GitRepo(factory.GitRepoParams{
				Provider: factory.GitRepoProvider(cfg.GitRepo.Provider),
				Host:     cfg.GitRepo.Host,
				Token:    cfg.GitRepo.Token,
			})
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			issueTracker, err := factory.IssueTracker(factory.IssueTrackerParams{
				Provider: factory.IssueTrackerProvider(cfg.IssueTracker.Provider),
				Host:     cfg.IssueTracker.Host,
				Username: cfg.IssueTracker.Username,
				Key:      cfg.IssueTracker.Key,
			})
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			cmd, err := command.New(command.NewParams{
				IssueTracker:   issueTracker,
				GitRepo:        gitRepo,
				VersionControl: &git.Client{},
			})
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			err = cmd.Start(c.Args().First(), cfg.GitRepo.Org, cfg.GitRepo.Repository)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			return err
		},
	}
)

// CLI Flags
// All flags that can be passed to the CLI
var (
	configPathFlag = &cli.StringFlag{
		Name:    "config-path",
		Aliases: []string{"c"},
		Usage:   "task config path yaml",
	}
)

// main function, here is where the magic begins
func main() {
	app := &cli.App{
		Name:        "task",
		Description: "TODO",
		Usage:       "TODO",
		Flags: []cli.Flag{
			configPathFlag,
		},
		Commands: []*cli.Command{
			startCMD,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("invalid command")
		os.Exit(1)
	}
}

// this function will read from a config file and override with any flags passed
func config(path string) configuration {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("invalid yaml file path")
		os.Exit(1)
	}

	c := configuration{}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		fmt.Println("invalid yaml to unmarshal")
		os.Exit(1)
	}

	return c
}
