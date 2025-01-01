CREATE TABLE projects (
    id character varying NOT NULL,
    name character varying(20) NOT NULL,
    created timestamp without time zone NOT NULL,
    enabled boolean NOT NULL
);

CREATE TABLE languages (
    id character varying NOT NULL,
    name character varying(20) NOT NULL,
    created timestamp without time zone NOT NULL,
    enabled boolean NOT NULL
);


CREATE TABLE resources (
    id character varying NOT NULL,
    project character varying(20) NOT NULL,
    language character varying(20) NOT NULL,
    sem_ver character varying(50) NOT NULL,
    "values" jsonb,
    created timestamp without time zone NOT NULL,
    enabled boolean NOT NULL
);


ALTER TABLE ONLY projects
    ADD CONSTRAINT projects_pkey PRIMARY KEY (id);

ALTER TABLE ONLY languages
    ADD CONSTRAINT languages_pkey PRIMARY KEY (id);

ALTER TABLE ONLY resources
    ADD CONSTRAINT resources_pkey PRIMARY KEY (id);


ALTER TABLE ONLY resources
    ADD CONSTRAINT unique_project_language_semver UNIQUE (project, language, sem_ver);

ALTER TABLE ONLY resources
    ADD CONSTRAINT fk_project FOREIGN KEY (project) REFERENCES projects(id);

ALTER TABLE ONLY resources
    ADD CONSTRAINT fk_language FOREIGN KEY (language) REFERENCES languages(id);
