all: lint vet mole 

FILES := $$(find . -name '*.go' | grep -vE 'vendor') 

golint:
	go get github.com/golang/lint/golint  

lint: golint
	@for path in $(SOURCE_PATH); do echo "golint $$path"; golint $$path; done;

clean:
	@rm -rf bin

vet:
	go tool vet $(FILES) 2>&1
	go tool vet --shadow $(FILES) 2>&1

mole:
	go build 
                            
moleDocker:                  
	CGO_ENABLED=0 go build -o ./bin/Mole_Docker main.go



