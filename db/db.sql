create
    unlogged table if not exists messages
(
    first_user  int not null,
    second_user int not null,
    weight      int not null,
    primary key (first_user, second_user)
);
