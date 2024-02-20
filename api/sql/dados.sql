INSERT into usuarios (nome, nick, email, senha, esporte, anosExperiencia, possuiPatrocinio)
 values 
 ("João", "joao", "joao@gmail.com", '123', "futebol", 5, true),
 ("Maria", "maria", "maria@hotmail.com, '123', "volei", 3, false),
 ("José", "jose", "josefa@yahoo.com", '123', "basquete", 7, true);

INSERT into seguidores (usuario_id, seguidor_id)
    values 
    (1, 2), --user follows user2
    (1, 3), --user follows user4
    (2, 1), --user2 follows user
    (3, 1); --user4 follows user


INSERT INTO publicacoes (titulo, conteudo, autor_id, midia)
    VALUES
    ("Futebol é vida", "Vamos jogar futebol", 1, NULL),
    ("Volei é vida", "Vamos jogar volei", 2, NULL),
    ("Basquete é vida", "Vamos jogar basquete", 3, NULL);