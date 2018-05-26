package main

import
(
	"encoding/json"
	"errors"
	"fmt"
	"new/http"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
)

var (
	API_KEY = 
)

type Request struct{
	ID int 'json:"id"'
}

type MovieDBResponse struct{
	Movies []Movie 'json:"results"'
}

type Movie struct{
	Title string 'json:"title"'
	Description string 'json:"overview"'
	Cover string 'json:"poster_path"'
	ReleaseData string 'json:"release_data"'
}

func Handler(request Request) ([]Movie, error){
}