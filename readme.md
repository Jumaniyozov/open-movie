<h1>Migrations</h1>

    ### In this command:
    The -seq flag indicates that we want to use sequential numbering like 0001, 0002, ... for the migration files (instead of a Unix timestamp, which is the default).
    The -ext flag indicates that we want to give the migration files the extension .sql.
    The -dir flag indicates that we want to store the migration files in the ./migrations directory (which will be created automatically if it doesn’t already exist).
    The name create_movies_table is a descriptive label that we give the migration files to signify their contents.

<ul>
    <li>Create movies table: <code>migrate create -seq -ext=.sql -dir=./migrations create_movies_table</code></li>
    <li>Create relative check constraints: <code>migrate create -seq -ext=.sql -dir=./migrations add_movies_check_constraints</code></li>
    <li>Run migrations: <code>migrate -path=./migrations -database=$EXAMPLE_DSN up</code></li>
    <li>Check Version of migrations: <code>migrate -path=./migrations -database=$EXAMPLE_DSN version</code></li>
    <li>Migrate up or down to a specific version: <code>migrate -path=./migrations -database=$EXAMPLE_DSN goto 1</code></li>
    <li>Use the down command to roll-back by a specific number of migrations: <code>migrate -path=./migrations -database =$EXAMPLE_DSN down 1</code></li>
    <li>Apply all down migrations: <code>migrate -path=./migrations -database=$EXAMPLE_DSN down</code></li>

    ### generally prefer to use the goto command to perform roll-backs (as it’s more explicit about the target version) and reserve use of the down command for rolling-back all migrations 
</ul>    
