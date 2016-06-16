package database

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/tebben/marvin/go/marvin/models"
	"log"
	"path"
	"runtime"
	"time"
)

var open bool
var dbName = "marvin.db"
var moduleBucketName = "module"

type Database struct {
	bolt *bolt.DB
}

func (db *Database) Open() error {
	var err error
	_, filename, _, _ := runtime.Caller(0) // get full path of this file
	dbFile := path.Join(path.Dir(filename), dbName)
	config := &bolt.Options{Timeout: 1 * time.Second}
	db.bolt, err = bolt.Open(dbFile, 0600, config)
	if err != nil {
		log.Fatal(err)
	}

	open = true

	return nil
}

func (db *Database) Close() {
	open = false
	db.bolt.Close()
}

func (db *Database) InsertModuleSettings(module models.MarvinModule) error {
	if !open {
		return fmt.Errorf("db must be opened before saving!")
	}
	err := db.bolt.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(moduleBucketName))
		settings := module.GetSettings()
		if settings == nil {
			return fmt.Errorf("No settings in module: %s", module.GetName())
		}

		enc, err := json.Marshal(settings)
		if err != nil {
			return err
		}

		if err != nil {
			return fmt.Errorf("could not encode module settings %s: %s", module.GetName(), err)
		}
		err = b.Put([]byte(module.GetName()), enc)
		return err
	})
	return err
}

func (db *Database) ModuleSettingsExist(module models.MarvinModule) (bool, error) {
	if !open {
		return false, fmt.Errorf("db must be opened before reading!")
	}

	err := db.bolt.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(moduleBucketName))
		exist := b.Get([]byte(module.GetName()))
		if exist != nil {
			return nil
		}

		return fmt.Errorf("Settings not found!")
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

func (db *Database) GetModuleSettings(module models.MarvinModule, settings interface{}) error {
	if !open {
		return fmt.Errorf("db must be opened before reading!")
	}

	err := db.bolt.View(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket([]byte(moduleBucketName))
		k := []byte(module.GetName())

		err = json.Unmarshal(b.Get(k), &settings)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		fmt.Printf("Could not get module ID %s", module.GetName())
		return err
	}

	return nil
}
