/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package create

import (
	"fmt"
	"github.com/spf13/cobra"
	cp "github.com/otiai10/copy"
)

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use:   "go",
	Short: "",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
		if (len(args) == 1) {
			filePath := args[0];
			err := cp.Copy("goTemplate", filePath);
			if (err != nil) {
				panic(err)
			}
			fmt.Println("Go backend template copied into: " + filePath)
		} else {
			fmt.Println("go called")
		}
	},
}

func init() {
}
