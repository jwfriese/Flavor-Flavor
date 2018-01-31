# Flavor-Flavor
How easy is it to make these acceptance tests pass in your language and
framework of choice?

## Initialize your environment

You'll only need to do this stuff one time.

1. Install the latest version of [Go](https://golang.org/dl/)
1. Pull this repo down into your `GOPATH`:
```
go get github.com/jwfriese/flavor-flavor
```
1. Get [ginkgo](https://onsi.github.io/ginkgo/) and
[gomega](https://onsi.github.io/gomega/):
```
go get github.com/onsi/ginkgo/ginkgo`
go get github.com/onsi/gomega/...
```
1. Set up your database. This repo provides a script that will spin up a
Postgres instance in a container using Docker. Run it using the following
command from the project's root:
```
./script/init-db.sh
```
You can also choose to set up your own Postgres instance if you like. Just make
sure that it has a user name `flavor-flavor` with a password of `moreflava` so
that the test script can reset the DB before running.

## Running the tests

1. Ensure your API is running on `localhost:8181`.
1. Run using the test script by running the following command from the
 project's root:
```
./test.sh
```
The script takes two actions in sequence: 1) It resets the database; 2) It runs
the tests.  
