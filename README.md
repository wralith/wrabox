## Project based on Alex Edwards' Let's Go book.

### To start server

- Start mysql server and add information's below

  ` go run $(ls -1 cmd/web/*.go | grep -v _test.go) --dsn="<dbUserName>:<dbPassword>@/<dbName>?parseTime=true" `

- Or with defaults...

  `go run $(ls -1 cmd/web/*.go | grep -v _test.go)`


<small>I don't know if there is better way</small>  
<small>Update: Okay i know now but i will implement later, i want to see how weird it can get</small>

### For tests  
`go test ./... -v`
