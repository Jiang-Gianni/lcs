-- drop table question;
-- drop table hint;
-- drop table editor;

create table question(
    id integer primary key not null,
    question_id text not null,
    link text not null,
    title text not null,
    title_slug text not null,
    is_paid_only boolean not null,
    difficulty text not null,
    content text not null
);

create table hint(
    id integer primary key not null,
    question_id text not null,
    hint text not null
);

create table editor(
    id integer primary key not null,
    question_id text not null,
    lang text not null,
    lang_slug text not null,
    code text not null
);
