package main

//import (
	////"fmt"
	////"encoding/json"
	//"github.com/boltdb/bolt"
//)

//func DBSavePost (p JSONPost) (int, error) {
	//db, err := bolt.Open(CurrentConfig.DB, 0644, nil)
	//if err != nil {
		//panic(err)
	//}
	////defer db.Close
	
	//newKey := getNewKey(db, "FlexPosts")
	
	//p.ID = newKey
	
	//err = db.Update(func(tx *bolt.Tx) error {
        //bucket, err := tx.CreateBucketIfNotExists("FlexPosts")
        //if err != nil {
            //return err
        //}

        //err = bucket.Put(key, value)
        //if err != nil {
            //return err
        //}
        //return nil
    //})
//}

//func getNewKey(db, bucket string) int {
	//err = db.Read(func(tx *bolt.Tx) error {
        //keyBucket, err := tx.CreateBucketIfNotExists("LastKeys")
        //if err != nil {
            //return err
        //}

        //err = keyBucket.Get(bucket, value)
        //if err != nil {
            //return err
        //}
        //return nil
    //})
//}
