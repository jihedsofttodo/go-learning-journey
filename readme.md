# Golang Beginners

## What Go ? Why Go ? How it is different ?

- developed at Google in 2007.
- open-sourced in 2009.

## Go Use Cases ?

- ### Evolution of Infrastructure:


  - infrastructure changed a lot:
    - Multicode processors.
    - Cloud Infrastructure.
    - Big Networked Computaion Clusters
  - scalable & distributed.
  - dynamic
  - more capacity

  -> Existing programming languages did not fully take advantage of it.

  -> basically you had applications that would just execute one task at a time in order but with infrastructure improvements it was possible to now write applications that would execute tasks in parallel , do multiple things at once.

  this way making the application faster and more user-friendly a simple example is you are using google drive you may be uploading or downloading files and folders but you can keep navigating back and forth in the UI. ( downloading , uploading , navigating UI all happen in parallel ). this is the concept of **multithreading**.

  -> this is make application fast , may also cause some issues , for example in google docs , many users can work on the same document at the same time ,

  when you have two users changing and adding stuff at the same time to the same document this should work smoothly without one user overriding all the changes that another user is making.

  whenn things are processed in parallel is a booking system for buying tickets or booking a hotel etc

  two users are trying to book at last remaining ticket at the same time of course this should work in a way that no double booking happens and this concept is called concurrency and needs to be handled by developers in code.they must write code that prevents any conflicts between the tasks when multiple tasks running in parallel and updating the same data.

  - **Concurrency** is about dealing with lots of things at once.
  - Developers need to write code to prevent conflicts.
- ### Multi-Core Concurrency Support

many languages do have features for implementing such applications however the code can get pretty complex and handling and preventing the concurrency issues can be pretty hard and with complexity of course there is always a higher risk of human errors and that's where the main purpose and difference of go comes into the picture so go was designed exactly for that purpose to make writing multi-threaded concurrent applications that take advantage of the new performance infrastructure much easier:

* Go was **designed to run on multiple cores** and built to support concurrency.
* Concurrency in Go in **cheap** and **easy**.

-> main use case of go or what it's best used for is writing applications that need to be very performant and will run on the modern scaled and distributed infrastructure with hundreds and thousands of servers typically on a cloud platform.

Go Language Characterics:

Go attempt to combine both:

* Simple and readable syntax of a dynamically typed language like Python.
* Efficiency and safety of a lower-level, statically typed language like C++.

Go is used on the server side or backend side of the applications and these types of applications can range from microservices and web applications to database services and so on. in fact many cloud technologies or technologies that run on modern cloud environments are actually written in go like docker hashicorp vault, kubernetes , cockroach db

Go advantages :

- simple syntax : easy to learn, read and write code.
  language features left out on purpose for simplicity ( Goal: easy to maintain over time )
- Fast build time, start up and run
- requires fewer resources.
- compiled language. ( compiles into single binary ( machine code ) )
- Faster than interpreted languages , like Python
- Consistent across different OS.

## Local Setup - Install GO & Editor

1 - Install Go : go distribution actually comes with a Go CLI tool.

- https://go.dev/doc/install
  - type go in the terminal and check if there is an output to make sure that you installed go compiler correclty.

2 - Install an IDE - Editor for writing code ( VS Code ).

3 - Install Go extension.

## First Program with Go

if you try to put an instruction for example Print('Hello World'); , an error appears that complains about go.mod

the first thing we need to fix is to make our go application into a project and for that we need to basically initialize it , that's the error that we get here and to do that

using the terminal

`$>go mod init  booking-app`

- Initialized a go.mod file.
- ```go
  module booking-app

  go 1.23.0

  ```
- Describes the module : with name/module path and go version used in the program.
- The module path is also the import path ( e.g import "github.com/jihed/booking-app")

the issue is fixed still to fix the package issue:

