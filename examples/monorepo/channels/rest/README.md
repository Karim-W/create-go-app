# rest

The rest package contains the REST API for the application. It is a `Gin-Gonic`
Web Framework application. You can change this to whatever you want.

## Subpackages

- `middlewares`
  - this is where you put your gin middlewares to be imported into the server
- `handlers`
  - REST API handlers for each resource

## Imports

- `{project}/internal/config`
  - this is where you put your config package and usually where the port is
  defined for the server
