package seed

var MIGRATIONS = []string {
	`delete from catalogs;`,
	`delete from slots;`,
	`insert into catalogs("name") values ('The god father'), ('Inception'), ('Prestige');`,
	`insert into slots("date") values (now()), (now() + interval  '1 day'), (now() + interval  '2 day');`,

}
