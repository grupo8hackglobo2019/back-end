package model

type Config struct {
    ElasticSearchUrl string `env:"ELASTICSEARCH_URL" envDefault:"Slomek"`
}

type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}