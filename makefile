name=nyantan

build:
	go build -o ${name}

run:
	echo "\e[32mStarting Server!\e[0m"
	./${name}

reflex:
	reflex -R '\.git' -r '\.go' -s -- make build run
