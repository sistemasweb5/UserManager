-- Tables
-- Table: category
CREATE TABLE category (
    id UUID NOT NULL,
    rol VARCHAR NOT NULL,
    CONSTRAINT category_pk PRIMARY KEY (id)
);

-- Table: specialty
CREATE TABLE specialty (
    id UUID NOT NULL,
    name VARCHAR NOT NULL,
    clientId UUID NOT NULL,
    CONSTRAINT specialty_pk PRIMARY KEY (id)
);

-- Table: work_schedule
CREATE TABLE workSchedule (
    id UUID NOT NULL,
    startTime VARCHAR NOT NULL,
    endTime VARCHAR NOT NULL,
    CONSTRAINT work_schedule_pk PRIMARY KEY (id)
);
-- Table: client
CREATE TABLE client (
    id UUID NOT NULL,
    name VARCHAR NOT NULL,
    emailAddress VARCHAR NOT NULL,
    categoryId UUID NOT NULL,
    workScheduleId UUID NOT NULL,
    CONSTRAINT client_pk PRIMARY KEY (id)
);
