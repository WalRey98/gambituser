package bd

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"go.mod/models"
	"go.mod/tools"
)

// SignUp inserta un nuevo usuario en la base de datos, verificando que no exista duplicado
func SignUp(sig models.SignUp) error {
	fmt.Println("Comienza el registro de usuario")

	// Conectar a la base de datos
	err := DbConnect()
	if err != nil {
		return fmt.Errorf("Error al conectar a la base de datos: %v", err)
	}

	defer Db.Close()

	// Verificar si el correo ya está registrado
	var count int
	err = Db.QueryRow("SELECT COUNT(*) FROM users WHERE User_Email = ?", sig.UserEmail).Scan(&count)
	if err != nil {
		fmt.Printf("Error al verificar si el email existe: %v\n", err)
		return err
	}

	if count > 0 {
		return fmt.Errorf("El email %s ya está registrado", sig.UserEmail)
	}

	// Inserción usando una consulta preparada para evitar SQL Injection
	sentencia := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES (?, ?, ?)"
	fmt.Println("Sentencia preparada: ", sentencia)

	_, err = Db.Exec(sentencia, sig.UserEmail, sig.UserUUID, tools.FechaMySQL())
	if err != nil {
		if sqlErr, ok := err.(*mysql.MySQLError); ok && sqlErr.Number == 1062 {
			fmt.Println("Error: El email ya está registrado.")
			return fmt.Errorf("El email %s ya está registrado", sig.UserEmail)
		}
		fmt.Printf("Error al ejecutar sentencia SQL: %v\n", err)
		return err
	}

	fmt.Println("SignUp > Ejecución exitosa")
	return nil
}
