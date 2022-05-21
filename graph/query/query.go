package query

var GetAllItems string = `
	SELECT id, name
	FROM item
`

var GetItem string = `
	SELECT name
	FROM item
	WHERE id = $1
`