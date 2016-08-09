# Name of binary output:
BINARY=client

# Values of VERSION and BUILD :
VERSION=1.0.0
BUILD=`date +%FT%T%z`
#BUILD=`git rev-parse HEAD`

# Setup the -ldflags option for go build here:
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

# Builds the project:
build:
	go build ${LDFLAGS} -o ${BINARY}
	
# Install project :
install:
	go install ${LDFLAGS}

# Clean project :
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi 