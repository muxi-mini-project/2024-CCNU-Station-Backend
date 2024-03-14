package model

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"guizizhan/pkg/CCNU"
	"time"
)

// Student
type Student struct {
	StuID          string    `gorm:"primaryKey" json:"stuid"`           //学号
	Password       string    `json:"password"`                          //密码
	RealName       string    `json:"realName"`                          //真实姓名
	Nickname       string    `json:"nickname" gorm:"default:华师学子"`      //昵称
	Grade          string    `json:"grade"`                             //年级
	College        string    `json:"college"`                           //学院
	Gender         string    `json:"gender"`                            //性别
	HeadImage      string    `json:"headimage"`                         //头像
	Age            int       `json:"age" gorm:"default:18"`             //年龄
	Sign           string    `json:"sign" gorm:"default:study hard"`    //签名
	FriendsNumber  int       `json:"friends_number" gorm:"default:0"`   //好友数
	FanNumber      int       `json:"followers_number" gorm:"default:0"` //粉丝数
	FollowerNumber int       `json:"follower_number" gorm:"default:0"`  //关注数
	PostNumber     int       `json:"post_number" gorm:"default:0"`      //发布数
	SchoolDate     time.Time `json:"date"`                              //进入校园的日期
	StayDate       int       `json:"stay_date" gorm:"default:1"`        //来华天数
	MBTI           string    `json:"mbti" gorm:"default:empty"`         //mbti
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func CreateStudent(ID string, password string, db *gorm.DB) {
	var gender, college, realname, grade string

	//信息处理和获取
	//password = md5.MakePassword(password)
	grade = Getgrade(ID)
	info := CCNU.GetUserNameAndCollegeAndGender(CCNU.CCNULogin(ID, password))
	gender = info["Gender"]
	college = info["College"]
	realname = info["Name"]

	var s = Student{
		StuID:      ID,
		Password:   password,
		Grade:      grade,
		Gender:     gender,
		College:    college,
		RealName:   realname,
		SchoolDate: time.Now(),
	}

	db.Create(&s)
}

func (s Student) AddStudent(ID string, password string, db *gorm.DB) {
	fmt.Println("starting add student")
	//创建成就
	CreateAchievement(ID, db)
	//创建学生
	CreateStudent(ID, password, db)
	fmt.Println("创建成功")
}

func (student Student) Updateinformation(nickname, sign, mbti string, age, daysDifference int, date time.Time, db *gorm.DB) {
	student.Nickname = nickname
	student.Age = age
	student.SchoolDate = date
	student.StayDate = -daysDifference
	student.Sign = sign
	student.MBTI = mbti
	db.Save(&student)
}

func Updateheadimage(db *gorm.DB, stuid string, url string) {
	db.Model(&Student{}).Where(&Student{StuID: stuid}).Updates(&Student{HeadImage: url})
}

func FindStudfromID(id string, db *gorm.DB) (Student, bool) {
	fmt.Println("Finding")
	var student Student
	result := db.Model(&Student{}).Where(&Student{StuID: id}).First(&student) //再根据学生的ID来查找对应的学生
	fmt.Println("ending")
	if errors.Is(result.Error, gorm.ErrRecordNotFound) { //检查一下是否有这个学生
		return student, false //如果找不到就返回false
	} else if result.Error != nil {
		// 处理其他数据库错误，如果需要的话
		// 例如，记录错误日志或返回错误响应
		return student, false
	} else {
		return student, true //找到了就返回true}
	}

}

func UpdateDate(db *gorm.DB, stuid string) {
	// Convert date to time.Time
	var stu Student
	var date time.Time

	db.Model(&Student{}).Where(&Student{StuID: stuid}).First(&stu)
	date = stu.SchoolDate
	today := time.Now()

	// 计算天数差
	daysDifference := int(date.Sub(today).Hours() / 24)

	stu.StayDate = -daysDifference

	db.Save(&stu)

}

func Getgrade(stuid string) string {
	grade := stuid[:4]
	return grade
}
