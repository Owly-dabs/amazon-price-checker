/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/Owly-dabs/amazon-price-checker/pkg/scraper"
	"github.com/spf13/cobra"
)

// Define flag variables here
var urlLink string
var maxWidth int = 10

func Truncate(s string, maxWidth int) string {
	if len(s) > maxWidth {
		return s[:maxWidth-3] + "..."
	}
	return s
}

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		itemName, err := scraper.GetItemName(urlLink)
		if err != nil {
			log.Fatalf("Error with scraping price from Amazon: %s", err)
		}

		price, err := scraper.GetPrice(urlLink)
		if err != nil {
			log.Fatalf("Error with scraping price from Amazon: %s", err)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)
		fmt.Fprintln(w, "Item Name\tPrice")
		fmt.Fprintf(w, "%s\t%s", Truncate(itemName, 20), price)
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	checkCmd.Flags().StringVarP(&urlLink, "url", "l", "", "Amazon URL of item")

}
