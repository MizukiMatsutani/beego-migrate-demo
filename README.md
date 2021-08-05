# This is a demo application for Beego's Issue.

The target issue is [here](https://github.com/beego/beego/issues/4712).

## This demo is created by the following procedure.

1. Install Bee. `go get -u github.com/beego/bee/v2`
1. Create a new project with the bee command. `bee new beego-migrate-demo`
1. Create migration file. `bee generate migration create_tests_table`
1. Added create table statement to `20210805_223243_create_tests_table.go` file

## To check, follow the steps below.

1. `docker compose up -d`
2. `bee migrate --conn="demo-user:secret@tcp(127.0.0.1:3306)/demo-db?charset=utf8"`

```bash
$ bee migrate --conn="demo-user:secret@tcp(127.0.0.1:3306)/demo-db?charset=utf8"
______
| ___ \
| |_/ /  ___   ___
| ___ \ / _ \ / _ \
| |_/ /|  __/|  __/
\____/  \___| \___| v2.0.2
2021/08/05 23:11:41 INFO     ▶ 0001 Using 'mysql' as 'driver'
2021/08/05 23:11:41 INFO     ▶ 0002 Using '/xxxxx/beego-migrate-demo/database/migrations' as 'dir'
2021/08/05 23:11:41 INFO     ▶ 0003 Running all outstanding migrations
2021/08/05 23:11:41 INFO     ▶ 0004 Creating 'migrations' table...
2021/08/05 23:11:44 INFO     ▶ 0005 |> 2021/08/05 23:11:42.731 [I]  start upgrade CreateTestsTable_20210805_223243
2021/08/05 23:11:44 INFO     ▶ 0006 |> 2021/08/05 23:11:42.731 [I]  exec sql:
2021/08/05 23:11:44 INFO     ▶ 0007 |>  create table tests (
2021/08/05 23:11:44 INFO     ▶ 0008 |>          id bigint auto_increment primary key
2021/08/05 23:11:44 INFO     ▶ 0009 |>          , name varchar(20) not null
2021/08/05 23:11:44 INFO     ▶ 0010 |>  )
2021/08/05 23:11:44 INFO     ▶ 0011 |>
2021/08/05 23:11:44 INFO     ▶ 0012 |> 2021/08/05 23:11:42.774 [I]  end upgrade: CreateTestsTable_20210805_223243
2021/08/05 23:11:44 INFO     ▶ 0013 |> 2021/08/05 23:11:42.774 [I]  total success upgrade: 1  migration
2021/08/05 23:11:44 SUCCESS  ▶ 0014 Migration successful!
```

3. `mysql -u demo-user demo-db -h 127.0.0.1 -P 3306 -psecret`
4. Check the status of the table.

```bash
mysql> show tables;
+-------------------+
| Tables_in_demo-db |
+-------------------+
| migrations        |
| tests             |
+-------------------+
2 rows in set (0.01 sec)

mysql> select * from migrations\G
*************************** 1. row ***************************
       id_migration: 1
               name: CreateTestsTable_20210805_223243
         created_at: 2021-08-05 23:11:42
         statements:
        create table tests (
                id bigint auto_increment primary key
                , name varchar(20) not null
        )

rollback_statements: NULL
             status: update
1 row in set (0.00 sec)

mysql> desc tests;
+-------+-------------+------+-----+---------+----------------+
| Field | Type        | Null | Key | Default | Extra          |
+-------+-------------+------+-----+---------+----------------+
| id    | bigint      | NO   | PRI | NULL    | auto_increment |
| name  | varchar(20) | NO   |     | NULL    |                |
+-------+-------------+------+-----+---------+----------------+
2 rows in set (0.01 sec)
```

5. Modify `20210805_223243_create_tests_table.go` and add fields.

```diff
func (m *CreateTestsTable_20210805_223243) Up() {
	m.SQL(`
	create table tests (
		id bigint auto_increment primary key
		, name varchar(20) not null
+       , added_fileld varchar(20)
	)
	`)
}
```

6. `bee migrate refresh --conn="demo-user:secret@tcp(127.0.0.1:3306)/demo-db?charset=utf8"`

```bash
$ bee migrate refresh --conn="demo-user:secret@tcp(127.0.0.1:3306)/demo-db?charset=utf8"
______
| ___ \
| |_/ /  ___   ___
| ___ \ / _ \ / _ \
| |_/ /|  __/|  __/
\____/  \___| \___| v2.0.2
2021/08/05 23:17:56 INFO     ▶ 0001 Using 'mysql' as 'driver'
2021/08/05 23:17:56 INFO     ▶ 0002 Using '/xxx/beego-migrate-demo/database/migrations' as 'dir'
2021/08/05 23:17:56 INFO     ▶ 0003 Refreshing all migrations
2021/08/05 23:18:01 INFO     ▶ 0004 |> 2021/08/05 23:17:57.675 [I]  start reset: CreateTestsTable_20210805_223243
2021/08/05 23:18:01 INFO     ▶ 0005 |> 2021/08/05 23:17:57.675 [I]  exec sql: DROP TABLE if exists tests
2021/08/05 23:18:01 INFO     ▶ 0006 |> 2021/08/05 23:17:57.711 [I]  end reset: CreateTestsTable_20210805_223243
2021/08/05 23:18:01 INFO     ▶ 0007 |> 2021/08/05 23:17:57.711 [I]  total success reset: 1  migration
2021/08/05 23:18:01 INFO     ▶ 0008 |> 2021/08/05 23:17:59.718 [I]  total success upgrade: 0  migration
2021/08/05 23:18:01 SUCCESS  ▶ 0009 Migration successful!
```

7. `mysql -u demo-user demo-db -h 127.0.0.1 -P 3306 -psecret`
8. Check the status of the table.

```bash
mysql> show tables;
+-------------------+
| Tables_in_demo-db |
+-------------------+
| migrations        |
+-------------------+
1 row in set (0.00 sec)

mysql> select * from migrations\G
*************************** 1. row ***************************
       id_migration: 1
               name: CreateTestsTable_20210805_223243
         created_at: 2021-08-05 23:17:57
         statements:
        create table tests (
                id bigint auto_increment primary key
                , name varchar(20) not null
        )

rollback_statements: DROP TABLE if exists tests
             status: rollback
1 row in set (0.01 sec)
```

## Expected state.

- The state field of the migrations table is `update`.
- The tests table has been recreated and added_fileld has been added.
