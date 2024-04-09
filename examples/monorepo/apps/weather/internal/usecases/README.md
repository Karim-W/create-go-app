# useevents

Package describes the useevents of the applications and handles the business
logic of the application

To create a new usecase you need to do the following:

- create a new interface in the `useevents` package's root
- create a subpackage in the `useevents` package with the name of the interface
- create a new implementation of that interface in the `useevents/{name}` package
- call the usecase's initializer/constructor in the `cmd/main.go` file
