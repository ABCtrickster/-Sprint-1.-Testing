// Задание 1 (Сдвиг)
package main

import "fmt"

func Rotate(data []int, pos int) []int {
    length := len(data)
    result := make([]int, length)
    
    for i := 0; i < length; i++ {
        newPosition := (pos + i) % length
        if newPosition < 0 {
            newPosition += length
        }        
        result[newPosition] = data[i]
    }
    
    return result
}

func main() {
    data := []int{1, 2, 3, 4, 5, 6, 7}
    pos := 3
    rotated := Rotate(data, pos)
    fmt.Println(rotated)
}

//Задание 2(Стек) 
package main

import (
	"errors"
	"fmt"
)

type Stack[T any] []T

func (s *Stack[T]) Push(key T) {
	*s = append(*s, key)
}

func (s *Stack[T]) Pop() (T, error) {
	if len(*s) == 0 {
		return zeroValue[T](), errors.New("стек пуст")
	}

	index := len(*s) - 1
	value := (*s)[index]
	*s = (*s)[:index]

	return value, nil
}

func zeroValue[T any]() T {
	var zero T
	return zero
}

func main() {
	stack := Stack[int]{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	for {
		value, err := stack.Pop()
		if err != nil {
			break
		}
		fmt.Println(value)
	}
}

//Задание 3 (CSV)
package main

import (
    "encoding/csv"
    "errors"
    "fmt"
    "os"
    "strconv"
)

func SumUp(filepath, colname string) (int, error) {
    file, err := os.Open(filepath)
    if err != nil {
        return 0, err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    columns, err := reader.Read()
    if err != nil {
        return 0, err
    }

    colIndex := -1
    for i, name := range columns {
        if name == colname {
            colIndex = i
            break
        }
    }

    if colIndex == -1 {
        return 0, errors.New("column not found")
    }

    sum := 0
    for {
        row, err := reader.Read()
        if err != nil {
            break
        }

        num, err := strconv.Atoi(row[colIndex])
        if err != nil {
            return 0, err
        }

        sum += num
    }

    return sum, nil
}

func main() {
    filepath := "data.csv"
    colname := "Number"

    sum, err := SumUp(filepath, colname)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Sum:", sum)
    }
}

//Задание 4(JSON)
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
)

func CompareJSON(json1, json2 []byte) (bool, error) {
    var obj1, obj2 interface{}
    err := json.Unmarshal(json1, &obj1)
    if err != nil {
        return false, fmt.Errorf("failed to unmarshal json1: %v", err)
    }

    err = json.Unmarshal(json2, &obj2)
    if err != nil {
        return false, fmt.Errorf("failed to unmarshal json2: %v", err)
    }

    marshaled1, err := json.Marshal(obj1)
    if err != nil {
        return false, fmt.Errorf("failed to marshal obj1: %v", err)
    }

    marshaled2, err := json.Marshal(obj2)
    if err != nil {
        return false, fmt.Errorf("failed to marshal obj2: %v", err)
    }

    return bytes.Equal(marshaled1, marshaled2), nil
}

func main() {
    json1 := []byte(`{"key1": "value1", "key2": "value2"}`)
    json2 := []byte(`{"key2": "value2", "key1": "value1"}`)
    json3 := []byte(`{"key1": "value1", "key2": "value3"}`)

    equal, err := CompareJSON(json1, json2)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("json1 is equal to json2:", equal)

    equal, err = CompareJSON(json1, json3)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("json1 is equal to json3:", equal)
}