# Introduction
The MovieAppAWS is Go Language Based Project where we will be fetching movie related data using AWS Lanbda.

# Installation
git 
Akshats-MacBook-Pro:MovieAppAWS akshatsharma$ go version
go version go1.10.2 darwin/amd64

Step 1 : 
```
go get github.com/aws/aws-lambda-go/lambda #Go Get was earlier GoInstall
```

Step 2: 
```
go get github.com/stretchr/testify
```

Step 3: 
        ```
        go list ... #Verify all packages by Using Command
        ```
        Or
        ```
        history | grep "go get" #Check the History of installed packages
        ```
Step 4: Create main.go file

Step 5:
```
set GOOS=linux
set GOARCH=amd64
go build -o main main.go #To generate Binary
```

for windows  use addtional command - go.exe get -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip

For Mac use command -

```
zip -r main.zip main
```
or
```
zip main.zip main
```

Reference :
1. https://github.com/aws/aws-lambda-go