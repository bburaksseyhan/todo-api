[![Go Doc](https://godoc.org/github.com/gogf/gf?status.svg)](https://godoc.org/github.com/gogf/gf)
[![Go Report Card](https://goreportcard.com/badge/github.com/bburaksseyhan/todo-api)](https://goreportcard.com/report/github.com/bburaksseyhan/todo-api)
[![GitHub stars](https://img.shields.io/github/stars/bburaksseyhan/todo-api)](https://github.com/bburaksseyhan/todo-api/stargazers)

# todo-api

run on docker :ship:
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

used packages :package:

- [ ] Web framework: go get -u github.com/gin-gonic/gin
- [ ] Read Configuration file : go get -u github.com/spf13/viper
- [ ] Logging : go get -u github.com/sirupsen/logrus
- [ ] Cassandra : go get -u github.com/gocql/gocql

documentation: [Related Post](https://dev.to/bseyhan/cassandra-golang-551a). :pencil2: :book:
