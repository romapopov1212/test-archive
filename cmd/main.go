package main

import (
	"encoding/json"
	"fmt"
	"justTest/internal"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: pm [create|update] <file>")
		return
	}
	
	command := os.Args[1]
	file := os.Args[2]
	
	switch command {
	case "create":
		doCreate(file)
	case "update":
		doUpdate(file)
	default:
		fmt.Println("Unknown command")
	}
}

func doCreate(configPath string) {
	var pf internal.PacketFile
	
	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	
	if err := json.Unmarshal(data, &pf); err != nil {
		fmt.Println("Invalid JSON:", err)
		return
	}
	
	// [!] Для простоты: используем только строковые пути
	var files []string
	for _, tgt := range pf.Targets {
		switch v := tgt.(type) {
		case string:
			files = append(files, v)
		case map[string]interface{}:
			if path, ok := v["path"].(string); ok {
				files = append(files, path)
			}
		}
	}
	
	archive := fmt.Sprintf("%s-%s.tar.gz", pf.Name, pf.Version)
	fmt.Println("Создаём архив:", archive)
	err = internal.CreateArchive(files, archive)
	if err != nil {
		fmt.Println("Ошибка архивации:", err)
		return
	}
	
	// пример — отправка на сервер (захардкожено для теста)
	err = internal.UploadFile("server:22", "user", "password", archive, "/remote/path/"+archive)
	if err != nil {
		fmt.Println("Ошибка загрузки по SSH:", err)
		return
	}
	fmt.Println("Архив успешно отправлен")
}

func doUpdate(packagesPath string) {
	// TODO: реализовать чтение packages.json, сравнение версий, загрузку и распаковку
	fmt.Println("Функция update пока не реализована")
}
