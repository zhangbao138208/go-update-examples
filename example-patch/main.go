package main

import (
	"net/http"

	"github.com/inconshreveable/go-update"
	"flag"
	"fmt"
)

var version string

func doUpdate(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = update.Apply(
		resp.Body, update.Options{Patcher: update.NewBSDiffPatcher()})
	if err != nil {
		// error handling
	}
	return err
}

func main() {
	updateFlag := flag.Bool(
		"update", false,
		"Update to specified version")

	versionFlag := flag.String(
		"version", "",
		"Update to this version")

	flag.Parse()

	if *updateFlag {
		url := fmt.Sprintf(
			"http://localhost:8080/static/patch-%v-%v",
			version,
			*versionFlag)
		fmt.Println(fmt.Sprintf("Updating to %v", url))
		doUpdate(url)
	}

	fmt.Println(fmt.Sprintf("version %v", version))
}


