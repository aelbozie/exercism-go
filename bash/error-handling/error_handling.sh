#!/usr/bin/env bash

main () {
	if (( "$#" != 1 ))
	then
		echo "Usage: error_handling.sh <person>"
		exit 1
	else
		echo "Hello, $1"
		exit 0
	fi
}

main "$@"
