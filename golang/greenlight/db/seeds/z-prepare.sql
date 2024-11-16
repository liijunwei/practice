insert into user_permissions
select id,(select  id from permissions where code = 'movies:read') from users;

insert into user_permissions
values(
  (select id from users where email = 'Emiliano_Cruickshank@example.com'),
  (select id from permissions where code = 'movies:write')
);

select email,array_agg(permissions.code) as permissions
from permissions
inner join user_permissions on user_permissions.permission_id = permissions.id
inner join users on user_permissions.user_id = users.id
where users.status = 'activated'
group by email;
