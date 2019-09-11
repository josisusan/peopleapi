# Peopleapi

First Golang API using [Granitic](http://www.granitic.io/)

## API includes following routes:
- Index route to list all the people info (/people GET)
- Create route to create a person (/people POST)
- Update route to make changes (/people/<person-uuid> PUT)

## DataStore:
- Using CSV File
- Using Local Dynamodb

## Getting Started:

### Using DynamoDB
To get started with dynamodb locally, there are couple of commands in the Makefile. So, in console run these commands:

1. Install dynamo using [Docker](https://docs.docker.com/install/).
  
      ` $ make install-dynamo `

2. Start dynamo with shared db and inMemory configuration
  
      ` $ make start-dynamo `
  
3. Create table in the dynamodb

      ` $ make local-table-create `
  
4. Download the dependencies

      ` $ go mod download `

5. Use following command to start the server

      ` $ grnc-yaml-bind && go build && ./peopleapi `


### Using CSV Store
Goto common.yml configuration and change the following lines
```yaml
  storeMechanism:
    type: stores.DynamodbStore
    Name: conf:Service.DynamoStore
    Region: us-west-2
    Endpoint: http://localhost:8000
```

To

```yaml
  storeMechanism:
    type: stores.FileStore
    Name: conf:Service.CSVStore
```

## Todo
- [ ] Implement CSV Update functionality
- [ ] Add validation on Update API
- [ ] Add tests
- [ ] Refactor

## Thank You
Thank you all mentors who guided me through Go lang and Granitic
- [Barun Thapa](https://github.com/barunthapa)
- [Samit Ghimire](http://github.com/samit22)
- [Rojesh Shrestha](https://github.com/rojesh)
- [Milap Neupane](https://github.com/milap-neupane)
