# back-end

O projeto visa utilizar de melhor maneira os dados gerados pelo CartolaFC

O modelo segue o seguinte exemplo:

```
type Message struct {
	ID			string			`json:"_id,omitempty"`
	Text        string          `json:"text"`
	CreatedAt   time.Time       `json:"createdAt"`
	User        User            `json:"user"`
}
```

```

type User struct {
	ID			int				`json:"_id,omitempty"`
	Name        string          `json:"name"`
	Avatar      string          `json:"avatar"`
}

```

## a API

O projeto consiste em dois endpoints:

1. O websocket
2. Envio de mensagem por método POST

### 1 - O websocket

Exemplo de chamada:

```
echo "{\"_id\": \"9\",\"text\": \"hey you\",\"createdAt\": \"2012-04-23T18:25:43.511Z\",\"user\": {\"_id\": 5,\"name\": \"Brahma\",\"avatar\": \"https://i.imgur.com/Y5ie5sg.png\"}}" | websocat ws://grupo-oito.herokuapp.com/ws

```
### 1 - Envio de mensagem por método POST

Exemplo de chamada:

```
POST http://grupo-oito.herokuapp.com/sendMessage

{
	"_id": "9",
	"text": "hey you",
	"createdAt": "2012-04-23T18:25:43.511Z",
	"user": {
		"_id": 5,
		"name": "Brahma",
		"avatar": "https://i.imgur.com/Y5ie5sg.png"
	}
}

```
