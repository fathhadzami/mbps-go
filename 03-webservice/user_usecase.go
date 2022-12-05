package user

import (
	"database/sql"
	"webservice/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10/translations/id"
)

type UserUsecase struct {
	db *sql.DB
}

func NewUserUsecase(db *sql.DB) *UserUsecase {
	return &UserUsecase{db: db}
}

func (uu UserUsecase) Register(c *gin.Context) {
	var user domain.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid input",
		})
		return
	}

	if user.Email == "" {
		c.JSON(400, gin.H{
			"message": "Email must be required",
		})
		return
	}

	if user.Password == "" {
		c.JSON(400, gin.H{
			"message": "Password must be required",
		})
		return
	}

	if user.Name == "" {
		c.JSON(400, gin.H{
			"message": "Name must be required",
		})
		return
	}

	if user.Department_Id == 0 {
		c.JSON(400, gin.H{
			"message": "Department_Id must be required",
		})
		return
	}

	// check department is exist
	query := `SELECT * FROM department where id = ?;`
	row := uu.db.QueryRow(query, user.Department_Id)

	var dpt domain.Department

	if err := row.Scan(&dpt.ID, &dpt.Name); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(400, gin.H{
				"message": "Department_Id not found",
			})
			return
		}
	}

	// insert into user
	query := `INSERT INTO user(ID,Name) VALUES('ID','Name')`
	row := uu.db.QueryRow(query, user.Department_Id)

	var dpt domain.Department

	err := db.Exec(&dpt.ID, &dpt.Name); if err != nil {
			return 0, 0, fmt.errorf("domain.Department: %v", err)			
		}

		}

	// buat token

	// buat response

	c.JSON(200, gin.H{
		"department": dpt,
	})
}
