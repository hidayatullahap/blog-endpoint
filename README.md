[Toolbox]: <https://gobuffalo.io/en/docs/db/toolbox>
[config]: <https://gobuffalo.io/en/docs/db/configuration>

### Installation

blog-endpoint using Gobuffalo [Toolbox] for migration 

If you are not install Toolbox yet, run this command below to get soda CLI. Run it inside your $GOPATH
```sh
$ go get github.com/gobuffalo/pop/...
$ go install github.com/gobuffalo/pop/soda
```

Copy .env.example to .env, edit it for your database configuration
```sh
$ cp .env.example .env
```

Copy database.yml.example to database.yml 
```sh
$ cp database.yml.example database.yml
```

or you can config your database by yourself

```sh
$ soda g config -t mysql
```
Click here to see [config] documentation.

Finaly run this command to migrate your database
```sh
$ soda migrate up
```

If you want to skip installing soda CLI you can dump your table with schema.sql inside /migrations

Click here to see full [Toolbox] documentation.