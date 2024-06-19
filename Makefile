run = nodemon -e go --exec go run main.go --signal SIGTERM

run:
	$(call run)