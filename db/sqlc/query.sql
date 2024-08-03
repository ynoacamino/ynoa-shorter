-- name: GetPublicShorters :many
SELECT
  *
FROM
  url
WHERE
  public = true
ORDER BY
  updated_at DESC;

-- name: GetPrivateShorters :many
SELECT
  *
FROM
  url
WHERE
  user_id = $1
ORDER BY
  updated_at DESC;

-- name: CreateShorter :one
INSERT INTO url (
  short_url,
  original_url,
  user_id,
  public
) VALUES (
  $1,
  $2,
  $3,
  $4
)
RETURNING *;

-- name: UpdateShorter :one
UPDATE url
SET
  short_url = $2,
  public = $3,
  updated_at = CURRENT_TIMESTAMP
WHERE
  url_id = $1
RETURNING *;

-- name: DeleteShorter :one
DELETE FROM url
WHERE
  url_id = $1
RETURNING *;