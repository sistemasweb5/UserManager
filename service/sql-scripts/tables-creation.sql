-- Tables

-- Table: category
CREATE TABLE category (
    id UUID NOT NULL,
    rol VARCHAR NOT NULL,
    CONSTRAINT category_pk PRIMARY KEY (id)
);

-- Table: "customer"
CREATE TABLE client (
    id UUID NOT NULL,
    name VARCHAR NOT NULL,
    email_address VARCHAR NOT NULL,
    category_id UUID NOT NULL,
    CONSTRAINT user_pk PRIMARY KEY (id)
);
