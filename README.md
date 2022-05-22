# echo-rater
rate limit demo for echo, custom limiter by request header and others

assume we have 2 URLs :

- /test/maxQPS_10
- /test/maxQPS_20

Here is how we define it in code:

```go

conf := []PathLimiter{
  {
    Path: "/test/maxQPS_10",
    Max:  float64(10),
  },
  {
    Path: "/test/maxQPS_20",
    Max:  float64(20),
  },
}
```

Use GatewayLimitMiddleware make ratelimiter

```go
for _, c := range conf {
  e.GET(c.Path, TestHandler, GatewayLimitMiddleware(c.Max))
}
```
