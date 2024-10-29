CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


INSERT INTO category VALUES
(uuid_generate_v4(), 'applicant'),
(uuid_generate_v4(), 'worker');

INSERT INTO client VALUES
('aaaaaaaa-1111-1111-1111-111111111111', 'Slim Shady', 'shady@not-calling.com', 'b1a1c1a1-1111-1111-1111-111111111111'),
('bbbbbbba-1111-1111-1111-111111111111', 'Rain man', 'rain-man@not-calling.com', 'b1a1c1a1-1111-1111-1111-111111111112');
