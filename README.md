# backend Task
this is the implementation of the task provided by Hussain as a test to be a candidate for a backend developer position.
## Table of contents
* [General info](#general-info)
* [Code Structure](#code-structure)
* [Future Work](#future-work)
## General info
This project created on golang as a backend language a long with mux, jwt and dotenv packages 
The Project simply built for the following points:
* API for User registration and login, using JWT
* Twitter Search API: An API that you send a search query, and then it searches through API calls in twitter.com, and returns the first 50 results
`this point has been replaced with another API due twitter new policies`
* An API that when sending it the tweets from the previous results, saves it in a database
* An API that returns the results saved in the database, with pagination
 
## Code Structure 
The Structure I used is similar to MVC patters except the V part (view) and the structure distrbuted as following:
* `Controllers`: this folder hold all the functionality for the each model
* `Models`: this folder hold all the data Schemas and structs 
* `db`: this folder hold the connection with the database
* `helper`: this folder hold the hlper functions such as the jwt token generator 

`PS: Each folder is a package`

## Future Work
* add bcrypt to the project to encrypt the user password
* create middleware for the refresh token 
* create middleware for authorize the client requests

