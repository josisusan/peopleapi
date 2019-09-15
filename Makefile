install-dynamo:
	docker pull amazon/dynamodb-local

start-dynamo:
	docker run -p 8000:8000 amazon/dynamodb-local -jar DynamoDBLocal.jar -inMemory -sharedDb

local-table-create:
	aws dynamodb create-table \
			--table-name PeopleInfo \
			--attribute-definitions \
					AttributeName=UUID,AttributeType=S \
			--key-schema \
					AttributeName=UUID,KeyType=HASH \
			--provisioned-throughput \
					ReadCapacityUnits=10,WriteCapacityUnits=5 \
			--endpoint-url http://localhost:8000

local-table-status:
	aws dynamodb describe-table --table-name PeopleInfo --endpoint-url http://localhost:8000 | grep TableStatus