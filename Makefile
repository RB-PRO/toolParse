all: run

run:
	go run main.go

push:
	git push git@github.com:RB-PRO/toolParse.git

pull:
	git pull git@github.com:RB-PRO/toolParse.git

pushW:
	git push https://github.com/RB-PRO/toolParse.git

pullW:
	git pull https://github.com/RB-PRO/toolParse.git