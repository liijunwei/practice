-- name: ListTODOs :many
SELECT
  id,
  title
FROM
  todos;

-- name: CreateTODO :exec
INSERT INTO
  todos(title)
VALUES
  (@title);

-- name: DeleteTODO :exec
DELETE FROM
  todos
WHERE
  id = @id;