* in Go everything is organized into packages
* also go need to know where to start the program ? where is the entrypoint?
  * main function that go will look for whenever you execute your go application.
  * Print() is a function that comes from a go package a built-in package called fmt or format we have to explicitly import any packages from which we're using the function:
    fmt is a part of  [standard library](https://pkg.go.dev/std)
  * * ```go
      import "fmt";
      fmt.Print("hello FT")
      ```

execute go application:
`$>go run main.go`

## Variables & Constants

to define a variable use the var keyword.

`var conferenceName = "Go Conference"`

in go unlinke many other language when we define a variable with a certain value and we don't use it in the code we get an error " ... declared but not used ". also you get the same error with package that package gets imported but not used which is a very good reminder to clean up your code.

to define the constants that are like variables, except that their values cannot be changed.

`const conferenceTcikets = 50`

## Formatting Output

whenever we're printing our text mixed with variables we can use a function called printf from the fmt package.

`fmt.Printf("Welcome to %v booking application \n", conferenceName)`

`%v` is the default format but you have other specific format also available if you want the values to be displayed differently and you can see this whole list in the go doc https://pkg.go.dev/fmt@go1.23.2

## Data Types in GO

```go
var userName
nameuserName = "Jamil"
```

in go all values have data types.

* When we create a variable or constant and assign a value to it immediately on the same line go can imply the data type based on the value so go knows that this is a variable type for String and this is a constant and variable types for integers.
* When we do not assign a value immediately go doesn't know what type of value you are going to store so it asks you to explicitly define a type to make your code
  * More Robust, reduces the likelihood of errors.
  * Helps developers to catch type mismatches sooner ( at compile time ).
* In GO we have uint8 ( u for unsigned ) , uint16..64 and int8...64 ( 8..64 correspond the length of integer and u for positive numbers ) and we have float and comlex for real numbers.
* One more thing in terms of variable definition in go ,  we have an alternative syntax for creating a variable and assigning it a value directly "Syntactic Suger"
  * instead of `conferenceName string = "Go programming "`
  * `we can get rid of var keyword as well as the type and before the equal sign just add a column: `**`conferenceName := "Go Programming"`**
  * note that this alternative syntax you cannot declare constants, it only applies to variables.

## Getting User Input

in order to read user input we use another function from the same format package called Scan.

fmt.Scan(userName) : save user's value in "userName" variable

after run that instruction we were not able to enter anything and the value of user is empty.

-> before the user input variable we need to add what's called pointer.

* A Pointer is **a variable** that points to the memory address of another that references the actual value.( in Go is called special variable )

e.g:

```go
fmt.Print(remainingTickets) // this will displays the value of the variable .
fmt.Print(&remainingTickets)   // this will displays  the address memory of the variable
```


"fmt" package gives us different functions : Formatted Input and Output ( I / O ):

* Print Messages.
* Collect User Input.
* Write into a File.

"fmt" package gives us different functions : Formatted Input and Output ( I / O ):

* Print Messages.
* Collect User Input.
* Write into a File.

## Arrays & Slices

arrays and slices are commonly used data types in go applications.

* array in go has fixed size and the type:
  `var bookins = [50]string  ( we can't have an array with mixed values )`
* to add or access an elements in array use the index:
  bookings[0] = "New Value"

-> we have an issue array what if we don't know the size when creating it? this is when slice comes on

**slice** is a list that is more **dynamic in size, where we don't need to specify a size at the beginning. and it is automatically expands when new elements are added.**

Slices in Go is an abstraction of an Array ( use the array under the hood).

Slices are more flexible and powerful: variable-length or get an sub-array of its own.

Slices are also index-based and have a size, but is resized when needed.

to create a slice:
var bookings []string

in order to add an element into slice we don't use the index instead we need to use append(slice, value), this will adds the element(s) at the end of the slice

## Loops

loops are simplified in Go, you only have the "for loop"

**for-each loop**

```
for i, element := range bookings{    }
```

range iterates over elements for different data structures.( not only arrays and slices ).

* _ ( blank identifier ) to ignore a variable you don't want to use.

**for loop condition:**
for condition {} : by default the condition is always true.

## Functions:

```go
func func_name(){}
func func_name(arg type ...) {}
func func_name(arg type) type {}
```

* Go function can return multiple values:

```go
func valideUserInput(firstNamestring, lastNamestring, emailstring, userTicketsuint, remainingTicketsuint) (bool, bool, bool) {
   isValidName := len(firstName) >= 2 && len(lastName) >= 2
   isValidEmail := strings.Contains(email, "@")
   isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
   returnisValidName, isValidEmail, isValidTicketNumber
}
```

we can define variables that are shared among multiple functions so variables that are accessible both for main and other functions as well ouput without having to pass them arround. it makes sense to create those variables in a place that lets multiple functions including the main have access to them and these are called **package level variables** and these are variables defined outside from all the functions.

```go
const conferenceTickets = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{}
func main(){...}
```
