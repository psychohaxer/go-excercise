# go-exercise

This repository contains several exercises in Go programming language for learning and practicing programming concepts.

## Files

1. `exercise-1.go`: Contains the implementation for detecting and counting palindromic numbers. Status: Done
2. `exercise-2.go`: Contains the implementation for sorting objects with a specific data structure. Status: Done
3. `exercise-3.go`: Contains the implementation for finding the missing number from a string. Status: Work in Progress

Each file has its own unit tests to verify its functionality.

## Webservice

The implementation of a webservice can be found in the `main.go` file. The webservice accepts input and performs data processing using the functions from the exercise files mentioned above. The webservice provides relevant output based on the request.

## Docker Image

In this repository, there is also a `Dockerfile` that is used to build a Docker Image. The Docker Image will include all the necessary files and dependencies to run the webservice and perform data processing.

## Usage

1. Make sure you have Go installed on your machine.
2. Clone this repository to your local directory.
3. To run the webservice, open a terminal in the repository directory and execute the command:
   ```
   go run main.go
   ```
4. Open a web browser and access `http://localhost:8080` to access the webservice.


## Note

- The file `exercise-3.go` is still in development (Work in Progress). Please stay tuned for updates regarding this file.
