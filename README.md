## Airline

### Create a reservation
There should exist a sourceairport and targetairport in the flight list
````
curl localhost:8081/reservations | jq 
{
  "reservation": [
    {
      "id": 2,
      "paymentsuccessful": false,
      "passenger": {
        "firstname": "tony",
        "lastname": "rådyr",
        "passportnumber": 12345,
        "dateofbirth": "201292",
        "email": "",
        "creditcardnumber": 0
      },
      "flightinfo": {
        "flightnumber": "BOEING777",
        "operatingairlines": "Emirates",
        "sourceairport": "Oslo",
        "targetairport": "Cancun"
      }
    }
  ]
}

````
### Create a payment for an existing reservation. paymentid is the same as id in reservation

````
curl -X POST localhost:8081/payment/pay -d '{"paymentid": 2}'
````

### Update the reservation in order to set payment to true

````
curl -X PUT localhost:8081/reservation/2 -d '{"id": 2}'
````

### List the reservations made

````
curl -X GET localhost:8081/reservations | jq
````

### Get the available flights

````
curl -X GET localhost:8081/test | jq
````
