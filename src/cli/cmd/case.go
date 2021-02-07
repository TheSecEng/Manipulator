package cmd

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"

	m "github.com/theseceng/manipulator/manipulator"

	aw "github.com/deanishe/awgo"
	"github.com/spf13/cobra"
)

var caseCmd = &cobra.Command{
	Use:   "case",
	Short: "Convert the case of a string",
	RunE:  caseRun,
}

func caseRun(cmd *cobra.Command, args []string) error {
	query := strings.Join(args, " ")

	log.Printf("query=%s", query)

	// Call self with "check" command if an update is due and a check
	// job isn't already running.
	if wf.UpdateCheckDue() && !wf.IsRunning(updateJobName) {
		log.Println("Running update check in background...")

		cmd := exec.Command(os.Args[0], "-check")
		if err := wf.RunInBackground(updateJobName, cmd); err != nil {
			log.Printf("Error starting update check: %s", err)
		}
	}

	// Only show update status if query is empty.
	if query == "" && wf.UpdateAvailable() {
		wf.Configure(aw.SuppressUIDs(true))

		wf.NewItem("Update available!").
			Subtitle("↩ to install").
			Autocomplete("workflow:update").
			Valid(false).
			Icon(iconAvailable)
	}

	results := make(map[string]string)
	results["Uppercase"] = m.HandleUpper(query)
	results["Lowercase"] = m.HandleLower(query)
	results["Snakecase"] = m.HandleSnake(query)
	results["CamelCase"] = m.HandleCamel(query)
	results["LowerCamelCase"] = m.HandleLowerCamel(query)
	results["TitleCase"] = m.HandleTitleCase(query)

	for key, result := range results {
		wf.NewItem(result).
			Subtitle(fmt.Sprintf("%s Encoded String of %s", key, query)).
			Arg(result).
			UID(fmt.Sprintf("%d", rand.Int())).
			Valid(true)
	}

	// Send results to Alfred
	wf.SendFeedback()
	return nil
}

func init() {
	rootCmd.AddCommand(caseCmd)
}
