package main

import(
	"os"
	"encoding/csv"
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"log"
	"io"
)

type Address struct {
	Id string 
	Firstname string 
	Lastname string  
	Email string    
	Phonenumber string  
}

type Address_Id struct {
	Id string 
}


//this function is used to delete an entry have a specific ID
func deletebyid(x Address_Id,book []Address)[]Address{
	x1 :=x.Id
	s :=-1

	//s is the slice's index of the entry that will be deleted 
	for i:=0;i<len(book);i++{
		if x1 == book[i].Id {
			s = i
			break
		}
	}
	if s == -1 {
		temp:=[]Address{Address{"notfound","notfound","notfound","notfound","notfound"}}[:]
		return temp
	}


	
	return append(book[:s],book[s+1:]...)
}

//this function is used to return an entry have a specific ID
func readbyid(x Address_Id,book []Address)Address{
	x1 :=x.Id
	var s int
	s = -1

	//s is the slice's index of the entry that will be deleted 
	for i:=0;i<len(book);i++{
		if x1 == book[i].Id {
			s = i
			break
		}
	}

	if s == -1{
		return Address{"notfound","notfound","notfound","notfound","notfound"}
	}
	return book[s]
}


//this function is used to replace an existing entry with another entry having same Id
func modifybyid(x Address,book []Address)[]Address{
	s := -1
	x1 :=x.Id

	for i:=0;i<len(book);i++{
		if x1 == book[i].Id {
			s = i
			book[i] = x
			break
		}
	}

	if s == -1 {
		temp:=[]Address{Address{"notfound","notfound","notfound","notfound","notfound"}}[:]
		return temp
	}
	return book
}


//write csv to structure
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

//read csv file to structure
func readcsv(filename string) (data []Address){

    // Open CSV file
    f, err := os.Open(filename)
	if err != nil {
			log.Fatal(err)
	}

    defer f.Close()

    lines, err := csv.NewReader(f).ReadAll()
    if err != nil {
        panic(err)
    }
    var book []Address

    for _, line := range lines {
		tempvalue:=Address{line[0],line[1],line[2],line[3],line[4]}
        book = append(book,tempvalue)
    }
return book
} 


func main() {

	//book is the resource;
	msg1 :="the addressbook is empty"
	msg2 :="A new entry was appended"
	msg3 :="An entry was deleted"
	msg4 :="An entry was modified"
	
	var book []Address
	book = readcsv("myaddressbook.csv")

	//API for Read the whole book
	readall:=func(w http.ResponseWriter,r *http.Request){
		if len(book)>0 {
			j,_:=json.Marshal(book)
			w.Write(j)
		}else{

		//explictly report the empty addressbook
		j,_:=json.Marshal(msg1)
		w.Write(j)
			}
	}

	//API for append a record
	appendentry:=func(w http.ResponseWriter,r *http.Request){
		var record Address
		b,_:=ioutil.ReadAll(r.Body)
		err:=json.Unmarshal(b,&record)
			if err != nil {
			log.Fatal(err)
	}

		book = append(book,record)
		j,_:=json.Marshal(msg2)
		w.Write(j)
	}

	deleteentry:=func(w http.ResponseWriter,r *http.Request){
		var id_D Address_Id
		b,_:=ioutil.ReadAll(r.Body)
		err:=json.Unmarshal(b,&id_D)
			if err != nil {
			log.Fatal(err)
	}

		fmt.Print(err)
		book = deletebyid(id_D,book)
		fmt.Print(id_D)
		if book[0].Id == "notfound"{
			msg3 = "notfound"
		}
		j,_:=json.Marshal(msg3)
		w.Write(j)
		
	}

	modifyentry:=func(w http.ResponseWriter,r *http.Request){		
		var record Address
		b,_:=ioutil.ReadAll(r.Body)
		err:=json.Unmarshal(b,&record)
			if err != nil {
			log.Fatal(err)
	}

		book = modifybyid(record,book)
			if book[0].Id == "notfound"{
			msg4 = "notfound"
		}
	
		j,_:=json.Marshal(msg4)
		w.Write(j)
	}


	readentry:=func(w http.ResponseWriter,r *http.Request){
		var id_D Address_Id
		b,_:=ioutil.ReadAll(r.Body)
		_=json.Unmarshal(b,&id_D)
		s := readbyid(id_D,book)
		j,_:=json.Marshal(s)		
		w.Write(j)
	}

	
	downloadaddressbook:=func(w http.ResponseWriter,r *http.Request){
		name:="addressbook.csv"
		writecsv(book,name)
	    http.ServeFile(w, r, name)
	}

	//upload the addressbook
	//after finishing, the book variable is refreshed
	uploadaddressbook:=func(w http.ResponseWriter,r *http.Request){
	file, header, err := r.FormFile("file")

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	defer file.Close()

	out, err := os.Create(header.Filename)
	if err != nil {
		fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
		return
	}

	defer out.Close()

	// write the content to the file
	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	    fmt.Fprintf(w, "File uploaded successfully: ")
		fmt.Fprintf(w, header.Filename)
		
		book = nil
		book = readcsv(header.Filename)
		writecsv(book,"myaddressbook.csv")
	}
	
	http.HandleFunc("/readall",readall)
	http.HandleFunc("/appendentry",appendentry)
	http.HandleFunc("/deleteentry",deleteentry)
	http.HandleFunc("/modifyentry",modifyentry)
	http.HandleFunc("/readentry",readentry)
	http.HandleFunc("/downloadaddressbook",downloadaddressbook)
	http.HandleFunc("/uploadaddressbook",uploadaddressbook)
	http.ListenAndServe(":80", nil)

}
