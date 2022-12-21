package models

import "mime/multipart"

type CreateUser struct {
	Nama        string `json:"nama" form:"nama"`
	NoTelp		string `json:"no_telp" form:"no_telp"`
	Email       string `json:"email" form:"email"`
	Role    	string `json:"role" form:"role"`
	Password    string `json:"password" form:"password"`
	CreatedAt   string `json:"created_at" form:"created_at"`
}

type UserCreate struct {
	IdUser		string `json:"id_user" form:"id_user"`
	Nama        string `json:"nama" form:"nama"`
	NoTelp		string `json:"no_telp" form:"no_telp"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	Role    	string `json:"role" form:"role"`
	CreatedAt   string `json:"created_at" form:"created_at"`
}

type GetUser struct {
	IdUser			string `json:"id_user" form:"id_user"`
	Search       	string `json:"search" form:"search"`
	Email       	string `json:"email" form:"email"`
	Password    	string `json:"password" form:"password"`
	Limit       	string `json:"limit" form:"limit"`
	Offset       	string `json:"offset" form:"offset"`
}

type UserGet struct {
	IdUser       	string `json:"id_user" form:"id_user"`
	Nama         	string `json:"nama" form:"nama"`
	Email      		string `json:"email" form:"email"`
	Password    	string `json:"password" form:"password"`
	NoTelp         	string `json:"no_telp" form:"no_telp"`
	Role    		string `json:"role" form:"role"`
	CreatedAt       string `json:"created_at" form:"created_at"`
	UpdatedAt       string `json:"updated_at" form:"updated_at"`
	DeletedAt       string `json:"deleted_at" form:"deleted_at"`
}

type UpdateUser struct {
	IdUser       	string `json:"id_user" form:"id_user"`
	Token    		string `json:"token" form:"token"`
	Nama         	string `json:"nama" form:"nama"`
	Role    		string `json:"role" form:"role"`
	Email      		string `json:"email" form:"email"`
	NoTelp         	string `json:"no_telp" form:"no_telp"`
	Password    	string `json:"password" form:"password"`
	CreatedAt       string `json:"created_at" form:"created_at"`
	UpdatedAt       string `json:"updated_at" form:"updated_at"`
}

type DeleteUser struct {
	IdUser 		string `json:"id_user" form:"id_user"`
	DeletedAt  	string `json:"deleted_at" form:"deleted_at"`
}

type UserToken struct {
	IdToken     *string `json:"id_token" form:"id_token"`
	FcmToken    string  `json:"fcm_token" form:"fcm_token"`
	IdUser      *string `json:"id_user" form:"id_user" binding:"-"`
	Email       string `json:"email" form:"email" binding:"-"`
	Password    string  `json:"password" form:"password" binding:"-"`
}

type CreateFile struct {
	NamaFile	      	string `json:"nama_file" form:"nama_file"`
	File		    	*multipart.FileHeader `json:"file" form:"file"`
	CreatedAt        	string `json:"created_at" form:"created_at"`
}

type FileCreate struct {
	IdFile				string `json:"id_file" form:"id_file"`
	NamaFile	      	string `json:"nama_file" form:"nama_file"`
	File		    	string `json:"file" form:"file"`
	CreatedAt        	string `json:"created_at" form:"created_at"`
}

type DeleteFile struct {
	IdFile				string `json:"id_file" form:"id_file"`
	File		    	string `json:"file" form:"file"`
	DeletedAt        	string `json:"deleted_at" form:"deleted_at"`
}

type GetFile struct {
	IdFile				string `json:"id_file" form:"id_file"`
	Search       		string `json:"search" form:"search"`
	CreatedAt        	string `json:"created_at" form:"created_at"`
	Limit     			*int64 `json:"limit" form:"limit"`
	Offset    			*int64 `json:"offset" form:"offset"`
}

type FileGet struct {
	IdFile				string `json:"id_file" form:"id_file"`
	NamaFile	      	string `json:"nama_file" form:"nama_file"`
	File		    	string `json:"file" form:"file"`
	CreatedAt        	string `json:"created_at" form:"created_at"`
	DeletedAt        	string `json:"deleted_at" form:"deleted_at"`
}

type GetText struct {
	File		    	string `json:"file" form:"file"`
	IdFile				string `json:"id_file" form:"id_file"`
}

type GetTextlive struct {
	File		    	*multipart.FileHeader `json:"file" form:"file"`
}

type TextGet struct {
	Text		    	string `json:"text" form:"text"`
}