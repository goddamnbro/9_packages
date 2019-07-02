package main

import (
	"9_packages/homework"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var cmdCompress = &cobra.Command{
		Use:   "compress [string to compress]",
		Short: "Compress archive your string to a short one",
		Long:  "Compress archive your string to a short one",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			rawString := strings.Join(args, "")
			fmt.Printf("Compressed: %s \n", homework.CompressWord(rawString))
		},
	}

	var cmdGetPopularWords = &cobra.Command{
		Use:   "popular [string]",
		Short: "Popular is looking for the most popular words in a string",
		Long:  "Popular is looking for the most popular words in a string",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			rawString := strings.Join(args, " ")
			popular := homework.GetFrequencyWords(rawString)
			fmt.Printf("Polular: %s \n", strings.Join(popular, ", "))
		},
	}

	var rootCmd = &cobra.Command{}
	rootCmd.AddCommand(cmdCompress)
	rootCmd.AddCommand(cmdGetPopularWords)
	rootCmd.Execute()
}
