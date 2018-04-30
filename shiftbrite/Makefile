#! /usr/bin/make 

.PHONY: help clean build teardown upload monitor

help:  ## Print the help documentation
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build:  ## Build the code
	platformio run

upload:  ## Upload the code
	platformio run --target upload

clean: ## Clean the library
	platformio run --target clean

teardown: ## Remove all virtualenv files
	rm -rf .pioenvs/
	rm -rf .piolibdeps/

monitor:  ## Monitor the serial port
	platformio device monitor --baud 115200

default: help
