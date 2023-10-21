package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Note struct {
	Index int
	Body  string
}

type AllNotes struct {
	Notes []Note
}

/*
getDb retrieves the path to the database file, creating it if it doesn't exist.
It returns the database file path and any encountered error.
*/
func getDb() (string, error) {
	dbName := os.Getenv("HOME") + "/.gotepad_db.json"
	_, err := os.OpenFile(dbName, os.O_CREATE, 0644)
	if err != nil {
		return "", err
	}
	return dbName, nil
}

/*
getUserCommand parses command-line arguments and returns the user command and its associated value (if any).
The first argument is assumed to be the user command, and the second argument is its value.
If no value is provided, the second return value will be an empty string.
*/
func getUserCommand() (string, string) {
	var userCommand, value string
	for i, arg := range os.Args {
		if i == 1 {
			userCommand = arg
		} else if i == 2 {
			value = arg
		}
	}
	return userCommand, value
}

func addNote(note, db string) {
	// implement me
	fmt.Println("adding notes")
}

/*
listNotes reads and lists notes from the specified database file.
It reads the JSON data from the file, unmarshals it into a struct, and prints the notes.
It returns an error if there is a problem reading or unmarshaling the data.
*/
func listNotes(db string) error {
	f, err := os.ReadFile(db)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var data AllNotes
	err = json.Unmarshal(f, &data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for i, note := range data.Notes {
		fmt.Println(i, note.Index, note.Body)
	}
	return nil
}

func main() {

	// get db
	db, err := getDb()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// get command from user
	userCommand, value := getUserCommand()

	// handle command
	switch userCommand {
	case "add":
		addNote(value, db)
	case "ls":
		listNotes(db)
	default:
		fmt.Println("usage ./gotepad <add>|<ls>|<rm> <value>")
	}
}
