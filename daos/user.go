package daos

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"user_microservices/databases"
	"user_microservices/helper"
	"user_microservices/middleware"
	"user_microservices/models"
)

type User struct {
	helper helper.Helper
}

func (m *User) UserCreate(params models.CreateUser) (models.UserCreate, error) {

	user := models.UserCreate{}

	user.IdUser = m.helper.StringWithCharset()
	user.Nama = params.Nama
	user.NoTelp = params.NoTelp
	user.Email = params.Email
	user.Password,_ = EncryptPassword(params.Password)
	user.Role = params.Role
	user.CreatedAt = m.helper.GetTimeNow()

	err := databases.DatabaseBinus.DB.Table("user").Create(&user).Error

	if err != nil {
		return models.UserCreate{}, err
	}

	return user, nil
}

func (m *User) UserGet(params models.GetUser) ([]models.UserGet, error) {

	user := []models.UserGet{}

	err := databases.DatabaseBinus.DB.Table("user")
	if params.IdUser != "" {
		err = err.Where("id_user = ?", params.IdUser)
	}
	if params.Email != "" {
		err = err.Where("email = ?", params.Email)
	}

	if params.Search != "" {
		err = err.Where("u.nama ilike '%" + params.Search + "%' OR u.email ilike '%" + params.Search + "%' OR u.no_telp ilike '%" + params.Search + "%'")
	}

	if params.Limit != "" {
		limits, _ := strconv.Atoi(params.Limit)
		err = err.Limit(limits)
	}

	if params.Offset != "" {
		offsets, _ := strconv.Atoi(params.Offset)
		err = err.Offset(offsets)
	}

	err = err.Find(&user)

	errx := err.Error


	if errx != nil {
		return []models.UserGet{}, errx
	}

	return user, nil
}

func (m *User) UserUpdate(params models.UpdateUser) ([]models.UserGet, error) {

	user := models.UpdateUser{}
	getuser := []models.UserGet{}

	user.UpdatedAt = m.helper.GetTimeNow()
	user.Nama = params.Nama
	user.Email = params.Email
	user.NoTelp = params.NoTelp
	user.Role = params.Role
	if params.Password != "" {
		user.Password,_ = EncryptPassword(params.Password)
	}

	err := databases.DatabaseBinus.DB.Table("user").Where("id_user = ?", params.IdUser).Update(&user).Error

	if err != nil {
		return []models.UserGet{}, err
	}

	paramuser := models.GetUser{}
	paramuser.IdUser = params.IdUser
	getuser,errx := m.UserGet(paramuser)
	if errx != nil {
		return []models.UserGet{}, errx
	}
	return getuser, nil

}

func (m *User) UserDelete(params models.DeleteUser) (models.DeleteUser, error) {

	user := models.DeleteUser{}

	user.DeletedAt = m.helper.GetTimeNow()

	err := databases.DatabaseBinus.DB.Table("user").Where("id_user = ?", params.IdUser).Update(&user).Error

	if err != nil {
		return models.DeleteUser{}, err
	}

	return user, nil

}

func (m *User) LoginCheck(params models.UserToken) error {

	checkakun := models.UserGet{}
	var check bool

	check = databases.DatabaseBinus.DB.Table("user").
		Where("email = ?", params.Email).Find(&checkakun).RecordNotFound()

	if check == true {
		err := errors.New("Email Tidak Ditemukan")
		return err
	}

	return error(nil)

}

func (m *User) Signin(params models.UserToken) ([]models.UserGet, string, error) {

	userGet := models.GetUser{}
	userRead := []models.UserGet{}
	updateToken := models.UpdateUser{}
	var token string
	var er error

	err := m.LoginCheck(params)

	if err != nil {
		return userRead, "", err
	}

	if params.Email != "" {
		userGet.Email = params.Email
	}
	if params.Password != "" {
		userGet.Password = params.Password
	}

	userRead, err = m.UserGet(userGet)

	if err != nil {
		return userRead, "", err
	}

	//if userRead[0].Verifikasi == "N" {
	//	err = errors.New("Akun Anda di Nonaktifkan, Tidak Dapat di Akses")
	//	return userRead, "", err
	//}

	//token, er := m.helper.GetToken(userRead[0].IdUser)
	password,_ := DecryptPassword(userRead[0].Password)

	if userRead[0].Email == params.Email && params.Password == password  {
		fmt.Println("cocok")
		token, er = middleware.CreateAuth(userRead[0].IdUser, "user", "none", "none")

		if er != nil {
			return userRead, "", er
		}

		updateToken.Token = token

		err = databases.DatabaseBinus.DB.Table("user").Where("id_user = ?", userRead[0].IdUser).Update(&updateToken).Error

		if err != nil {
			return userRead, "", err
		}
	}else {
		fmt.Println("email atau password tidak cocok")
		//return userRead, "", err
	}



	return userRead, token, nil

}

