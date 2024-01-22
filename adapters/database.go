
package adapters

import (
    "gorm.io/gorm"
	"fmt"
    "gorm.io/driver/mysql"
)

// *gorm.DB is the database connection object
func NewDatabase() (*gorm.DB, error) {
	//Database connection parameters
    username := "root"
    password := "1234"
    host := "localhost"
    databaseName := "rails_studentDashboard_development"
    
    //Create the DSN (Data Source Name)- contains all the information needed to connect to the database.
    dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",username, password, host, databaseName)
    return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
