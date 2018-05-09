# Exercise 5


Something is adding an additional "/" to each traverse loop.
Hence it will traverse `www.example.com/`
Then will try `www.example.com//` etc



## Cases

```go
./ -- Ignore. Return base site
./about/ -- Go to base + about
```


## Outputting as XML

Create a struct specifying the shape.
Add tags to fill.

See struct for adding a namespace

**Example**
```xml
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>http://www.example.com/</loc>
  </url>
  <url>
    <loc>http://www.example.com/dogs</loc>
  </url>
</urlset>
```

## Getting the domain

Domain can be obtained from the response object of the get request.

```go
resp, err := http.GET("www.example.com")

reqURL := resp.request.URL

baseURL := &url.URL{
  Scheme: reqUrl.Scheme,
  Host: reqURL.Host,
}
```


## The seen map

Can use `seen := make(map[string]struct{})`, rather than a map[string]bool. An empty struct takes up no memory.


## BFS.

Currently this doesn't use a BFS.

BFS would create two queues. One of links to explore. A second of the next round of links to explore.
Once the first queue is spent, the second queue becomes the primary and a new secondary queue is created

