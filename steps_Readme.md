Hi.

This is a webservice for documenting the API for Receipt Processor.

Here are the Steps to run this webservice

1. Clone this repository:

```
git clone https://github.com/NavyaSreeKv11/receipt-processor-challenge.git
```

2. Build the Docker image using following command in the terminal

```
docker build -t receipt_processor .
```

3. Run the Docker container:

```
docker run -dp 8080:8080 receipt_processor
```

4. You can check the APIs using Postman or cUrl using command or any online curl editor
```
curl -X POST -H "Content-Type: application/json" -d '{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": 6.49
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": 12.25
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": 1.26
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": 3.35
    },{
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": 12.00
    }
  ],
  "total": 35.35
}' http://localhost:8080/receipts/process

```

5. Save the 'id' retrieved from the above curl and replace {id} with that id number in this command.

```
curl -X GET http://localhost:8080/receipts/{id}/points
```
