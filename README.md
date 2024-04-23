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
      docker build  -t libarary-1 --build-arg FILEDIR="./config" --build-arg FILENAME="library.local.yaml" --build-arg TARGETFILENAME="library.yaml" .

      docker rm -f libraryservice
      docker run --name libraryservice -p 8081:8081 -it libarary-1
     
     
      docker build  -t libarary-k8 --build-arg FILEDIR="./config" --build-arg FILENAME="library.yaml" --build-arg TARGETFILENAME="library.yaml" .
       ```
