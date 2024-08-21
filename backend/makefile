name=nihility

files=$(wildcard src/*.go)

build:
	go build -o ${name} ${files}

run:
	echo -e "\e[32mStarting Server!\e[0m"
	./${name}

reflex:
	reflex -R '\.git' -r '\.go' -s -- make build run
