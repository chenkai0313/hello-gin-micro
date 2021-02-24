package app

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
)

func InitEs() *elasticsearch.Client {
	es, err := elasticsearch.NewClient(
		elasticsearch.Config{
			Addresses: []string{Config.Es.Dns},
		},
	)
	if err != nil {
		fmt.Println("connect es fail ", "ERROR: Unable to create client:"+err.Error())
	}
	return es
}
