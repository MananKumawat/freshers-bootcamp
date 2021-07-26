package Models

type Student struct {

	Id      	uint   `json:"id" gorm:"primary_key"`
	Name    	string `json:"name"`
	LastName	string `json:"lastname"`
	DOB   		string `json:"dob"`
	Address 	string `json:"address"`
	Subjects    []Subject
}

type Subject struct {

	StudentId 			uint
	Id   				uint   `gorm:"primary_key"`
	Name 				string `json:"name"`
	Marks 				uint   `json:"marks"`
}

func (b *Student) TableName() string {
	return "student"
}