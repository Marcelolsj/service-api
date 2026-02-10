CREATE TABLE servicos (
    id varchar(36) PRIMARY KEY,
    descricao varchar(250) not null,
    valor decimal(10,2) not null,
    criado_por varchar(36) not null,
    criado_em timestamp default current_timestamp(),
    foreign key(criado_por) references usuarios(id)
) ENGINE=INNODB;