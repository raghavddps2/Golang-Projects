1. How do we run the code that we are writing ?
    To run the code in our project, we simply write go run hello.go.
    This is all we have to do to run a go code

    go CLI - 
            go build - compiles a go file
            go run - compiles and executes files
            go fmt - formats all  the code in each file of the current directory
            go install - compiles and installs a package
            go get - This downloads the raw source code of someone else's package
            go test - Runs the tests associated with the code.

2. What does the line package main is doing for us ?

    Package basically groups various files, so that we can use them together
    Inside go, there are 2 kinds of packages
    1. Executable type
        Executable types are the one that generate a file that we can run
        The word main is used to make an executable type package.
    2. Reusable type
        Code basically used as helpers. Good place to put reusable logic
        Compiling a non-main package gives nothing.

3. What does import "fmt" mean in the go code ?

    To use the funtionalites defined by others, what we can do is simply import these packages
    Standard library packages are: fmt,math,io etc..

4. What is the func thing ?

    func basically means function, these are just like functions in any other programming language.
    func -> Keyword
    main(args) -> function name and how we pass arguments
    {} -> inside this we write logic

5. How is the go code organized ?
    package, import statements, main and other functions.