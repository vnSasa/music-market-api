# music-market-api
creating a music database with the possibility of adding your favorite compositions to your personal cabinet

## task-1

- create a project structure

- implement functionality for user registration and authentication

- add the HTML files to the corresponding functionality

## task-2

- create a functionality for the administrator (adding an artist and a song and new administrators to the database)

- add the HTML files to the corresponding functionality

## task-3

- create a functionality for the user (adding a song to the user`s playlist and toplist)

- create a functionality to display the rating of a song (If a user adds a song to their library, the song's rating increases by 1. If a user adds a song to their top list, the song's rating increases by 1 again. So we will get a real rating of songs, based on the number of additions to users' libraries)

- add the HTML files to the corresponding functionality

## Before starting the project, make migrations:

- migrate -path ./schema -database 'mysql://sasa:110513@tcp(127.0.0.1:3306)/music-market' up