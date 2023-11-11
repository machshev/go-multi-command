# Go based multi command

There are plenty of libraries that provide Cli args parsing with subcommands. In Python for example `click` seems quite popular, and for go there is for example `cobra`. These allow multiple git like sub commands within the same executable.

But what if you want a more plugin based approach like `git` and the `go` command itself. Where you can extend your application using any executable on `$PATH`. This repo contains a concept demonstrator go application that will do just that.

```
go build .
```

This will give an executable called `mc` which does the marshalling. Just running the `./mc` command will provide the usage help. If we run `./mc test` then we get an error saying it can't find the "test" subcommand.

For demonstration purposes we can add the current directory to `$PATH`.

```
export PATH=$PATH:$PWD
```

Then we can run the following commands, executing the plug-in subcommands as if they were part of the main command executable..

```
mc
mc bash
mc bash 1 2 3 4
mc python
mc python -c 123 ABC
```
