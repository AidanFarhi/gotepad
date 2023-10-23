package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// This represents a note.
type Note struct {
	Index int
	Body  string
}

// This represents all of the notes.
type AllNotes struct {
	Notes []Note
}

// getDb retrieves the path to the database file, creating it if it doesn't exist.
// It returns the database file path and any encountered error.
func getDb() (string, error) {
	dbName := os.Getenv("HOME") + "/.gotepad_db.json"
	_, err := os.OpenFile(dbName, os.O_CREATE, 0644)
	if err != nil {
		return "", err
	}
	return dbName, nil
}

// getUserCommand parses command-line arguments and returns the user command and its associated value (if any).
// The first argument is assumed to be the user command, and the second argument is its value.
// If no value is provided, the second return value will be an empty string.
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

// getNotes loads all notes from the json file.
func getNotes(note, db string) (*AllNotes, error) {
	f, err := os.ReadFile(db)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var data AllNotes
	err = json.Unmarshal(f, &data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &data, nil
}

// addNote adds a new note to the json file.
func addNote(newNote string, notes *AllNotes, db string) {
	// TODO: remove the hardcoded 3 here and implement some sort of check index logic
	notes.Notes = append(notes.Notes, Note{Index: 3, Body: newNote})
	data, _ := json.MarshalIndent(notes, "", " ")
	os.WriteFile(db, data, 0644)
}

// listNotes reads and lists notes from the specified database file.
// It reads the JSON data from the file, unmarshals it into a struct, and prints the notes.
// It returns an error if there is a problem reading or unmarshaling the data.
func listNotes(notes *AllNotes) {
	for _, note := range notes.Notes {
		fmt.Println(note.Index, note.Body)
	}
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
		notes, _ := getNotes(value, db)
		addNote(value, notes, db)
	case "ls":
		notes, _ := getNotes(value, db)
		listNotes(notes)
	default:
		fmt.Println("usage ./gotepad <add>|<ls>|<rm> <value>")
	}
}
