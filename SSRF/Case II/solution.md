[ssrf.go](ssrf.go) only checks whether the input string contains the substring `vercel.app` anywhere. It does not check the actual hostname. Because of that, an attacker can bypass this in multiple ways:

* `https://x.vercel.app@example.com` (typical open-redirect style bypass)
* `https://x.vercel.app.example.com` (subdomain-based bypass)
and many more ways.

This ensures the supplied URL contains the substring `vercel.app`, but the URL is not actually pointing to `vercel.app`. Instead, it is pointing to a third-party domain or internal ip (for further exploitation).

![SSRF01](images/ssrf02.png)

* Go to [Case III](/SSRF/Case%20III/)