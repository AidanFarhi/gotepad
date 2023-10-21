package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Note struct {
	Index string
	Body  string
}

type AllNotes struct {
	Notes []Note
}

func getDb() (string, error) {
	dbName := os.Getenv("HOME") + "/.gotepad_db.json"
	_, err := os.OpenFile(dbName, os.O_CREATE, 0644)
	if err != nil {
		return "", err
	}
	return dbName, nil
}

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

func listNotes(db string) error {
	// TODO: fix this function
	f, err := os.ReadFile(db)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var data AllNotes
	json.Unmarshal(f, &data)
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
