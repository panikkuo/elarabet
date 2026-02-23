CREATE OR REPLACE FUNCTION add_user(
    _username VARCHAR(256), 
    _email VARCHAR(256), 
    _password VARCHAR(256), 
    _name VARCHAR(256))
RETURNS TABLE(
    id_ uuid, 
    username_ VARCHAR(256)) AS $$
BEGIN
    INSERT INTO users(username, email, password, name)
    VALUES(_username, _email, _password, _name);

    RETURN QUERY
    SELECT id, username
    FROM users
    WHERE username = _username;
END;
$$ LANGUAGE plpgsql