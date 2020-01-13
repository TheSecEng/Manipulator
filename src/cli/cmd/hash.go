package cmd

import (
	"fmt"
	"log"
	m "manipulator/manipulator"
	"math/rand"
	"os"
	"os/exec"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/spf13/cobra"
)

var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Hash a given string into various formats",
	RunE:  hashRun,
}

func hashRun(cmd *cobra.Command, args []string) error {
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
	results["Sha1"] = m.HandleSha1(query)
	results["Sha256"] = m.HandleSha256(query)
	results["Sha512"] = m.HandleSha512(query)
	results["MD5"] = m.HandleMD5(query)

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
	rootCmd.AddCommand(hashCmd)
}
