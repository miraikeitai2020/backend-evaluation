package database

import(
	"fmt"
	// import gorm library
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/miraikeitai2020/backend-evaluation/pkg/config"
)

const(
	DB_DRIVER_MYSQL = "mysql"

	EMOTION_NUM_1 = "happiness"
	EMOTION_NUM_2 = ""
	EMOTION_NUM_3 = "exciting"
	EMOTION_NUM_4 = "sadness"
	// SQL Query Format
	QUERY_FORMAT_UPDATE_TEMPLATE = "UPDATE `emotions` SET %s = %s %s0.001 WHERE id = ?"
)

type DB struct {
	DB	*gorm.DB
}

func Init() (*DB, error) {
	token := config.GetConnectionToken()
	for i:=0; i<5; i++ {
		if db, err := gorm.Open(DB_DRIVER_MYSQL, token);err == nil {
			return &DB{
				DB: db,
			}, nil
		}
		config.BackOff(i)
	}
	return nil, fmt.Errorf("Faild connection database")
}

func (db *DB) UpdateSpotEmotionValue(id string, emotion int, stat bool) {
	query := makeQuery(emotion, stat)
	db.DB.Exec(query, id)
}

func makeQuery(emotion int, stat bool) string {
	column := setColumn(emotion)
	symbol := "+"
	if !stat {
		symbol = "-"
	}
	return fmt.Sprintf(QUERY_FORMAT_UPDATE_TEMPLATE, column, column, symbol)
}

func setColumn(emotion int) string {
	var column string
	switch emotion {
	case 1:
		column = EMOTION_NUM_1
	case 2:
		column = EMOTION_NUM_2
	case 3:
		column = EMOTION_NUM_3
	case 4:
		column = EMOTION_NUM_4
	}
	return column
}
