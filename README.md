# Weather CLI

A little CLI application that can quickly tell me
the weather when I start my terminal in the morning.

## About 
This CLI tool is built with Go and uses the Google Maps Weather API. It requires an 
API key. It currently pulls a key from a file that is stored outside of this repository. 
You can make a `weather-api-key.txt` file inside of a `keys` directory that is located 
one directory above the compiled app.
```
.
|-keys/
|   |_weather-api-key.txt
|-weather-cli/
    |_weather

```

## Build
You can build you own version of this app with the following: 
```
go build -o weather main.go
```

## Running
You can use the compiled 'weather' file that is in this repository, or you can build
it yourself. 

### Flags

| Flag     | Required     | Description       |
| -------- | ------------ | ----------------- |
|  -city   | Yes          | Retreive the weather for this city. |
|  -hours  | No           | Number of hours of weather to retreive. Default, 12. Max, 24 |


### Example 
`./weather -city Cincinnati -hours 24`

If you add `weather` to your PATH, then you'll be able to run 
'weather -city Cincinnati' from anywhere in your terminal.

