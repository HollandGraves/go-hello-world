/*



GO NOTE:
To run a file:
	a. go to the file's path in the terminal
	b. type "go run fileName.go"
	c. e.g.
		- go run main.go
		- This command will run the program and print "Hi there!"

GO NOTE:
Commands found in the Go CLI (Command Line Interface):

	a. go build

		- NOTE:
		- simply compiles file
		- e.g.
		- go build main.go
		- in this case the cose from the file main.go in the directory selected through the terminal will ONLY compile

		- NOTE:
		- if you look in the terminal at your file path where you compiled through build, and look through ls, there will be a new file
		- e.g.
		- main    main.go
		- this main file is the compiled code and will appear with the extension .exe for windows. This file is an executable file, despite the OS used

		- NOTE:
		- the way to run the file through the CLI is with:
		- ./main
		- this will run the file just like the following CLI: go run main.go

	b. go run

		- takes up to a handful of files
		- compiles those files
		- and runs those files immediately

	c. go fmt

		- formats all the code in each file in the current directory

	d. go install

		- Compiles and "installs" a package that is a

	e. go get
	f. go test



*/