Start update server

    cd ./server
    go run ./main.go &


# Full update

    cd example

Run example

    go run -ldflags "-X 'main.version=1.0'" ./main.go
    
Build two versions of the app
    
    go build -o ./ex-1.0 \
        -ldflags "-X main.version=1.0" ./main.go

    go build -o ./ex-2.0 \
        -ldflags "-X 'main.version=2.0'" ./main.go
        
    rm ../server/static/ex-2.0
        
    cp ./ex-2.0 ../server/static

Update

    ./ex-1.0
    
    ./ex-2.0
    
    ./ex-1.0 -update -version 2.0
    
    ./ex-1.0
    

# Patch update

    cd example-patch

Build two versions of the app

Create binary patch

    brew install bsdiff
    
    bsdiff ./ex-1.0 ./ex-2.0 ./patch-1.0-2.0
    
    rm ../server/static/patch-1.0-2.0
    
    cp ./patch-1.0-2.0 ../server/static
    
Update

    




