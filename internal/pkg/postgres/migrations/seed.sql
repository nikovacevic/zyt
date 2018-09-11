INSERT INTO users (email, password) VALUES
    ('niko@zyt.works', crypt('password', gen_salt('bf', 8)));

INSERT INTO events (user_id, start_time, end_time, duration, name, notes) 
    SELECT id, now(), null, null, 'Work on ZYT API', null FROM users WHERE email='niko@zyt.works';