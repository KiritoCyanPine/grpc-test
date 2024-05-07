# grpc-test Golang

A simple example of implementing grpc service via protobuf and `protoc` tool. 🚀

## Running

### Server 
Run the following command on the project terminal.

```bash
$ export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
$ make go-gen
$ make run-server
```

### Client 
Run the following command on the project terminal.

```bash
$ make run-client
```

### TODO :
* [x] create a simple service with necessary files and tools
* [ ] create a docker file and run server code from there
* [ ] write test cases to run everytime docker image is created to perform operations on whole service
* [ ] write github actions to Automate Tests on each merge to main branch
