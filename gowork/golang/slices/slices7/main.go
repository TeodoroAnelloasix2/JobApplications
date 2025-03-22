/*
Slice of Slices
Slices can hold other slices, effectively creating a matrix, or a 2D slice.
rows := [][]int{}

# Assignment

We support various visualization dashboards on Textio that display message analytics to our users.
The UI for our graphs and charts is built on top of a grid system. Let's build some grid logic.
Complete the createMatrix function. It takes a number of rows and columns and returns a 2D slice of integers.
The value of each cell is i * j where i and j are the indexes of the row and column respectively.
Basically, we're building a multiplication chart.
*/
package main

func createMatrix(rows, cols int) [][]int {

	mult_chart := make([][]int, rows)
	for i := 0; i < rows; i++ {
		//crear cada fila
		row := make([]int, cols)
		for j := 0; j < cols; j++ {
			row[j] = i * j
		}
		mult_chart[i] = row //agregar la fila a la matriz
	}
	return mult_chart
}

/*
func main() {
	matrix := createMatrix(3, 4)
	for _, row := range matrix {
		fmt.Println(row)
	}
}
*/
