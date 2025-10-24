В директории проекта запустить:


```
docker build -t hakaton_mongo . 
```

```
docker run -d --name hakaton_mongo -p 27017:27017 -v mongo_data:/data/db -v mongo_logs:/var/log/mongodb -v C:/mongo-backups:/backups hakaton_mongo
```

