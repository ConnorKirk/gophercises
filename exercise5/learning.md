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