build:
	go build ./cmd/AD_Post/

create_db:
	go run tools/BuildDB.go

run: build create_db
	./AD_Post

test:
	go test ./db
	go test ./cmd/AD_Post/

clean:
	rm -rf AD_Post ad.db
