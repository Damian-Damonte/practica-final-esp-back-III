package pacientes

const (
	QueryGetAllPacientes  = `SELECT id, apellido, nombre, domicilio, dni, fecha_alta FROM pacientes`
	QueryGetPacientesById = `SELECT id, apellido, nombre, domicilio, dni, fecha_alta FROM pacientes WHERE id = ?`
	QuertyInsertPaciente  = `INSERT INTO pacientes (apellido, nombre, domicilio, dni, fecha_alta) VALUES (?,?,?,?,?);
	`
)
