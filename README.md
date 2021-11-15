# todo-api

[![Go Report Card](https://goreportcard.com/badge/github.com/bburaksseyhan/todo-api)](https://goreportcard.com/report/github.com/bburaksseyhan/todo-api)

```
 docker-compose up -d  
```

```
 docker exec -it cassandra cqlsh; 
```

<img width="1388" alt="Screen Shot 2021-11-08 at 15 46 57" src="https://user-images.githubusercontent.com/60069987/140744650-0afc90ff-19f0-4746-a39c-3d56131d7b6f.png">

```
CREATE KEYSPACE todo WITH replication = {'class':'SimpleStrategy', 'replication_factor' : 1};
```

```
use todo

create table todo (id text primary key, title text, content text);

select * from todo
```
