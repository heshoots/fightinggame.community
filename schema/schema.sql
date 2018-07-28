create extension if not exists postgis;

create table scenes (
  id serial primary key,
  name varchar(50),
  location GEOMETRY(Point, 26910)
);

create index scene_gix on scenes using gist (location);

create table contactpoints (
  id serial primary key,
  scene_id integer references scenes(id), 
  twitter text,
  discord text,
  facebook text 
);

create table games (
  id serial primary key,
  scene_id integer references scenes(id), 
  street_fighter_v boolean
);
