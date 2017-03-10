package main

import (
    "bytes"
    "fmt"
    "os"
    "strconv"
    "time"
)

func main() {
    num, _  := strconv.Atoi(os.Args[1])
    convert := Convert(num)

    fmt.Printf("Arabic: %d\n", num)
    fmt.Printf("Roman: %v", convert)
}

func Convert(num int) string {
    defer timeTrack(time.Now(), "Convert")

    var buffer bytes.Buffer
    decimalMap := getDecimalMap()
    romanMap   := getRomanMap()

    for index := 0; index < len(decimalMap); index++ {
        for decimalMap[index] <= num {
            buffer.WriteString(romanMap[index])
            num -= decimalMap[index]
        }
    }

    return  buffer.String()
}

func getRomanMap() [13]string {

    return [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
}

func getDecimalMap() [13]int {

    return [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
}

func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    fmt.Printf("%s took %s\n", name, elapsed)
}