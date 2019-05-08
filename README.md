# cl
`cl` is a tool that takes one or more group(s) of color/pattern(regular expression)
and applies these rules on the given input(a file or piped through stdin) by
coloring the part of the input that matches the given pattern with the given color.

## Install
### From source
#### Install go
Follow the [install instructions](https://golang.org/doc/install) from golang official website.
#### Clone repository
```
git clone https://github.com/aburdulescu/cl.git
```
#### Build the code
```
go build
```
#### Install the binary
##### Copy the binary to a directory that it's in your `PATH` variable
```
sudo cp cl /usr/local/bin/
sudo ldconfig
```
##### Install the binary to a working golang [workspace](https://golang.org/doc/code.html#Workspaces)
```
go install
```
