package main

import (
	//"fmt"
	//"encoding/json"
	"github.com/boltdb/bolt"
	"strconv"
)

func DBSavePost (p FlexPost) ([]byte, error) {
	db, err := bolt.Open(CurrentConfig.DB, 0644, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	//p.ID = newKey
	
	newKey := []byte("")
	
	err = db.Update(func(tx *bolt.Tx) error {
        bucket, err := tx.CreateBucketIfNotExists([]byte("FlexPosts"))
        if err != nil {
            return err
        }

		//encode struct to json here
		value := ""
		
		key, _ := bucket.NextSequence()
		newKey := []byte(strconv.FormatUint(key, 64))
        if err != nil {
            return err
        }

        err = bucket.Put(newKey, []byte(value))
        if err != nil {
            return err
        }
        return nil
    })
    return newKey, err
}
