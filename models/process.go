package models

import (
	"fmt"
	"log"
	"time"
	"timeMonitorServer/global"
	"timeMonitorServer/types"
)

func FindUserIdByUserName(username string) (uint, error) {
	var res types.UserModel
	global.Mdb.Where("username = ?", username).First(&res)

	if res.Id == 0 {
		return 0, fmt.Errorf("<UNK>")
	}

	return res.Id, nil
}

func FindProcessId(form types.ProcessModel, userId uint) types.ProcessModel {
	var res types.ProcessModel
	global.Mdb.Where("process = ? AND date = ? AND hour = ? AND user_id = ?", form.Process, form.Date, form.Hour, userId).Find(&res)
	return res
}

// FindTitleByIdAndTitle 获取title
func FindTitleByIdAndTitle(id uint, title string) types.TitleModel {
	var res types.TitleModel
	global.Mdb.Where("process_id = ? AND title = ?", id, title).Find(&res)
	return res
}

// InsertAllProcessAndTitle 插入
func InsertAllProcessAndTitle(form []types.UploadForm, userId uint) {
	db := global.Mdb

	// 循环处理 process
	for i := 0; i < len(form); i++ {
		j := types.ProcessModel{
			Process: form[i].Process,
			UserId:  userId,
			Date:    form[i].Time.Format("2006-01-02"),
			Hour:    uint8(form[i].Time.Hour()),
		}

		res := FindProcessId(j, userId)

		if res.Id == 0 {
			// 未查询到 插入该数据
			db.Create(&j)

			j1 := FindProcessId(j, userId)

			n := types.TitleModel{
				ProcessId: j1.Id,
				Title:     form[i].Title,
				Time:      1,
			}

			db.Create(&n)
		} else {
			// 查询到 根据ID进行处理title
			t := FindTitleByIdAndTitle(res.Id, form[i].Title)

			if t.Id == 0 {
				// 未查询到title 新增

				n := types.TitleModel{
					ProcessId: res.Id,
					Title:     form[i].Title,
					Time:      1,
				}

				db.Create(&n)
			} else {
				// 查询到 增加秒
				db.Model(&t).Where("id = ?", t.Id).Update("time", t.Time+1)
			}
		}
	}
}

func FindAllByDay(userId uint, day string) []types.ProcessModel {
	db := global.Mdb

	var res []types.ProcessModel

	// 使用Preload加载关联的Titles数据
	err := db.Where("date = ? AND user_id = ?", day, userId).
		Preload("Titles").
		Find(&res).Error

	if err != nil {
		log.Fatalf("Failed to query data: %v", err)
	}

	return res
}

func EditUserTime(username string, password string, dailyTime uint, everyTime uint) error {
	db := global.Mdb

	var user types.UserModel

	limit := db.Where("username = ? AND password = ?", username, password).First(&user)
	if limit.RowsAffected == 0 {
		return fmt.Errorf("<UNK>")
	}

	user.EveryTime = everyTime
	user.DailyTime = dailyTime

	db.Save(&user)

	return nil
}

func FindDailyLimit(username string) uint {
	var form types.UserModel
	global.Mdb.Where("username = ?", username).First(&form)

	return form.DailyTime
}

func FindLimit(userId uint) types.UserModel {
	var form types.UserModel
	global.Mdb.Where("id = ?", userId).First(&form)

	return form
}

func ComputedAll(username string) (uint, error) {

	db := global.Mdb

	id, err := FindUserIdByUserName(username)
	if err != nil {
		return 0, err
	}

	var all uint

	targetDate := time.Now().Format("2006-01-02")

	sqlQuery := `
		SELECT SUM(t.time) as 'all'
		FROM processes p
		JOIN titles t ON t.process_id = p.id
		WHERE p.date = ? AND p.user_id = ?;
	`
	// 注意：在 GORM Raw 方法中，参数替换符是 `?`。GORM 会安全地转义和替换。
	err = db.Raw(sqlQuery, targetDate, id).Scan(&all).Error
	if err != nil {
		return 0, err
	}

	return all, nil

}

func FindTitleClass() []types.TitleClassModel {
	db := global.Mdb

	var form []types.TitleClassModel

	db.Find(&form)

	return form
}

func FindLastTime(userId uint) time.Time {
	db := global.Mdb

	var maxUpdateTime time.Time
	// Using Joins and Select for MAX(t.update_time)
	sqlQuery := `
				SELECT MAX(t.update_time)
				FROM processes p , titles t
				WHERE p.user_id = ?
				`
	// 注意：在 GORM Raw 方法中，参数替换符是 `?`。GORM 会安全地转义和替换。
	err := db.Raw(sqlQuery, userId).Scan(&maxUpdateTime).Error
	if err != nil {
		return time.Time{}
	}

	return maxUpdateTime
}

func UpdateAllDailyTime() error {
	db := global.Mdb

	sqlQuery := `
             UPDATE users
             SET daily_time = every_time,
                 update_time = ?
             `
	// 使用 db.Exec 而不是 db.Raw 来执行不返回结果的 DML (Data Manipulation Language) 语句，如 UPDATE。
	// db.Exec 返回的 Result 包含 RowsAffected。
	// time.Now() 作为参数传递，GORM 会安全地将其转换为数据库时间格式。
	result := db.Exec(sqlQuery, time.Now()) // 将当前时间作为参数传递

	if result.Error != nil {
		return result.Error
	}

	return nil
}
