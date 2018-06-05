# Validator for Go

Parse validation definition JSON file, and check object.

## Predefine Types and Validator

Keyword | Type
-|-
_string | String
_json | JSON format string
_in | Value must in on of the Array

## Validation Definition Sample

``` json
{
  "headers": {
    "Content-Type": "application/json"
  },
  "requestContext": {
    "authorizer": {
      "jwt_info": {
        "_type": "json",
        "app_id": {
          "_type": "string"
        },
        "username": {
          "_type": "string"
        },
        "scope": {
          "_type": "string",
          "_in": ["tracmo"]
        }
      }
    }
  },
  "body": {
    "_type": "json",
    "device_id": {
      "_type": "string"
    }
  }
}
```

## Sample Input Validation Definition

## Reference

1. [Data validation library for golang](https://github.com/gima/govalid)