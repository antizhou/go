package es

import (
	"strings"
	"sync"
	"time"

	"antizhou.com/go/config"
	"github.com/cihub/seelog"
	"github.com/olivere/elastic"
)

var (
	clientMu sync.Mutex
	client   *elastic.Client
)

type Log struct {
	Service   string    `json:"service"`
	Metric    string    `json:"metric"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
	Instance  string    `json:"instance"`
	CreatedAt time.Time `json:"created_at"`
	DC        string    `json:"dc"`
}

func GetClient() (*elastic.Client, error) {
	clientMu.Lock()
	defer clientMu.Unlock()

	if client != nil {
		return client, nil
	}

	url := config.Get("es_url")

	seelog.Info("es url is: ", url)

	urls := strings.Split(url, ",")
	options := []elastic.ClientOptionFunc{
		elastic.SetURL(urls...),
		elastic.SetBasicAuth(config.Get("username"), config.Get("password")),
		elastic.SetSniff(false),
	}
	c, err := elastic.NewClient(options...)

	seelog.Info("creating es client for ", urls)

	if err != nil {
		seelog.Infof("creating es client error: ", err)
		return nil, err
	}

	client = c

	return client, nil
}
