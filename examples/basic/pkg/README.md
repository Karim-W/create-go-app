# pkg

Package for Libary code That is allowed to be imported by other applications
and libraries.

This is where the bulk of your code will go.

## Subpackages

- Adapters
  - Any clients or connection that your application needs to establish to
  *External* services
- Domains
  - Models,contracts,entities to be used by the application stuff here should
    *not* import anything kinda following the clean architecture pattern but not
    fully
- Infra
  - applicaiton infrastucture like logging, metrics, tracing, etc.
- Services
  - Any services that your application needs to have to perform jobs, basically
      interfaces that perform some sort of functionality like the context factory
- Repositories
  - Package That implements the Repository Pattern for your application
- Usecases
  - Package describes the usecases of the applications and handles the business
    logic of the application

For more info on this package refer to [this guide](https://github.com/golang-standards/project-layout/tree/master/pkg)
