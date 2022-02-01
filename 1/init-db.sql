CREATE TABLE technologies (
  name    VARCHAR(255),
  details VARCHAR(255)
);
insert into technologies values (
  'Go', 'An open source programming language that makes it easy to build simple and efficient software.'
);
insert into technologies values (
  'JavaScript', 'A lightweight, interpreted, or just-in-time compiled programming language with first-class functions.'
);
insert into technologies values (
  'MySQL', 'A powerful, open source object-relational database'
);

CREATE TABLE album (
  id    VARCHAR(255),
  title VARCHAR(255),
  artist VARCHAR(255),
  price VARCHAR(255)
);
insert into album values ('1', 'Blue Train','John Coltrane','59.5');
  
insert into album values ('2', 'Jeru', 'Gerry Mulligan', '46.6' );

insert into album values ( '3', 'MySQL', 'tutoral','free');

