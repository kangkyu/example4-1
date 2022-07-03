package main

import (
	"encoding/json"
	"fmt"
	"os"
	"log"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Movie struct {
	ID string 
	Name string
}

func main() {


	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		// config.WithCredentialsProvider(aws.AnonymousCredentials{}),
	)
	if err != nil {
		log.Fatal(err)
	}

	movies, err := readMovies("movies.json")
	if err != nil {
		log.Fatal(err)
	}

	for _, movie := range movies {
		fmt.Println("Inserting:", movie.Name)
		err = insertMovie(cfg, movie)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func readMovies(fileName string) ([]Movie, error) {
	movies := make([]Movie, 0)

	data, err := os.ReadFile(fileName)
	if err != nil {
		return movies, err
	}

	err = json.Unmarshal(data, &movies)
	if err != nil {
		return movies, err
	}

	return movies, nil
}

func insertMovie(cfg aws.Config, movie Movie) error {
	itemValue := make(map[string]types.AttributeValue, 1)
	itemValue["ID"] = &types.AttributeValueMemberS{ Value: movie.ID }
	itemValue["Name"] = &types.AttributeValueMemberS{ Value: movie.Name }

	svc := dynamodb.NewFromConfig(cfg)
	_, err := svc.PutItem(
		context.TODO(),
		&dynamodb.PutItemInput{
			TableName: aws.String("movies"),
			Item: itemValue,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
