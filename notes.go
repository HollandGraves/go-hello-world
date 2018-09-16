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



GO NOTE:
package main

	a. anything with "package main" is the code of your go file that compiles into an executable file

		- the thing to note about package main files is that mutiple files in the same folder use package main

		- also to note, ANYTIME YOU MAKE A PACKAGE MAIN FILE THERE MUST BE A FUNCTION CALLED "func main() {code goes here e.g.}" within the file

			- func main() will be called the moment the package main file is ran after being compiled

	b. anything with "package any-other-name-here-besides-package-main" is a importable file, a dependency that can be called into another go file (i.e. a reusable file)

		- "package any-other-name-here-besides-package-main" files are outside the folder the package main files are found in and are imported in as noted before

		- if a "package any-other-name-here-besides-package-main", or any file for that matter, is in the same folder as a file. YOU DO NOT NEED TO USE THE

			- IMPORT SYNTAX TO UTILIZE FUNCTION FROM THAT "package any-other-name-here-besides-package-main" FILE!!!

		- a good use for reusable files is to create a function within a file, import the file, and call the function



GO NOTE:
import other packages

	a. when you use the syntax e.g. import "fmt" you are importing the package named fmt, in this case the fmt package is a standard library package of golang

		- to access a function within the standard library import, within the functio of, let's say, your func main() you must call the package and use dot notation
			to access the function from the package

	b. some examples of standard library packages:
		- debug
		- math
		- encoding
		- fmt
		- io
		- crypto


*/