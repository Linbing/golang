package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"library"
	"mp"
)

var lib *library.MusicManager

var id int = 1
var ctrl, signal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
		}
	case "add":
		{
			if len(tokens) == 6 {
				id++
				lib.Add(&library.MusicEntry{strconv.Itoa(id),
					tokens[2], tokens[3], tokens[4], tokens[5]})
			} else {
				fmt.Println("USAGE: lib add <name> <artist> <source> <type>")
			}
		}
	case "remove":
		{
			if len(tokens) == 3 {
				//lib.Remove(tokens[2])
				fmt.Println(" ****** remove ******* ")
			} else {
				fmt.Println("USAGE: lib remove <name> ")
			}
		}
	default:
		fmt.Println("unrecognized lib command : ", tokens[1])
	}
}

func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}

	e := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("The Music", tokens[1], "does not exit.")
		return
	}

	mp.Play(e.Source, e.Type)
}

func start_music() {
	fmt.Printf(" ---------\n")
	fmt.Println(`Enter following commands to control the player:
				lib list -- View the exiting music lib
				lib add <name><artise><source><type> --Add a music to the music lib
				lib remove  <name> --Remove the specified music from lib
				play <name> --Play the specified music
				`)
	lib = library.NewMusicManager()
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command ->: ")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)
		if line == "q" || line == "e" {
			break
		}

		tokens := strings.Split(line, " ")

		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		} else {
			fmt.Println("Unrecongined command : ", tokens[0])
		}
	}

}
