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

	// Inserción usando una consulta preparada para evitar SQL Injection
	sentencia := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES (?, ?, ?)"
	fmt.Println("Sentencia preparada: ", sentencia)

	_, err = Db.Exec(sentencia, sig.UserEmail, sig.UserUUID, tools.FechaMySQL())
	if err != nil {
		fmt.Printf("Error al ejecutar sentencia SQL: %v\n", err)
		return err
	}

	fmt.Println("SignUp > Ejecución exitosa")
	return nil
}
