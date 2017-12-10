# router
gin wrapper to use http.Handler and http.HandlerFunc

# example

```
func HomePage(w http.ResponseWriter, r *http.Request) {
	c := router.ContextW(w)
	...
}

func init() {
	mux := router.New()
	mux.Any("/home.html",
		HomePage)
}
```