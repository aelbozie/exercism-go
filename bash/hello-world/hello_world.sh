#!/usr/bin/env bash

# This option will make the script exit when there is an error
set -o errexit
# This option will make the script exit when it tries to use an unset variable
set -o nounset

   main() {

   input=${@:-"World"}

   echo "Hello, $input!"

   exit 0
}
   main "$@"
