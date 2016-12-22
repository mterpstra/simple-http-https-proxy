# simple-http-https-proxy
Very simple proxy to allow http to https traffic with logging.

## Work in Progress ##
When looking for a simple way to pass HTTP traffic to a proxy and have
that proxy forward it to an HTTPS service, I did not find any very
simple solutions out there.  It seems as nginx could do something like 
it, but again, it was too much effort.  This is a very simple, with 
no error checking, but working version of http to https proxy.

## Build Instructions ##
go build

## Usage ##
./simple-http-https-proxy {Your Host Name Here}

## What it does ##
1. Listens on localhost:8080
2. Accepts requests
3. Dumps the rquest to the console
4. Replace Host header with Argv[1]
5. Replace http with https
6. Sends the request to the server at Argv[1]
7. Dumps server response
8. Sends response back to caller
