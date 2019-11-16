package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Animal an interface with three main behaviours
type Animal interface {
	Eat() string   // Eat prints out the Animal's kind of food it eats
	Move() string  // Move prints out the Animal's kind of locomotion it uses
	Speak() string // Speak prints out the Animal's kind of noise it produces
}

// ======================================================
// COW
// ======================================================

// A Cow Animal
type Cow struct {
	name string
}

// Associated Methods

// Eat prints out the Animal's kind of food it eats
func (c Cow) Eat() string {
	return "grass"
}

// Move prints out the Animal's kind of locomotion it uses
func (c Cow) Move() string {
	return "walk"
}

// Speak prints out the Animal's kind of noise it produces
func (c Cow) Speak() string {
	return "moo"
}

// Name prints out the Animal's Name
func (c Cow) Name() string {
	return c.name
}

// ======================================================
// BIRD
// ======================================================

// A Bird Animal
type Bird struct {
	name string
}

// Associated Methods

// Eat prints out the Animal's kind of food it eats
func (c Bird) Eat() string {
	return "worms"
}

// Move prints out the Animal's kind of locomotion it uses
func (c Bird) Move() string {
	return "fly"
}

// Speak prints out the Animal's kind of noise it produces
func (c Bird) Speak() string {
	return "peep"
}

// Name prints out the Animal's Name
func (c Bird) Name() string {
	return c.name
}

// ======================================================
// SNAKE
// ======================================================

// A Snake Animal
type Snake struct {
	name string
}

// Associated Methods

// Eat prints out the Animal's kind of food it eats
func (c Snake) Eat() string {
	return "mice"
}

// Move prints out the Animal's kind of locomotion it uses
func (c Snake) Move() string {
	return "slither"
}

// Speak prints out the Animal's kind of noise it produces
func (c Snake) Speak() string {
	return "hsss"
}

// Name prints out the Animal's Name
func (c Snake) Name() string {
	return c.name
}

// ===============================================================
// Main
// ===============================================================
func main() {

	showPurposeAndInstructions()

	am := make(map[string]Animal)

	for {
		command, name, typeOrInfo := getUserRequest()

		switch command {
		case "newanimal":
			err := createNewAnimal(name, typeOrInfo, am)
			if err != nil {
				fmt.Printf("%s. Please try again\n", err)
			} else {
				fmt.Println("Created it!")
			}
		case "query":
			response, err := doQueryAction(name, typeOrInfo, am)
			if err != nil {
				fmt.Printf("%s. Please try again\n", err)
			} else {
				fmt.Println(response)
			}
		default:
			fmt.Printf("Command '%s' unknown. Please enter 'newanimal' or 'query' commands\n", command)
		}

	}
}

func createNewAnimal(theName, theType string, db map[string]Animal) error {

	if _, existing := db[theName]; existing {
		return errors.New("there is already an animal stored with the same name")
	}

	switch theType {
	case "cow":
		db[theName] = Cow{theName}
	case "bird":
		db[theName] = Bird{theName}
	case "snake":
		db[theName] = Snake{theName}
	default:
		return errors.New("type of animal is not in the list of accepted types")
	}

	return nil
}

func doQueryAction(theName, theQueryName string, db map[string]Animal) (string, error) {

	animal, existing := db[theName]

	if !existing {
		return "", errors.New("there is no animal stored with this name")
	}

	switch theQueryName {
	case "eat":
		return animal.Eat(), nil
	case "move":
		return animal.Move(), nil
	case "speak":
		return animal.Speak(), nil
	default:
		return "", errors.New("type of animal is not in the list of accepted types")
	}
}

func showPurposeAndInstructions() {
	pai := `
Purpose:
	This program allows you to create a set of animals and to get information about those animals. 
	Each animal has a name and can be either a cow, bird, or snake.   
	With each command, you can either create a new animal of one of the three types, 
	or you can request information about an animal that you have already created. 
	Each animal has a unique name, defined by you but note that you can define animals of a chosen type, 
	but the types of animals are restricted to either cow, bird, or snake. 

Instructions:
	This program will present you with a prompt, “>”, to indicate that you can type a request. 
	This program will accept one command at a time, print out a response, and print out a new prompt on a new line. 
	This program will continue in this loop forever until you press <CTRL+C>. 
	
	Every command must be either a “newanimal” command or a “query” command.

	Each “newanimal” command must be a single line containing 3 strings:
	- The first string is “newanimal”. 
	- The second string is an arbitrary string which will be the name of the new animal. 
	- The third string is the type of the new animal, either “cow”, “bird”, or “snake”. 
	The program will process each newanimal command by creating the new animal and printing “Created it!” on the screen.

	Each “query” command must be a single line containing 3 strings. 
	- The first string is “query”. 
	- The second string is the name of the animal. 
	- The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”. 
	The program will process each query command by printing out the requested data.
	`
	fmt.Println(pai)
}

func getUserRequest() (string, string, string) {

	in := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">")
		in.Scan()
		scanned := in.Text()
		fields := strings.Fields(scanned)

		if len(fields) != 3 {
			fmt.Println("Bad request. Remember three strings are requested")
			continue
		} else {
			return strings.ToLower(fields[0]), strings.ToLower(fields[1]), strings.ToLower(fields[2])
		}
	}
}
