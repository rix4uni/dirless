package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/rix4uni/dirless/banner"
)

// Define structure for the regex config
type RegexConfig struct {
	Name     string `json:"name"`
	Match    string `json:"match"`
	Flags    string `json:"flags"`
	Maxcount int    `json:"maxcount"`
	Minkeep  int    `json:"minkeep"`
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Define the flags
	verbose := flag.Bool("verbose", false, "Enable verbose output.")
	noColor := flag.Bool("nc", false, "Disable colors in CLI output.")
	matchFile := flag.String("match", "", "Path to custom match.json file.")
	silent := flag.Bool("silent", false, "Silent mode.")
	versionFlag := flag.Bool("version", false, "Print the version of the tool and exit.")
	flag.Parse()

	if *versionFlag {
		banner.PrintBanner()
		banner.PrintVersion()
		return
	}

	if !*silent {
		banner.PrintBanner()
	}

	// ANSI color codes for highlighting
	const (
		highlightStart = "\033[1;31m" // Red bold
		highlightEnd   = "\033[0m"    // Reset
	)

	// Locate the match.json file
	var matchFilePath string
	if *matchFile != "" {
		// Use the custom match.json path provided via -match flag
		matchFilePath = *matchFile
	} else {
		// Check default locations
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home directory:", err)
			return
		}
		homeConfigPath := fmt.Sprintf("%s/.config/dirless/match.json", homeDir)
		pwdConfigPath := "match.json"

		// Check if the file exists in either default location
		if _, err := os.Stat(homeConfigPath); err == nil {
			matchFilePath = homeConfigPath
		} else if _, err := os.Stat(pwdConfigPath); err == nil {
			matchFilePath = pwdConfigPath
		} else {
			fmt.Println("Error: match.json not found in default locations.")
			return
		}
	}

	// Read the JSON file containing the regex configurations
	data, err := ioutil.ReadFile(matchFilePath)
	if err != nil {
		fmt.Println("Error reading match.json:", err)
		return
	}

	// Parse the JSON into regex configurations
	var regexConfigs []RegexConfig
	err = json.Unmarshal(data, &regexConfigs)
	if err != nil {
		fmt.Println("Error unmarshalling match.json:", err)
		return
	}

	// Read the input from stdin
	inputBytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	lines := strings.Split(string(inputBytes), "\n")

	// Use a map to track the domains and their random IDs
	seenDomains := make(map[string]string)

	// Track URLs that do not match any regex
	unmatchedURLs := make(map[string]bool)

	// Iterate over each regex config
	for _, config := range regexConfigs {
		// Compile the regular expression
		compiledRegex, err := regexp.Compile(config.Match)
		if err != nil {
			fmt.Println("Error compiling regex:", err)
			continue
		}

		// Iterate over each line in the input and apply the regex
		for _, line := range lines {
			if line == "" {
				continue
			}

			if compiledRegex.MatchString(line) {
				// Highlight the matched portion in the line
				var highlightedLine string
				if *noColor {
					highlightedLine = line
				} else {
					highlightedLine = compiledRegex.ReplaceAllString(line, highlightStart+"$0"+highlightEnd)
				}

				// Parse the URL to get the domain
				parsedURL, err := url.Parse(line)
				if err != nil {
					fmt.Println("Error parsing URL:", err)
					continue
				}

				// Get the domain and assign a random ID if not already assigned
				domain := parsedURL.Host
				randID, exists := seenDomains[domain]
				if !exists {
					randID = fmt.Sprintf("id%04d", rand.Intn(10000))
					seenDomains[domain] = randID
				}

				// Print the match or ignored status
				if *verbose {
					if exists {
						fmt.Printf("[%s] [%s] [ignored] %s\n", config.Name, randID, highlightedLine)
					} else {
						fmt.Printf("[%s] [%s] %s\n", config.Name, randID, highlightedLine)
					}
				} else if !exists {
					fmt.Println(highlightedLine)
				}

				// Mark this line as matched
				unmatchedURLs[line] = false
			} else {
				// Mark this line as unmatched
				if _, alreadyChecked := unmatchedURLs[line]; !alreadyChecked {
					unmatchedURLs[line] = true
				}
			}
		}
	}

	// Print unmatched URLs
	if *verbose {
		for line, isUnmatched := range unmatchedURLs {
			if isUnmatched && line != "" {
				fmt.Printf("[unmatched] %s\n", line)
			}
		}
	}
}
