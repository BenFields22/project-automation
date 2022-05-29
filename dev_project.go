package main

import (
	"fmt"
	"runtime"
    "errors"
    "os"
    "log"
)

func createFile(folder string,fileName string){
    path, err := os.Getwd()
    if err != nil {
        panic(err)
    }
    file, err := os.Create(path+"/"+folder+"/"+fileName)
    if err != nil {
        panic(err)
    }
    fmt.Println("Created file: "+path+"/"+folder+"/"+fileName)
    file.Close()
    if err != nil {
        panic(err)
    }

}
        
func createFolder(folderName string){
    path := folderName 
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}

func writeMakeFile(folder string){
    path, err := os.Getwd()
    fi, err := os.OpenFile(path+"/"+folder+"/Makefile", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
    if err != nil {
        panic(err)
    }
    defer fi.Close()
    _, err2 := fi.WriteString("default:\n\tg++ source.cpp --std=c++11 -o run\nrunner:\n\tcat ./input.txt | ./run\nclean:\n\trm ./run")

    if err2 != nil {
        log.Fatal(err2)
    }
}

func writeCppFile(folder string){
    path, err := os.Getwd()
    fi, err := os.OpenFile(path+"/"+folder+"/source.cpp", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
    if err != nil {
        panic(err)
    }
    defer fi.Close()
    _, err2 := fi.WriteString("#include <iostream>\n\nint main(int argc, char** argv){\n\n\treturn 0;\n}\n")

    if err2 != nil {
        log.Fatal(err2)
    }
}

func executeProgram() {
    fmt.Println("Enter the folder name: ")
    var folder string
    fmt.Scanln(&folder)
    createFolder(folder)
    createFile(folder,"Makefile")
    createFile(folder,"input.txt")
    createFile(folder,"source.cpp")
    writeMakeFile(folder)
    writeCppFile(folder)
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		executeProgram()
	}
}
