# CRUD with go and postgres

## Prerequisites

The project uses go language hence you need to install it. You can install golang from this url -> <mark>[Install GoLang](https://go.dev/doc/install)</mark>.
In this project I have used postgreSQL for database. Based on your requirements you can use any of the database you're comfortable with and can remove the postgreSQL dependencies and add your database dependencies to it. If you want to continue using postgreSQL just install it from this url -> <mark>[Install PostgrSQL](https://www.postgresql.org/download)</mark> and then read the below section.

### Setting up .env file

Now that you have installed all the required files and modules, you have to make a <mark>.env</mark> file and then copy the contents of <mark>.env.example</mark> file to <mark>.env</mark> file and change the enviroment variables value accordingly.

### Running the code

Firstly you need to install all the project dependencies. If you do not have a go.mod file run the following command to create one `go mod init <your_module-name>`. Once it is created then run this `go mod tidy` to install all the required dependencies.

At this point you have all the dependencies installed to run the project. Now this project also includes programmatic database migrations. So the project has 3 CMD commands to either run the migrations or the server. The commands are as follows:

1) `go run src/main.go migrateUp` - This will run all the up migrations.
2) `go run src/main.go migrateDown` - This will run all the down migrations.
3) `go run src/main.go run` - This will start the server.

The starting point of the project is in <mark>src</mark> folder named <mark>main.go</mark>.

## Enjoy playing with API's - make updates and hit a PULL Request to merge 🎉 