This time, [ssrf.go](ssrf.go) explicitly checks the hostname. So, if we directly use `https://example.com` or `https://<internal-ip>/InternalEndpoint`, it will fail since it does not match the expected hostname. This may look like a perfect fix, but it actually is not.

![SSRF04](images/ssrf04.png)

**Payload:**
```
http://localhost:8080/?ssrf=https://vercel-open-redirect.vercel.app/api/redirect?url=https://example.com
```

In this case:
* The hostname extracted is `vercel-open-redirect.vercel.app`
* The hostname check passes
* The request proceeds
* The second `302` redirect to `https://example.com` is never evaluated by the hostname validation. 

Server fetches (`example.com` or internal IPs) → Server returns the fetched response back to the user’s browser. Hence, bypassed! 

* Ready for the next one? Go to [Case IV](/SSRF/Case%20IV/)