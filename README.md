# chatty

To run the server:

- In the root directory: `go build`

To run the client:
- In ./tool/chatty-cli : `go build`

The server runs on port :8081

TODO: create a docker image so it can run on any port

I did not vendor the tests: if you want to run them install ginkgo and
testify/mocks and run `go
test`

```
go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega
go get go get github.com/stretchr/testify
```

TODO: more negative test cases

TODO: better url parsing

TODO: refine @mentions to not mistakenly parse emails

TODO: reverse lookup of emoticons ??? FUTURE FEATURE MAYBE?
