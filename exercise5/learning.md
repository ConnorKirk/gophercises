# Exercise 5


Something is adding an additional "/" to each traverse loop.
Hence it will traverse `www.example.com/`
Then will try `www.example.com//` etc



## Cases

```go
./ -- Ignore. Return base site
./about/ -- Go to base + about
```