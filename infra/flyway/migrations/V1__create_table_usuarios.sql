CREATE TABLE usuarios (
    id varchar(36) PRIMARY KEY,
    nome varchar(250) not null,
    login varchar(250) not null,
    email varchar(250) not null unique,
    senha varchar(250) not null,
    tipo_usuario varchar(50) not null,
    criado_em timestamp default current_timestamp()
) ENGINE=INNODB;