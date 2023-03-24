package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/fatih/color"
	"github.com/robfig/cron/v3"
)

// Request represents a curl request to the SAIME's website and returns a boolean value
// indicating whether the page is avialable or not.
func Request() bool {
	url := "https://siic.saime.gob.ve"
	curl := exec.Command("curl", url)

	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	_, err := curl.Output()
	if err != nil {
		fmt.Printf("La pagina del SAIME se encuentra: %s\n", red("Offline"))
		return false
	} else {
		fmt.Printf("La pagina del SAIME se encuentra: %s\n", green("Online"))
		return true
	}
}

func main() {
	// Create a job to check each 5 minutes if the SAIME's website is available
	c := cron.New()
	_, err := c.AddFunc("@every 5m", func() {
		Request()
	})

	if err != nil {
		log.Fatal(err)
	}

	Request()
	c.Run()
}
