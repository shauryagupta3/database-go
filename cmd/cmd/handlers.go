package cmd

import (
	"fmt"
	"os"
)

func handleMetaCommands(cmd string) string {
	switch cmd {
	case ".exit":
		fmt.Println("Exiting DB")
		os.Exit(0)
	default:
		return fmt.Sprintf("unrecognized meta-data command : %s \n", cmd)
	}
	return ""
}
func handleSQLCommands(cmd string) string {
	switch cmd {
	case "CREATE":
		return fmt.Sprintf("Exiting DB %s", cmd)
	default:
		return fmt.Sprintf("unrecognized SQL command : %s \n", cmd)
	}
	return ""
}
