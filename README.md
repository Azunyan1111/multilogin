# MultiLogin
## What’s MultiLogin?
It is SaaS.
Provide login function to the service.

Service developers establish links and get user information by calling API.

### Use technology 
```
Go lang (echo)
Go Test
MVC model
MySQL + (ormapper:gorm)
Redis
Docker-compose
Makefile
```

### Security

```
CSRF Token
Cookie secure
HMAC-SHA256 (API)
Response (example XSS Protection Header)
https
```

## Get Start
Only the beginning.
```
make setup
```
Set docker-compose environment variable

sendgrid is mail sender service.<a href="https://app.sendgrid.com">https://app.sendgrid.com</a>

docker-conpose.yml 
```
export SENDGRID_API_KEY="your SENDGRID API KEY" &&
```
If you do not use it please change this return to nil

sendMail.go
```
	if err != nil {
		return err // return nil
	}
```
## How to use
Start up the server
```
make start
```
Access to the local server

<a href="http://localhost:8080">http://localhost:8080</a>

### Service developer example
GetName API example for python.

```main.py
# -*- coding: utf-8 -*-

import hashlib
import hmac
import urllib2
from datetime import datetime

if __name__ == '__main__':
    # String to append to 'authorization' in header. Separate by commas when it becomes different information
    authorization = ""
    # MultiLogin URL
    url = "http://localhost:8080"
    # API Request Parameter The URL to actually send the request is `http:#localhost:8040/api/user/name?uuid=26d2983e-3d5a-421c-bf6f-d4608025e555`
    param = "/api/user/name?uuid=26d2983e-3d5a-421c-bf6f-d4608025e555"
    # 36 character token and secret displayed on My Page after registering service
    token = "your token"
    secret = "your secret"
    # Unix time when sending request
    timeStr = datetime.now().strftime('%s')

    # Declare ML's proprietary authentication method
    authorization += "MLAuth1.0" + ","
    # Write a token
    authorization += "token=" + token + ","
    # Write Unix time
    authorization += "time=" + timeStr + ","

    # Create parameter string by connecting parameter URL and Unix time with a comma
    join = param + "," + timeStr
    # Generate hash
    hash = hmac.new(secret, join, hashlib.sha256)
    # Describe the generated hash
    authorization += "signature=" + hash.hexdigest() + ","

    req = urllib2.Request(url + param)
    # Header setting
    req.add_header('authorization', authorization)

    res = urllib2.urlopen(req)
    r = res.read()
    print r
```

Examples of Golang and Python and PHP

<a href="https://github.com/Azunyan1111/ml-example">https://github.com/Azunyan1111/ml-example</a>
