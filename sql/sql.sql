CREATE DATABASE IF NOT EXISTS devbook;
USE  devbook;

DROP TABLE IF EXISTS   publicacoes;
DROP TABLE IF EXISTS  seguidores;
DROP TABLE IF EXISTS   usuarios;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null,
    criadoEm timestamp default current_timestamp()
)ENGINE=INNODB;

CREATE TABLE seguidores(
    usuario_id int not null,
    foreign key(usuario_id)
    references usuarios(id)
    on delete cascade,

    seguidor_id int not null,
    foreign key(seguidor_id)
    references usuarios(id)
    on delete cascade,

    primary key(usuario_id, seguidor_id)

)ENGINE=INNODB;

CREATE TABLE publicacoes(
    id int auto_increment primary key,
    titulo varchar(50) not null,
    conteudo varchar(300) not null,

    autor_id int not null,
    foreign key (autor_id)
    references usuarios(id)
    on delete cascade,

    curtidas int default 0,
    cridaEm timestamp default current_timestamp
)ENGINE=INNODB;

