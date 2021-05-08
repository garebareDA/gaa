# gaa  
小さなWeb フレームワークです。  
# Example
```go
s := gaa.New()
s.Get("/", func(w http.ResponseWriter, r *http.Reqest, u url.) {
  log.Println!("Get!")
})
s.Run(":8080")
```
