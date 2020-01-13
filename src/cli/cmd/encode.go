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

var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encode a given string into various formats",
	RunE:  encodeRun,
}

func encodeRun(cmd *cobra.Command, args []string) error {
	var encode bool

	query := strings.Join(args, " ")

	decode, err := cmd.Flags().GetBool("decode")
	if err != nil {
		return err
	}

	if !decode {
		encode = true
	} else {
		encode = false
	}

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

	if base64, err := m.HandleBase64(encode, query); err == nil && base64 != "" {
		results["Base64"] = base64
	} else {
		log.Println(err)
		// results["Base64"] = "Error processing Base64"
	}

	if url, err := m.HandleURL(encode, query); err == nil && url != "" {
		results["URL"] = url
	} else {
		log.Println(err)
		// results["URL"] = "Error processing URL"
	}

	if hex, err := m.HandleHex(encode, query); err == nil && hex != "" {
		results["Hex"] = hex
	} else {
		log.Println(err)
		// results["Hex"] = "Error processing Hex"
	}

	if binary, err := m.HandleBinary(encode, query); err == nil && binary != "" {
		results["Binary"] = binary
	} else {
		log.Println(err)
		// results["Binary"] = "Error processing Binary"
	}

	results["HTML"] = m.HandleHTML(encode, query)

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
	encodeCmd.Flags().BoolP("decode", "d", false, "Decode the provided string")
	rootCmd.AddCommand(encodeCmd)
}
