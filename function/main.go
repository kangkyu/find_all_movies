package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go-v2/config"
	// "github.com/aws/aws-sdk-go-v2/feature/rds/auth"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Movie entity
type Movie struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
}

// {"username":"dbuser","password":"password","engine":"postgres","host":"demo-here.c7ryv7ido6jm.us-west-2.rds.amazonaws.com","port":5432,"dbInstanceIdentifier":"demo-here"}
type SecretsList struct {
	UserName             string `json:"username"`
	Password             string `json:"password"`
	Engine               string `json:"engine"`
	Host                 string `json:"host"`
	Port                 int    `json:"port"`
	DBInstanceIdentifier string `json:"dbInstanceIdentifier"`
	ProxyEndpoint        string `json:"proxyendpoint`
}

var (
	secretName = "RDS-Proxy-Demo-Secret"
	// "arn:aws:secretsmanager:us-west-2:328321743722:secret:RDS-Proxy-Demo-Secret-7r5ixc"
	awsRegion = "us-west-2"
)

func findAll() (events.APIGatewayProxyResponse, error) {
	ctx := context.TODO()
	fmt.Println("Starting...")
	cfg, err := config.LoadDefaultConfig(ctx)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while retrieving AWS credentials",
		}, nil
	}

	fmt.Println("Fetching secrets...")
	svc := secretsmanager.NewFromConfig(cfg)

	input := &secretsmanager.GetSecretValueInput{
		SecretId: &secretName,
	}
	result, err := svc.GetSecretValue(ctx, input)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while fetching from secrets manager",
		}, nil
	}
	var secrets SecretsList
	json.Unmarshal([]byte(*result.SecretString), &secrets)

	// authenticationToken, err := auth.BuildAuthToken(
	// 	context.TODO(),
	// 	fmt.Sprintf("%s:%d", secrets.Host, secrets.Port),
	// 	awsRegion,
	// 	secrets.UserName, // Database Account
	// 	cfg.Credentials,
	// )
	// if err != nil {
	// 	panic("failed to create authentication token: " + err.Error())
	// }

	// postgres DNS string
	conString := fmt.Sprintf("%s://%s:%s@%s:%d/%s",
		secrets.Engine, secrets.UserName, secrets.Password,
		secrets.ProxyEndpoint, secrets.Port, secrets.DBInstanceIdentifier)

	// Use db to perform SQL operations on database
	fmt.Println("Setting up database...")
	db := setupDB(conString)

	// fetch all movies from the db
	fmt.Println("Querying...")
	rows, err := db.Queryx("SELECT id, name, cover, description FROM movies")

	fmt.Println("Reading rows...")
	var items []Movie
	// iterate over each row
	for rows.Next() {
		var movie Movie
		err = rows.StructScan(&movie)
		items = append(items, movie)
	}
	// check the error from rows
	err = rows.Err()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while running SELECT query",
		}, nil
	}

	fmt.Println("Responding...")
	response, err := json.Marshal(items)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		Body: string(response),
	}, nil
}

func main() {
	lambda.Start(findAll)
}

func setupDB(conString string) *sqlx.DB {
	fmt.Println("Connecting...")
	db, err := sqlx.Open("postgres", conString)
	if err != nil {
		panic("failed to open postgres db: " + err.Error())
	}

	fmt.Println("Ping! Ping!")
	err = db.Ping()

	return db
}
