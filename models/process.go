package models

import (
	"log"
	"time"
	"timeMonitorServer/global"
	"timeMonitorServer/types"
)

func FindProcessId(form types.ProcessModel) types.ProcessModel {
	var res types.ProcessModel
	global.Mdb.Where("process = ? AND date = ? AND hour = ?", form.Process, form.Date, form.Hour).Find(&res)
	return res
}

// FindTitleByIdAndTitle 获取title
func FindTitleByIdAndTitle(id uint, title string) types.TitleModel {
	var res types.TitleModel
	global.Mdb.Where("process_id = ? AND title = ?", id, title).Find(&res)
	return res
}

// InsertAllProcessAndTitle 插入
func InsertAllProcessAndTitle(form []types.UploadForm) {
	db := global.Mdb

	// 循环处理 process
	for i := 0; i < len(form); i++ {
		j := types.ProcessModel{
			Process: form[i].Process,
			Date:    form[i].Time.Format("2006-01-02"),
			Hour:    uint8(form[i].Time.Hour()),
		}

		res := FindProcessId(j)

		if res.Id == 0 {
			// 未查询到 插入该数据
			db.Create(&j)

			j1 := FindProcessId(j)

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

func FindAllByDay() []types.ProcessModel {
	db := global.Mdb

	day := time.Now().Format("2006-01-02")

	var res []types.ProcessModel

	// 使用Preload加载关联的Titles数据
	err := db.Where("date = ?", day).
		Preload("Titles").
		Find(&res).Error

	if err != nil {
		log.Fatalf("Failed to query data: %v", err)
	}

	return res
}
