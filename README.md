# StackIT challenge 
## Disclaimer 
* No incremental code pushes were made 
* No tests were written
* No relevant comments were added (given that highly unstable so far)
* Borrowed a lot of code/architecture from my [previous projects](https://github.com/yago-123/galeLB)
* **No validation was added when it comes to persisting data** (duplicates) 
* **No validation was added data** (email parsing, etc)
* Used Redis for "persistence"
  * Persistence is not guaranteed really 
  * Bad option for retrieving all users when it comes to performance
* **Relation between encoding and storage must be revisited** when it comes to coupling  
* Used these prompts 
  * [Handling errors for user registration and curl commands](https://chatgpt.com/share/6800d5be-705c-800a-91f2-0e39eb5da8ae)
  * [Filling the Redis persistence](https://chatgpt.com/share/6800d603-f844-800a-8322-72d1368621a9)
* Package `types` needs to be revisited, I would think twice about the design pattern of the app 

## Description
Running the app 
```bash 
$ make run 
```

# API v1
## Register new user 
```bash 
$ curl -X POST http://localhost:8080/api/v1/user \
  -H "Content-Type: application/json" \
  -d '{
    "FirstName": "Felix",
    "LastName": "Schneider",
    "Email": "felix@schneider.com"
  }'
```

## Retrieve all users registered 
```bash 
$ curl -X GET http://localhost:8080/api/v1/users 
```

## TODO
- [ ] Comment code
- [ ] Add proper logging 
- [ ] Add test cases 