func (m *User) FileCreate(params models.CreateFile) (models.FileCreate, error) {

	file := models.FileCreate{}

	var tgl = m.helper.GetTanggalNow()

	//fmt.Println(tgl)

	path := "/file/"
	pathFile := "./files/"+path
	ext := filepath.Ext(params.File.Filename)
	filename := strings.Replace(params.NamaFile," ","_", -1)+tgl+ext

	os.MkdirAll(pathFile, 0777)
	errx := m.helper.SaveUploadedFile(params.File, pathFile+filename)
	if errx != nil{
		return models.FileCreate{},errx
	}

	url := string(filepath.FromSlash(path+filename))

	file.IdFile = m.helper.StringWithCharset()
	file.NamaFile = params.NamaFile
	file.File = url
	file.CreatedAt = m.helper.GetTimeNow()


	err := databases.DatabaseBinus.DB.Table("file").Create(&file).Error

	if err != nil {
		return models.FileCreate{}, err
	}
	//as, err := os.Stat("files/test.py")
	//if os.IsNotExist(err) {
	//	//response.ApiMessage = "File script.py not found"
	//	//response.Data = err
	//	fmt.Println("File script.py not found", as)
	//	return models.FileCreate{}, err
	//}
	//fmt.Println("File script.py found")
	//
	//cmd := exec.Command("python3", "files/ekstrak.py", "file/"+params.NamaFile+tgl+".docx", "text/"+params.NamaFile+tgl+".txt")
	//_, err = cmd.Output()
	//if err != nil {
	//	//response.ApiMessage = "Execute script.py failed"
	//	//response.Data = err
	//	fmt.Println("Execute script.py failed")
	//	return models.FileCreate{}, err
	//}
	//response.ApiMessage = "Success"
	//response.Data = "Execute script.py success"
	//fmt.Println("Execute script.py success")

	return file, nil
}

func (m *User) FileDelete(params models.DeleteFile) (models.DeleteFile, error) {

	file := models.DeleteFile{}

	file.DeletedAt = m.helper.GetTimeNow()

	errx := m.helper.RemoveFile("files/"+params.File)
	if errx != nil{
		return file,errx
	}

	err := databases.DatabaseBinus.DB.Table("file").Where("id_file = ?", params.IdFile).Update(&file).Error

	if err != nil {
		return models.DeleteFile{}, err
	}

	return file, nil

}

func (m *User) FileGet(params models.GetFile) ([]models.FileGet, error) {

	file := []models.FileGet{}

	err := databases.DatabaseBinus.DB.Table("file").Order("created_at desc")

	if params.IdFile != "" {
		err = err.Where("id_file = ?", params.IdFile)
	}
	if params.CreatedAt != "" {
		err = err.Where("created_at::text like  ?", "%"+params.CreatedAt+"%")
	}
	if params.Search != "" {
		err = err.Where("nama_file ilike '%"+params.Search+"%'")
	}
	if params.Limit != nil {
		err = err.Limit(*params.Limit)
	}
	if params.Offset != nil {
		err = err.Offset(*params.Offset)
	}

	err = err.Find(&file)

	errx := err.Error


	if errx != nil {
		return []models.FileGet{}, errx
	}

	return file, nil
}

func (m *User) FileText(params models.GetText) (models.TextGet, error) {

	text := models.TextGet{}
	////dataFile := []
	//
	//paramfile := models.GetFile{}
	//paramfile.IdFile = params.IdFile
	//getfile,errx := m.FileGet(paramfile)
	//if errx != nil {
	//	return models.TextGet{}, errx
	//}
	////dataFile = getfile
	//fmt.Println("data", getfile)


	as, err := os.Stat("files/ekstrak.py")
	if os.IsNotExist(err) {
		//response.ApiMessage = "File script.py not found"
		//response.Data = err
		fmt.Println("File test.py not found", as)
		return models.TextGet{}, err
	}
	fmt.Println("File test.py found")

	cmd, err := exec.Command("python3", "files/ekstrak.py", "files/"+params.File).Output()
	//_, err = cmd.Output()
	if err != nil {
		//response.ApiMessage = "Execute script.py failed"
		//response.Data = err
		fmt.Println("Execute script.py failed")
		return models.TextGet{}, err
	}
	//fmt.Println(cmd)

	text.Text = string(cmd)
	fmt.Println("ini hasilnya :", text.Text)


	return text, nil
}

func (m *User) FileTextlive(params models.GetTextlive) (models.TextGet, error) {

	text := models.TextGet{}

	path := "/file/"
	pathFile := "./files/"+path
	ext := filepath.Ext(params.File.Filename)
	filename := "gettextlive"+ext

	os.MkdirAll(pathFile, 0777)
	errx := m.helper.SaveUploadedFile(params.File, pathFile+filename)
	if errx != nil{
		return models.TextGet{},errx
	}

	as, err := os.Stat("files/ekstrak.py")
	if os.IsNotExist(err) {
		//response.ApiMessage = "File script.py not found"
		//response.Data = err
		fmt.Println("File test.py not found", as)
		return models.TextGet{}, err
	}
	fmt.Println("File test.py found")

	cmd, err := exec.Command("python3", "files/ekstrak.py", "files/file/gettextlive.docx").Output()
	//_, err = cmd.Output()
	if err != nil {
		//response.ApiMessage = "Execute script.py failed"
		//response.Data = err
		fmt.Println("Execute script.py failed")
		return models.TextGet{}, err
	}
	//fmt.Println(cmd)

	text.Text = string(cmd)
	fmt.Println("ini hasilnya :", text.Text)


	return text, nil
}