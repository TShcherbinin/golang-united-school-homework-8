package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Arguments map[string]string

type itemInfo struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func findIdxById(items []itemInfo, id string) (int, bool) {
	for idx, v := range items {
		if v.Id == id {
			return idx, true
		}
	}
	return 0, false
}

func Add(item string, filename string, writer io.Writer) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666) // For read access.
	if err != nil {
		return fmt.Errorf("Add. os.OpenFile: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	byteData, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("Add. ioutil.ReadAll: %w", err)
	}

	var items []itemInfo
	if len(byteData) != 0 {
		// Данные в файле уже имеются
		err = json.Unmarshal(byteData, &items)
		if err != nil {
			return fmt.Errorf("Add. File json.Unmarshal: %w", err)
		}
	}

	var newItem itemInfo
	err = json.Unmarshal([]byte(item), &newItem)
	if err != nil {
		return fmt.Errorf("Add. json.Unmarshal: %w", err)
	}

	_, ok := findIdxById(items, newItem.Id)
	if ok {
		writer.Write([]byte(fmt.Sprintf("Item with id %s already exists", newItem.Id)))
		return nil
	}

	items = append(items, newItem)

	bytes, err := json.Marshal(&items)
	if err != nil {
		return fmt.Errorf("Add. json.Marshal: %w", err)
	}

	file.Seek(0, 0)
	if _, err := file.Write(bytes); err != nil {
		return fmt.Errorf("Add. file.Write: %w", err)
	}

	return nil
}

func List(filename string, writer io.Writer) error {
	file, err := os.OpenFile(filename, os.O_RDWR, 0666) // For read access.
	if err != nil {
		return fmt.Errorf("Add. os.OpenFile: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	byteData, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("Add. ioutil.ReadAll: %w", err)
	}

	writer.Write(byteData)
	//fmt.Print(string(byteData))

	return nil
}

func Remove(id string, filename string, writer io.Writer) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666) // For read access.
	if err != nil {
		return fmt.Errorf("Remove. os.OpenFile: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	byteData, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("Remove. ioutil.ReadAll: %w", err)
	}

	var items []itemInfo
	if len(byteData) != 0 {
		// Данные в файле уже имеются
		err = json.Unmarshal(byteData, &items)
		if err != nil {
			return fmt.Errorf("Remove. File json.Unmarshal: %w", err)
		}
	}

	foundIdx, ok := findIdxById(items, id)
	if !ok {
		writer.Write([]byte(fmt.Sprintf("Item with id %s not found", id)))
		return nil
	}

	items[foundIdx] = items[len(items)-1]
	items = items[:len(items)-1]

	bytes, err := json.Marshal(&items)
	if err != nil {
		return fmt.Errorf("Remove. json.Marshal: %w", err)
	}

	if err := ioutil.WriteFile(filename, bytes, 0644); err != nil {
		return fmt.Errorf("Remove. ioutil.WriteFile: %w", err)
	}

	return nil
}

func findById(id string, filename string, writer io.Writer) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666) // For read access.
	if err != nil {
		return fmt.Errorf("findById. os.OpenFile: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	byteData, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("findById. ioutil.ReadAll: %w", err)
	}

	var items []itemInfo
	if len(byteData) != 0 {
		// Данные в файле уже имеются
		err = json.Unmarshal(byteData, &items)
		if err != nil {
			return fmt.Errorf("findById. File json.Unmarshal: %w", err)
		}
	}

	foundIdx, ok := findIdxById(items, id)
	if ok {
		bytes, err := json.Marshal(&items[foundIdx])
		if err != nil {
			return fmt.Errorf("findById. json.Marshal: %w", err)
		}
		writer.Write(bytes)
		//fmt.Print(string(bytes))
	}

	return nil
}
