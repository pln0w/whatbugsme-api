# WhatBugsMe
WhatBugsMe is an open source project to create issue boards within your organisation. The main focuse of this project is to stay anonymous, for those, who want to report a case.

## Getting started

### How to run (Docker)

* Run  
`docker-compose up --build -d`

### How to run (native Go environment)

* Run  
`go run main.go`


### DEVELOPMENT MODE
Server has live-reload mode enabled. Use Dockerfile.dev in your _docker-compose.yml_ file or Dockerfile.prod if you only want to build and use API server.

Run `docker logs app -f` to see container logs, expect "_Server start listening on port 8888_" message. You may also `tail -f server.log`  

To force rebuild container, run  
`docker-compose up -d --force-recreate`  
This will destroy MongoDB container  data. Change volumes settings in _docker-compose.yml_ file, if you want to prevent it.

## API docs

[See documentation](API.md)

## Contributors
Frontend: **Szymon Półtorak**  
Backend: **Wojciech Pawlinów**

## About Software Brothers
WhatBugsMe was started by the team at [Software Brothers](https://rst-it.com/en).