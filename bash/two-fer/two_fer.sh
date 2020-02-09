#!/usr/bin/env bash
main() {
     
	input=${1:-"you"}
     	echo "One for $input, one for me."
	exit 0
}
  
main "$@"
