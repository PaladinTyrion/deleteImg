package main

import (
	"deleteimg/config"
	"deleteimg/models"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {

	MongoDBInfo := mgo.DialInfo{
		Addrs:    []string{config.Mongodbhosts},
		Database: config.Database,
		Username: config.Username,
		Password: config.Password,
		Timeout:  config.Timeout,
	}

	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	session, err := mgo.DialWithInfo(&MongoDBInfo)
	if err != nil {
		log.Fatal("CreateSession:", err)
	}
	//	Schedule the close to occur once the function returns.
	defer session.Close()

	//	Get && Operate my expect Collection
	c := session.DB(config.Database).C(config.Collection)

	//	Return the count of Collection
	count, err := c.Count()
	if err != nil {
		log.Fatal("Count Of Collection err: ", err)
	}
	fmt.Println(count)

	//	查询结果放入result
	var result []models.Trick
	query := bson.M{}

	//	Get all Query result && Print field of Screenshot
	err = c.Find(query).All(&result)
	if err != nil {
		log.Fatal("Fail to get Collection: ", err)
	}

	//	查询后的screenshot名放入preserveImg []string
	preserveImg := make([]string, 0)
	for k, v := range result {
		shotname := v.Screenshot
		fmt.Println("Num" + strconv.Itoa(int(k)) + "   " + shotname)
		preserveImg = append(preserveImg, shotname)
	}
	fmt.Println(preserveImg)

	//	操作路径，读取路径所含文件
	pathname := config.DeletePathName
	files, err := ioutil.ReadDir(pathname)

	if err != nil {
		log.Fatal(err)
	}

LABEL:
	//	遍历文件，每遍历一个，查询slice中是否含有该名，没有则删除文件
	for _, file := range files {
		if file.IsDir() {
			continue LABEL
		} else {
			filename := file.Name()
			for _, preImg := range preserveImg {
				if string(preImg) == string(filename) {
					continue LABEL
				}
			}
			err = os.Remove(pathname + filename)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(pathname, filename, "删除成功!")
		}
	}

}
