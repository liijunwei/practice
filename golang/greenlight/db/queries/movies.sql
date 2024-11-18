-- name: CreateMovie :one
insert into
  movies(title, year, runtime, genres)
values
  (@title, @year, @runtime, @genres) returning id,
  created_at,
  version;

-- name: GetMovieByID :one
select id,created_at,title,year,runtime,genres,version from movies where id = @id;

-- TODO add pagination and filtering back
-- name: GetAllMovies :many
select
  count(*) over(),
  id,
  created_at,
  title,
  year,
  runtime,
  genres,
  version
from
  movies;

-- name: UpdateMovie :exec
update movies
set title = @title, year = @year, runtime = @runtime, genres = @genres, version = version+1, updated_at = now()
where id = @id and version = @version
returning version;

-- name: DeleteMovie :exec
DELETE from movies where id = @id;
