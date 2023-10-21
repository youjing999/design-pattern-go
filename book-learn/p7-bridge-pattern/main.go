package main

import (
	"fmt"
	"math/rand"
)

type IDataFetcher interface {
	Fetch(sql string) []interface{}
}

// IDataExport abstraction
type IDataExport interface {
	Fetcher(fetcher IDataFetcher)
	Export(sql string) error
}

type MysqlFetcher struct {
	config string
}

func NewMysqlFetcher(config string) *MysqlFetcher {
	return &MysqlFetcher{
		config: config,
	}
}

func (mysql *MysqlFetcher) Fetch(sql string) []interface{} {
	fmt.Println("Fetch data from mysql source:" + mysql.config)
	data := make([]interface{}, 0)
	data = append(data, rand.Perm(10), rand.Perm(20))
	return data
}

type OracleFetcher struct {
	config string
}

func NewOracleFetcher(config string) *MysqlFetcher {
	return &MysqlFetcher{
		config: config,
	}
}

func (oracle *OracleFetcher) Fetcher(sql string) []interface{} {
	fmt.Println("Fetch data from mysql source:" + oracle.config)
	data := make([]interface{}, 0)
	data = append(data, rand.Perm(10), rand.Perm(20))
	return data
}

type CsvExporter struct {
	mFetcher IDataFetcher
}

func (ce *CsvExporter) Fetcher(fetcher IDataFetcher) {
	ce.mFetcher = fetcher
}

func (ce *CsvExporter) Export(sql string) error {
	rows := ce.mFetcher.Fetch(sql)
	fmt.Printf("CsvExporter.Export, got %v rows\n", len(rows))
	for i, v := range rows {
		fmt.Printf("  行号: %d 值: %s\n", i+1, v)
	}
	return nil
}

func NewCsvExporter(fetcher IDataFetcher) IDataExport {
	return &CsvExporter{
		mFetcher: fetcher,
	}
}

type JsonExporter struct {
	mFetcher IDataFetcher
}

func (ce *JsonExporter) Fetcher(fetcher IDataFetcher) {
	ce.mFetcher = fetcher
}

func (ce *JsonExporter) Export(sql string) error {
	rows := ce.mFetcher.Fetch(sql)
	fmt.Printf("Json.Export, got %v rows\n", len(rows))
	for i, v := range rows {
		fmt.Printf("  行号: %d 值: %s\n", i+1, v)
	}
	return nil
}
func NewJsonExporter(fetcher IDataFetcher) IDataExport {
	return &JsonExporter{
		mFetcher: fetcher,
	}
}

func main() {
	mFetcher := NewMysqlFetcher("mysql://127.0.0.1:3306")
	csvExporter := NewCsvExporter(mFetcher)
	err := csvExporter.Export("select * from xzq")
	if err != nil {
		fmt.Println("导出错误")
	}

	fmt.Printf("\n")
	fetcher := NewOracleFetcher("oracle://192.168.1.1")
	jsonExport := NewJsonExporter(fetcher)
	err = jsonExport.Export("select * from yj")
	if err != nil {
		fmt.Println("导出错误")
	}
}
