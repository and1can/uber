var pg = require('pg');
var connectionString = process.env.DATABASE_URL || 'postgres://localhost:5432/uber';

var client = new pg.Client(connectionString);
client.connect();
var query = client.query('CREATE TABLE history(id serial primary key, start integer, distance integer, reqid text, totalcharged integer)');
//var query = client.query('CREATE TABLE images(img bytea)');
var query = client.query('INSERT INTO history(id, start, distance, reqid, totalcharged) values ($1, $2, $3, $4, $5)', [1, 1468972800, 10, 'someid1', 15]);
var query = client.query('INSERT INTO history(id, start, distance, reqid, totalcharged) values ($1, $2, $3, $4, $5)', [2, 1469059200, 3, 'someid2', 5]);
var query = client.query('INSERT INTO history(id, start, distance, reqid, totalcharged) values ($1, $2, $3, $4, $5)', [3, 1469145600, 2, 'someid3', 3]);
query.on('end', function() { client.end(); }); 