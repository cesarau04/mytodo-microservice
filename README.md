mytodo-microservice
===================

About
-----
This is an implementation of a task manager using Golang and Firebase. This RESTful API is capable of:
- Lists tasks (ToDo's and done)
- Add a new task
- Delete a task
- Mark as done a task
- Get individual info about a task


Build
-----
```
make [build | run | clean]
```
If having problems in linux try using `make build-linux`

Usage
-----
```
./mytodo-microservice -port=9000
```
By default the program launches at `localhost:9000`.

Routes
======
