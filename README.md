# CompanyX coding challenge

This is the CompanyX coding challenge. It should be simple to understand as I 
have chosen not to use any third party libraries. All code was written between
10:00AM and 1:00PM Pacific time on January 8th, 2017.

## Up and running

* Ensure you have a working golang setup (`GOBIN`, `GOPATH` and `GOROOT` set
  up)
* Make sure this folder (`companyxchallenge`) is in `$GOPATH/src`
  (`$GOPATH/src/companyxchallenge`)
* `cd` into `$GOPATH/src/companyxchallenge/cmd/companyxweb`
* Run `go install` from the command line
* You should now have a binary called `companyxweb` you can run.
  Use `companyxweb -help` for help.

## Etc

This is all my own original work coded by myself, Jason Scheirer,
and the external resources used were:

* The Google doc with the specification for the API
* `curl` (to inspect API endpoints' output)
* The Go standard library documentation at https://golang.org/pkg/
* Visual Studio Code for developing

## Things I'd do if I had more time

* Tests (lots of tests)
* Logging beyond some prints
* Use something a little more robust than the standard library
* Perhaps do some caching so I'm not hitting every API endpoint every time
  (I noticed [this endpoint](http://api.icndb.com/jokes) returns _everything_,
  for example, so I could have fetched this once and kept it in memory, cutting
  out one more place of potential failure)
