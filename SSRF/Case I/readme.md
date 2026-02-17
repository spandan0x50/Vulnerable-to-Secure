Take a close look at `ssrf.go`. There’s a very simple code snippet that does the following:

1. Accepts a URL from the user via the `?ssrf` parameter.
2. Fetches that URL server-side.
3. Renders the fetched response.

Yes, a classic [Server-Side Request Forgery (SSRF)](https://portswigger.net/web-security/ssrf) scenario.

![SSRF](images/ssrf.png)
> User’s browser (`example.com`) → Server fetches (`example.com`) → Server returns the fetched response back to the user’s browser.

We'll look into multiple ways on how developers tries to fix this issues, why one solution may work on one scenario but not on other, and how attackers bypasses the fix. 