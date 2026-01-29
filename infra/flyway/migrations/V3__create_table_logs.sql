CREATE TABLE logs (
    id varchar(36) PRIMARY KEY DEFAULT (UUID()),
    log varchar(4000) not null,
    id_recurso varchar(36),
    criado_por varchar(36),
    criado_em timestamp default current_timestamp(),
    foreign key(criado_por) references usuarios(id)
) ENGINE=INNODB;