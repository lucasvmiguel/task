# Task

Task is a [CLI](https://en.wikipedia.org/wiki/Command-line_interface) that automates the software development process.

## Problem Statement

The software development process consists of several interactions with different softwares: reading an issue on Jira, creating a PR on Github, and so on.

It might be silly to spend time developing something to automate these processes, but the amount of time wasted could worth a solution like that.

## How it works

The CLI has different commands for different actions, below you can see how it works:

### Start to work on an issue

```
task start <ISSUE-ID>
```

Flow:
1. Assign the issue passed as a param on an issue tracker (eg: Jira) to you and also set to doing stage
2. Create a branch on the github project where you ran the `task` command
3. Create a pull request on the Git repository (eg: Github) with the information from the issue tracker and a tag with `work-in-progress`

### Finish to work on an issue

```
task finish <ISSUE-ID>
```

Flow:
1. Set the issue passed as param to verify stage on a issue tracker (eg: Jira)
2. Remove the label `work-in-progress` from the PR

## License

[MIT](LICENSE)