package subcommands

import "encoding/json"
import "errors"
import "flag"
import "io/ioutil"
import "os"
import "strings"

const readmeDescription = "Append licensing information to README."

var README = Subcommand{
	Tag:         "seller",
	Description: readmeDescription,
	Handler: func(args []string, paths Paths) {
		flagSet := flag.NewFlagSet("readme", flag.ExitOnError)
		silent := Silent(flagSet)
		flagSet.SetOutput(ioutil.Discard)
		flagSet.Usage = readmeUsage
		flagSet.Parse(args)
		var existing string
		data, err := ioutil.ReadFile("README.md")
		if err != nil {
			if os.IsNotExist(err) {
				existing = ""
			} else {
				os.Stderr.WriteString("Error reading README.md.\n")
				os.Exit(1)
			}
		} else {
			existing = string(data)
		}
		projectIDs, termsIDs, err := readEntries(paths.CWD)
		if err != nil {
			os.Stderr.WriteString(err.Error() + "\n")
			os.Exit(1)
		}
		if len(existing) > 0 {
			existing = existing + "\n\n"
		}
		existing = existing + "# Licensing"
		if len(projectIDs) == 0 {
			os.Stderr.WriteString("No License Zero project metadata in package.json.\n")
			os.Exit(1)
		}
		haveReciprocal := false
		haveNoncommercial := false
		haveParity := false
		haveProsperity := false
		for _, terms := range termsIDs {
			if terms == "noncommercial" {
				haveNoncommercial = true
			} else if terms == "reciprocal" {
				haveReciprocal = true
			} else if terms == "parity" {
				haveParity = true
			} else if terms == "prosperity" {
				haveProsperity = true
			}
		}
		multiple := twoOrMore([]bool{haveReciprocal, haveNoncommercial, haveParity, haveProsperity})
		var licenseScope string
		if multiple {
			licenseScope = "Some contributions to this package "
		} else {
			licenseScope = "This package "
		}
		summaries := []string{}
		availabilities := []string{}
		if haveReciprocal {
			summaries = append(
				summaries,
				licenseScope+
					"is free to use in open source under the terms of "+
					"the [License Zero Reciprocal Public License](./LICENSE).",
			)
			availabilities = append(
				availabilities,
				"Licenses for use in closed software "+
					"are available via [licensezero.com](https://licensezero.com).",
			)
		} else if haveNoncommercial {
			summaries = append(
				summaries,
				licenseScope+
					"is free to use for commercial purposes for a trial period under the terms of "+
					"the [License Zero Noncommercial Public License](./LICENSE).",
			)
			availabilities = append(
				availabilities,
				"Licenses for long-term commercial use "+
					"are available via [licensezero.com](https://licensezero.com).",
			)
		} else if haveParity {
			summaries = append(
				summaries,
				licenseScope+
					"is free to use in open source under the terms of "+
					"[Parity Public License](./LICENSE).",
			)
			availabilities = append(
				availabilities,
				"Licenses for use in closed software "+
					"are available via [licensezero.com](https://licensezero.com).",
			)
		} else if haveProsperity {
			summaries = append(
				summaries,
				licenseScope+
					"is free to use for commercial purposes for a trial period under the terms of "+
					"[The Prosperity Public License](./LICENSE).",
			)
			availabilities = append(
				availabilities,
				"Licenses for long-term commercial use "+
					"are available via [licensezero.com](https://licensezero.com).",
			)
		} else {
			os.Stderr.WriteString("Unrecognized License Zero project terms.\n")
			os.Exit(1)
		}
		existing = existing + "\n\n" + strings.Join(summaries, "\n\n")
		existing = existing + "\n\n" + strings.Join(availabilities, "\n\n")
		for _, projectID := range projectIDs {
			projectLink := "https://licensezero.com/projects/" + projectID
			badge := "" +
				"[" +
				"![licensezero.com pricing](" + projectLink + "/badge.svg)" +
				"]" +
				"[" + projectLink + "]"
			existing = existing + "\n\n" + badge
		}
		err = ioutil.WriteFile("README.md", []byte(existing), 0644)
		if err != nil {
			os.Stderr.WriteString("Error writing README.md.\n")
			os.Exit(1)
		}
		if !*silent {
			os.Stdout.WriteString("Wrote to README.md\n")
		}
		os.Exit(0)
	},
}

func twoOrMore(values []bool) bool {
	counter := 0
	for _, value := range values {
		if value {
			counter++
		}
		if counter == 2 {
			return true
		}
	}
	return false
}

func readEntries(directory string) ([]string, []string, error) {
	data, err := ioutil.ReadFile("package.json")
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil, errors.New("Could not read package.json.")
		} else {
			return nil, nil, err
		}
	}
	var existingMetadata struct {
		LicenseZero []struct {
			License struct {
				ProjectID string `json:"projectID"`
				Terms     string `json:"terms"`
			} `json:"license"`
			AgentSignature    string `json:"agentSignature"`
			LicensorSignature string `json:"licensorSignature"`
		} `json:"licensezero"`
	}
	err = json.Unmarshal(data, &existingMetadata)
	if err != nil {
		return nil, nil, errors.New("Could not parse package.json metadata.")
	}
	// TODO: Validate package.json metadata entries.
	var projectIDs []string
	var terms []string
	for _, entry := range existingMetadata.LicenseZero {
		projectIDs = append(projectIDs, entry.License.ProjectID)
		terms = append(terms, entry.License.Terms)
	}
	return projectIDs, terms, nil
}

func readmeUsage() {
	usage := readmeDescription + "\n\n" +
		"Usage:\n" +
		"  licensezero readme\n\n" +
		"Options:\n" +
		flagsList(map[string]string{
			"silent": silentLine,
		})
	os.Stderr.WriteString(usage)
	os.Exit(1)
}
