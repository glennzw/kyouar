# Kyouar
    
Kyouar (pronounced 'QR') is a tiny endpoint to create a QR code from a supplied URL.

For example:  
https://kyouar.herokuapp.com/u/www.google.com/ 
Returns QR image for http://www.google.com  
https://kyouar.herokuapp.com/u/https://www.google.com Returns QR code for https://www.google.com (optional leading http(s))

Rather than a raw image, the base64 encoding of the QR cann be returned by, suitable to be inserted into an img tag. This is done by calling /b/<url>, rather than /u/<url>. e.g:

```
curl https://kyouar.herokuapp.com/b/www.google.com/ 

data:image/png;base64, iVBORw0KGgoAAAANSUhEUgAAAQAAAAEAAQMAAABmvDolAAAABlBMVEX///8AAABVwtN+AAABWUlEQVR42uyYsZUDMQhE8VOgkBIoRa1daSpFJRAq2LdzD2GvrTuf8xOawMH6RzzBDNDW1tbWJ2HoSJ1yJ1Zu4l++1gJO+0lH6hmdiLhJsy8lHHCzMh0JPXdWVhIrVY0KUEbP0FGoyECy1+KNsybgfUF5vIe/G2d14JqTuWf9OEjXBl7kffHBIVcGTrp5ZcwuWBmNpJbJN1cAbldjoDOg4g+iUjDgJKvUmA+A1UnQSp0HSAQAp5UqweIBMZRIWpkGSAzACnUkHDYflBjAb0NZAIDnA3gmHn9LpfLsizAAjedwl0ojmX0zBmB1wkG+LRJbX6CWH3kyCvDqm2YW5dUvlgDu8hfhKWjuiyDAIxUD13FAUN9sB4sDzyOJbQfXnEQ44LE1j5x0V10ZMFtUgaXBsMC4BSl7LJ7XqCDAdSShhyuGBJ5zciwH4zhQ3g/S/wxsbW1F1HcAAAD//7lb0+Lrsw6oAAAAAElFTkSuQmCC
```

This result could be displayed like so:

```
<img src="data:image/png;base64, iVBORw0KGgoAAAANSUhEUgAAAQAAAAEAAQMAAABmvDolAAAABlBMVEX///8AAABVwtN+AAABWUlEQVR42uyYsZUDMQhE8VOgkBIoRa1daSpFJRAq2LdzD2GvrTuf8xOawMH6RzzBDNDW1tbWJ2HoSJ1yJ1Zu4l++1gJO+0lH6hmdiLhJsy8lHHCzMh0JPXdWVhIrVY0KUEbP0FGoyECy1+KNsybgfUF5vIe/G2d14JqTuWf9OEjXBl7kffHBIVcGTrp5ZcwuWBmNpJbJN1cAbldjoDOg4g+iUjDgJKvUmA+A1UnQSp0HSAQAp5UqweIBMZRIWpkGSAzACnUkHDYflBjAb0NZAIDnA3gmHn9LpfLsizAAjedwl0ojmX0zBmB1wkG+LRJbX6CWH3kyCvDqm2YW5dUvlgDu8hfhKWjuiyDAIxUD13FAUN9sB4sDzyOJbQfXnEQ44LE1j5x0V10ZMFtUgaXBsMC4BSl7LJ7XqCDAdSShhyuGBJ5zciwH4zhQ3g/S/wxsbW1F1HcAAAD//7lb0+Lrsw6oAAAAAElFTkSuQmCC">
```

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