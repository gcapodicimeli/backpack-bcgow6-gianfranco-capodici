/* 1. Mostrar el título y el nombre del género de todas las series.*/
SELECT series.title, genres.name
FROM movies_db.series series INNER JOIN movies_db.genres genres
ON series.genre_id = genres.id;

/* 2. Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.*/


/* 3. Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.*/
SELECT series.title, MAX(seasons.number) AS total
FROM movies_db.series series INNER JOIN movies_db.seasons seasons
ON series.id = seasons.serie_id
GROUP BY series.id;

/* 4. Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.*/
SELECT gen.name, COUNT(gen.name) AS "Cantidad de movies"
FROM movies_db.genres gen INNER JOIN movies_db.movies mov
ON gen.id = mov.genre_id
GROUP BY gen.name
HAVING COUNT(gen.name) >= 3;

/* 5. Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.*/
SELECT ac.first_name, ac.last_name
FROM movies_db.movies mo
INNER JOIN movies_db.actor_movie ac_mo ON mo.id = ac_mo.movie_id 
INNER JOIN movies_db.actors ac ON ac.id = ac_mo.id
WHERE mo.title LIKE "La Gue%"
GROUP BY ac.first_name, ac.last_name
