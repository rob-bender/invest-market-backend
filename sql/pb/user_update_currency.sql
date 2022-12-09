CREATE OR REPLACE FUNCTION user_update_currency(js json)
	RETURNS void
	LANGUAGE plpgsql
AS $function$
DECLARE
	_u users;
BEGIN
	SELECT *
	FROM users
	WHERE tele_id = (js->>'tele_id')::BIGINT
	INTO _u;

	IF _u.id ISNULL THEN
		RAISE EXCEPTION 'пользователь не найден';
	END IF;

	UPDATE users SET currency = js->>'currency' WHERE tele_id = (js->>'tele_id')::BIGINT;
END;
$function$
