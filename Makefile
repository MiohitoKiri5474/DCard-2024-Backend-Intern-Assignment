build:
	go build cmd/AD_Post/main.go

create_db:
	go run tools/BuildDB.go

run:
	go run cmd/AD_Post/main.go

test:
	go test ./db

clean:
	rm -rf build ad.db
