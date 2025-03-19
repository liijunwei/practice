-- name: ListShorturls :many
SELECT
  *
FROM
  shorturls;

-- name: CreateShorturl :one
INSERT INTO
  shorturls(original, shorturl)
VALUES
  (@original, @shorturl) RETURNING *;

-- name: GetOriginalByShort :one
SELECT
  *
FROM
  shorturls
WHERE
  shorturl = @shorturl;

-- TODO try this first, then optimize with bloom filter to see it's performance gain
-- name: CheckOriginalExists :one
SELECT
  *
FROM
  shorturls
WHERE
  original = @original;
