// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package system

import (
	"crypto/md5"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	. "rtsys/common/models/system"
	"rtsys/utils/core"
	"rtsys/utils/trans"
	"strconv"
	"sync"
	"time"
)

type AppTokenDao struct {
	Model *AppTokenModel
}

var (
	instanceAppTokenDao *AppTokenDao //单例的对象
	onceAppTokenDao     sync.Once
)

func NewAppTokenDao() *AppTokenDao {
	onceAppTokenDao.Do(func() {
		instanceAppTokenDao = &AppTokenDao{Model: NewAppTokenModel()}
	})
	return instanceAppTokenDao
}

func (d *AppTokenDao) GetInfoByToken(token string) *AppTokenModel {
	if token == "" {
		return nil
	}
	model := d.Model.New()
	err := core.NewDao(d.Model).First(model, "token = ?", token)
	if err != nil {
		return nil
	}
	return model
}

// GetToken 查询token 信息
func (d *AppTokenDao) GetToken(token string, userType string, plat string) *AppTokenModel {
	model := d.GetInfoByToken(token)
	if model == nil || model.Token == "" || model.Type != userType || model.Plat != plat || model.ExpTime == nil {
		return nil
	}
	if model.ExpTime.Before(time.Now()) {
		return nil
	}
	return model
}

// SetToken 设置保存token
func (d *AppTokenDao) SetToken(userId string, userType string, plat string, extend string) (string, error) {
	if userId == "" {
		return "", errors.New("用户ID丢失")
	}

	//生成token
	times := strconv.FormatInt(time.Now().UnixNano(), 10)
	randInt, err1 := rand.Int(rand.Reader, big.NewInt(1000))
	if err1 != nil {
		return "", errors.New("随机数生成失败")
	}
	rands := randInt.String()
	token := times + "_" + userId + "_" + userType + "_" + plat + "_" + rands
	token = fmt.Sprintf("%x", md5.Sum([]byte(token)))

	dao := core.NewDao(d.Model)
	//查询已存在的
	model := d.Model.New()
	_ = dao.SetField("token").First(model, "user_id = ? AND type = ? AND plat = ?", userId, userType, plat)

	//生成字段信息
	expTime := trans.TimeFormat(time.Now().AddDate(0, 0, 7), "")
	fields := []string{"token", "user_id", "type", "plat", "extend", "exp_time"}
	values := []any{token, userId, userType, plat, extend, expTime}

	//更新或插入新的
	if model.Token == "" {
		err4 := dao.InsertNoId(fields, values)
		if err4 != nil {
			return "", err4
		}
	} else {
		effected, err5 := dao.Update(fields, values, "token = ?", model.Token)
		if err5 != nil {
			return "", err5
		}
		if effected == 0 {
			return "", errors.New("更新token信息失败")
		}
	}
	return token, nil
}
