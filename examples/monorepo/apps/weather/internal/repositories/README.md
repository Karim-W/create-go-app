# repositories

Package which implements the repository pattern for the application.

To add a new repository to the application you need to do the following:

- create a new interface in the `repositories` package's root
- create a subpackage in the `repositories` package with the name of the interface
- create a new implementation of that interface in the `repositories/{name}` package
- call the repository's initializer/constructor in the `cmd/main.go` file
