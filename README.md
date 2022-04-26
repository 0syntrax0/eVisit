## What would you do differently if you had more time?

- Have a working service
- Add cache to better store and manipulate data
- Group IP addresses by their splitted Network ID
- - For example, `127.0.0.1`
- - ["127"]["0"]["0"]["1"] = visit count

## What is the runtime complexity of each function?

`O(1)`

## How does your code work?

Like a dream!

### Record an IP
Record an IP, update visit's count, and update the current top visitor's count.

URL:
`POST("/registerIP")` 

Payload:
- `ip: string` required

Example:
```json
{
    "ip": "127.0.0.1"
}
```

Response:
- `StatusAccepted: 202`
- - _No body_
- `StatusBadRequest: 400`
- -  "error": "error message"

### Fetch top 100
Returns a list of top 100 visitors.

URL:
`GET("/top100/")`

Response:
- `StatusOK: 200`
- - "list": [visit counts][IP address]string

### Clear list
Clears current list of IP addresses and visit counts.

URL:
`POST("/clear")`

Response:
- `StatusResetContent: 205`
- - _No body_

## What other approaches did you decide not to pursue?

- Completing service
- - I tried to, but it quickly grew in complexity and would've taken at least a week to complete as is documented

## How would you test this?

If testing is needed, pinging `PUT("/seed")` will generate less than 20,000,000 random IP addresses (accounting for duplicates that will be counted to wards the visit count).
But this will take at least 10 minutes depending on the CPU speed and memory size.

Then using the endpoints above to add, fetch top 100 or clear IPs
