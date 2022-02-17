package logic

import (
	"go-scaffold/dao/rds"
	"go-scaffold/models"
	"go-scaffold/pkg/jwt"
	"go-scaffold/pkg/snowflake"

	"go.uber.org/zap"
)

// SignUP 注册的业务逻辑
func SignUP(params *models.BodySignUP) (err error) {
	// 判断用户是否存在
	if err = rds.GetUserDB().CheckExistByEmail(params.Email); err != nil {
		return err
	}

	// 通过雪花算法生成UID
	id := snowflake.GenerateID()

	// 对密码进行加密
	encodePWD, err := rds.GetUserDB().EncryptPassword(params.Password)
	if err != nil {
		return err
	}

	// 构造用户实例
	user := models.User{
		BasicMode: models.BasicMode{ID: id},
		Email:     params.Email,
		Password:  encodePWD,
		Nickname:  params.Nickname,
		Gender:    params.Gender,
	}

	// 保存到数据库
	return rds.GetUserDB().Create(&user)
}

// SignIn 登录的业务逻辑
func SignIn(params *models.BodySignIn) (user *models.User, token string, err error) {
	// 查询用户
	user, err = rds.GetUserDB().GetByEmail(params.Email)
	if err != nil {
		return nil, "", err
	}

	// 验证密码
	err = rds.GetUserDB().ComparePassword(user.Password, params.Password)
	if err != nil {
		return nil, "", err
	}

	// 生成JWT Token
	token, err = jwt.GenerateToken(user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

// ListUser 获取用户列表
func ListUser(params *models.QueryUserInfoListNull) (data *models.DataUserInfoList, err error) {
	// 查询用户列表
	users, total, err := rds.GetUserDB().GetByConditions(params)
	if err != nil {
		return nil, err
	}

	// 格式化用户信息
	userInfos := make([]*models.UserInfo, 0, len(users))
	for _, v := range users {
		info := new(models.UserInfo)
		err = copyStruct(v, info)
		if err != nil {
			zap.L().Error("CopyStruct failed", zap.Error(err))
			continue
		}
		userInfos = append(userInfos, info)
	}

	// 返回数据
	data = &models.DataUserInfoList{
		Users:  userInfos,
		Total:  total,
		Limit:  params.Limit.Int64,
		Offset: params.Offset.Int64,
	}

	return data, nil
}

// GetUserByID 通过ID获取用户
func GetUserByID(id int64) (data interface{}, err error) {
	// 查询用户
	user, err := rds.GetUserDB().GetByID(id)
	if err != nil {
		return nil, err
	}

	// 格式化用户信息
	info := new(models.UserInfo)
	err = copyStruct(user, info)
	if err != nil {
		zap.L().Error("CopyStruct failed", zap.Error(err))
		return nil, err
	}

	// 返回数据
	return info, nil
}

// UpdateUserByID 通过ID更新用户
func UpdateUserByID(id int64, params *models.BodyUpdateUserNull) (data interface{}, err error) {
	// 判断用户是否存在，不存在则返回
	if err = rds.GetUserDB().CheckExistByID(id); err != nil {
		if err != rds.ErrorUserExist {
			return nil, err
		}
	}

	// 更新用户
	values := make(map[string]interface{})
	if params.Password.Valid {
		// 对密码进行加密
		encodePWD, err := rds.GetUserDB().EncryptPassword(params.Password.String)
		if err != nil {
			return nil, err
		}
		values["password"] = encodePWD
	}
	if params.Nickname.Valid {
		values["nickname"] = params.Nickname.String
	}
	if params.Gender.Valid {
		values["gender"] = params.Gender.Int64
	}

	err = rds.GetUserDB().UpdateByMap(id, values)
	if err != nil {
		return nil, err
	}

	// 查询用户
	user, err := rds.GetUserDB().GetByID(id)
	if err != nil {
		return nil, err
	}

	// 格式化用户信息
	userInfo := new(models.UserInfo)
	err = copyStruct(user, userInfo)
	if err != nil {
		zap.L().Error("CopyStruct failed", zap.Error(err))
		return nil, err
	}

	// 返回数据
	return userInfo, nil
}

// DeleteUserByID 通过ID删除用户
func DeleteUserByID(id int64) (err error) {
	// 判断用户是否存在，不存在则返回
	if err = rds.GetUserDB().CheckExistByID(id); err != nil {
		if err != rds.ErrorUserExist {
			return err
		}
	}

	// 删除用户
	err = rds.GetUserDB().DeleteByID(id)
	if err != nil {
		return err
	}

	// 返回数据
	return nil
}
