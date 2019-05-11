# {name}

{go:header}

* Cache arbitrary data with an expiration time
* 0 dependencies
* Less than 100 lines of code
* 100% test coverage

## Usage

```go
// New cache
c := cache.New(5 * time.Minute)

// Put something into the cache
c.Set("a", "b", 1 * time.Minute)

// Read from the cache
obj, found := c.Get("a")

// Convert the type
fmt.Println(obj.(string))
```

{go:footer}
