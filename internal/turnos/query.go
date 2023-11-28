package turnos

var (
	QueryGetAllTurnos = `SELECT t.id, t.descripcion, t.fecha_hora, o.*, p.* FROM turnos t
	INNER JOIN odontologos o ON t.odontologos_id = o.id
	INNER JOIN pacientes p ON t.pacientes_id = p.id`
	QueryGetTurnoById = `SELECT t.id, t.descripcion, t.fecha_hora, o.*, p.* FROM turnos t
	INNER JOIN odontologos o ON t.odontologos_id = o.id
	INNER JOIN pacientes p ON t.pacientes_id = p.id
	WHERE t.id = ?`
	QueryGetTurnoByPacienteDni = `SELECT t.id, t.descripcion, t.fecha_hora, o.*, p.* FROM turnos t
	INNER JOIN odontologos o ON t.odontologos_id = o.id
	INNER JOIN pacientes p ON t.pacientes_id = p.id
	WHERE p.dni = ?;`
)
