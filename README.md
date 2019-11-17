# SQL-Image

SQL-Image is a daemon-less docker image analyzer.

SQL-Image treat docker-images as DATABASES and image-layers as TABLES.

Inspired by https://github.com/wagoodman/dive and https://github.com/kashav/fsql.

## Syntax
~~~~sql
SHOW IMAGES;

USE my-image:v1;

SHOW LAYERS;

SELECT … FROM my-layer WHERE ...;
~~~~

![1](doc/SQL-Image-1.jpeg)

![2](doc/SQL-Image-2.jpeg)

## Note 
Only support linux overlay2 docker environment.