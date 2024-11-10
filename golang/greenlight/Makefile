
# export PGPASSWORD in shell before `make psql`
psql:
	@PGPASSWORD=${PGPASSWORD} psql -U greenlight -d greenlight
	# create database greenlight
	# \c greenlight;
	# create role greenlight with login password '{password}';
	# create extension if not exists citext;

show-pg-config-file:
	sudo -u postgres psql -c 'show config_file;'
	# /Library/PostgreSQL/17/data/postgresql.conf
