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
https://pkg.go.dev/search?q=mux

go env -w GO111MODULE=off
go env -w GO111MODULE=auto

### init the go tracking files
The go mod init command creates a go.mod file to track your code's dependencies
go mod init example.com/hello

### Notes
 * Redeclare is not allowed but shadowing is allowed
 * If you are variable start with small letter then its allowed only in current package
 * If you are variable name start with capital letter then its allowed to export to another package as well
 * constant can declare and with out use but variable should use if you are declare