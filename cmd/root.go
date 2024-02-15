/*
Copyright © 2024 Nikhil Bhutani
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

const SEARCH_API = ""

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "star-repos",
	Short: "Search for the most starred GitHub repositories within a date range.",
	Long: `A simple CLI application that fetches the most starred GitHub repositories within a date range. For example:
		   star-repos <DATE_FROM> <DATE_TO>`,
	Run: search,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}

func search(cmd *cobra.Command, args []string) {
	var url string
	t := time.Now()
	today := fmt.Sprintf("%v", t.Format("2006-01-02"))
	if len(args) <= 1 {
		url = fmt.Sprintf("https://api.github.com/search/repositories?q=created:>=%s&sort=stars&order=desc&page=1", today)
	} else if len(args) == 2 {
		fromDate := args[1]
		url = fmt.Sprintf("https://api.github.com/search/repositories?q=created:>=%v&sort=stars&order=desc&page=1", fromDate)
	} else {
		fromDate, toDate := args[1], args[2]
		url = fmt.Sprintf("https://api.github.com/search/repositories?q=created:%s..%s&sort=stars&order=desc&page=1", fromDate, toDate)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error creating a the request")
	}
	req.Header = http.Header{
		"Accept":               {"application/vnd.github+json"},
		"Authorization":        {"Bearer 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"},
		"X-GitHub-Api-Version": {"2022-11-28"},
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error calling the API.")
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var result map[string]interface{}
	json.Unmarshal(resBody, &result)

	// Marshal the map back into JSON with indentation
	prettyJSON, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	// Print the pretty JSON
	fmt.Println(string(prettyJSON))
}
