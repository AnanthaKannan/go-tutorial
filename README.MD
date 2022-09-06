### Version
    go version
### Env
    go env
### Run the application
    go run fileName.go
### Build the Application
    go build
### Panic
panic is just like throw error

### packages

All the package are htere [URL](https://pkg.go.dev/search?q=mux)  

To Instal the package use:
```
go get github.com/gorilla/mux
```
go env -w GO111MODULE=off
go env -w GO111MODULE=auto

### init the go tracking files
```
go mod init
```
The go mod init command creates a go.mod file to track your code's dependencies
go mod init example.com/mongoConnect

### find in the local path
```
go mod edit -replace example.com/greetings=../greetings
```
### Get all the modules 
```
go mod tidy
```
From the command prompt in the hello directory, run the go mod tidy command to synchronize the example.com/hello module's dependencies, adding those required by the code, but not yet tracked in the module.

### Build
```
go build filename.go
```
It is just build the file and generate the build file, if you are run the file, it will get execute. If you want to build the file, that that should be `package main`

### Format
```
go format
```
used to format all the file
### Test
```
go test
```
used to execute the test file
### Notes
 * Redeclare is not allowed but shadowing is allowed
 * If you are variable start with small letter then its allowed only in current package
 * If you are variable name start with capital letter then its allowed to export to another package as well
 * constant can declare and with out use but variable should use if you are declare
 * All the modules details are there int he go.mod file