create table if not exists users(
	email varchar(50) not null,
	username varchar(41) not null,
	sales int,
	PRIMARY KEY (email)
);

create table if not exists listings(
	email varchar(50) not null,
	listID int not null,
	title varchar(50),
	description varchar(100),
	img varchar(50),
	active boolean not null,
	category varchar(50),
	rarity varchar(15),
	foreign key (email) references users(email) on delete cascade,
	primary key (email, listID)
);

create table if not exists actions(
	aID int not null,
	time datetime, 
	description varchar(100),
	email varchar(50),
	primary key (aID)
);

create table if not exists chats(
	cID int not null,
	users varchar(110) not null,
	primary key (cID)
);

create table if not exists messages(
	mID int not null,
	cID int,
	message varchar(150) not null,
	sender varchar(50) not null,
	receiver varchar(50) not null,
	time datetime,
	primary key (mID),
	foreign key (cID) references chats(cID) on delete set null
);
