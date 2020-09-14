package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/lucasvmiguel/task/internal/command"
	"github.com/lucasvmiguel/task/internal/gitrepo"
	"github.com/lucasvmiguel/task/internal/gitrepo/github"
	"github.com/lucasvmiguel/task/internal/issuetracker"
	"github.com/lucasvmiguel/task/internal/issuetracker/jira"
	"github.com/lucasvmiguel/task/internal/versioncontrol/git"
	"gopkg.in/yaml.v2"

	"github.com/urfave/cli/v2"
)

// CLI Configuration
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
		PullRequest struct {
			New struct {
				Title       string `yaml:"title"`
				Description string `yaml:"description"`
			} `yaml:"new"`
		} `yaml:"pull-request"`
	} `yaml:"git-repo"`
}

// CLI Commands
var (
	startCMD = &cli.Command{
		Name:  "start",
		Usage: "start a task",
		Action: func(c *cli.Context) error {
			cfg := config(c.String("config-path"))
			spew.Dump(cfg)
			cmd, err := command.New(command.Config{
				IssueTracker:   issueTracker(cfg),
				GitRepo:        gitRepo(cfg),
				VersionControl: &git.Client{},
			})
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			err = cmd.Start(c.Args().First())
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			return err
		},
	}
)

// CLI Flags
var (
	configPathFlag = &cli.StringFlag{
		Name:    "config-path",
		Aliases: []string{"c"},
		Usage:   "task config path yaml",
	}
)

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

func gitRepo(cfg configuration) gitrepo.GitRepo {
	var gitRepo gitrepo.GitRepo
	switch cfg.GitRepo.Provider {
	case "github":
		gitRepo = &github.Client{Token: cfg.GitRepo.Token}
	default:
		fmt.Println("invalid issue tracker")
		os.Exit(1)
	}

	return gitRepo
}

func issueTracker(cfg configuration) issuetracker.IssueTracker {
	var issueTracker issuetracker.IssueTracker
	switch cfg.IssueTracker.Provider {
	case "jira":
		issueTracker = &jira.Client{
			Host:     cfg.IssueTracker.Host,
			Username: cfg.IssueTracker.Username,
			Key:      cfg.IssueTracker.Key,
		}
	default:
		fmt.Println("invalid issue tracker")
		os.Exit(1)
	}

	return issueTracker
}
