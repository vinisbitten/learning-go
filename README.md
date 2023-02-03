<div align="center">

![go banner](/src/go-banner.png)

<h1>Golang</h1>
</div>

This is a repository to help you understand basic Go syntax and some of its features.z
Since we're using go workspaces, there are 2 ways of running the examples.
Go workspaces allow you to create a multi-module environment.

If you want to run for example the first module "hello-world", you can use the two methods below:

```bash
# Run the file (if you are in repo root path)
go run examples/01-hello-world/main.go
# Run the module 
go run hello-world
```

![run print](/src/run-print.png)

You can also turn go files or modules into binnary files and then run it:

```bash
# build file
go build examples/01-hello-world/main.go

# build module
go build hello-world

# We can then execute the built binary directly
./hello-world
```

![build-print](/src/build-print.png)
