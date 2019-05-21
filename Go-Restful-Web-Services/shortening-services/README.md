# Shortening Services 

URL shortening services are built upon two things: 

* A string mapping algorithm to map long strings to short strings (Base62)
* A simple web server that redirects a short URL to the original URL 

## API Design 

|URL|REST Verb|Action|Success|Failure|
|-|-|-|-|-|
|/api/v1/new|POST|Create a shortened URL|200|500,404|


