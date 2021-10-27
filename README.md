# golang-simple-server
This project handles various requests to manipulate a key. Moreover it can backup and restore a key upon restarting. The app is also deployed at  [heroku.](https://immense-ravine-94498)
# Running
To run the project clone it to your local and go to the apps directory. Run 

``` 
go run main.go 

```
# Testing

App utilizes unit testing to simply detect any issues during the development process. To test the app on your local firstly run the server and then, run

```
go test
``` 

# Endpoints

- setKey 
- getKey 
- flush

## setKey
This endpoint only gets a put request. The key meant to be set is sent through the request and it basically sets the key to the given value. The request must contain a key value pair and the key should be "key". The request template can be seen below.

```
curl -X PUT -H "Content-Type: application/json"  -d '{"key": "ege"}' https://immense-ravine-94498.herokuapp.com/setKey
```

## getKey
This endpoint only gets a get request. Simply do a request at the API endpoint and receive the key if it is set.

```
curl -v https://immense-ravine-94498.herokuapp.com/getKey
```

## flush
This endpoint only gets a delete request. It simply flushes the key at the server.

```
curl -X DELETE https://immense-ravine-94498/flush
```

The template requests can be seen in the requests directory in the project. Run the scripts to see the results. Do not forget to change the domain of the endpoint according to where you run the server. If at your local use "localhost:PORT", if at heroku use the provided heroku [url.](https://immense-ravine-94498)