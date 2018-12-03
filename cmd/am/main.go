// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// OpenPitrix Identity Management service app.
package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	"openpitrix.io/iam/pkg/am/config"
	"openpitrix.io/iam/pkg/am/service"
	"openpitrix.io/iam/pkg/version"
	"openpitrix.io/logger"
)

func Main() {
	app := cli.NewApp()
	app.Name = "am"
	app.Usage = "am provides am service."
	app.Version = version.GetVersionString()

	app.UsageText = `am [global options] command [options] [args...]

EXAMPLE:
   am gen-config

   am info
   am can-do

   am list-role
   am list-role-binding

   am create-role
   am modifu-role
   am delete-role

   am create-binding
   am delete-binding

   am load-rbac
   am save-rbac

   am serve

   am tour`

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "config",
			Value:  "am-config.json",
			Usage:  "am config file",
			EnvVar: "OPENPITRIX_AM_CONFIG",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "debug",
			Usage:  "debug app",
			Hidden: true,

			Action: func(c *cli.Context) {
				fmt.Println(nil)
				fmt.Println(version.GetVersion())
				return
			},
		},

		{
			Name:  "serve",
			Usage: "run as drone service",
			Action: func(c *cli.Context) {
				cfg := config.MustLoadConf(c.GlobalString("config"))
				if cfg.AM.TlsEnabled {
					server, err := service.OpenServer(cfg.Mysql.DbType(), cfg.Mysql.GetUrl())
					if err != nil {
						logger.Criticalf(nil, "%v", err)
						os.Exit(1)
					}
					err = server.ListenAndServe(fmt.Sprintf(":%d", cfg.AM.Port))
					if err != nil {
						logger.Criticalf(nil, "%v", err)
						os.Exit(1)
					}
				} else {
					server, err := service.OpenServer(cfg.Mysql.DbType(), cfg.Mysql.GetUrl())
					if err != nil {
						logger.Criticalf(nil, "%v", err)
						os.Exit(1)
					}
					err = server.ListenAndServeTLS(
						fmt.Sprintf(":%d", cfg.AM.Port),
						cfg.AM.TlsCertFile,
						cfg.AM.TlsKeyFile,
					)
					if err != nil {
						logger.Criticalf(nil, "%v", err)
						os.Exit(1)
					}
				}
			},
		},

		{
			Name:  "tour",
			Usage: "show more examples",
			Action: func(c *cli.Context) {
				fmt.Println(tourTopic)
			},
		},
	}

	app.CommandNotFound = func(ctx *cli.Context, command string) {
		fmt.Fprintf(ctx.App.Writer, "not found '%v'!\n", command)
	}

	app.Run(os.Args)
}

const tourTopic = `
am gen-config
`
