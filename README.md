To start server  
`go run $(ls -1 cmd/web/*.go | grep -v _test.go)`  
  
<small>I don't know if there is better way</small>  
  
For tests  
`go test ./... -v`