# Library

## Create the executeble
- create a bin for linux
    ```
    GOOS=linux GOARCH=amd64 go build -o bin/library cmd/*go
    ```
- create a bin for mac
  ```
   go build -o bin/library cmd/*go
    ```
## Run the tests

  - run the test
  ```
   go test -v -race ./...
  ```
- run
  ```
    ./bin/library
  ```
## Build the image 
  1. Build the image and run it

       ``` shell
      docker build -t test-file --build-arg FILEDIR="./config" --build-arg FILENAME="ports.json" --build-arg SLEEP=1 .

      docker rm -f fileService
      docker run --name fileService -it test-file
       ```
