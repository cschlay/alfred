## Deployment

Build it, replace `env` with `set` and put to separate lines  if using windows.
```
env GOOS=linux GOARCH=386 go build main.go
```

Put the program to server and reverse proxy to `localhost:8000` .
