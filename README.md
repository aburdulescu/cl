# cl
`cl` is a tool that takes one or more group(s) of color/pattern(regular expression)
and applies these rules on the given input(a file or piped through stdin) by
coloring the part of the input that matches the given pattern with the given color.

## Install
### Prebuilt binary
#### Download tarball with prebuilt binary for your OS(E.g. for `GNU/Linux x86_64`)
```
curl -L -O https://github.com/aburdulescu/cl/releases/download/v0.1/cl_v0.1_linux-amd64.tar.gz
```
#### Optionally, verify the authenticity of the tarball
- downloaded SHA512 hash corresponding to the downloaded tarball
```
curl -L -O https://github.com/aburdulescu/cl/releases/download/v0.1/cl_v0.1_linux-amd64.tar.gz.sha512
```
- generate SHA512 hash of the downloaded tarball
```
sha512sum cl_v0.1_linux-amd64.tar.gz > cl_v0.1_linux-amd64.tar.gz.sha512.local
```
- check to see if they match(no output means everything is fine)
```
diff -u cl_v0.1_linux-amd64.tar.gz.sha512.local cl_v0.1_linux-amd64.tar.gz.sha512
```

#### Unpack tarball and put it somewhere in your PATH(e.g.: `/usr/local/bin/`)
```
tar xf cl_v0.1_linux-amd64.tar.gz
sudo cp cl /usr/local/bin/
```
#### Regenerate ld cache(otherwise your shell won't find the binary)
```
sudo ldconfig
```
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
