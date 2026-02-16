`ssrf.go` contains a **Server-Side Request Forgery (SSRF)** vulnerability.

{ssrf_1.png}
> User’s browser (`example.com`) → Server fetches (`example.com`) → Server returns the fetched response back to the user’s browser. User's browser doesn't fetch `example.com`.

There are multiple ways to mitigate this SSRF issue. One of the most effective approaches is implementing a strict allowlist using a `map`. For example:

```
allowedDomain := map[string]bool{
    "spandanpokhrel.com.np": true,
}
```

Then validate the parsed hostname before making the request:

```
if !allowedDomain[parseUrl.Hostname()] {
    return
}
```

With this check in place, the server will only fetch resources from explicitly whitelisted domains. This prevents attackers from abusing the functionality to access internal resources such as `localhost`, `http://169.254.169.254/latest/meta-data/`, or `http://169.254.169.254/metadata/instance`.