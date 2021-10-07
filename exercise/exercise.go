package main
import (
	"bufio"
	"fmt"
	"gifhelper"
	"image"
	"os"
	"strconv"
	//"strings"
	//"math"

)

// The data stored in a single cell of a field
type Cell struct {
	former_strategy string
	current_strategy  string //represents "C" or "D" corresponding to the type of prisoner in the cell
	score float64 //represents the score of the cell based on the prisoner's relationship with neighboring cells
}

// The GameBoard is a 2D slice of Cell objects
type GameBoard [][]Cell
//some definition
var Board GameBoard
var line []string
var column int
var row int

func read_image(txt string)string{
	filePath := txt
	file,err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error: something went wrong opening the file")
		fmt.Println("Probably you gave the wrong filename")
		return ""
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file) //关闭文本流

	reader := bufio.NewScanner(file)//读取文本
	for reader.Scan(){
	line = append(line,reader.Text())
		}
		//till now it's a line by line stored array
	num := len(line[0])
	row,err= strconv.Atoi (string(line[0][0:(num-1)/2]))
	column,err = strconv.Atoi(string(line[0][(num+1)/2:num]))
	return ""
}

func init_board()GameBoard{
	//we need paddings to simplify the process
	Board = make(GameBoard, row+2)
	// now we need to make the rows too
	for r := range Board {
		Board[r] = make([]Cell, column+2)
	}

	for i:=1;i<=row;i++ {
		for k := 0; k < column; k++ {
			Board[i][k+1].current_strategy = string(line[i][k])
			Board[i][k+1].former_strategy = string(line[i][k])
			Board[i][k+1].score = 0
		}
	}
	//next add paddings
	for i:=0;i<column+2;i++{
		Board[0][i].current_strategy = string('N')
		Board[0][i].former_strategy = string('N')
		Board[0][i].score = 0
		Board[row+1][i].current_strategy = string('N')
		Board[row+1][i].former_strategy = string('N')
		Board[row+1][i].score = 0
	}
	for i:=1;i<row+2;i++{
		Board[i][0].current_strategy = string('N')
		Board[i][0].former_strategy = string('N')
		Board[i][0].score = 0
		Board[i][column+1].current_strategy = string('N')
		Board[i][column+1].former_strategy = string('N')
		Board[i][column+1].score = 0
	}
	return Board
}

func countScore(r int,c int,b float64)float64{
    prisoner := Board[r][c].current_strategy
	var score[8] string
	var s float64 = 0
	score[0] = prisoner + Board[r-1][c-1].current_strategy
	score[1] = prisoner + Board[r-1][c].current_strategy
	score[2] = prisoner + Board[r-1][c+1].current_strategy
	score[3] = prisoner + Board[r][c-1].current_strategy
	score[4] = prisoner + Board[r][c+1].current_strategy
	score[5] = prisoner + Board[r+1][c-1].current_strategy
	score[6] = prisoner + Board[r+1][c].current_strategy
	score[7] = prisoner + Board[r+1][c+1].current_strategy
	if prisoner == string('C'){

		for i:=0;i<8;i++{
			if score[i] == string("CC"){
				s += 1
			}
		}
	}
	if prisoner == string('D'){

		for i:=0;i<8;i++{
			if score[i] == string("DC"){
				s += b
			}
		}
	}
	return s
}

func step_update(r int,c int)string{
	var group[9] float64
	var m string
	group[0] = Board[r-1][c-1].score
	group[1] = Board[r-1][c].score
	group[2] = Board[r-1][c+1].score
	group[3] = Board[r][c-1].score
	group[4] = Board[r][c].score
	group[5] = Board[r][c+1].score
	group[6] = Board[r+1][c-1].score
	group[7] = Board[r+1][c].score
	group[8] = Board[r+1][c+1].score
	//找最大值和最大值索引
	maxVal := group[0]
	maxValIndex := 0
	for i := 0; i < 9; i++ {
		//从第二个元素开始循环比较，如果发现有更大的数，则交换
		//这里用小于等于就换，意在尽可能换不会为0的区域
		if maxVal <= group[i] {
			maxVal = group[i]
			maxValIndex = i
		}
	}
	numTable := map[int]string{0: Board[r-1][c-1].former_strategy,
		                       1:  Board[r-1][c].former_strategy,
							   2:  Board[r-1][c+1].former_strategy,
							   3:  Board[r][c-1].former_strategy,
		                       4:  Board[r][c].former_strategy,
		                       5:  Board[r][c+1].former_strategy,
		                       6:  Board[r+1][c-1].former_strategy,
		                       7:  Board[r+1][c].former_strategy,
							   8:  Board[r+1][c+1].former_strategy}
	if numTable[maxValIndex] == string('C') {
		m = string('C')
	}
	if numTable[maxValIndex] == string('D') {
		m = string('D')
	}
	return m
}



func drawGameBoards(Board GameBoard,cellWidth int) Canvas {
	 	numRows := len(Board)
	 	numCols := len(Board[0])
	 	picture := CreateNewCanvas(cellWidth*numRows, cellWidth*numCols)
	 	picture.SetFillColor(MakeColor(0, 0, 255))
	 	picture.Clear()
	 	picture.SetFillColor(MakeColor(255, 0, 0))
	 	for i := 1; i < numRows-1; i++ {
	 		for j := 1; j < numCols-1; j++ {
	 			if Board[i][j].current_strategy == "D" {
	 				picture.MoveTo(float64(cellWidth)*float64(i), float64(cellWidth)*float64(j))
	 				picture.LineTo(float64(cellWidth)*float64(i+1), float64(cellWidth)*float64(j+1))
					picture.Stroke()
	 			}
	 		}
	 	}
	 	return picture
	 }

func main() {
	txt, b, steps := os.Args[1], os.Args[2], os.Args[3]
	num_b, _ := strconv.ParseFloat(b, 64)
	num_step, _ := strconv.Atoi(steps)

    read_image(txt)
	Board = init_board()
	//
	picture := drawGameBoards(Board,5)
	imageList := []image.Image{GetImage(picture)}
//循环
	for step :=0;step<num_step;step++{
		//算分
		for i:=1;i<1+row;i++{
			for k:=1;k<1+column;k++{
				Board[i][k].score = countScore(i,k,num_b)
			}
		}
        //更新
		for i:=1;i<1+row;i++{
			for k:=1;k<1+column;k++{
				Board[i][k].current_strategy = step_update(i,k)
			}
		}

        //旧版用新版覆盖再进下一次循环
		for i:=1;i<1+row;i++{
			for k:=1;k<1+column;k++{
				Board[i][k].former_strategy = Board[i][k].current_strategy
			}
		}
		picture = drawGameBoards(Board,5)
		imageList = append(imageList, GetImage(picture))
	}
	gifhelper.ImagesToGIF(imageList, "Prisoners")
	picture = drawGameBoards(Board,5)
	picture.SaveToPNG("Prisoners.png")

}