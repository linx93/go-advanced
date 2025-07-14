package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	err := readJsonFileToDb("./address.json")
	if err != nil {
		log.Println("读取失败:", err.Error())
		return
	}

}

// 读取省市区县三级行政区json数据到数据库中
func readJsonFileToDb(jsonFileDir string) error {
	file, err := os.ReadFile(jsonFileDir)
	if err != nil {
		return err
	}

	type model3 struct {
		Code string
		Name string
	}

	type model2 struct {
		Code     string
		Name     string
		Children []model3
	}

	type model1 struct {
		Code     string
		Name     string
		Children []model2
	}

	ms := make([]model1, 0)

	err = json.Unmarshal(file, &ms)
	if err != nil {
		return err
	}

	db, err := getDB("root", "root", "192.168.2.11", "customer", 33306)
	if err != nil {
		return err
	}

	//Level 1-省,2-市,3-区县

	rs := make([]*Regions, 0)

	for _, m1 := range ms {
		l1 := &Regions{ParentCode: "", Level: 1, Code: m1.Code, Name: m1.Name, ShortName: m1.Name}
		rs = append(rs, l1)
		if m1.Children != nil && len(m1.Children) > 0 {
			for _, m2 := range m1.Children {
				l2 := &Regions{ParentCode: m1.Code, Level: 2, Code: m2.Code, Name: m1.Name + "/" + m2.Name, ShortName: m2.Name}
				rs = append(rs, l2)
				if m2.Children != nil && len(m2.Children) > 0 {
					for _, m3 := range m2.Children {
						l3 := &Regions{ParentCode: m2.Code, Level: 3, Code: m3.Code, Name: m1.Name + "/" + m2.Name + "/" + m3.Name, ShortName: m3.Name}
						rs = append(rs, l3)
					}
				}
			}
		}
	}

	err = db.Model(&Regions{}).CreateInBatches(rs, 100).Error
	if err != nil {
		return err
	}

	return nil
}

func getDB(user, pwd, host, database string, port int) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pwd, host, port, database)), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

type Regions struct {
	ID         uint   `gorm:"primarykey" json:"id"`                 //主键ID
	ParentCode string `json:"parentCode" gorm:"column:parent_code"` //父级code
	Name       string `json:"name"`                                 //行政区划全名
	ShortName  string `json:"shortName"  gorm:"column:short_name"`  //行政区划简称
	Code       string `json:"code"`                                 //行政区划代码
	Level      uint8  `json:"level"`                                //1-省,2-市,3-区县
	Sort       uint   `json:"sort"`                                 //排序字段
	Status     uint8  `json:"status"`                               //0-禁用,1-启用
}

func (Regions) TableName() string {
	return "sys_regions"
}
