# Phone Number Normalizer

[original task](https://github.com/gophercises/phone)

The program will iterate through the database and normalize all phone numbers in the database.

```
1234567890       ->   not_updated
123 456 7891     ->   1234567891
(123) 456 7892   ->   1234567892
(123) 456-7893   ->   1234567893
123-456-7894     ->   1234567894
123-456-7890     ->   1234567890
1234567892       ->   not_updated
(123)456-7892    ->   1234567892
```

## Configuration

The configuration file is stored in the ./config/config.yml file. You can change its path by checking the `cfg` flag.

Before changing the location of the configuration file, you should update the information in ./internal/app/migrate.go.

To normalize a phone number, use a regular expression that clears all non-numeric characters (\D).

## DB schema

All DB migrations are located in the /migrations directory.

Phones table:

```sql
TABLE public.phones
(
    id character varying(36) NOT NULL,
    "number" character varying(15) NOT NULL,
    CONSTRAINT phones_pkey PRIMARY KEY (id)
)
```

## DB migration

To init db migration you should build app with tags `migrate`:

```bash
cd cmd/app
go build -tags migrate
```

## Run options

### Linux
#### Build command:
```bash
cd cmd/app
go build -o phonenorm -tags migrate
```

#### Lunch command:
```bash
./phonenorm
```

### Windows
#### Build command:
```bash
cd cmd/app
go build -o phonenorm.exe -tags migrate
```

#### Lunch command:
```bash
.\phonenorm.exe
```
