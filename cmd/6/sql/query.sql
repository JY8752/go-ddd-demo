-- name: CrateUser :one
insert into users(id, first_name, last_name)
values(?, ?, ?)
returning *;

-- name: FindUser :one
select * from users where id = ?;

-- name: FindUserByName :one
select * from users where first_name = ? and last_name = ?;

-- name: UpdateUserName :one
update users set first_name = ?, last_name = ? where id = ? returning *;

-- name: DeleteUser :one
delete from users where id = ? returning *;