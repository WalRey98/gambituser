package bd

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"go.mod/models"
	"go.mod/tools"
)

// SignUp inserta un nuevo usuario en la base de datos sin verificar duplicados en el correo electrónico
func SignUp(sig models.SignUp) error {
	fmt.Println("Comienza el registro de usuario")

	// Conectar a la base de datos
	err := DbConnect()
	if err != nil {
		return fmt.Errorf("Error al conectar a la base de datos: %v", err)
	}

	defer Db.Close()

	// Verificamos que los valores esten correctos
	fmt.Println("UserEmail:", sig.UserEmail)
	fmt.Println("UserUUID", sig.UserUUID)

	// Preparamos la consulta con valores dinámicos
	stmt, err := Db.Prepare("INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println("Error al preparar la consulta:", err)
		return err
	}
	defer stmt.Close() // Cerramos el statement al final

	// Ejecutamos la inserción
	_, err = stmt.Exec(sig.UserEmail, sig.UserUUID, tools.FechaMySQL())
	if err != nil {
		fmt.Println("Error al ejecutar la inserción:", err)
		return err
	}

	fmt.Println("SingUp > Ejecución Exitosa")
	return nil
}
