package bd

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"go.mod/models"
	"go.mod/secretm"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

// ReadSecret obtiene las credenciales de la base de datos desde AWS Secrets Manager
func ReadSecret() error {
	var err error
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		return fmt.Errorf("Error al leer el secreto: %v", err)
	}
	return nil
}

// DbConnect conecta a la base de datos usando las credenciales obtenidas de AWS Secrets Manager
func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Printf("Error al abrir la conexión a la base de datos: %v\n", err)
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Printf("Error al hacer ping a la base de datos: %v\n", err)
		return err
	}

	fmt.Println("Conexión exitosa a la base de datos!")
	return nil
}

// ConnStr construye la cadena de conexión a MySQL
func ConnStr(claves models.SecretRDSJson) string {
	dbUser := claves.Username
	authToken := claves.Password
	dbEndpoint := claves.Host
	dbName := "gambit"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println("Cadena de conexión: ", dsn)
	return dsn
}
