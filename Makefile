.PHONY: dev dev-server dev-sveltekit mock-cameras build build-rpi build-sveltekit sync sync-server sync-setup

#-----------------------------------------------------
# DEV
#-----------------------------------------------------

dev:
	@gochange -k -i '**/*.go' -- go run ./

.DEFAULT_GOAL := dev