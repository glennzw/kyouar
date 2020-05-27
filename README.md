# Kyouar
    
Kyouar (pronounced 'QR') is a tiny endpoint to create a QR code from a supplied URL.

For example:  
https://kyouar.herokuapp.com/u/www.google.com/ 
Returns QR image for http://www.google.com  
https://kyouar.herokuapp.com/u/https://www.google.com Returns QR code for https://www.google.com (optional leading http(s))

Basic URL validation is performed, but a lot of leeway is given. A response of "Invalid URL!" with a 400 status code is returned on bad URLs, or a 500 on failure to generate a QR code. 

## Why, tho?

Useful for dyamically creating QR codes, e.g. in HTML source:

```
<html>
    Please scan the QR code below to get a free Bitcoin!</br>

    <img src='https://kyouar.herokuapp.com/u/https://i.gifer.com/ZDRD.gif>
</html>
```

![alt text](example.png "Free Bitcoin!")

## Todo

Add JSON endpoint.  
Add parameters for customizing QR code.