package subcommands

import "flag"

func doNotOpenFlag(flagSet *flag.FlagSet) *bool {
	return flagSet.Bool("do-not-open", false, "Do not open checkout page.")
}

func openFlag(flagSet *flag.FlagSet) *bool {
	return flagSet.Bool("open", false, openLine)
}

func idFlag(flagSet *flag.FlagSet) *string {
	return flagSet.String("id", "", idLine)
}

func offerIDFlag(flagSet *flag.FlagSet) *string {
	return flagSet.String("offer", "", offerIDLine)
}

func priceFlag(flagSet *flag.FlagSet) *uint {
	return flagSet.Uint("price", 0, priceLine)
}

func relicenseFlag(flagSet *flag.FlagSet) *uint {
	return flagSet.Uint("relicense", 0, relicenseLine)
}

func noRelicenseFlag(flagSet *flag.FlagSet) *bool {
	return flagSet.Bool("no-relicense", false, noRelicenseLine)
}

func silentFlag(flagSet *flag.FlagSet) *bool {
	return flagSet.Bool("silent", false, silentLine)
}
