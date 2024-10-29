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
