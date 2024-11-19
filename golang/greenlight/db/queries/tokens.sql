-- name: CreateScopedUserToken :exec
insert into tokens(hash, user_id, expire_at, scope)
values(@hash,@user_id,@expire_at,@scope);

-- name: DeleteScopedUserTokens :exec
delete from tokens where scope = @scope and user_id = @user_id;
