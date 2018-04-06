### Design motivation  

From my understanding, the REST philosophy modeled a software demand into two parts: resources(nouns) and the actions(verbs) that 
used these resources. Usually the resources are specificed by the URI and the actions are specificed by the http methods. ("GET","POST",etc.) 

In our case, assuming a structure "book" in the server:

Resources: addressbook 
Actions: download,upload,delete entry,add entry,...

Since there is only one resource----the address book, we can omit it without anyside effects. Thus, in the server.go, the URI/API maps are 
defined as:

	http.HandleFunc("/readall",readall)    //"GET"
	http.HandleFunc("/appendentry",appendentry) // "POST" 
	http.HandleFunc("/deleteentry",deleteentry) //"POST" 
	http.HandleFunc("/modifyentry",modifyentry) //"POST" 
	http.HandleFunc("/readentry",readentry)//"POST" 
	http.HandleFunc("/downloadaddressbook",downloadaddressbook) //"POST" 
	http.HandleFunc("/uploadaddressbook",uploadaddressbook) //"POST" 
 
Instead of the http methods, descriptive names are used to identify the actions. The actions for entries were decoupled into two stages: 
first,  identify the specific entry using the Id receiving from the http "POST" method; second, the actions on the specific entry.

### Data structure

Address entry:
type Address struct {
	Id string 
	Firstname string 
	Lastname string  
	Email string    
	Phonenumber string  
}

Id entry:
type Address_Id struct {
	Id string 
}

The whole address book is held in a slice:
	var book []Address

### API document

Below 7 functions are self explained by its name.

    151:	readall:=func(w http.ResponseWriter,r *http.Request){}
    164:	appendentry:=func(w http.ResponseWriter,r *http.Request){}
    177:	deleteentry:=func(w http.ResponseWriter,r *http.Request){}
    196:	modifyentry:=func(w http.ResponseWriter,r *http.Request){}		
    214:	readentry:=func(w http.ResponseWriter,r *http.Request){}
    224:	downloadaddressbook:=func(w http.ResponseWriter,r *http.Request){}
    232:	uploadaddressbook:=func(w http.ResponseWriter,r *http.Request){}

### Utility functions document

     27://this function is used to delete an entry have a specific ID
     28:func deletebyid(x Address_Id,book []Address)[]Address{}

     49://this function is used to return an entry have a specific ID
     50:func readbyid(x Address_Id,book []Address)Address{}

     70://this function is used to replace an existing entry with another entry having same Id
     71:func modifybyid(x Address,book []Address)[]Address{}

     //write csv file to the slice
     92:func writecsv(data []Address,filename string) {}

     //read csv file and save the contents to the slice
     115:func readcsv(filename string) (data []Address){}
### Conclusion
In this work, we investigated the "flat" approach to the REST API and microservice design. There could be other approachs. 

For example, we can treate the entries of the address book as the second resource. Instead of design 7 APIs for a single resource, we can 
assign the 7 required actions to two resources. This approach maybe more favorable for larger projects with many required actions/services.

