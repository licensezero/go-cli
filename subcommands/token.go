package subcommands

import "flag"
import "github.com/licensezero/cli/data"
import "os"

const tokenDescription = "Set your licensezero.com licensor ID and access token."

var Token = Subcommand{
	Description: tokenDescription,
	Handler: func(args []string, paths Paths) {
		flagSet := flag.NewFlagSet("token", flag.ExitOnError)
		licensorID := flagSet.String("licensor-id", "", "Licensor ID")
		flagSet.Usage = tokenUsage
		flagSet.Parse(args)
		if *licensorID == "" {
			tokenUsage()
		}
		token := SecretPrompt("Token: ")
		newLicensor := data.Licensor{
			LicensorID: *licensorID,
			Token:      token,
		}
		existingLicensor, _ := data.ReadLicensor(paths.Home)
		if existingLicensor != nil && *existingLicensor != newLicensor {
			if !Confirm("Overwrite existing licensor info?") {
				os.Exit(0)
			}
		}
		err := data.WriteLicensor(paths.Home, &newLicensor)
		if err != nil {
			os.Stderr.WriteString("Could not write licensor file.\n")
			os.Exit(1)
		} else {
			os.Exit(0)
		}
	},
}

func tokenUsage() {
	usage := tokenDescription + "\n\n" +
		"Usage:\n" +
		"  licensezero token --licensor-id ID\n\n" +
		"Options:\n" +
		"  --licensor-id ID  Licensor ID (UUID).\n"
	os.Stderr.WriteString(usage)
	os.Exit(1)
}
