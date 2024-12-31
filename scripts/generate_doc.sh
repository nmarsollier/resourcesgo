set -e

rm -rf ./docs
swag fmt
swag init 
