drop table messages;
create table messages(id integer not null primary key autoincrement, content text not null, room_id integer);
