# luhn-service
Simple API built using Go to validate identification numbers according to the [Luhn Algorithm](https://en.wikipedia.org/wiki/Luhn_algorithm)

## Example Request
```
/validate?number=123
```

## Example Response
```
{
  "valid": "true"
}
```
