package cmd

import (
	"fmt"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "genee",
		Short: "A microservice project generator written in golang",
		Long:  `A microservice project generator written in golang`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	//rootCmd.PersistentFlags().StringVar(
	//	&cfgFile, CMD_CONFIG_L, "", "config file (default is $HOME/.cobra.yaml)")

	rootCmd.PersistentFlags().StringP(CMD_AUTHOR_LONG, CMD_AUTHOR_SHORT, CMD_AUTHOR_VALUE, CMD_AUTHOR_USAGE)
	rootCmd.PersistentFlags().StringVarP(&userLicense, CMD_LICENSE_LONG, CMD_LICENSE_SHORT, "", CMD_LICENSE_USAGE)

	//rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	//if err := viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author")); err != nil {
	//	errExit(err)
	//}
	//if err := viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper")); err != nil {
	//	errExit(err)
	//}
	//viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	//viper.SetDefault("license", "apache")

	//rootCmd.AddCommand(addCmd)
	//rootCmd.AddCommand(initCmd)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			errExit(err)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
