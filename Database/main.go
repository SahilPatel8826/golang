package main

import (
	"encoding/json"
	"fmt"
)

const version = "1.0.0"

type(
	Logger interface{
		Fatal(string, ...interface{})
		Error(string, ...interface{})
		Warn(string, ...interface{})
		Info(string, ...interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})

	}
	Driver interface{
		mutex sync.Mutex
		mutexes map[string]*sync.Mutex
		dir string
		log Logger
	}
)
type Options struct{
	Logger
}

func New(dir string,options *Option)(*Driver,error){
	dir = filepath.clean(dir)

	opts:= Options{}

	if options!=nil{
		opts=*options
	}

	if opts.Logger==nil{
		opts.Logger=lumber.NewConsoleLogger((lumber.INFO))
	}
	driver := Driver{
		dir:dir,
		mutexes:make(map[string]*sync.Mutex),
		log:opts.Logger,
	}
	if _,err:=os.Stat(dir);err==nil{
		opts.Logger.Debug("Database directory exists:",dir)
		return &driver,nil
	}
	opts.Logger.Info("Creating database directory:",dir)
	return &driver,os.MkdirAll(dir,0755)

}

func(d *driver) Write(collection,resource string,v interface{})error{
	if collection==""{
		return fmt.Errorf("collection name cannot be empty")
	}
	if resource==""{
		return fmt.Errorf("resource name cannot be empty")
	}
	mutex:=d.getOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()	

	dir :=filepath.Join(d.dir,collection)
	fnlPath:=filepath.Join(dir,resource+".json")
	tempPath:=fnlPath+".tmp"

	if err:=os.MkdirAll(dir,0755);err!=nil{
		return err
	}

	b,err:=json.MarshalIndent(v,"","\t")
	if err!=nil{
		return err
	}	

	b = append(b,byte('\n'))
	if err:=ioutil.WriteFile(tempPath,b,0644);err!=nil{
		return err
	}
}
func (d *Driver)Read(collection,resource string,v interface{})error{
	if collection==""{
		return fmt.Errorf("collection name cannot be empty")
	}
	if resource=="" {
		return fmt.Errorf("resource name cannot be empty")
		
	}
	record:=filepath.Join(d.dir,collection,resource)

	if _,err:=stat(record);err!=nil{
		return err
	}

	b,err:=ioutil.ReadFile(record + ".json")
	if err!=nil{
		return err
	}

	return json.Unmarshal(b,&v)


}
func (d *Driver)ReadAll(collection string)([]string,error){
	if collection==""{
		return nil,fmt.Error{"Missimg Collection - unable to read"}

	}
	dir:=filepath.Join(d.dir,collection)

	if _,err:=stat(dir);err!=nil{
		return nil,err
	}

	files,_:=ioutil.ReadDir(dir)

	var records []string

	for _,file:=range files{
		b,err:=ioutil.ReadFile(filepath.Join(dir,file.Name()))
		if err!=nil{
			return nil,err
		}
		records=append(records,string(b))

	}
	return records,nil
	
}

func Delete()error{

}
func(d *Driver) getOrCreateMutex(collection string)*sync.Mutex{
   d.mutex.Lock()
   defer d.mutex.Unlock()

	m,ok:=d.mutexes[collection]

	if !ok{
		m=&sync.Mutex{}
		d.mutexes[collection]=m
	}
}
func stat(path string)( fi os.FileInfo,err error){
	if fi,err=os.Stat(path);os.IsNotExist(err){
		fi,err= os.Stat(path+".json")
	}
	return
}
type Address struct {
	City    string
	State   string
	Country string
	Pincode json.Number
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address string
}

func main() {
	dir := "./"

	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	employees := []User{
		{"John", "25", 233444333, "Myrl Tech", Address{"bangalore", "karnataka", "india", "410013"}},
		{"Steve", "33", 233444333, "Google", Address{"bangalore", "karnataka", "india", "410013"}},
		{"Robert", "29", 233444333, "Microsoft", Address{"bangalore", "karnataka", "india", "410013"}},
		{"Vance", "29", 233444333, "Facebook", Address{"bangalore", "karnataka", "india", "410013"}},
		{"Neo", "31", 233444333, "Remote Teams", Address{"bangalore", "karnataka", "india", "410013"}},
		{"Albert", "32", 233444333, "Dominate", Address{"bangalore", "karnataka", "india", "410013"}},
	}

	for _,value:=range employees{
		db.Write("users",value.Name,User{
			Name: value.Name,
			Age:value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,

		})
	}

	records,err:= db.ReadAll("users")
	if err!=nil{
		fmt.Println("Error",err)
	}
	fmt.Println(records)


	alluser:=[]User{}
	for _,record:=range records{
		employeesFound:=User{}
		json.Unmarshal([]byte(record),&employeesFound)
		alluser=append(alluser,employeesFound)
	}	
	fmt.Println("All Users:",alluser)

	// if err := db.Delete("user","john"); err != nil {
	// 	fmt.Println("Error", err)
	// }
    // if err := db.Delete("users",""); err != nil {
	// 	fmt.Println("Error", err)
	// }

}
