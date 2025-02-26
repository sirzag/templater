package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sirzag/templater/internal/cmd_resolver"
	"github.com/sirzag/templater/internal/utils"
	"github.com/labstack/gommon/log"
	"github.com/urfave/cli/v3"
)

// var commands []*cli.Command = make([]*cli.Command, 0)

func main() {
	log.SetHeader("${time_rfc3339} ${level} ${short_file}:${line} ▶")

	templateDir := utils.GetTemplateDir()
	hasTemplates := utils.CheckTemplates(templateDir)
	if !hasTemplates {
		fmt.Printf("Please add template files to: %s\n", templateDir)
		os.Exit(0)
	}

	commands, err := cmd_resolver.BuildCommands()
	if err != nil {
		log.Fatal(err)
	}

	cmd := &cli.Command{
		Commands: commands,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

// {
// 				Name:    "add",
// 				Aliases: []string{"a"},
// 				Usage:   "add a task to the list",
// 				Action: func(ctx context.Context, cmd *cli.Command) error {
// 					fmt.Println("added task: ", cmd.Args().First())
// 					return nil
// 				},
// 			},
// 			{
// 				Name:    "complete",
// 				Aliases: []string{"c"},
// 				Usage:   "complete a task on the list",
// 				Action: func(ctx context.Context, cmd *cli.Command) error {
// 					fmt.Println("completed task: ", cmd.Args().First())
// 					return nil
// 				},
// 			},
// 			{
// 				Name:    "template",
// 				Aliases: []string{"t"},
// 				Usage:   "options for task templates",
// 				Commands: []*cli.Command{
// 					{
// 						Name:  "add",
// 						Usage: "add a new template",
// 						Action: func(ctx context.Context, cmd *cli.Command) error {
// 							fmt.Println("new task template: ", cmd.Args().First())
// 							return nil
// 						},
// 					},
// 					{
// 						Name:  "remove",
// 						Usage: "remove an existing template",
// 						Action: func(ctx context.Context, cmd *cli.Command) error {
// 							fmt.Println("removed task template: ", cmd.Args().First())
// 							return nil
// 						},
// 					},
// 				},
// 			},
