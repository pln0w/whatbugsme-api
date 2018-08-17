# API end points docs

### Register organisation
```json
"request": {
    "method": "POST",
    "header": [],
    "body": {
        "mode": "formdata",
        "formdata": [{
            "key": "name",
            "value": "{{organisation_name}}",
            "description": "",
            "type": "text"
        }]
    },
    "url": {
        "raw": "localhost:8888/organisation",
        "host": [
            "localhost"
        ],
        "port": "8888",
        "path": [
            "organisation"
        ]
    },
    "description": "localhost:8080/organisation"
}
```

#### Responses:  
* `status: 200`
```json
{
    "id": "5b76b95488987601f64b10dd",
    "name": "SB",
    "created_at": "2018-08-17T12:02:28.731016Z"
}
```

* `status: 409`
```json
{
    "message": "Organisation already exist",
    "status": "fail"
}
```

### Search for organisation
```json
"request": {
    "method": "GET",
    "header": [],
    "body": {},
    "url": {
        "raw": "localhost:8888/organisation?name={{organisation_name}}",
        "host": [
            "localhost"
        ],
        "port": "8888",
        "path": [
            "organisation"
        ],
        "query": [{
            "key": "name",
            "value": "{{organisation_name}}"
        }]
    },
    "description": "localhost:8888/organisation?name={{organisation_name}}"
}
```

#### Responses:
* `status: 200`
```json
{
    "id": "5b76b95488987601f64b10dd",
    "name": "SB",
    "created_at": "2018-08-17T12:02:28.731Z"
}
```

* `status: 404`
```json
{
    "message": "Organisation not found",
    "status": "fail"
}
```
  

### Sign up
```json
"request": {
    "method": "POST",
    "header": [
        {
            "key": "Content-Type",
            "value": "application/x-www-form-urlencoded"
        }
    ],
    "body": {
        "mode": "urlencoded",
        "urlencoded": [
            {
                "key": "username",
                "value": "{{username}}",
                "description": "",
                "type": "text"
            },
            {
                "key": "email",
                "value": "{{email}}",
                "description": "",
                "type": "text"
            },
            {
                "key": "password",
                "value": "{{password}}",
                "description": "",
                "type": "text"
            },
            {
                "key": "organisation",
                "value": "{{organisation_id}}",
                "description": "",
                "type": "text"
            }
        ]
    },
    "url": {
        "raw": "localhost:8888/sign-up",
        "host": [
            "localhost"
        ],
        "port": "8888",
        "path": [
            "sign-up"
        ]
    },
    "description": "http://localhost:8080/sign-up"
}
```

#### Responses:
* `status: 200`
```json
{
    "id": "5b76c61b8898760201722f61",
    "username": "test1",
    "email": "test@rst-it.com",
    "password": "8d65075ace8bf7ce270305b0d88d5f34",
    "organisation": "5b76b95488987601f64b10dd",
    "token": "d092091d1e4c3dba73e5d6faf443a825",
    "created_at": "2018-08-17T12:56:59.3616693Z"
}
```

* `status: 409`
```json
{
    "message": "User exists",
    "status": "fail"
}
```

### Login
```json
"request": {
    "method": "POST",
    "header": [],
    "body": {
        "mode": "formdata",
        "formdata": [
            {
                "key": "email",
                "value": "{{email}}",
                "description": "",
                "type": "text"
            },
            {
                "key": "password",
                "value": "{{password}}",
                "description": "",
                "type": "text"
            },
            {
                "key": "organisation",
                "value": "{{organisation_id}}",
                "description": "",
                "type": "text"
            }
        ]
    },
    "url": {
        "raw": "localhost:8888/login",
        "host": [
            "localhost"
        ],
        "port": "8888",
        "path": [
            "login"
        ]
    },
    "description": "http://localhost:8080/login"
}
```

#### Responses:
* `status: 200`
```json
{
    "status": "success",
    "token": "0dec17e2b220120c18a65457f1bea062",
    "username": "test"
}
```

* `status: 422`
```json
{
    "message": "Login failed",
    "status": "fail"
}
```


### Get topics
```json
"request": {
    "method": "GET",
    "header": [
        {
            "key": "X-Auth-Token",
            "value": "{{token}}"
        }
    ],
    "body": {},
    "url": {
        "raw": "http://localhost:8888/organisation/{{organisation_id}}/topics",
        "protocol": "http",
        "host": [
            "localhost"
        ],
        "port": "8888",
        "path": [
            "organisation",
            "{{organisation_id}}",
            "topics"
        ]
    },
    "description": "http://localhost:8080/{{organisation_id}}/topics"
}
```

#### Responses:
* `status: 200`
```json
[
    {
        "id": "5b76c8be8898760201722f64",
        "content": "Example content",
        "votes_up": 0,
        "votes_down": 0,
        "organisation": "5b76b95488987601f64b10dd",
        "created_at": "2018-08-17T13:08:14.295Z"
    },
    {
        "id": "5b76c8bb8898760201722f63",
        "content": "Example content",
        "votes_up": 0,
        "votes_down": 0,
        "organisation": "5b76b95488987601f64b10dd",
        "created_at": "2018-08-17T13:08:11.849Z"
    },
]
```  
  

