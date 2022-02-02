# install postgresql

```bash
brew search postgresql
brew install postgresql@12

# add these to ~/.zshrc
export PATH="/usr/local/opt/postgresql@12/bin:$PATH"
export LDFLAGS="-L/usr/local/opt/postgresql@12/lib"
export CPPFLAGS="-I/usr/local/opt/postgresql@12/include"
export PKG_CONFIG_PATH="/usr/local/opt/postgresql@12/lib/pkgconfig"

brew services start postgresql@12
```

## if using `gem pg`

```bash
gem install pg -- --with-pg-config=/usr/local/opt/postgresql@12/bin/pg_config
```

# install pgadmin

https://www.postgresql.org/ftp/pgadmin/pgadmin4/v6.4/macos/

# init user/role/db

```bash
psql postgres

CREATE ROLE postgres WITH LOGIN PASSWORD '123456';
ALTER ROLE postgres CREATEDB;
```

TODO

