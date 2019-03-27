mytodo-microservice
===================

About
-----
This is an implementation of a task manager using Golang and Firebase. This RESTful API is capable of:
- List tasks (ToDo's and done)
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
By default the program launches at `localhost:9000`, Ex. if your run `make run`
```
./mytodo-microservice -port=9000
```

Routes
======

GET - List tasks
----------------
The route used is `/tasks`
```
curl http://localhost:9000/tasks
```

POST - Add task
---------------
The route used is `/tasks`
```
curl -X POST -H "Content-Type: application/json" --data '{"title": "a title", "desc": "a desc", "date": "01/12/19"}' http://localhost:9000/tasks

```

DELETE - Delete task
--------------------
The route used is `/tasks`
```
curl -X DELETE -H "Content-Type: text/plain" --data '"TASKID"' http://localhost:9000/tasks

```

PUT - Marked as done
--------------------
The route used is `/tasks`
```
curl -X PUT -H "Content-Type: text/plain" --data '"TASKID"' http://localhost:9000/tasks
```

GET - Get a task info
---------------------
The route used is `/tasks/{id}`
```
curl http://localhost:9000/tasks/TASKID
```

