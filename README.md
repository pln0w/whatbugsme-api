# WhatBugsMe
WhatBugsMe is an open source project to create issue boards within your organisation. The main focuse of this project is to stay anonymous, for those, who want to report a case.

## Getting started

### Basic Installation

#### Using Docker environment
1. Run  
`docker-compose up --build -d`

2. Server has live-reload mode
`docker logs app -f`  
and expect "_Server start listening on port 8080_" message 

To force rebuild container, run  
`docker-compose up -d --force-recreate`


#### Native Go environment
1. Install Docker locally  
2. `go run main.go`


## Contributors
Frontend: **Szymon Półtorak**  
Backend: **Wojciech Pawlinów**

## About Software Brothers
WhatBugsMe was started by the team at [Software Brothers](https://rst-it.com/en).