# URL shortner backend


Url shortner backend implementation. 
Basic functionalities like 
1. Short to long url
    ```For this, user just try to access the given url and we will have to redirect the user to new url. 
    url ex: www.domain.com/123123 => www.facebook.com/avinash.kumar123123
    
    curl: curl http://www.domain.com/123123
        redirect 
    ```


2. Long url to short 
    ```
    For this, user comes to the page and types another url. And we return a copyable response. 
    url ex: www.facebook.com/avinash.kumar123123 => www.domain.com/123123
    ```

Ex: 
www.facebook.com/avinash.kumar123123 => www.domin.com/123123

And vice versa. 

DB: It uses mysql db for storing the data related to short and long urls. 

Tables
```
url_shortner
___________________
id        | uuid/string
short_url | string (indexed based on this)
long_url  | string
```



How does the flow work in FE? 

Lets assume our domain for website is like this. www.urlshortner.com

now the portal opens the given url only. 

Now when a user goes and put a new url, they get a url like www.urlshortner.com/{xxxxx}
When we get a call to such urls, internally it will have to call the backend with the same url, since we are building it for backend call. 

But internal URL for backend will not be same as that of frontend right? 
