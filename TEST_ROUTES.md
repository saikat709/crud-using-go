## TESTING routes
Run any version then try the following on terminal. \
Make sure you have CURL installed.

### 1. Get all todos
```bash
curl http://localhost:3000/todos/
```
or 
```bash
curl GET http://localhost:3000/todo/
```

### 2. Geta route of ID, say 3
```bash
curl http://localhost:3000/todo/3
```
or 
```bash
curl GET http://localhost:3000/todo/3
```

### 3. Create a TODO 
```bash
curl -X POST http://localhost:3000/todo \
      -H "Content-Type: application/json" \
      -d '{"id":1, "completed":false," body":"New Body"}'
```

### Update a TODO
```bash
curl -X PUT http://localhost:3000/todo/3 \
      -H "Content-Type: application/json" \
      -d '{"completed":true," body":"New Body Updated"}'
```

### Delete a TODO
```bash
curl -X DELETE http://localhost:3000/todo/3
```

