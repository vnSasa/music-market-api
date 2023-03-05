# music-market-api
creating a music database with the possibility of adding your favorite compositions to your personal cabinet

## task-1

- create a project structure

- implement functionality for user registration and authentication

- add the HTML files to the corresponding functionality

## Before starting the project, start Redis:

- docker run -d -p 6379:6379 redis

## and make migrations:

- migrate -path ./schema -database 'mysql://sasa:110513@tcp(127.0.0.1:3306)/music-market' up