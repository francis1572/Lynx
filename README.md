# Lynx
Label Lab system Back-end api via go lang

## Before Started
### Most important
1. Please create your own branch for your features. If you want to merge your codes into `main` branch, please create a new pull request and contact project owner or other team members to review your codes.
2. If you need dramatic refactoring, please discuss with code owners.
### For Mac users
1. Make sure you have golang installed:
      * Install through homebrew: `brew install golang`
      * Check whether golang is successfully installed: `go version`
2. Go environment:
      * <b>Step 1</b>: In your terminal (take zsh for example), create a directory specially for Go projects, you have to put and run your Go projects in this directory. For example, in your Documents directory: <br>
      `$ mkdir go-projects`
      * <b>Step 2</b>: Set Go Paths environment variables. Feel free to choose any one from `Vim version` and `terminal version`.<br>
        * Vim version:<br>
            1. terminal command: ```$ vim ~/.zshenv```<br>
            2. In your vim: <br>
            ```
            export GOROOT=/usr/local/go
            export GOPATH=~/Documents/go-projects
            export PATH=$PATH:$GOPATH/bin
            export PATH=$PATH:$GOROOT/bin
            ```
            * Remind: GOROOT would be where you installed your golang. Can be checked by `$ which go`<br>
        * Terminal version:
            * Termianl commands:
              ```
              $ echo 'export GOPATH=~/Documents/go-projects' >> ~/.zshenv
              $ echo 'export GOROOT=/usr/local/go' >> ~/.zshenv
              $ echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.zshenv
              $ echo 'export PATH=$PATH:$GOROOT/bin' >> ~/.zshenv
              ```
        Check whether you have successfully finished this step by typing follwoing command in a new terminal window.(no matter which version you choose to set variables.)<br>
        `$ echo $GOPATH` (should show `~/Documents/go-projects`) (can deduce the rest from this)
      * <b>Step 3</b>: Create go workspace by command:<br>
      `$ mkdir -p $GOPATH $GOPATH/src $GOPATH/pkg $GOPATH/bin`
      * <b>Step 4</b>: Clone or create your Go projects in `$GOPATH/src`. Since this repo have absolute directory `github.com/...`, your have to create a directory named `github.com` in `$GOPATH/src`, and then clone this repo in `github.com`.
      * <b>Step 5</b>: Install packages. You can use `$ go install` command to check which packages you should install before run the project. And use `$ go get [packages routes]` to install the packeages.
      * <b>Step 6</b>: Run this project, use `go run .`
      * <b>Step 7</b>: To check whether server is successfully running, you can try: `$ curl -X GET 'http://127.0.0.1:9090'` in a new terminal window.
      * [Ref](https://sourabhbajaj.com/mac-setup/Go/README.html)
4. Some file/directory notes:
      * `main.go`: main function that run this server.
      * `route.go`: routes of restful api can be found here.
      * `/controller`: main logics here.
      * `/service`: connect to db and do CRUDs.
      * `/models`: db schemas
      * `/db`: deal with mongodb connection.
      * others: are not important.

## Execute

* To run the server, you can use
```
go mod download
go run .
```
* If you need live-reload, install `gin`: `$ go get github.com/codegangsta/gin`. Then run this repo by: `$ gin -i run .`. [Ref](https://github.com/codegangsta/gin)

## Connect to Linux Server (temp)
To check the error log or CI/CD process.
* make ssh connection: `ssh frank@140.112.107.121` and enter the password.
* attach to tmux session: `tmux attach -t golang_server`
* if Server is shutdown, run: `bash server.sh` to re-run the server.

## Tutorial Recommandation

* [**Basic GO Lang**](https://michaelchen.tech/golang-programming/write-first-program/)
* [**Package in GO Lang**](https://calvertyang.github.io/2019/11/12/a-beginners-guide-to-packages-in-golang/)
* [**Build Web Application with Golang**](https://willh.gitbook.io/build-web-application-with-golang-zhtw/)
* [**Connection with mongoDB**](https://zhuanlan.zhihu.com/p/144308830)
