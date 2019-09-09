package stores

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DynamodbStore struct {
	Name string
}

type Item struct {
	UUID   string
	PName  string
	Age    string
	Gender string
}

func (fs DynamodbStore) Write(content string) (string, error) {
	ses, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-west-2"),
		Endpoint: aws.String("http://localhost:8000"),
	})
	if err != nil {
		return "Dynamo Connection Issue", err
	}
	db := dynamodb.New(ses)

	data := strings.Split(content, ",")
	uid := strconv.FormatInt(generateUUID(), 10)
	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"UUID": {
				S: aws.String(uid),
			},
			"PName": {
				S: aws.String(data[0]),
			},
			"Age": {
				S: aws.String(data[1]),
			},
			"Gender": {
				S: aws.String(data[2]),
			},
		},
		TableName: aws.String(fs.Name),
	}
	_, err = db.PutItem(input)

	if err != nil {
		return "Failed to insert data in dynamo", err
	}

	return (uid + ", " + content), nil
}

func (fs DynamodbStore) Read() ([][]string, error) {
	ses, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-west-2"),
		Endpoint: aws.String("http://localhost:8000"),
	})
	if err != nil {
		return [][]string{}, err
	}
	db := dynamodb.New(ses)
	input := &dynamodb.ScanInput{
		TableName: aws.String(fs.Name),
	}
	result, err := db.Scan(input)
	if err != nil {
		return [][]string{}, err
	}
	var r [][]string
	for _, i := range result.Items {
		item := Item{}
		err = dynamodbattribute.UnmarshalMap(i, &item)
		if err != nil {
			fmt.Println("Got error on unmarshalling:")
			return [][]string{}, err
		}
		r = append(r, []string{item.UUID, item.PName, item.Age, item.Gender})
	}
	return r, nil
}

// Find finds the record in dynamo
func (fs DynamodbStore) Find(uid string) ([]string, error) {
	ses, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-west-2"),
		Endpoint: aws.String("http://localhost:8000"),
	})
	if err != nil {
		return []string{}, err
	}
	db := dynamodb.New(ses)
	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(fs.Name),
		Key: map[string]*dynamodb.AttributeValue{
			"UUID": {
				S: aws.String(uid),
			},
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return []string{}, err
	}

	item := Item{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	return []string{item.UUID, item.PName, item.Age, item.Gender}, nil
}

// Update updates db
func (fs DynamodbStore) Update(uid string, u map[string]string) error {
	ses, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-west-2"),
		Endpoint: aws.String("http://localhost:8000"),
	})
	if err != nil {
		return err
	}
	db := dynamodb.New(ses)
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":n": {
				S: aws.String(u["Name"]),
			},
			":a": {
				S: aws.String(u["Age"]),
			},
			":g": {
				S: aws.String(u["Gender"]),
			},
		},
		TableName: aws.String(fs.Name),
		Key: map[string]*dynamodb.AttributeValue{
			"UUID": {
				S: aws.String(uid),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set PName = :n, Age = :a, Gender = :g"),
	}
	_, err = db.UpdateItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
