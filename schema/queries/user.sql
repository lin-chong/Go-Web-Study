-- name: GetUserByUserId :one
SELECT * FROM `user` WHERE id = ?;