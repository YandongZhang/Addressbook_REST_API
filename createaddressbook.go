package main

import(
	"os"
	"fmt"
	"encoding/csv"
	"log"
	"strconv"

)

type Address struct {
	Id string 
	Firstname string 
	Lastname string  
	Email string    
	Phonenumber string  
}

//write csv
func writecsv(data []Address,filename string) {
    file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
	}

    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()
	
    for _, value := range data {
		tempvalue:=[]string{value.Id,value.Firstname,value.Lastname,value.Email,value.Phonenumber}
        err := writer.Write(tempvalue)
			if err != nil {
			log.Fatal(err)
	}

    }
}

func int2str(x int)string{
	return fmt.Sprintf("%d",x)
}

func createcsvfile(size int,name string){
	var book []Address
	var entry Address
	for i:=0;i<size;i++{
		entry = Address{int2str(i),"firstname"+int2str(i),"lastname"+int2str(i),int2str(i)+"@"+"email","phone"+int2str(i)}
		book = append(book,entry)
	}
	writecsv(book,name)
}

func main(){
		N1 := 100

		N1,_ =strconv.Atoi(os.Args[1])

	createcsvfile(N1,"myaddressbook.csv")

}
