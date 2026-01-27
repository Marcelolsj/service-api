CREATE TABLE usuarios (
    id int auto_increment primary key,
    nome varchar(250) not null,
    email varchar(100) not null unique,
    senha varchar(50) not null,
    criado_em timestamp default current_timestamp()
) ENGINE=INNODB;