package main

import (
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli"
)

// Version set at compile-time
var Version string

func main() {
	app := cli.NewApp()
	app.Name = "xmpp plugin"
	app.Usage = "xmpp plugin"
	app.Action = run
	app.Version = Version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "xmpp.host",
			Usage:  "xmpp host",
			EnvVar: "PLUGIN_XMPP_HOST,XMPP_HOST",
		},
		cli.StringFlag{
			Name:   "xmpp.jid",
			Usage:  "xmpp jid",
			EnvVar: "PLUGIN_XMPP_JID,XMPP_JID",
		},
		cli.StringFlag{
			Name:   "xmpp.password",
			Usage:  "xmpp password",
			EnvVar: "PLUGIN_XMPP_PASSWORD,XMPP_PASSWORD",
		},
		cli.StringSliceFlag{
			Name:   "xmpp.to",
			Usage:  "send message to user",
			EnvVar: "PLUGIN_XMPP_TO,XMPP_TO",
		},
		cli.StringSliceFlag{
			Name:   "message",
			Usage:  "xmpp message",
			EnvVar: "PLUGIN_MESSAGE",
		},
		cli.StringFlag{
			Name:   "repo.owner",
			Usage:  "repository owner",
			EnvVar: "DRONE_REPO_OWNER",
		},
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "commit.sha",
			Usage:  "git commit sha",
			EnvVar: "DRONE_COMMIT_SHA",
		},
		cli.StringFlag{
			Name:   "commit.branch",
			Value:  "master",
			Usage:  "git commit branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
		},
		cli.StringFlag{
			Name:   "commit.author",
			Usage:  "git author name",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.StringFlag{
			Name:   "commit.author.email",
			Usage:  "git author email",
			EnvVar: "DRONE_COMMIT_AUTHOR_EMAIL",
		},
		cli.StringFlag{
			Name:   "commit.message",
			Usage:  "commit message",
			EnvVar: "DRONE_COMMIT_MESSAGE",
		},
		cli.StringFlag{
			Name:   "build.event",
			Value:  "push",
			Usage:  "build event",
			EnvVar: "DRONE_BUILD_EVENT",
		},
		cli.IntFlag{
			Name:   "build.number",
			Usage:  "build number",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.StringFlag{
			Name:   "build.status",
			Usage:  "build status",
			Value:  "success",
			EnvVar: "DRONE_BUILD_STATUS",
		},
		cli.StringFlag{
			Name:   "build.link",
			Usage:  "build link",
			EnvVar: "DRONE_BUILD_LINK",
		},
		cli.StringFlag{
			Name:   "build.tag",
			Usage:  "build tag",
			EnvVar: "DRONE_TAG",
		},
		cli.Float64Flag{
			Name:   "job.started",
			Usage:  "job started",
			EnvVar: "DRONE_JOB_STARTED",
		},
		cli.Float64Flag{
			Name:   "job.finished",
			Usage:  "job finished",
			EnvVar: "DRONE_JOB_FINISHED",
		},
		cli.StringFlag{
			Name:   "env-file",
			Usage:  "source env file",
			EnvVar: "ENV_FILE",
		},
	}

	app.Run(os.Args)
}

func run(c *cli.Context) error {
	if c.String("env-file") != "" {
		_ = godotenv.Load(c.String("env-file"))
	}

	plugin := Plugin{
		Repo: Repo{
			Owner: c.String("repo.owner"),
			Name:  c.String("repo.name"),
		},
		Build: Build{
			Tag:      c.String("build.tag"),
			Number:   c.Int("build.number"),
			Event:    c.String("build.event"),
			Status:   c.String("build.status"),
			Commit:   c.String("commit.sha"),
			Branch:   c.String("commit.branch"),
			Author:   c.String("commit.author"),
			Email:    c.String("commit.author.email"),
			Message:  c.String("commit.message"),
			Link:     c.String("build.link"),
			Started:  c.Float64("job.started"),
			Finished: c.Float64("job.finished"),
		},
		Config: Config{
			Host:       c.String("xmpp.host"),
			Jid:        c.String("xmpp.jid"),
			Password:   c.String("xmpp.password"),
			To:         c.StringSlice("xmpp.to"),
			Message:    c.StringSlice("message"),
		},
	}

	return plugin.Exec()
}
