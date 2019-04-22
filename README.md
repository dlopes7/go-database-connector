# go-database-connector

The idea is that this is a binary that can connect and query multiple databases

For now, oracle is supported.

Usage: 

```
LD_LIBRARY_PATH=/path/to/oracle_client_libraries ./go-database-connector -hostname=localhost -username=david -password=password -sid=ORCLPDB1 -query="SELECT CUSTOMER_NAME, CUSTOMER_ID, CUSTOMER_AGE, CITY FROM myschema.CUSTOMERS"
```

Response:

```json
{
   "error": false,
   "errorMessage": "",
   "rows": [
      {
         "columns":[
            {
               "name":"CUSTOMER_NAME",
               "value":"David"
            },
            {
               "name":"CUSTOMER_ID",
               "value":"1"
            },
            {
               "name":"CUSTOMER_AGE",
               "value":"30"
            },
            {
               "name":"CITY",
               "value":"Sao Paulo"
            }
         ]
      },
      {
         "columns":[
            {
               "name":"CUSTOMER_NAME",
               "value":"Ady"
            },
            {
               "name":"CUSTOMER_ID",
               "value":"2"
            },
            {
               "name":"CUSTOMER_AGE",
               "value":"33"
            },
            {
               "name":"CITY",
               "value":"Sao Paulo"
            }
         ]
      },
      {
         "columns":[
            {
               "name":"CUSTOMER_NAME",
               "value":"Andre"
            },
            {
               "name":"CUSTOMER_ID",
               "value":"3"
            },
            {
               "name":"CUSTOMER_AGE",
               "value":"0"
            },
            {
               "name":"CITY",
               "value":"Sao Paulo"
            }
         ]
      }
   ]
}
```