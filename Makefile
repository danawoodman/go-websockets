.PHONY: dev 

dev:
	@gochange -k -i '**/*.go' -- go run ./

.DEFAULT_GOAL := dev