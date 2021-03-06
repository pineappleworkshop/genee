package cmd

import (
	"errors"
	"fmt"

	"genee/services"

	"github.com/spf13/cobra"
)

var (
	projectCmd = &cobra.Command{
		Use:   "project",
		Short: "Generates a microservice project from a template",
		Long:  "Generates a microservice project from a template",
		Run: func(cmd *cobra.Command, args []string) {
			template, err := cmd.Flags().GetString(CMD_PROJECT_TEMPLATE_LONG)
			if err != nil {
				errExit(err)
			}
			if template == "" {
				err := errors.New("the `template / -t` flag is required")
				errExit(err)
			}

			destination, err := cmd.Flags().GetString(CMD_PROJECT_DESTINATION_LONG)
			if err != nil {
				errExit(err)
			}
			if destination == "" {
				err := errors.New("the `destination / -d` flag is required")
				errExit(err)
			}

			config, err := cmd.Flags().GetString(CMD_PROJECT_CONFIG_LONG)
			if err != nil {
				errExit(err)
			}
			if config == "" {
				err := errors.New("the `config / -c` flag is required")
				errExit(err)
			}

			project(template, destination, config)
		},
	}
)

func init() {
	projectCmd.Flags().StringP(
		CMD_PROJECT_TEMPLATE_LONG, CMD_PROJECT_TEMPLATE_SHORT, "", CMD_PROJECT_TEMPLATE_USAGE)
	projectCmd.Flags().StringP(
		CMD_PROJECT_DESTINATION_LONG, CMD_PROJECT_DESTINATION_SHORT, "", CMD_PROJECT_DESTINATION_USAGE)
	projectCmd.Flags().StringP(CMD_PROJECT_CONFIG_LONG, CMD_PROJECT_CONFIG_SHORT, "", CMD_PROJECT_CONFIG_USAGE)
	rootCmd.AddCommand(projectCmd)
}

func project(template, destination, config string) {
	fmt.Println(fmt.Sprintf("Generating a project from following directory: %s", template))
	fmt.Println(fmt.Sprintf("The generated project will be in the following directory: %s", destination))
	fmt.Println(fmt.Sprintf("Using the following configuration: %s", config))

	conf, err := services.ParseConfig(config)
	if err != nil {
		errExit(err)
	}

	if err := conf.SearchReplaceVars(); err != nil {
		errExit(err)
	}

	dirs, files, err := services.ParseTemplateDirectory(template)
	if err != nil {
		errExit(err)
	}

	if err := services.GenerateRoot(destination); err != nil {
		errExit(err)
	}

	if err := services.GenerateDirs(destination, dirs); err != nil {
		errExit(err)
	}

	if err := services.GenerateFiles(conf, template, destination, files); err != nil {
		errExit(err)
	}

	if err := services.RunCommands(destination, conf.Commands); err != nil {
		errExit(err)
	}
}
