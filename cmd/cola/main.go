package main

import (
	"log"
)

import (
	"github.com/spf13/cobra"
)

import (
	"github.com/wang1309/cola-go/cmd/cola/v1/internal/project"
)

var rootCmd = &cobra.Command{
	Use:     "cola",
	Short:   "Cola: An elegant toolkit for Go DDD microservices.",
	Long:    `Cola: An elegant toolkit for Go DDD microservices.`,
	Version: release,
}

func init() {
	rootCmd.AddCommand(project.CmdNew)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
