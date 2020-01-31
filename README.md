# lolo-sdk-go
Official Lolo SDK for the Go programming language.

## Usage
```
client, err := lolo.NewClient(apiKey)
if err != nil {
  log.Fatalf("fatal: %s\n", err)
}

app := lolo.App{
  Name: "My app",
  Description: "Hello Gopher!",
}

err = client.CreateApp(&app)
if (err != nil) {
  log.Fatalf("fatal: %s\n", err)
}
log.Printf("Created app %s", app.Id);
```
