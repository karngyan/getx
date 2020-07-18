# GetX

## gist

- Given url of a webpage, server fetches the html, dumps it to a file and serves it to the client.
- If the fetch fails, server sends back the expected path of the file to the client and spawns 
a go routine at the same time which retries fetching the webpage every hour, `retryLimit` number of times.
 
 
##  milestones

- [x] fetch html by url and dump it to a file
- [ ] retry go routine
- [ ] implement a task queue (may be persist it)
- [ ] frontend



## run

`go run main.go`

- Once the server is on, try the following request:

 ```
curl --location --request POST 'http://localhost:7771/pagesource' \
--header 'Content-Type: application/json' \
--data-raw '{
    "uri": "https://google.com",
    "retryLimit": 3
}'
```

- Sample Response: 201 Created

```
{
    "id": "d27913f6-c8ea-11ea-81ea-9828a617c8db",
    "uri": "https://google.com",
    "sourceUri": "files/d27913f6-c8ea-11ea-81ea-9828a617c8db.html",
    "retryLimit": 3
}
```

- visit pagesource at: `localhost:7771/sourceUri`
- SourceUri is also in the `Location` header of the response

 ## todo
 
 - Retries go routine
 - Task queue
 - Frontend
 
 