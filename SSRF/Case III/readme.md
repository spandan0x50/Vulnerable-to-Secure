This time in [ssrf.go](ssrf.go), the case is a bit different. We're now parsing the URL, extracting the hostname & then comparing it to a certain value eg, `vercel.app`. This simply means `https://x.vercel.app@example.com` or `https://x.vercel.app.example.com` will no longer work as a bypass because neither of these domain's hostname is `vercel.app`, so they'll fail the `if` condition. 

So, the question is: is it still possible to bypass? If yes, how? 

Need Help? Go to [Solution](/SSRF/Case%20III/solution.md)
Already found the bypass? Go to [Case IV](/SSRF/Case%20IV/)