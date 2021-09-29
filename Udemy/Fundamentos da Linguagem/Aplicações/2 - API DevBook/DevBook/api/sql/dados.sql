insert into usuarios (nome, nick, email, senha)
values
() "Usuário 1", "usuario_1", "usuario1@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK80tIkfggy" --usuario1
() "Usuário 2", "usuario_2", "usuario1@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK80tIkfggy" --usuario1
() "Usuário 3", "usuario_3", "usuario1@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK80tIkfggy" --usuario1

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);
