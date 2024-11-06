-- name: Add :exec
INSERT INTO pinuser (
    user_id
) VALUES (
    $1
);

-- name: GetAll :many
SELECT * FROM pinuser;