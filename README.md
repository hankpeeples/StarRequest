# StarRequest

StarRequest is a Postman alternative run strictly from the command line. Requests are configured via either `*.sr.json`, `*.sr.yaml`, ...(more formats to come?).

## JSON Request config layout

Requests configured via JSON need to be written in the following manner.

```json
{
  "requests": [
    {
      "name": "Example Request",
      "url": "http://example.com",
      "method": "GET"
    }
  ]
}
```
