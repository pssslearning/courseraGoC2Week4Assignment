/*
Write a program which allows the user to create a set of animals and to get information about those animals. Each animal has a name and can be either a cow, bird, or snake. With each command, the user can either create a new animal of one of the three types, or the user can request information about an animal that he/she has already created. Each animal has a unique name, defined by the user. Note that the user can define animals of a chosen type, but the types of animals are restricted to either cow, bird, or snake.
https://www.coursera.org/learn/golang-functions-methods/peer/gurxW/module-4-activity/submit
*/

package main

import (
       "os"
       "fmt"
       "bufio"
       "strings"
)

type Animal interface {
     Eat()
     Move()
     Speak()
}

type Cow struct {
     food string
     locomotion string
     noise string
     }

type Bird struct {
     food string
     locomotion string
     noise string
     }

type Snake struct {
     food string
     locomotion string
     noise string
     }

func ( c Cow ) Eat() {
     fmt.Println( c.food );
}

func ( c Cow ) Move() {
     fmt.Println( c.locomotion );
}

func ( c Cow ) Speak() {
     fmt.Println( c.noise );
}

func ( b Bird ) Eat() {
     fmt.Println( b.food );
}

func ( b Bird ) Move() {
     fmt.Println( b.locomotion );
}

func ( b Bird ) Speak() {
     fmt.Println( b.noise );
}

func ( s Snake ) Eat() {
     fmt.Println( s.food );
}

func ( s Snake ) Move() {
     fmt.Println( s.locomotion );
}

func ( s Snake ) Speak() {
     fmt.Println( s.noise );
}

func getInputReturnString() string {
     // This function is to take just one word as input
     // and return it after converting to a needed datatype if required.
     scanner := bufio.NewScanner( os.Stdin );
     scanner.Scan();
     inStr := scanner.Text();

     return inStr;
}

func main() {
     var animalInfo map[string]Animal;
     animalInfo = make(map[string]Animal);
     for {
     	 fmt.Printf( ">" )
     	 inpStr := getInputReturnString();
     	 inpArray := strings.Split( inpStr, " " );
     	 ques := inpArray[ 0 ];
	 name := inpArray[ 1 ];
	 info := inpArray[ 2 ];
	 if ques == "newanimal" {
	    if info == "cow" {
	       	    animalInfo[ name ] = Cow{ food: "grass", locomotion:"walk", noise:"moo"};
	    } else if info == "bird" {
	            animalInfo[ name ] = Bird{ food: "worms", locomotion:"fly", noise:"peep"};
	    } else if info == "snake" {
	    	    animalInfo[ name ] = Snake{ food: "mice", locomotion:"slither", noise:"hsss"};
	    }
	    fmt.Println( "Created it!" )
	 } else if ques == "query" {
	   	if info == "eat" {	
		animalInfo[ name ].Eat();
	    } else if info == "move" {
	       animalInfo[ name ].Move()
	    } else if info == "speak" {
	       animalInfo[ name ].Speak()
	    }
	 }
     	 
     }
}