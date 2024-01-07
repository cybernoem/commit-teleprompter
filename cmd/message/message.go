package message

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/exp/slices"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var baseKeywords = []string{
	"feat",
	"fix",
	"docs",
	"style",
	"refactor",
	"test",
	"chore",
	"build",
	"ci",
}

var Cmd = &cobra.Command{
	Use:     "message",
	Short:   "Message templates",
	Aliases: []string{"msg", "m"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println(generateCommitExample(args[0]))
		} else {
			listTemplates()
		}
	},
}

func listTemplates() {
	for _, item := range baseKeywords {
		fmt.Println(generateCommitExample(item))
	}
}

func generateCommitExample(keyword string) string {
	if !slices.Contains(baseKeywords, keyword) {
		log.Fatalf("Error: Unkown keyboard: %s", keyword)
		os.Exit(404)
	}

	var issueKey = viper.Get("IssueKey")
	if viper.Get("IssueKey") != "" {
		issueKey = fmt.Sprintf("%s - ", issueKey)
	}
	return fmt.Sprintf("%s: %s", keyword, issueKey)
}
