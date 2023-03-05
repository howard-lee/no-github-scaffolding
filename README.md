# no-github-scaffolding
golang code scaffolding for lan-based version control developers, or 
developers do not push code to git/github, but can still develop large complex code base with different package.

To develop golang application with packages in different sub-folders, use 
  replace (package name, version) => ./(subfolder-name) 

for example, in this code
  replace api v1.0.0 => ./api/v1
  replace server v1.0.0 => ./internal/grpc/server

(api and server package should be added as part of the 'require' in go.mod file )

This folder-scaffolding with code are extracted from other project with many portion trimmed down, show case the references among packages and how they tied together through go.mod.  (Please be awared earlier version of golang - prior to 1.15 might not work)

To run the applicaiton, 

(1) cd to /cmd/grpcclient, then type

    go test . -v

and it will show the followings:

=== RUN   TestGrpcServerGetWalletsWithValidParameters

 ... TestGrpcServerGetWalletsWithValidParameters invoked
yyyy/mm/nn hh:mm:ss Listening on localhost:13436
    grpc_server_test.go:93: get wallet operation ok,  got 0, want 0
    grpc_server_test.go:93: get wallet operation ok,  got 404, want 404
    grpc_server_test.go:93: get wallet operation ok,  got 0, want 0
--- PASS: TestGrpcServerGetWalletsWithValidParameters (0.01s)
=== RUN   TestGetWalletMethodCallsDataStoreWithDifferentQueries

 ... TestGetWalletMethodCallsDataStoreWithDifferentQueries invoked
    grpc_server_test.go:120: get wallet operation ok,  got 0, want 0
    grpc_server_test.go:120: get wallet operation ok,  got 404, want 404
    grpc_server_test.go:120: get wallet operation ok,  got 0, want 0
--- PASS: TestGetWalletMethodCallsDataStoreWithDifferentQueries (0.00s)
PASS
ok      blocktype/cmd/grpcClient   


(2.1) start a grpc server: open another terminal window, 
  cd to /cmd/grpcserver

  and type 'go run main.go' to start the grpc server (without data)

  it will show the followings:

  grpcBlkServer invoked  yyyy-mm-dd hh:min:ss.nnnnnn
2023/03/04 16:44:19 Listening on :8080

  (The grpc server is running and listen to port 8080)


(2.2) to launch a client to the server, back to the terminal window opened in (1)
 ( it should still be in folder /cmd/grpcclient)

 type 'go run main.go' to launch grpc client that will connect to the server

 it will show the followings:

 grpcBlkClient main invoked  yyyy-mm-dd hh:mm:ss.nnnnnn 
Get Balance completed successfully, account = , balance = 0 , response code = 404, message=id should be non-zero value 

  (This call does not get valid result because there is no valid backend database/datastore provided.  In 1. testcase, the cachestore was injected and work fine.)

Then look at the packages such as api/server, they are all 'included locally' without need to pull from external internet code repository - such as github.

When you work locally - having concern of security, or working on a LAN based version-control system, you can take these to help develop complex code base/packages in an local/LAN environments.


