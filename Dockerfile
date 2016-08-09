
# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang
MAINTAINER Tomer Golany <tomer.golany@gmail.com>

# Copy the local package files to the container's workspace.
ADD . /go/src/webAppExample

# Build the webapp command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install webAppExample

# Run the webapp command by default when the container starts.
ENTRYPOINT webAppExample

# Document that the service listens on port 8080.
EXPOSE 8080