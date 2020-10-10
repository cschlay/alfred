## Deployment

Build it
```
docker-compose exec web env GOOS=linux GOARCH=386 go build -o bin/main main.go 
```

Put the program `main` to server, `chmod +x main` if not already and reverse proxy to `localhost:8000` .
