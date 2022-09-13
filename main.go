package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/corenting/cookies/internal/decryption"
	"github.com/corenting/cookies/internal/entities"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:            "cookies",
		Usage:           "decrypt cookies content",
		Action:          decryptCommand,
		Version:         entities.AppVersion,
		Compiled:        time.Now(),
		UsageText:       "cookies [global options] cookie_to_decrypt",
		HideHelpCommand: true,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// The main command is the decrypt command tries to identify the cookie provided, decrypt it and print the result to the user.
func decryptCommand(context *cli.Context) error {
	if context.Args().Len() != 1 {
		return cli.Exit("You need to provide a cookie to decrypt.", 1)
	}

	decryptedCookie, err := decryption.DecryptCookie(context.Args().First())
	if err != nil {
		return cli.Exit("Error decoding the provided cookie, maybe it's not supported?", 1)
	}
	fmt.Println(decryptedCookie)

	return nil
}
