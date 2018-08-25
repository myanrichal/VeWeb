package migrations

//Build initial empty database

import (
	"github.com/rubenv/sql-migrate"
)

func CreateInitial() *migrate.Migration {
	create_initial := migrate.Migration {
		Id: "1", 
		Up: []string {
			`
			CREATE TABLE ve_inventory (
				id int,
				name varchar(255),
				condition varchar(500),
				dateOfPurchase datetime
				slug varchar(255),
				category int
				PRIMARY KEY (id)
			) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
			`,
			`
			CREATE TABLE ve_people (
				id int,
				name varchar(255),
				qualification int,
				wavier varchar(255),
				PRIMARY KEY (id)
			) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
			`,
			`
			CREATE TABLE ve_people_to_inventory (
				id int,
				inv_id int,
				person_id int,
				date_out datetime,
				date_in datetime,
				condition_in varchar(500),
				condition_out varchar(500)
				PRIMARY KEY (id)
			) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
			`,
		},
		Down: []string{
			` DROP TABLE ve_inventory; `,
			` DROP TABLE ve_people; ` ,
			` DROP TABLE ve_people_to_inventory; `,
		},
	}

	return &create_initial
}