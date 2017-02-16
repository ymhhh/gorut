all:
	go build
	@mv gorut /usr/local/bin/

clean:
	@rm -f /usr/local/bin/gorut
	go clean