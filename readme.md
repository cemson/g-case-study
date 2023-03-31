# Getir GoLang Case Study
This is a Golang application that handles a provided mongodb read and saving and reading from an in-memory db.

The application is deployed to an EC2 instance and you can publicly reach and call endpoints from the address below
```
http://ec2-3-83-238-113.compute-1.amazonaws.com:11923
```

## Setup
Before starting the application, please make sure to set up the conf.json file with the necessary configuration parameters.

## Running the Application
To run the application, under application folder use the following command:

```
go get .
go run . prod
```

## Endpoints
The following endpoints are available in the application:

### Health-check Endpoint

This is the health-check endpoint for checking the application up and running

```http request
curl --location 'http://127.0.0.1:11923/test'
```

Sample response
```json
{
    "message": "g-case-study rest api is well alive..."
}
```

### MongoDB Read Endpoint

This is the mongodb read endpoint. A sample request with all fields of body included is below

```http request
curl --location 'http://127.0.0.1:11923/records' \
--header 'Content-Type: application/json' \
--data '{
    "startDate": "2016-01-28",
    "endDate": "2018-02-02",
    "minCount": 2800,
    "maxCount": 3000
}'
```

Sample response

```json
{
    "AdditionalData": null,
    "Data": [
        {
            "key": "TAKwGc6Jr4i8Z487",
            "createdAt": "2017-01-28T01:22:14.398Z",
            "totalCount": 2800
        },
        {
            "key": "NAeQ8eX7e5TEg7oH",
            "createdAt": "2017-01-27T08:19:14.135Z",
            "totalCount": 2900
        }
    ],
    "Message": "Request handled successfully",
    "IsCompletedSuccessfully": true,
    "IsValidationError": false,
    "NotValidFields": null,
    "IsUnexpectedError": false
}
```
### In-Memory Save Endpoint

This is the endpoint that saves the produced data to in-memory db.

Sample request:
```http request
curl --location 'http://127.0.0.1:11923/keyval' \
--header 'Content-Type: application/json' \
--data '{
    "key": "2",
    "value": "some_data"
}'
```

Sample response:
```json
{
    "AdditionalData": null,
    "Data": {
        "Key": "2",
        "Value": "some_data"
    },
    "Message": "Request handled successfully",
    "IsCompletedSuccessfully": true,
    "IsValidationError": false,
    "NotValidFields": null,
    "IsUnexpectedError": false
}
```

### In-Memory Read Endpoint
This is the endpoint for getting the stored data from previous request.

Sample request:
```http request
curl --location 'http://127.0.0.1:11923/keyval?key=2'
```

Sample response:
```json
{
    "AdditionalData": null,
    "Data": {
        "Key": "2",
        "Value": "some_data"
    },
    "Message": "Request handled successfully",
    "IsCompletedSuccessfully": true,
    "IsValidationError": false,
    "NotValidFields": null,
    "IsUnexpectedError": false
}
```