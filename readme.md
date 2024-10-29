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
