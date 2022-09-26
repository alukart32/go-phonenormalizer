WITH _data (id, number) AS (
	VALUES
	 %s
)
UPDATE public.phones AS phone
SET number = _data.number
FROM _data
WHERE phone.id = _data.id