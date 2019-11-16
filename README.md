# courseraGoC2Week4Assignment

Assigment Week #4 - Functions, Methods, and Interfaces in Go - INTERFACES FOR ABSTRACTION  

- [Assignment](https://www.coursera.org/learn/golang-functions-methods/home/week/4)


- [Course: Functions, Methods, and Interfaces in Go](https://www.coursera.org/learn/golang-functions-methods)
  
## Instructions

The goal of this assignment is to write a GoLang routine that allows users to create a set of animals and then get information about those animals.

**Write** a program which allows the user to create a set of animals and to get information about those animals. Each animal has a name and can be either a cow, bird, or snake. With each command, the user can either create a new animal of one of the three types, or the user can request information about an animal that he/she has already created. Each animal has a unique name, defined by the user. Note that the user can define animals of a chosen type, but the types of animals are restricted to either cow, bird, or snake. The following table contains the three types of animals and their associated data.

| Animal	| Food eaten	| Locomotion method	| Spoken sound |
| :-------- | :------------ | :---------------- | :----------- |
| cow	    | grass	        | walk	            | moo          |
| bird	    | worms	        | fly	            | peep         |
| snake	    | mice          | slither	        | hsss         |


Your program should present the user with a prompt, “>”, to indicate that the user can type a request.    
Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line.     
Your program should continue in this loop forever.     

Every command from the user must be either a **“newanimal”** command or a **“query”** command:

Each **“newanimal”** command must be a single line containing **three strings**. 
- The first string is **“newanimal”**. 
- The second string is an arbitrary string which will be the **name of the new animal**. 
- The third string is the **type** of the new animal, either **“cow”**, **“bird”**, or **“snake”**. 

Your program should process each newanimal command by creating the new animal and printing **“Created it!”** on the screen.

Each **“query”** command must be a single line containing **3 strings**. 
- The first string is **“query”**. 
- The second string is the **name** of the animal. 
- The third string is the **name of the information requested** about the animal, either **“eat”**, **“move”**, or **“speak”**. 
Your program should process each query command by printing out the requested data.

Define an interface type called **Animal** which describes the methods of an animal. Specifically, the **Animal** interface should contain the methods **Eat()**, **Move()**, and **Speak()**, which take no arguments and return no values. The **Eat()** method should print the animal’s food, the **Move()** method should print the animal’s locomotion, and the **Speak()** method should print the animal’s spoken sound. Define three types **Cow**, **Bird**, and **Snake**. For each of these three types, define methods **Eat()**, **Move()**, and **Speak()** so that the types **Cow**, **Bird**, and **Snake** all satisfy the **Animal** interface. When the user creates an animal, create an object of the appropriate type. Your program should call the appropriate method when the user issues a query command.

## My assignment

Source code at `animals/animals.go`

## Sample compilation and first test execution

```sh
devel1@vbxdeb10mate:~$ cd $GOPATH/src/github.com/pssslearning/courseraGoC2Week4Assignment/animals/
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/courseraGoC2Week4Assignment/animals$ go run animals.go 

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
	
>newanimal lucera cow
Created it!
>newanimal tweetee bird
Created it!
>newanimal cobry snake
Created it!
>query lucera eat
grass
>query lucera move
walk
>query lucera speak
moo
>query tweetee eat
worms
>query tweetee move
fly
>query tweetee speak
peep
>query cobry eat
mice
>query cobry move
slither
>query cobry speak
hsss
>newanimal another cat
type of animal is not in the list of accepted types. Please try again
>query another eat
there is no animal stored with this name. Please try again
>delete tweetee bird
Command 'delete' unknown. Please enter 'newanimal' or 'query' commands
>one two
Bad request. Remember three strings are requested
>^Csignal: interrupt

```