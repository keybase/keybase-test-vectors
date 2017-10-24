# Keybase Test Vectors

This repository contains test vectors used as test inputs for sigchain parsers.

## User Sigchains

User sigchains are converted from a DSL in
`chains/inputs/*.{iced,cson}` into `chains/*.json` by the
[forge](https://github.com/keybase/node-forge-sigchain). Some vectors
are specified in json directly and begin their life in
`chains/*.json`.

Then they go through a couple more generator scripts and end up as go tests.

## Team Sigchains

Team sigchains are converted from a DSL in `teamchains/inputs/*.iced`
into `teamchains/*.json` by the
[forge](https://github.com/keybase/node-forge-sigchain). Then those
JSON files are loaded directly by go tests.

```
$ cd teamchains

# Rebuild all json tests:
$ ./compile.sh inputs/*.iced

# Rebuild one test and watch for when any input iced or the forge change:
$ export FOCUS=member_duplicate
$ ls $GOPATH/src/github.com/keybase/node-forge-sigchain/build-stamp inputs/*.iced | entr -r sh -c "forge-sigchain --pretty --team -f iced < inputs/$FOCUS.iced | tee $FOCUS.json"

# Run the client go tests and watch for when any input json change
$ cd $GOPATH/src/github.com/keybase/client/go/teams
# first collect coverage for all the tests, so that coverage will be complete but we don't have to run them all.
$ go test --coverprofile=/tmp/teams-coverage-online.out
# then run the go tests and make a coverage report
$ ls $GOPATH/src/github.com/keybase/keybase-test-vectors/teamchains/*.json ./*.go | entr -r sh -c "echo starting && go test -v --coverprofile=/tmp/teams-coverage.out -run "^TestUnits$" 2>&1 | tee /tmp/test.log ; tail -n +2 /tmp/teams-coverage-online.out >> /tmp/teams-coverage.out && go tool cover -html=/tmp/teams-coverage.out -o /tmp/teams-coverage.html"

# Run one client go test
# do the same as above with this env var set
$ export KEYBASE_TEAM_TEST_SELECT="member_duplicate.json"

# Run all the client tests with no coverage
$ ls $GOPATH/src/github.com/keybase/keybase-test-vectors/teamchains/*.json ./*.go | entr -r sh -c "go test -v -run "^TestUnits$" 2>&1 | tee /tmp/test.log" | mark
```
