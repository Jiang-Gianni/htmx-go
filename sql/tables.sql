-- name: CreateDummy :exec
create table if not exists dummy(
    id integer primary key,
    name text not null
);

-- name: GetDummies :many
select * from dummy;