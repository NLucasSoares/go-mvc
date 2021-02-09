# Example

## Goal

Show how to do MVC simple application in go using gorilla mux.

## Request

### Request with invalid home name causing not found error

`curl http://127.0.0.1:8080/home/ff` =>

```json
{
	"code": 1234,
	"message": "VIEW: Home not found: not found"
}
```

### Request with valid home name using `Home` view

`curl http://127.0.0.1:8080/home/first` =>

```json
{
	"name": "first",
	"address": "44 bis rue de puiseux 95490 VAUREAL"
}
```
