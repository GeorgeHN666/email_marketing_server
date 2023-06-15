## Build Binary file
build:
	@ echo Building Binary
	@ go build -o ./bin/app.exe ./app/main.go 
	@ echo Binary built 

go: build
	@ echo Starting binary 
	@ ./bin/app.exe 
	@ echo App successfuly started
