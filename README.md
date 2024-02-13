# Knapsack Problem (modified)

## Problem Statement
Imagine for a moment that one of our product lines ships in various pack sizes: 
- 250 items
- 500 items
- 1000 items
- 2000 items

Our customers can order any number of these items through our website, but they will always only be given complete packs. 
1. Only whole packs can be sent. Packs cannot be broken open. 
2. Within the constraints of Rule 1 above, send out no more items than necessary to fulfil the order. 
3. Within the constraints of Rules 1 & 2 above, send out as few packs as possible to fulfil each order. 

*(Please note, rule #2 takes precedence over rule #3)*

To further illustrate the rules above, please consider this custom pack size example:
- 23 items
- 31 items
- 53 items

Items Order : 263
Correct Number of packs: 2x23, 7x31
Incorrect answer: 5x53

## For running and testing the code locally
### Running the code
The code is written in Go. To run the code, please use the following command:
```bash
make run
```

### Running the tests
```bash
make test
```

### Running the linter
```bash
make lint
```

## Source code is deployed to server and running for testing on online environment

### API Endpoint - For creating items (pack sizes)
items - comma separated values, for example: 23,31,53. You can change the values as per your requirement.
```curl
curl -X POST --location "http://134.122.99.1:8080/api/items" \
    -H "Content-Type: application/json" \
    -d '{
          "items": "23,31,53"
        }'
```

### API Endpoint - For creating order
263 is ordered items for the following request. Then endpoint returns the number of packs we need to ship to the customer.
```curl
curl -X GET --location "http://134.122.99.1:8080/api/calculate/263"
```

### API Endpoint - For getting all items
```curl
curl -X GET --location "http://134.122.99.1:8080/api/items"
```