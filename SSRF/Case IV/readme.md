From the [previous solution](/SSRF/Case%20III/solution.md):
> * The second `302` redirect to `https://example.com` is never evaluated by the hostname validation. 

This time, we're checking such multi-step redirections too:
```
	finalURL:=resp.Request.URL
	finalH := finalURL.Hostname()
	if finalH=="vercel.app" || strings.HasSuffix(finalH, ".vercel.app"){}
```

![SSRF01](images/ssrf01.png)

So, the question is: 
> is it still possible to hit/bypass & fetch `example.com`? 
> If yes, how?

Hint: there are 2 different attack scenarios.

* Need help or want to verify your answer? Go to [Solution](/SSRF/Case%20IV/solution.md)
* Already found the bypass? Go to [Case IV](/SSRF/Case%20V/)