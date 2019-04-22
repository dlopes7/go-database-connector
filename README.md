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
   "error":false,
   "errorMessage":"",
   "rows":[
      {
         "columns":[
            {
               "index":0,
               "name":"CUSTOMER_NAME",
               "value":"David"
            },
            {
               "index":1,
               "name":"CUSTOMER_ID",
               "value":"1"
            },
            {
               "index":2,
               "name":"CUSTOMER_AGE",
               "value":"30"
            },
            {
               "index":3,
               "name":"CITY",
               "value":"Sao Paulo"
            }
         ]
      },
      {
         "columns":[
            {
               "index":1,
               "name":"CUSTOMER_ID",
               "value":"2"
            },
            {
               "index":2,
               "name":"CUSTOMER_AGE",
               "value":"33"
            },
            {
               "index":3,
               "name":"CITY",
               "value":"Sao Paulo"
            },
            {
               "index":0,
               "name":"CUSTOMER_NAME",
               "value":"Ady"
            }
         ]
      },
      {
         "columns":[
            {
               "index":0,
               "name":"CUSTOMER_NAME",
               "value":"Andre"
            },
            {
               "index":1,
               "name":"CUSTOMER_ID",
               "value":"3"
            },
            {
               "index":2,
               "name":"CUSTOMER_AGE",
               "value":"0"
            },
            {
               "index":3,
               "name":"CITY",
               "value":"Sao Paulo"
            }
         ]
      }
   ]
}
```