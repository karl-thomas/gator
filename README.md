# Gator

you'll need to have a postgres server up and running, as well as go installed on your machine. 

```sh
go install github.com/karl-thomas/gator
```

create ~/.gatorconfig.json with the following structure:  
```json
{
  "db_url": "connection string to your gator db",
  "username": ""
}
```

should be able to run gator from the command line now. 


first you need to add your name and the feeds you want to follow.
```sh
gator register <your name>

gator addfeed <feed name> <feed url>
```

then you can run the following to fetch the feeds on a loop

```sh
gator agg <timeout like 5s or 1m>
```

in a another terminal you can see your most recent posts

```sh 
gator browse <limit>
```


