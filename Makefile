PORT = 5000

server:
	go run server.go $(PORT)

client:
	go run client.go