### Add topic
```json
"request": {
    "method": "POST",
    "header": [
        {
            "key": "X-Auth-Token",
            "value": "{{token}}"
        }
    ],
    "body": {
        "mode": "formdata",
        "formdata": [
            {
                "key": "content",
                "value": "CONTENT",
                "description": "",
                "type": "text"
            }
        ]
    },
    "url": {
        "raw": "http://localhost:8888/organisation/{{organisation_id}}/topics",
        "protocol": "http",
        "host": [
            "localhost"
        ],
        "port": "8888",
        "path": [
            "organisation",
            "{{organisation_id}}",
            "topics"
        ]
    },
    "description": "http://localhost:8080/organisation/{{organisation_id}}/topics"
}
```


#### Responses:
* `status: 200`
```json
{
    "id": "5b76b9b288987601f64b10df",
    "content": "Example content",
    "votes_up": 0,
    "votes_down": 0,
    "organisation": "5b76b95488987601f64b10dd",
    "created_at": "2018-08-17T12:04:02.4283777Z"
}
```


### Vote up
```json
"request": {
    "method": "POST",
    "header": [
        {
            "key": "X-Auth-Token",
            "value": "{{token}}"
        }
    ],
    "body": {
        "mode": "formdata",
        "formdata": [
            {
                "key": "vote_type",
                "value": "1",
                "description": "",
                "type": "text"
            }
        ]
    },
    "url": {
        "raw": "http://localhost:8888/organisation/{{organisation_id}}/topic/{{created_topic_id}}/votes",
        "protocol": "http",
        "host": [
            "localhost"
        ],
        "port": "8888",
        "path": [
            "organisation",
            "{{organisation_id}}",
            "topic",
            "{{created_topic_id}}",
            "votes"
        ]
    },
    "description": "http://localhost:8080/{{organisation_id}}/topics/{{created_topic_id}}/vote"
}
```

#### Responses
* `status: 200`
```json
{
    "id": "5b76cb048898760201722f65",
    "content": "Example content",
    "votes_up": 1,
    "votes_down": 0,
    "organisation": "5b76b95488987601f64b10dd",
    "created_at": "2018-08-17T13:17:56.858Z"
}
```

* `status: 409`
```json
{
    "message": "User already voted",
    "status": "fail"
}
```

### Vote down
```json
"request": {
    "method": "POST",
    "header": [
        {
            "key": "X-Auth-Token",
            "value": "{{token}}"
        }
    ],
    "body": {
        "mode": "formdata",
        "formdata": [
            {
                "key": "vote_type",
                "value": "-1",
                "description": "",
                "type": "text"
            }
        ]
    },
    "url": {
        "raw": "http://localhost:8888/organisation/{{organisation_id}}/topic/{{created_topic_id}}/votes",
        "protocol": "http",
        "host": [
            "localhost"
        ],
        "port": "8888",
        "path": [
            "organisation",
            "{{organisation_id}}",
            "topic",
            "{{created_topic_id}}",
            "votes"
        ]
    },
    "description": "http://localhost:8080/{{organisation_id}}/topics/{{created_topic_id}}/vote"
}
```

#### Responses
* `status: 200`
```json
{
    "id": "5b76cb048898760201722f65",
    "content": "Example content",
    "votes_up": 0,
    "votes_down": 1,
    "organisation": "5b76b95488987601f64b10dd",
    "created_at": "2018-08-17T13:17:56.858Z"
}
```

* `status: 409`
```json
{
    "message": "User already voted",
    "status": "fail"
}
```

### Get topic votes
```json
"request": {
    "method": "GET",
    "header": [
        {
            "key": "X-Auth-Token",
            "value": "{{token}}"
        }
    ],
    "body": {},
    "url": {
        "raw": "http://localhost:8888/organisation/{{organisation_id}}/topic/{{created_topic_id}}/votes",
        "protocol": "http",
        "host": [
            "localhost"
        ],
        "port": "8888",
        "path": [
            "organisation",
            "{{organisation_id}}",
            "topic",
            "{{created_topic_id}}",
            "votes"
        ]
    },
    "description": "http://localhost:8080/organisation/{{organisation_id}}/topic/{{created_topic_id}}/votes"
}
```

#### Responses
* `status: 200`
```json
[
    {
        "id": "5b76b9ef88987601f64b10e0",
        "user_token": "85e8da1489d4026d876b82a0d5ea6041",
        "topic": "5b76b9b288987601f64b10df",
        "vote_type": 1,
        "created_at": "2018-08-17T12:05:03.81Z"
    },
    {
        "id": "5b76b9ef88987601f64b10e0",
        "user_token": "8278da1489d4026d876b82a0d5ea77sa",
        "topic": "5b76b9b288987601f64b10df",
        "vote_type": -1,
        "created_at": "2018-08-17T12:12:05.42Z"
    }
]
```

### Fails
* Wrong organisation or topic in URL  
`status: 422`
```json
{
    "message": "Invalid {{param}} ID",
    "status": "fail"
}
```

* Empty or missing params  
`status: 422`
```json
{
    "message": "Parameter {{param}} cannot be empty",
    "status": "fail"
}
```

* Wrong auth token  
`status: 401`
```json
{
    "status": 401,
    "message": "Not authorized"
}
```