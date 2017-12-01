package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newRootCommand() *cobra.Command {

	root := newGoMakeCommand()
	root.AddCommand(
		newAllCommand(),
		newCleanCommand(),
		newGoFmtCommand(),
		newTargetsCommand(),
		newHelpCommand(),
		newTestCommand(),
	)

	return root
}

func main() {
	err := newRootCommand().Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
