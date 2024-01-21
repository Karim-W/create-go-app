# usecases

Package describes the usecases of the applications and handles the business
logic of the application

To create a new usecase you need to do the following:

- create a new interface in the `usecases` package's root
- create a subpackage in the `usecases` package with the name of the interface
- create a new implementation of that interface in the `usecases/{name}` package
- call the usecase's initializer/constructor in the `cmd/main.go` file
