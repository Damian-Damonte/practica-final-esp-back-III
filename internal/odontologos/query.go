package odontologos

var (
	QueryGetAllOdontologos = `SELECT id, apellido, nombre, matricula FROM odontologos`
	QueryGetOdontologoById = `SELECT id, apellido, nombre, matricula FROM odontologos WHERE id = ?`
	QuertyInsertOdontologo = `INSERT INTO odontologos (apellido, nombre, matricula) VALUES (?,?,?)`
)