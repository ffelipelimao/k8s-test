## Dockerfile helper

#### Why use workdir
The `golang` container image sets the GOPATH directory to /go and the working directory to /go too.
Following GOPATH layout, source code is expected to be in directories under /go/src.
`/go/bin` is where binaries are installed by default, and `/go/pkg` include things like the module cache.
With modules, it is entirely possible to work outside of the GOPATH (as you most likely are doing).
In which case, a common pattern is to choose your own working directory,
eg `WORKDIR /workspace` and work in there.