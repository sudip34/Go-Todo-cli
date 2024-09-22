Todo app in GO
===
## create a module inside the direcotry with a name 

>create go.mod 
```
go mod init <name>

```

## run a go module

> inside the go.mod directory
```
go run ./

```


## run a simple go file

```
go run <file.go>
```  


## To printing up the TODO list use external libray `github.com/aquasecurity/table`

- 1. openup an external terminal

```
go get github.com/aquasecurity/table

```
this will update the ***go.mod*** and generate ***go.sum***



## Use the command.go or cmdline

>in `command.go` file we have following options 

|Option | Input Type|
|---|---|
| Add   | string |
| Del   | int |
| Edit  | string|
| Toggle| int| 
| List  | bool|


use command as:
```
go run ./ -<command>

e.g. go run ./ -List

```