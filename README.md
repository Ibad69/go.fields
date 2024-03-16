the project architecture / structure is designed in a way that is suitable for independent development and with good go dev practices.

our code initializes from the cmd api main.go file, from where it starts and references to a config start file the config start file has a struct to all of the properties,
which in turn starts all of our database and registers our routes of all the services / endpoints that we may have, for example a user service is seperately kept as a package so that it's routes and every logic is done seperately.


go get -u github.com/go-chi/chi/v5