# Go Calculator Bot

This simple example bot written in Go Lang performs calculations; it will keep
taking in numbers until the user sends in an operator which will then trigger
the bot to perform the operation on all numbers it saved and send back the output

This example is meant to showcase the features and strengths of the Go programming
languages and many of the approaches used here can be improved and optimised.

# Installing

This project uses [GoVendor](https://github.com/kardianos/govendor) to handle
dependencies:

```
go get -u github.com/kardianos/govendor
```

Before building run `govendor install +local`

# Before running

Be sure to set the appropriate environment variables, this includes the 
verification token you set in your facebook developer dashboard and the page
access token:

```
export GOBOT_VERIFICATION_TOKEN=VERIFICATION_TOKEN
export GOBOT_PAGE_ACCESS_TOKEN=PAGE_ACCESS_TOKEN
```
