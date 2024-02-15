
# GitHub CLI App to fetch most starred repos within a date range

Simple CLI that uses GitHub's Search REST API to fetch the most starred repositories within a given date range(optional)

## Technologies uses

- Go
- Cobra



## Usage/Examples

After cloning the repository run the following command to build the CLI.
```bash
go build -o bin/star-repos.exe
```
The ```.exe```extension is used only to run this CLI on Windows. To build the CLI on Linux use the following command:
```
go build -o bin/star-repos
```
A reference to build for other OS environments [here.](https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures)

The next step is to run the CLI using
A simple CLI application that fetches the most starred GitHub repositories within a date range. For example:
                  ``` star-repos <DATE_FROM> <DATE_TO>```

Usage:
  ```star-repos [flags]```

Flags:
  ```-h, --help   help for star-repos```
