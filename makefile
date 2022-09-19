build:
	go build -o internet_checker_go internet_checker_go.go

run:
	go run internet_checker_go.go --url https://www.google.com --retry 4

run_icmp:
	go run internet_checker_go.go --url https://www.google.com --retry 4 --icmp