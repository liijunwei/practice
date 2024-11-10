help:
	go run cmd/api/* -h

# export PGPASSWORD in shell before `make psql`
psql:
	@PGPASSWORD=${PGPASSWORD} psql -U greenlight -d greenlight
	# OR
	# psql $GREENLIGHT_DB_DSN
	#
	# create database greenlight
	# \c greenlight;
	# create role greenlight with login password '{password}';
	# create extension if not exists citext;

show-pg-config-file:
	sudo -u postgres psql -c 'show config_file;'
	# /Library/PostgreSQL/17/data/postgresql.conf

install-tools:
	brew install postgres
	brew install golang-migrate
