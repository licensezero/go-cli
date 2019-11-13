package subcommands

import "flag"
import "licensezero.com/cli/api"
import "licensezero.com/cli/data"
import "io/ioutil"

const sponsorDescription = "Sponsor relicensing."

// Sponsor starts an offer sponsorship transaction.
var Sponsor = &Subcommand{
	Tag:         "buyer",
	Description: sponsorDescription,
	Handler: func(args []string, paths Paths) {
		flagSet := flag.NewFlagSet("sponsor", flag.ExitOnError)
		doNotOpen := doNotOpenFlag(flagSet)
		offerID := offerIDFlag(flagSet)
		id := idFlag(flagSet)
		flagSet.SetOutput(ioutil.Discard)
		flagSet.Usage = sponsorUsage
		flagSet.Parse(args)
		if *offerID == "" && *id == "" {
			sponsorUsage()
		}
		if *offerID != "" && *id != "" {
			sponsorUsage()
		}
		if *offerID != "" {
			*id = *offerID
		}
		if !validID(*id) {
			invalidID()
		}
		identity, err := data.ReadIdentity(paths.Home)
		if err != nil {
			Fail(identityHint)
		}
		location, err := api.Sponsor(identity, *id)
		if err != nil {
			Fail("Error sending sponsor request: " + err.Error())
		}
		openURLAndExit(location, doNotOpen)
	},
}

func sponsorUsage() {
	usage := sponsorDescription + "\n\n" +
		"Usage:\n" +
		"  licensezero sponsor --id ID\n\n" +
		"Options:\n" +
		flagsList(map[string]string{
			"id ID": idLine,
		})
	Fail(usage)
}
