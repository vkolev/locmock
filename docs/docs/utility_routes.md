# Utility routes

Utility routes include helper routes that are preconfigured to responde
dynamically.

## Ip

Returns the requesters IP

**Route:**

- `https://localhost:8080/l/ip`
- Method: `GET`
- Query parameters:
  - *ipv6* - value `true`

**Response:**

```text
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
Date: Mon, 21 Aug 2023 05:10:14 GMT
Content-Length: 9
Connection: close

127.0.0.1
```

## Ping

Returns a simple `pong`

**Route:**

- `https://localhost:8080/l/ping`
- Method: `GET`
- Query parameters:
  - None

**Response:**

```text
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
Date: Mon, 21 Aug 2023 05:10:00 GMT
Content-Length: 4
Connection: close

pong
```

## Person

Returns a random generated user profile

**Route:**

- `https://localhost:8080/l/person`
- Method: `GET`
- Query parameters:
    - gender - `[male, female]`

Response:

```json
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Mon, 21 Aug 2023 05:09:37 GMT
Content-Length: 1079
Connection: close

{
  "gender": "female",
  "name": {
    "first": "Ella",
    "last": "Martin",
    "title": "Mrs"
  },
  "location": {
    "street": "43 Adams Ct",
    "city": "Kingsbridge",
    "state": "West Virginia",
    "postcode": 62599
  },
  "email": "ella.martin@example.com",
  "login": {
    "username": "Biteheather",
    "password": "Foxwax",
    "salt": "1WuFJDMCMU7tThqe",
    "md5": "5a2d6d06c17dd26a907fd78a6e6de7b3",
    "sha1": "mpszm9kqrFb11yVX7ErumMRZDkE=",
    "sha256": "D1qOLZoH-EEcB_4r2iGe4b5UvB76imsSFU8e6c9RM64="
  },
  "dob": "Sunday 24 Sep 2023",
  "registered": "Sunday 26 Feb 2023",
  "phone": "+266 5 7 4861 86755",
  "cell": "+  054 0 4743844998",
  "id": {
    "name": "SSN",
    "value": "355-3-8562"
  },
  "picture": {
    "large": "https://randomuser.me/api/portraits/women/2.jpg",
    "medium": "https://randomuser.me/api/portraits/med/women/2.jpg",
    "thumbnail": "https://randomuser.me/api/portraits/thumb/women/2.jpg"
  },
  "nat": "US"
}
```

## User-Agent

Returns a random User-Agent string

**Route:**

- `https://localhost:8080/l/user-agent`
- Method: `GET`
- Query parameters:
    - None

**Response:**

```text
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
Date: Mon, 21 Aug 2023 05:09:15 GMT
Content-Length: 98
Connection: close

Opera/10.61 (J2ME/MIDP; Opera Mini/5.1.21219/19.999; en-US; rv:1.9.3a5) WebKit/534.5 Presto/2.6.30
```

## UUID

Returns a random `UUID` - per default it will be UUIDv4

**Route:**

- `https://localhost:8080/l/uuid`
- Method: `GET`
- Query parameters:
  - *type* - one of `[v1, v3, v4, v5]`

**Response:**

```text
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
Date: Mon, 21 Aug 2023 05:33:45 GMT
Content-Length: 36
Connection: close

20747749-c446-4bd6-8468-477c52c923b8
```

## Form Request

Accepts a Form or Multipart-Form request and returns the requested parameters as JSON

**Route:**

- `https://localhost:8080/l/form`
- Method: `POST`
- Headers: 
  - Multipart: `Content-Type: multipart/form-data; charset=utf-8;`
  - Normal: `Content-Type: application/x-www-form-urlencoded; charset=utf-8`

**Response:**

```json
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Mon, 21 Aug 2023 05:38:02 GMT
Content-Length: 547
Connection: close

{
    "file": null,
    "form": {},
    "headers": {
        "Authorization": [
            "Testing123"
        ],
        "Connection": [
            "close"
        ],
        "Content-Length": [
            "0"
        ],
        "Content-Type": [
            "application/x-www-form-urlencoded; charset=utf-8"
        ],
        "Cookie": [
            "JSESSIONID=440E8B50A752FEEBBD979C4A63D57239"
        ],
        "User-Agent": [
            "RapidAPI/4.2.0 (Macintosh; OS X/13.4.1) GCDHTTPRequest"
        ]
    },
    "method": "POST"
}
```

## Redirect

Redirect a request to a given URL

## Gzip

Returns a gzip response

## Generic echo

Users the path parameters to validate the request method and returns 
the given body, query parameters back as json response.

