
# Tunnel to the raspi to access the mysql db.
tunnel:
	 ssh -N -L 3336:127.0.0.1:3306 pi@10.0.0.3

build:
	 go build -o ingest ./cmd/ingest
	 go build -o aggregate ./cmd/aggregate

start:
	make aggregate

aggregate:
	 go run ./cmd/aggregate

ingest:
	 go run ./cmd/ingest

persist:
	cp tlds.json ../v01-hugo/static/js/hackernewsstats.json

magic:
	make aggregate && make persist
