CREATE SCHEMA "public";
CREATE TABLE public."user"
(
  id           SERIAL PRIMARY KEY,
  username     VARCHAR(20) NOT NULL,
  password_md5 VARCHAR(32),
  token        VARCHAR(100),
  note         VARCHAR(2000)
);
CREATE UNIQUE INDEX user_id_uindex
  ON public."user" (id);

CREATE TABLE public.report
(
  id            SERIAL PRIMARY KEY NOT NULL,
  summerization VARCHAR(2000),
  date          TIMESTAMP,
  fk_user_id    INT,
  CONSTRAINT report_user_id_fk FOREIGN KEY (fk_user_id) REFERENCES "user" (id)
)

CREATE UNIQUE INDEX report_id_uindex
  ON public.report (id)

CREATE TABLE public.todo
(
  id           SERIAL PRIMARY KEY,
  content      VARCHAR(300) NOT NULL,
  status       INT,
  fk_report_id INT,
  CONSTRAINT todo_report_id_fk FOREIGN KEY (fk_report_id) REFERENCES report (id)
);
CREATE UNIQUE INDEX todo_id_uindex
  ON public.todo (id);

CREATE TABLE public."group"
(
  id   SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL
);
CREATE UNIQUE INDEX group_id_uindex
  ON public."group" (id);

CREATE TABLE public.participate
(
  fk_user_id  INT,
  fk_group_id INT,
  CONSTRAINT participate_user_id_fk FOREIGN KEY (fk_user_id) REFERENCES "user" (id),
  CONSTRAINT participate_group_id_fk FOREIGN KEY (fk_group_id) REFERENCES "group" (id)
);
CREATE INDEX participate_fk_user_id_index
  ON public.participate (fk_user_id);
CREATE INDEX participate_fk_group_id_index
  ON public.participate (fk_group_id);

CREATE TABLE public.comment
(
  id           SERIAL PRIMARY KEY,
  content      VARCHAR(500),
  date         TIMESTAMP,
  fk_user_id   INT,
  fk_report_id INT,
  CONSTRAINT comment_user_id_fk FOREIGN KEY (fk_user_id) REFERENCES "user" (id),
  CONSTRAINT comment_report_id_fk FOREIGN KEY (fk_report_id) REFERENCES report (id)
);
CREATE UNIQUE INDEX comment_id_uindex
  ON public.comment (id);
