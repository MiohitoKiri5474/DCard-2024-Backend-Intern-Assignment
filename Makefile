build:
	go build ./cmd/AD_Post/

create_db:
	go run tools/BuildDB.go

run:
	go run cmd/AD_Post/main.go

test:
	go test ./db

clean:
	rm -rf AD_Post ad.db
