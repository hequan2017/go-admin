package restful

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var (
	// dbs 全局的数据库操作句柄
	dbs *sql.DB
)

// SQLInit 初始化数据库操作句柄，这里要提供:
// driverName string: 数据库类型，例如mysql、sqlite等，参考github.com/go-sql-driver/mysql官方介绍
// dataSourceName string: 数据库地址，参考github.com/go-sql-driver/mysql官方介绍
// MaxOpenConns int: 最大缓存连接数，这个数值包含了MaxIdleConns
// MaxIdleConns int：预备的最大空闲连接数
func SQLInit(driverName, dataSourceName string, maxOpenConns, maxIdleConns int) error {
	if dbs == nil {
		var err error
		if dbs, err = sql.Open(driverName, dataSourceName); err != nil {
			return err
		}

		dbs.SetMaxOpenConns(maxOpenConns)
		dbs.SetMaxIdleConns(maxIdleConns)
	}
	return nil
}

func sqlCheckParam(param string) error {
	/*if strings.Contains(param, "where") {
		return errors.New("can not have where")
	}

	if strings.Contains(param, "and") {
		return errors.New("can not have and")
	}

	if strings.Contains(param, "or") {
		return errors.New("can not have or")
	}

	if strings.Contains(param, "=") {
		return errors.New("can not have =")
	}*/

	if strings.Contains(param, ";") {
		return errors.New("can not have ;")
	}

	return nil
}

// SQLInsert 增加一条数据
// tableName string: 操作的表名
// data []byte: 需要更新的内容，用string转换后是json格式
func SQLInsert(tableName string, data []byte) (int64, error) {
	if err := sqlCheckParam(tableName); err != nil {
		return 0, err
	}

	var f []map[string]interface{}
	err := json.Unmarshal(data, &f)
	if err != nil {
		return 0, err
	}

	var res sql.Result

	for _, data := range f {
		var sqlset string

		for k, v := range data {
			if sqlset != "" {
				sqlset += ","
			}

			switch vv := v.(type) {
			case string:
				sqlset += k + "='" + vv + "'"
			case int:
				sqlset += k + "=" + strconv.Itoa(vv)
			case float64:
				sqlset += k + "=" + strconv.FormatFloat(vv, 'f', -1, 64)
			default:
				fmt.Println(k, "is of a type I don't know how to handle")
			}
		}

		stmt, err := dbs.Prepare("INSERT " + tableName + " set " + sqlset)
		if err != nil {
			return 0, err
		}

		res, err = stmt.Exec()
		if err != nil {
			return 0, err
		}
	}

	return res.LastInsertId()
}

