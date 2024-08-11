package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Start() {
	fmt.Println("Starting database")
	for {
		var dbName string

		// get dbname
		if dbName == "" {
			dbName = "database-go"
		}

		printDbName(dbName)

		reader := bufio.NewScanner(os.Stdin)
		commandVal, _, _ := strings.Cut(input(reader), ";")

		if CheckMetaCommands(commandVal[0]) {
			result := handleMetaCommands(commandVal)
			printDbName(dbName)
			output(result)
			continue
		}

		result := handleSQLCommands(commandVal)
		printDbName(dbName)
		output(result)
	}

}

func printDbName(dbName string) {
	fmt.Printf("%s> ", dbName)
}

func input(r *bufio.Scanner) string {

	var text string
	for r.Scan() {
		line := r.Text()
		if strings.Contains(line, ";") {
			text += line
			break
		}
		fmt.Print("> ")
		text += line
		if CheckMetaCommands(text[0]) {
			break
		}
	}
	return text
}

func output(output string) {
	fmt.Println("out : ", output)
}

func CheckMetaCommands(commandFirstRune byte) bool {
	return commandFirstRune == '.'
}
