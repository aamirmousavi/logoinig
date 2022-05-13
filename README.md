# Logoing
logoing is a simple example for login and generate secure token with jwt and gin packages and postgres as our database.
in this project we are not using token with cookies it's just a simple example you can develop it yourself.
note: don't recive your toekn in query parameters of your users request.


# Contents
- [Install](#install)
- [Requirements](#requirements)
- [Testing](#testing)

# Install

```
git clone github.com/aamirmousavi/logoinig
```
```
cd logoinig
```
and run the app
```
go run cmd/main.go
```
also you can build and run the build

# Requirements

Install and run posgres 
you can change postgres dataSourceName in internal/services/database/postgredb/pgresdb.go line 18

also you need to create a table with user as name and
```
id          int
username    text
password    text
firstname   text
lastname    text
```
as our columns
and insert one or more rows to your table
# Testing
I created a user with 'aamirmousavi' as username and '85856969A' as password
we send the request
```
curl --location --request POST 'localhost:8080/login' \
--form 'username="aamirmousavi"' \
--form 'password="85856969A"'
```
out:
```
token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImFhbWlybW91c2F2aSIsImV4cCI6MTY1MjQ4MzMwMH0.oE5qtBNzoOHOdG8NrjUbjQ4_I8dMuM9AGBHQbRD3jJs
```
now we check our token
```
curl --location --request GET 'localhost:8080/admin?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImFhbWlybW91c2F2aSIsImV4cCI6MTY1MjQ4MzMwMH0.oE5qtBNzoOHOdG8NrjUbjQ4_I8dMuM9AGBHQbRD3jJs'
```
out:
```
hello aamirmousavi.
```