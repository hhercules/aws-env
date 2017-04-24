package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-ini/ini"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func credentialsFilename() string {
	return filepath.Join(userHomeDir(), ".aws", "credentials")
}

func configFilename() string {
	return filepath.Join(userHomeDir(), ".aws", "config")
}

func userHomeDir() string {
	homeDir := os.Getenv("HOME") // *nix
	if len(homeDir) == 0 {       // windows
		homeDir = os.Getenv("USERPROFILE")
	}

	return homeDir
}

func main() {
	fmt.Println("So it begins")

	cfg, err := ini.Load(credentialsFilename())
	check(err)

	// sections := cfg.Sections()
	profiles := cfg.SectionStrings()

	// fmt.Println(sections)
	// fmt.Printf("%v", profiles)

	for _, profile := range profiles {
		if profile != "DEFAULT" {
			fmt.Printf("%s\n", profile)
		}
	}
}
