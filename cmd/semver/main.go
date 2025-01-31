package main

import (
	"fmt"
	"os"

	"github.com/adlandh/semver-cli/internal/commands"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "semver",
		Short: "Semantic version cli",
		Long:  "Semantic version tool helper to validate and increase versions semantically",
	}
)

func init() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName(".semver")
	_ = viper.ReadInConfig()

	rootCmd.AddCommand(commands.NewInitCommand().Cmd())
	rootCmd.AddCommand(commands.NewUpVersionCommand().Cmd())
	rootCmd.AddCommand(commands.NewGetCommand().Cmd())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