// SQLUpdate 更新一条数据
// tableName string: 操作的表名
// where string: 过滤条件，就是where后面跟着的部分
// data []byte: 需要更新的内容，用string转换后是json格式
func SQLUpdate(tableName, where string, data []byte) (int64, error) {
	if err := sqlCheckParam(tableName + where); err != nil {
		return 0, err
	}

	var f map[string]interface{}
	err := json.Unmarshal(data, &f)

	var sqlset string

	for k, v := range f {
		if sqlset != "" {
			sqlset += ","
		}

		switch vv := v.(type) {
		case string:
			sqlset += k + "='" + vv + "'"
		case int:
			sqlset += k + "=" + strconv.Itoa(vv)
		case float64:
			sqlset += k + "=" + strconv.FormatFloat(vv, 'f', -1, 64)
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

	stmt, err := dbs.Prepare("UPDATE " + tableName + " set " + sqlset + " where " + where)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// SQLDelete 根据where条件删除数据
// tableName string: 操作的表名
// where string: 过滤条件，就是where后面跟着的部分
func SQLDelete(tableName, where string) (int64, error) {
	if err := sqlCheckParam(tableName + where); err != nil {
		return 0, err
	}

	if dbs == nil {
		return 0, errors.New("gogo sql not init")
	}

	//删除数据
	stmt, err := dbs.Prepare("DELETE from " + tableName + " where " + where)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// sqlQueryTable 从数据库中查询到的数据，这里是以数组方式存储的，需要做二次转换
func sqlQueryTable(feilds, tableName, where, order string, offset, count int) ([]*sql.ColumnType, int, [][]interface{}, int, error) {
	if dbs == nil {
		return nil, 0, nil, 0, errors.New("gogo sql not init")
	}

	if feilds == "" {
		feilds = "*"
	}

	sqlstr := "select " + feilds + " from " + tableName
	if where != "" {
		sqlstr += " where " + where
	}
	if order != "" {
		sqlstr += " order by "
		if strings.HasPrefix(order, "-") {
			sqlstr += string([]byte(order)[1:]) + " desc"
		} else {
			sqlstr += order + " asc"
		}
	}
	if offset < 0 {
		offset = 0
	}
	if count <= 0 {
		count = 20
	}
	sqlstr += " limit " + strconv.Itoa(offset) + "," + strconv.Itoa(count)

	rows, err := dbs.Query(sqlstr)
	if err != nil {
		return nil, 0, nil, 0, err
	}

	columnsType, _ := rows.ColumnTypes()
	columnsLen := len(columnsType)

	queryData := make([][]interface{}, count)
	queryCount := 0

	for rows.Next() {
		queryData[queryCount] = make([]interface{}, columnsLen)
		for a := 0; a < columnsLen; a++ {
			switch columnsType[a].DatabaseTypeName() {
			case "TINYINT":
				{
					queryData[queryCount][a] = new(int8)
				}
			case "SMALLINT":
				{
					queryData[queryCount][a] = new(int16)
				}
			case "MEDIUMINT":
				{
					queryData[queryCount][a] = new(int32)
				}
			case "INT":
				{
					queryData[queryCount][a] = new(sql.NullInt64)
				}
			case "INTEGER":
				{
					queryData[queryCount][a] = new(int32)
				}
			case "BIGINT":
				{
					queryData[queryCount][a] = new(int64)
				}
			case "FLOAT":
				{
					queryData[queryCount][a] = new(float32)
				}
			case "DOUBLE":
				{
					queryData[queryCount][a] = new(float64)
				}
			default:
				{
					queryData[queryCount][a] = new(sql.NullString)
				}
			}
		}

		if err = rows.Scan(queryData[queryCount]...); err != nil {
			return nil, 0, nil, 0, err
		}

		queryCount = queryCount + 1
	}

	return columnsType, columnsLen, queryData, queryCount, nil
}

// sqlGetValues 根据结构体中指向实际数据的指针获取出数据，并存储到另一张表中返回
func sqlGetValues(pvs []interface{}, columnsType []*sql.ColumnType, columnsLen int) map[string]interface{} {

	result := make(map[string]interface{}, columnsLen)

	for a := 0; a < columnsLen; a++ {
		switch s := pvs[a].(type) {
		case *int8:
			result[columnsType[a].Name()] = *s
		case *int16:
			result[columnsType[a].Name()] = *s
		case *int32:
			result[columnsType[a].Name()] = *s
		case *int64:
			result[columnsType[a].Name()] = *s
		case *float32:
			result[columnsType[a].Name()] = *s
		case *float64:
			result[columnsType[a].Name()] = *s
		case *string:
			result[columnsType[a].Name()] = *s
		case *sql.NullInt64:
			result[columnsType[a].Name()] = *s
		case *sql.NullString:
			result[columnsType[a].Name()] = *s
		}
	}
	return result
}

// 这里返回的是原始数组的基础上加上了字段名标识
func sqlQuery(columnsType []*sql.ColumnType, columnsLen int, queryData [][]interface{}, queryCount int) ([]map[string]interface{}, error) {

	jsondata := make([]map[string]interface{}, queryCount)
	for k1, v1 := range queryData {
		if k1 >= queryCount {
			break
		}

		jsondata[k1] = sqlGetValues(v1, columnsType, columnsLen)
	}

	return jsondata, nil
}

func sqlQueryByTinyIntMap(columnName string, columnsType []*sql.ColumnType, columnsLen int, queryData [][]interface{}, queryCount int) (map[int8]map[string]interface{}, error) {

	jsondata := make(map[int8]map[string]interface{}, queryCount)
	for k1, v1 := range queryData {
		if k1 >= queryCount {
			break
		}

		for a := 0; a < columnsLen; a++ {
			if columnsType[a].Name() == columnName {
				if value, ok := v1[a].(*int8); ok {
					jsondata[*value] = sqlGetValues(v1, columnsType, columnsLen)
				}
				break
			}
		}
	}

	return jsondata, nil
}

func sqlQueryBySmallIntMap(columnName string, columnsType []*sql.ColumnType, columnsLen int, queryData [][]interface{}, queryCount int) (map[int16]map[string]interface{}, error) {

	jsondata := make(map[int16]map[string]interface{}, queryCount)
	for k1, v1 := range queryData {
		if k1 >= queryCount {
			break
		}

		for a := 0; a < columnsLen; a++ {
			if columnsType[a].Name() == columnName {
				if value, ok := v1[a].(*int16); ok {
					jsondata[*value] = sqlGetValues(v1, columnsType, columnsLen)
				}
				break
			}
		}
	}

	return jsondata, nil
}

func sqlQueryByIntMap(columnName string, columnsType []*sql.ColumnType, columnsLen int, queryData [][]interface{}, queryCount int) (map[int32]map[string]interface{}, error) {

	jsondata := make(map[int32]map[string]interface{}, queryCount)
	for k1, v1 := range queryData {
		if k1 >= queryCount {
			break
		}

		for a := 0; a < columnsLen; a++ {
			if columnsType[a].Name() == columnName {
				if value, ok := v1[a].(*int32); ok {
					jsondata[*value] = sqlGetValues(v1, columnsType, columnsLen)
				}
				break
			}
		}
	}

	return jsondata, nil
}

func sqlQueryByBigIntMap(columnName string, columnsType []*sql.ColumnType, columnsLen int, queryData [][]interface{}, queryCount int) (map[int64]map[string]interface{}, error) {

	jsondata := make(map[int64]map[string]interface{}, queryCount)
	for k1, v1 := range queryData {
		if k1 >= queryCount {
			break
		}

		for a := 0; a < columnsLen; a++ {
			if columnsType[a].Name() == columnName {
				if value, ok := v1[a].(*int64); ok {
					jsondata[*value] = sqlGetValues(v1, columnsType, columnsLen)
				}
				break
			}
		}
	}

	return jsondata, nil
}

func sqlQueryByFloatIntMap(columnName string, columnsType []*sql.ColumnType, columnsLen int, queryData [][]interface{}, queryCount int) (map[float32]map[string]interface{}, error) {

	jsondata := make(map[float32]map[string]interface{}, queryCount)
	for k1, v1 := range queryData {
		if k1 >= queryCount {
			break
		}

		for a := 0; a < columnsLen; a++ {
			if columnsType[a].Name() == columnName {
				if value, ok := v1[a].(*float32); ok {
					jsondata[*value] = sqlGetValues(v1, columnsType, columnsLen)
				}
				break
			}
		}
	}

	return jsondata, nil
}

func sqlQueryByDoubleMap(columnName string, columnsType []*sql.ColumnType, columnsLen int, queryData [][]interface{}, queryCount int) (map[float64]map[string]interface{}, error) {

	jsondata := make(map[float64]map[string]interface{}, queryCount)
	for k1, v1 := range queryData {
		if k1 >= queryCount {
			break
		}

		for a := 0; a < columnsLen; a++ {
			if columnsType[a].Name() == columnName {
				if value, ok := v1[a].(*float64); ok {
					jsondata[*value] = sqlGetValues(v1, columnsType, columnsLen)
				}
				break
			}
		}
	}

	return jsondata, nil
}

func sqlQueryByStringMap(columnName string, columnsType []*sql.ColumnType, columnsLen int, queryData [][]interface{}, queryCount int) (map[string]map[string]interface{}, error) {

	jsondata := make(map[string]map[string]interface{}, queryCount)
	for k1, v1 := range queryData {
		if k1 >= queryCount {
			break
		}

		for a := 0; a < columnsLen; a++ {
			if columnsType[a].Name() == columnName {
				if value, ok := v1[a].(*string); ok {
					jsondata[*value] = sqlGetValues(v1, columnsType, columnsLen)
				}
				break
			}
		}
	}

	return jsondata, nil
}

func sqlGetColumnType(columnsType []*sql.ColumnType, columnsLen int, valueName string) string {

	for a := 0; a < columnsLen; a++ {
		if columnsType[a].Name() == valueName {
			return columnsType[a].DatabaseTypeName()
		}
	}

	return ""
}

// SQLQueryByMap 将查询到的数据，按照指定字段的值做为索引构建map并返回
// columnName string: 作为索引的字段名称
// feilds string: 查询需要获取哪些字段的值，就是select后面跟着的部分，一般用"*"
// tableName string: 查询的表名
// where string: 过滤条件，就是where后面跟着的部分
// order string: 排序条件，就是order by后面跟着的部分。默认是ASC排序，除非"-"开头则DESC排序
// offset string: limit后面逗号相隔的两个数值，前者就是offset，后者就是count
// count string: limit后面逗号相隔的两个数值，前者就是offset，后者就是count
func SQLQueryByMap(columnName, feilds, tableName, where, order string, offset, count int) (interface{}, error) {
	if err := sqlCheckParam(columnName + feilds + tableName + where + order); err != nil {
		return 0, err
	}

	columnsType, columnsLen, queryData, queryCount, err := sqlQueryTable(feilds, tableName, where, order, offset, count)
	if err != nil {
		return nil, err
	}

	if queryCount == 0 {
		return "", errors.New("0")
	}

	if columnName == "" {
		return sqlQuery(columnsType, columnsLen, queryData, queryCount)
	}

	switch sqlGetColumnType(columnsType, columnsLen, columnName) {
	case "TINYINT":
		return sqlQueryByTinyIntMap(columnName, columnsType, columnsLen, queryData, queryCount)
	case "SMALLINT":
		return sqlQueryBySmallIntMap(columnName, columnsType, columnsLen, queryData, queryCount)
	case "MEDIUMINT":
		return sqlQueryByIntMap(columnName, columnsType, columnsLen, queryData, queryCount)
	case "INT":
		return sqlQueryByIntMap(columnName, columnsType, columnsLen, queryData, queryCount)
	case "INTEGER":
		return sqlQueryByIntMap(columnName, columnsType, columnsLen, queryData, queryCount)
	case "BIGINT":
		return sqlQueryByBigIntMap(columnName, columnsType, columnsLen, queryData, queryCount)
	case "FLOAT":
		return sqlQueryByFloatIntMap(columnName, columnsType, columnsLen, queryData, queryCount)
	case "DOUBLE":
		return sqlQueryByDoubleMap(columnName, columnsType, columnsLen, queryData, queryCount)
	}

	return sqlQueryByStringMap(columnName, columnsType, columnsLen, queryData, queryCount)
}
