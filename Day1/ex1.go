package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type Matrix struct{
	rows int `json:"rows"`
	columns int `json:"columns"`
	elements [][]int `json:"elements"`
}

func MatrixInit(row, column int) Matrix { // Matrix Initialization
	var matrix Matrix
	matrix.rows = row
	matrix.columns = column
	matrix.elements = make([][]int, row)
	for i:=0; i<row; i++{
		matrix.elements[i] = make([]int, column)
	}
	return matrix
}

func (matrix Matrix) GetNumberofRows() int { // Get number of rows of a matrix
	return matrix.rows
}

func (matrix Matrix) GetNumberofColumns() int { // Get number of columns of a matrix
	return matrix.columns
}

func (matrix *Matrix)  set(row, column, value int)  { // Set element on a particular position
	if row<0 || column<0 || row>=matrix.rows || column>=matrix.columns {
		err := errors.New("Index out of Bounds")
		log.Fatal(err)
	}
	matrix.elements[row][column] = value
}

func (FirstMatrix Matrix) add(SecondMatrix Matrix) Matrix { // Add two matrices
	row := FirstMatrix.rows
	column := FirstMatrix.columns
	FinalMatrix := MatrixInit(row, column)
	for i:=0; i<row; i++ {
		for j:=0; j<column; j++ {
			FinalMatrix.elements[i][j] = FirstMatrix.elements[i][j] + SecondMatrix.elements[i][j]
		}
	}
	return FinalMatrix
}

func (matrix Matrix) JsonMarshal() { // print in json format
	m, err := json.Marshal(matrix.elements)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(m))
}

func main() {
	matrix1 := MatrixInit(3, 3)
	matrix2 := MatrixInit(3,3)
	for i:=0; i<3; i++{
		for j:=0; j<3; j++ {
			matrix1.set(i, j, i+j)
			matrix2.set(i, j, 2*(i+j))
		}
	}

	matrix3 := matrix1.add(matrix2)

	fmt.Println(matrix3.rows)

	matrix1.JsonMarshal()
	matrix2.JsonMarshal()
	matrix3.JsonMarshal()

}

