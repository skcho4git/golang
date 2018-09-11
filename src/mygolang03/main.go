//OS Package
package main

import (

	"fmt"
	"os"
	"log"
	
)

func main() {
	
	//fmt.Println("hogehoge")
	
	//os.Exit(0)	
	
	/*fmt.Printf("length=%d\n", len(os.Args))
	for _,v := range os.Args {
		
		fmt.Println(v)
	}*/
	
	
	
	f, err := os.Open("sample.txt")
	
	if err != nil {
		
		fmt.Println("File Open Error!!")
		log.Fatal(err)
	}
	
	bs := make([]byte, 128)
	
	//n, err := f.Read(bs)	
	n, err := f.ReadAt(bs,10)
	
	//offset, err := f.Seek(10, os.SEEK_SET);
	//offset, err := f.Seek(-2, os.SEEK_CUR);
	offset, err := f.Seek(0, os.SEEK_END);
	
	fmt.Println(n)
	fmt.Println(offset)
	
	//fi:os.FileInfo
	/*fi, err := f.Stat()
	
	fmt.Println(fi.Name())
	fmt.Println(fi.Size())
	fmt.Println(fi.Mode())
	fmt.Println(fi.ModTime())
	fmt.Println(fi.IsDir())*/
	
	
	f2, err := os.Create("hoge.txt")
	fi, _ := f2.Stat()
	
	fmt.Println(fi.Name())
	fmt.Println(fi.Size())
	fmt.Println(fi.Mode())
	fmt.Println(fi.ModTime())
	fmt.Println(fi.IsDir())
	
	f2.Write([]byte("Hello World!\n"))
	f2.WriteAt([]byte("Golang\n"),7)
	f2.Seek(0, os.SEEK_END)
	f2.WriteString("Yeah!!")
	
	
	
	dir,_ := os.Getwd()
	
	fmt.Println(dir)
	
	f3, _ := os.Open(".")
	
	fis, _ := f3.Readdir(0)
	
	for _,fi3 := range fis {
		
		
		fmt.Println(fi3.Name())
		
		/*if fi3.IsDir() {
		
		    fmt.Println(fi3.Name())			
			
		}*/
	}
	
	os.Mkdir("hoge",0755)
	
	fmt.Println(os.TempDir())
	fmt.Println(os.Hostname())
	
	
	f.Close()
	f2.Close()
	f3.Close()
	
	os.Rename("hoge.txt", "fuo.txt")
	//os.Remove("hoge.txt")
	
	
	for _,v := range os.Environ() {
		
		fmt.Println(v)
	}
	
	fmt.Println(os.Getenv("path"))
	
	if path, ok := os.LookupEnv("path"); ok {
		
		fmt.Println(path)
	} else {
		
		fmt.Println("no path")
	}
	
	
	
	
	
	//defer f.Close()
	//defer f2.Close()


}