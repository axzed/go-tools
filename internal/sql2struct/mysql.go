package sql2struct

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 数据模型信息
type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}

// 数据库基本信息
type DBInfo struct {
	DBType   string // 数据库类型
	Host     string // host
	UserName string // 使用者
	Password string // 密码
	Charset  string // 字符集
}

// 数据表中列的基本信息
type TableColumn struct {
	ColumnName    string // 列名
	DataType      string // 数据类型,只包含数据类型
	IsNullable    string // 是否为空
	ColumnKey     string // 数据的键设定
	ColumnType    string // 列的数据类型,包含其它信息(精度、长度、是否无符号)
	ColumnComment string // 列的注释信息
}

// 声明初始化方法 NewDBModel 和三个核心的结构体对象: DBModel 是整个数据库连接的核心对象
// DBInfo 用于存储mysql的一些基本信息
// TableColumn 存储列的信息
func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{DBInfo: info}
}

// 连接数据库
func (m *DBModel) Connect() error {
	var err error
	s := "%s:%s@tcp(%s)/information_schema?" + "charset=%s&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		s,
		m.DBInfo.UserName,
		m.DBInfo.Password,
		m.DBInfo.Host,
		m.DBInfo.Charset,
	)
	// sql.Open: 1.驱动名称(如mysql) 2.数据库连接信息dsn
	// 注: 必须导入 github.com/go-sql-driver/mysql 进行mysql驱动程序的初始化,不然会报错
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	if err != nil {
		return err
	}
	return nil
}

// 获取表中的列的信息
func (m *DBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	query := "SELECT COLUMN_NAME, DATA_TYPE, COLUMN_KEY, " +
		"IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT " +
		"FROM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?"
	rows, err := m.DBEngine.Query(query, dbName, tableName)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("没有数据")
	}
	defer rows.Close()

	var columns []*TableColumn
	for rows.Next() {
		var column TableColumn
		err := rows.Scan(&column.ColumnName, &column.DataType, &column.ColumnKey, &column.IsNullable, &column.ColumnType, &column.ColumnComment)
		if err != nil {
			return nil, err
		}
		columns = append(columns, &column)
	}
	return columns, nil
}

// 表字段映射(go 与 mysql 中字段不一致)
var DBTypeToStructType = map[string]string{
	"int":        "int32",
	"tinyint":    "int8",
	"smallint":   "int",
	"mediumint":  "int64",
	"bigint":     "int64",
	"bit":        "int",
	"bool":       "bool",
	"enum":       "string",
	"set":        "stirng",
	"varchar":    "string",
	"char":       "string",
	"tinytext":   "string",
	"mediumtext": "string",
	"text":       "string",
	"longtext":   "string",
	"blob":       "string",
	"tinyblob":   "string",
	"mediumblob": "string",
	"longblob":   "string",
	"date":       "time.Time",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
	"time":       "time.Time",
	"float":      "float64",
	"double":     "float64",
}
