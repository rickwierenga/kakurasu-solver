package main

import (
  "fmt"
  "math/rand"
)

func randBool() bool {
  return rand.Float32() < 0.5
}

type Board struct {
  data [][]*bool
  solRows []int
  solCols []int
}

func NewBoard(n int) Board {
  // Generate solution board first, then calculate rows.
  sols :=  make([][]bool, n)
  for i := 0; i < n; i++ {
    sols[i] = make([]bool, n)
  }

  for i := 0; i < n; i++ {
    for j := 0; j < n; j++ {
      sols[i][j] = randBool()
    }
  }

  // calc solutions
  solRows := make([]int, n)
  solCols := make([]int, n)
  for i := 0; i < n; i++ {
    rowSum := 0
    colSum := 0
    for j := 0; j < n; j++ {
      if sols[i][j] {
        rowSum += j + 1
      }
      if sols[j][i] {
        colSum += j + 1
      }
    }
    solRows[i] = rowSum
    solCols[i] = colSum
  }

  // Generate random empty board
  data := make([][]*bool, n)
  for i := 0; i < n; i++ {
    data[i] = make([]*bool, n)
  }
  for i := 0; i < n; i++ {
    for j := 0; j < n; j++ {
      data[i][j] = nil
    }
  }

  return Board{data: data, solRows: solRows, solCols: solCols}
}

func (b Board) Print() {
  n := len(b.solRows)

  // Print row header
  for i := 0; i < n; i++ {
    if i == 0 {
      fmt.Printf(" ")
    }
    fmt.Printf(" %d", i + 1)
  }
  fmt.Println()

  // Print rows
  for i := 0; i < n; i++ {
    fmt.Printf("%d", i + 1)

    for j := 0; j < n; j++ {
      var sign string
      if b.data[i][j] == nil {
        sign = "?"
      } else if *b.data[i][j] {
        sign = "+"
      } else {
        sign = "-"
      }
      fmt.Printf(" %s", sign)
    }
    fmt.Printf(" %d", b.solRows[i])
    fmt.Println()
  }

  // Print sol cols
  for i := 0; i < n; i++ {
    if i == 0 {
      fmt.Printf(" ")
    }
    fmt.Printf(" %d", b.solCols[i])
  }

  fmt.Println()
}

func (b Board) Validate() bool {
  n := len(b.solCols)
  for i := 0; i < n; i++ {
    rowSum := 0
    colSum := 0
    for j := 0; j < n; j++ {
      if b.data[i][j] != nil && *b.data[i][j] {
        rowSum += j + 1
      }
      if b.data[j][i] != nil && *b.data[j][i] {
        colSum += j + 1
      }
    }
    if b.solRows[i] != rowSum {
      return false
    }
    if b.solCols[i] != colSum {
      return false
    }
  }

  return true
}
