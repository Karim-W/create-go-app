# cmd

The `cmd` is the entrypoint package for the application. Here we define the
main function and the command line interface.

this sucks, beacuse any relative imports will be relative to the cmd package
not the root of the project.

> if you load an .env/json/yaml/toml/ini files here, you kinda have to put them
either here or import them relative to here

refer to this guide for [more info](https://github.com/golang-standards/project-layout/blob/master/cmd/README.md)

> p.s. i put the server here for now bc to me it's not really a package, rather
an entrypoint to the application. i think it's fine to put it here for now.
