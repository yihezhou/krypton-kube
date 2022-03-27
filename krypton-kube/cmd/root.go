package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:	"krypton-kube",
	Short:	"The krypton cli to interactive with k8s.",
	Long:	"The krypton cli to interactive with kubernetes API server.",
	Run:	func(cmd *cobra.Command, args []string) {
		fmt.Println("This is krypton-kube root cmd,add -h for help.")
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
