# echo-rater
rate limit demo for echo, custom limiter by request header and others

assume we have 2 URLs :

- /test/maxQPS_10
- /test/maxQPS_20

Here is how we define it in code:

conf := []PathLimiter{
  {
    Path: "/test/max1",
    Max:  float64(10),
  },
  {
    Path: "/test/max2",
    Max:  float64(2),
  },
}

Use GatewayLimitMiddleware make ratelimiter

for _, c := range conf {
  e.GET(c.Path, TestHandler, GatewayLimitMiddleware(c.Max))
}
