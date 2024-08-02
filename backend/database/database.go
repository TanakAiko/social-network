package database

import "database/sql"

type Database struct {
	Db *sql.DB
}

// func InitDb(name, path string) (*Database, error) {
// 	_, err := os.Stat(path + name)
// 	if os.IsNotExist(err) {
// 		file, err := os.Create(path + name)
// 		if err != nil {
// 			return fmt.Errorf("failed to create database file: %v", err)
// 		}
// 		file.Close()
// 	}

// 	if _, err := os.Stat(path + "migrates"); os.IsNotExist(err) {
// 		err := os.Mkdir(path+"migrates", 0755)
// 		if err != nil {
// 			return fmt.Errorf("failed to create migrates directory: %v", err)
// 		}
// 	}

// 	o.Db, err = sql.Open("sqlite3", path+name)
// 	if err != nil {
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	PATH = path
// 	return nil
// }

func (d *Database) CreateDatabase() error {

	return nil
}

func (d *Database) Close() {

}
