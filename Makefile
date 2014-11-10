setup:
	go get github.com/go-sql-driver/mysql
	go get github.com/jessevdk/go-flags
	go build
	go install
clean:
	rm -rf ./q_show_grants
	rm -rf $(GOPATH)/bin/q_show_grants
