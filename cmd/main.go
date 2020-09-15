package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/lucasvmiguel/task/internal/command"
	"github.com/lucasvmiguel/task/internal/gitrepo"
	"github.com/lucasvmiguel/task/internal/gitrepo/github"
	"github.com/lucasvmiguel/task/internal/issuetracker"
	"github.com/lucasvmiguel/task/internal/issuetracker/jira"
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

			cmd, err := command.New(command.Config{
				IssueTracker:   issueTracker(cfg),
				GitRepo:        gitRepo(cfg),
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

// returns a git repository
func gitRepo(cfg configuration) gitrepo.GitRepo {
	var gitRepo gitrepo.GitRepo
	switch cfg.GitRepo.Provider {
	case "github":
		gitRepo = &github.Client{}
	default:
		fmt.Println("invalid git repo")
		os.Exit(1)
	}

	err := gitRepo.Authenticate(cfg.GitRepo.Host, cfg.GitRepo.Token)
	if err != nil {
		fmt.Println("failed to authenticate git repo")
		os.Exit(1)
	}

	return gitRepo
}

// returns a issue tracker
func issueTracker(cfg configuration) issuetracker.IssueTracker {
	var issueTracker issuetracker.IssueTracker
	switch cfg.IssueTracker.Provider {
	case "jira":
		issueTracker = &jira.Client{}
	default:
		fmt.Println("invalid issue tracker")
		os.Exit(1)
	}

	err := issueTracker.Authenticate(cfg.IssueTracker.Host, cfg.IssueTracker.Username, cfg.IssueTracker.Key)
	if err != nil {
		fmt.Println("failed to authenticate issue tracker")
		os.Exit(1)
	}

	return issueTracker
}
