package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	_ "github.com/lib/pq"
)

func main() {
	// 接收参数
	outputFile := flag.String("output", "schema.sql", "Path to output file")
	dbHost := flag.String("host", "localhost", "PostgreSQL host")
	dbPort := flag.String("port", "5432", "PostgreSQL port")
	dbUser := flag.String("user", "admin", "PostgreSQL user")
	dbPassword := flag.String("password", "admin123", "PostgreSQL password")
	dbName := flag.String("dbname", "postgres", "PostgreSQL database name")
	flag.Parse()

	// 构建 PostgreSQL 连接字符串
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", *dbHost, *dbPort, *dbUser, *dbPassword, *dbName)
	log.Println(connStr)
	// 打开数据库连接
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// 获取所有表名
	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'")
	if err != nil {
		log.Fatalf("Failed to query table names: %v", err)
	}
	defer rows.Close()

	log.Println("outputFile: ", *outputFile)

	// 打开输出文件
	file, err := os.Create(*outputFile)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer file.Close()

	// 使用 pg_dump 获取每个表的创建语句
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			log.Fatalf("Failed to scan table name: %v", err)
		}

		// 使用 pg_dump 提取表定义
		cmd := exec.Command("pg_dump", "--schema-only", "--table", tableName, *dbName)
		cmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", *dbPassword))
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("Failed to dump table schema for %s: %v\nOutput: %s", tableName, err, string(output))
		}

		// 写入文件
		_, err = file.WriteString(fmt.Sprintf("%s\n", string(output)))
		if err != nil {
			log.Fatalf("Failed to write to file: %v", err)
		}
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating rows: %v", err)
	}

	fmt.Printf("Schema exported to %s\n", *outputFile)
}
