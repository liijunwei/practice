-- name: GetAllForUser :many
select perm.code
from permissions perm
inner join user_permissions user_perm on user_perm.permission_id = perm.id
inner join users u on user_perm.user_id = u.id
where u.id = @user_id;

-- name: AddForUser :exec
insert into user_permissions(user_id, permission_id)
select @user_id, permissions.id from permissions where permissions.code = ANY(@codes::text[]);
