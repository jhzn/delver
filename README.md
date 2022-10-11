## What does it do?

Delver makes the command line interface for starting `dlv` the same as the one used in `go test`

### Example

Say you're using this when developing/testing:
```shell
go test -v -count=1 -run '^TestMyFunc$' ./pkg/api/tests
```
And want to switch to debugging it with `dlv`. You can't just swap the program and have it work.
```diff
-go test -v -count=1 -run '^TestMyFunc$' ./pkg/api/tests
+dlv test -v -count=1 -run '^TestMyFunc$' ./pkg/api/tests
```

Delver fixes this. Now you can do:
```shell
delver test -v -count=1 -run '^TestMyFunc$' ./pkg/api/tests
```
You'll get dropped into `dlv` like so:
```
delver is running cmd:
[dlv test --build-flags./pkg/api/tests -- -test.v -test.count=1 -test.run '^TestMyFunc$' ./pkg/api/tests]

Type 'help' for list of commands.
(dlv)
```

## Instructions

### Install

```shell
go install github.com/jhzn/delver@latest
```

### Develop

```shell
git clone https://github.com/jhzn/delver
go build -o delver ./...
```

## Note
This is very early software and not rigorously tested.

If you find any issue PR:s are very welcome :)
