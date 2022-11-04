#Ejericio 1
/* 1. Con la base de datos “movies”, se propone crear una tabla temporal llamada “TWD” y guardar en la misma las temporadas de “The Walking Dead”. (La cambié un poco xd)*/
#CREATE TEMPORARY TABLE TWD2 (Season varchar(40), Serie varchar(40));

INSERT INTO TWD2 SELECT season.title, series.title
FROM movies_db.seasons season
INNER JOIN movies_db.series series ON season.serie_id = series.id
WHERE series.title LIKE "The Wal%";

SELECT * FROM TWD2;

#Ejericio 2
/* 1. En la base de datos “movies”, seleccionar una tabla donde crear un índice y luego chequear la creación del mismo. */
EXPLAIN SELECT * FROM movies_db.movies WHERE length=120 /*Sin indice busca en las 21 filas, y con indice en las 6 que trae*/