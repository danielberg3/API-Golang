insert into usuarios(nome, nick, email, senha)
values
("usuario1", "usuario_1", "usuario_1@gmail.com", "$2a$10$hVUPMGFBojzpVelI2MqiAOBWWm7Y7bMbxDK7bAwB9XWV9WKLHFGa."),
("usuario2", "usuario_2", "usuario_2@gmail.com", "$2a$10$hVUPMGFBojzpVelI2MqiAOBWWm7Y7bMbxDK7bAwB9XWV9WKLHFGa."),
("usuario3", "usuario_3", "usuario_3@gmail.com", "$2a$10$hVUPMGFBojzpVelI2MqiAOBWWm7Y7bMbxDK7bAwB9XWV9WKLHFGa.");

insert into seguidores(usuario_id, seguidor_id)
values
(1, 3),
(3, 1),
(2, 1);

insert into publicacoes(titulo, conteudo, autor_id)
values
("Publicação do usuário 1", "Essa é a publicação do usuário 1! oba!", 1),
("Publicação do usuário 2", "Essa é a publicação do usuário 2! oba!", 2),
("Publicação do usuário 3", "Essa é a publicação do usuário 3! oba!", 3);