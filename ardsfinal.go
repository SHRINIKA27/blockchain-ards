package main

import (
    "strconv"
    "math/big"
    "log"
    "encoding/csv"
    "os"
    "database/sql"
    "fmt"
    "math/rand"
    "time"
    _ "github.com/go-sql-driver/mysql"
)

type oxidata struct {
    
    ID string   
    HGB string  
    SAO2 string   
    SVO2 string  
    PAO2 string  
    PVO2 string  
    CO string  
}
var (
    SNO1 string
    ID1 string
    HGB1 string
    SAO21 string
    SVO21 string
    PAO21 string
    PVO21 string
    CO1 string
    DO21 string
    VO21 string
    ards string
)
var remainder = big.NewInt(0);
var pk = big.NewInt(0);
var ten1 = big.NewInt(0);
var cao2_range1 = big.NewInt(0);
var cvo2_range1 = big.NewInt(0);
var co_range1 = big.NewInt(0);
var co_range2 = big.NewInt(0);

const MAX int = 10000000;
var get_key int =0;
var arr_key int =0;
var arr_key1 int =3;
var get_keyvalues int =0;
var a1 int =0;    
var b1 int =0;
var j int =0;

var do2_range  = make([]*big.Int, 0)
var vo2_range  = make([]*big.Int, 0)
var hgb_range  = make([]*big.Int, 0)
var pkey  = make([]*big.Int, 0)

var key_hard = make([]int,0)
var keylist = make([]int,0)
var list = make([]int,0)

func GetData() {
    db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/")
    if err != nil {
    fmt.Println(err.Error())
    } else {
    fmt.Println("Database connected successfully...")
    }

    _,err = db.Exec("CREATE DATABASE ards2")
    if err != nil {
    fmt.Println(err.Error())
    } else {
    fmt.Println("Successfully created database..")
    }

    _,err = db.Exec("USE ards2")
    if err != nil {
    fmt.Println(err.Error())
    } else {
    fmt.Println("DB selected successfully..")
    }

    stmt, err := db.Prepare("CREATE Table details(SNO int NOT NULL AUTO_INCREMENT,ID int NOT NULL, HGB varchar(50), SAO2 varchar(30), SVO2 varchar(50), PAO2 varchar(30), PVO2 varchar(50), CO varchar(30), PRIMARY KEY (SNO));")
    if err != nil {
    fmt.Println(err.Error())
    }

    _, err = stmt.Exec()
    if err != nil {
    fmt.Println(err.Error())
    } else {
    fmt.Println("Table created successfully..")
    }

    stmt1, err1 := db.Prepare("CREATE Table gortis(SNO int NOT NULL AUTO_INCREMENT,ID int NOT NULL, HGB varchar(50), SAO2 varchar(30), SVO2 varchar(50), PAO2 varchar(30), PVO2 varchar(50), CO varchar(30), DO2 VARCHAR(100), VO2 varchar(100),PRIMARY KEY (SNO));")
    if err1 != nil {
    fmt.Println(err1.Error())
    }

    _, err1 = stmt1.Exec()
    if err1 != nil {
    fmt.Println(err1.Error())
    } else {
    fmt.Println("Gorti Table created successfully..")
    } 

    stmt2, err2 := db.Prepare("CREATE Table eanalysis(SNO int NOT NULL AUTO_INCREMENT,ID int NOT NULL, HGB varchar(50), SAO2 varchar(30), SVO2 varchar(50), PAO2 varchar(30), PVO2 varchar(50), CO varchar(30), DO2 VARCHAR(100), VO2 varchar(100),ards varchar(10),PRIMARY KEY (SNO));")
    if err2 != nil {
    fmt.Println(err1.Error())
    }

    _, err2 = stmt2.Exec()
    if err2 != nil {
    fmt.Println(err2.Error())
    } else {
    fmt.Println("eanalysis Table created successfully..")
    } 

    stmt3, err3 := db.Prepare("CREATE Table doutput(SNO int NOT NULL AUTO_INCREMENT,ID int NOT NULL, HGB varchar(50), SAO2 varchar(30), SVO2 varchar(50), PAO2 varchar(30), PVO2 varchar(50), CO varchar(30), DO2 VARCHAR(100), VO2 varchar(100),ards varchar(10),PRIMARY KEY (SNO));")
    if err3 != nil {
    fmt.Println(err3.Error())
    }

    _, err3 = stmt3.Exec()
    if err3 != nil {
    fmt.Println(err3.Error())
    } else {
    fmt.Println("doutput Table created successfully..")
    } 

    defer db.Close()
    csvFile, err := os.Open("oximeter_dataset25.csv")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Successfully Opened CSV file")
    defer csvFile.Close()
    
    csvLines, err := csv.NewReader(csvFile).ReadAll()
    if err != nil {
        fmt.Println(err)
    }   
    var i int
    i = 0
    for _, line := range csvLines {
        oximeter_data := oxidata{
            ID: line[0],
            HGB:  line[1],
            SAO2: line[2],
            SVO2:  line[3],
            PAO2: line[4],
            PVO2:  line[5],
            CO: line[6],
        }
        
        id := oximeter_data.ID
        hgb := oximeter_data.HGB
        sao2 := oximeter_data.SAO2
        svo2 := oximeter_data.SVO2
        pao2 := oximeter_data.PAO2
        pvo2 := oximeter_data.PVO2
        co := oximeter_data.CO

        if i==0 {
            i = i + 1
        } else {
        fmt.Println(id + " " + hgb + " " + sao2+ " " + svo2+ " " + pao2+ " " + pvo2+ " " + co)
        insert, err := db.Prepare("INSERT INTO details ( ID , HGB , SAO2 , SVO2 , PAO2 , PVO2 , CO ) VALUES ( ? , ? , ? , ? , ? , ? , ? )")
        
        if err != nil {
            panic(err.Error())
        }
        
        insert.Exec(id,hgb,sao2,svo2,pao2,pvo2,co)
        defer insert.Close()
    }
    }
    fmt.Println("Insertion into database completed!...")

    //function-2
    
}


func Encryption_Initialization() {
           
            //prime get  
            get_key=get_key+1;
            get_keyvalues=get_key+1;
            key_generation := new(big.Int)  
            
            key_generation.Mul(big.NewInt(int64(list[get_key])),big.NewInt(int64(list[get_keyvalues])));
            fmt.Println(get_key);
            pkey=append(pkey,big.NewInt(int64(list[get_keyvalues])));
            //random key
            var result int =key_hard[arr_key];
            arr_key+=1;
            
            //power 
            power_values := new(big.Int)  

            power_values.Exp((big.NewInt(int64(list[get_keyvalues]))),big.NewInt(int64(list[get_key])),nil);
            //final power
            final_power:= new(big.Int)  

            final_power.Mul(power_values,big.NewInt(int64(result)));

            //mod
            remainder.Mod(final_power,key_generation);
    }

    func Encryption_process(val string) (ncrypted_unit *big.Int) {

            id1:= new(big.Int)  
            ncrypted_unit = new(big.Int)  
            id1,ok:= id1.SetString(val, 10)
            if !ok {
             fmt.Println("SetString: error")
             return
            }
            ncrypted_unit.Add(id1,remainder);
            return            
    }


func Encryption() {
    
    db1, err1 := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/")
    if err1 != nil {
    fmt.Println(err1.Error())
    } else {
    fmt.Println("Database connection established...")
    }

    _,err1 = db1.Exec("USE ards2")
    if err1 != nil {
    fmt.Println(err1.Error())
    } else {
    fmt.Println("DB selected successfully..")
    }
    //random number generation

    get_key =1000+rand.Intn(400);
        var i int 
        for  i = 0;i<MAX;i++ {
                      
                    var key int = 150+rand.Intn(200000);
                    if key>0 {
                        
                        key_hard = append(key_hard,key);
                        }       
        }
    //  prime number gneration

    var beg=1500
    var end=20000 
    var n int 
    var j int 
            for  n = beg; n <= end; n++ {
            var prime = true;
            for  j = 2; j <= n / 2; j++ {
                if n % j == 0 && n != j {
                    prime = false;
                }
            }
            if prime {
                list = append(list,n);
            }
            }    


    db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/")
    if err != nil {
    fmt.Println(err.Error())
    } else {
    fmt.Println("Database connection established...")
    }

    _,err = db.Exec("USE ards2")
    if err != nil {
    fmt.Println(err.Error())
    } else {
    fmt.Println("DB selected successfully..")
    }

    fmt.Println("Fetching from database...")
    rows, err := db.Query("select SNO , ID , HGB , SAO2 , SVO2 , PAO2 , PVO2 , CO from details")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    for rows.Next() {
        err := rows.Scan(&SNO1, &ID1, &HGB1, &SAO21, &SVO21, &PAO21 , &PVO21 , &CO1)
        if err != nil {
            log.Fatal(err)
        }
        log.SetFlags(0)

    //converting into float for cao2 calculation
    f1, _ := strconv.ParseFloat(HGB1, 32)
    f2, _ := strconv.ParseFloat(SAO21, 32)
    f3, _ := strconv.ParseFloat(PAO21, 32)
    f4, _ := strconv.ParseFloat(SVO21, 32)
    f5, _ := strconv.ParseFloat(PVO21, 32)


    //cao2 calculation
    cao2calc :=(1.34*f1*(f2/100))+(f3*0.0031)
    fmt.Print("CAO2 :")
    fmt.Println(int64(cao2calc))
    cao2s := strconv.Itoa(int(cao2calc))

    //cvo2 calculation
    cvo2calc :=(1.34*f1*(f4/100))+(f5*0.0031)
    fmt.Print("CVO2 :")
    fmt.Println(int64(cvo2calc))


    //cao2-cvo2
    cao2calc1 :=int64(cao2calc)
    cvo2calc1 :=int64(cvo2calc)
    dif := cao2calc1 - cvo2calc1
    fmt.Print("DIFF :")
    fmt.Println(int64(dif))
    difs := strconv.Itoa(int(dif))

    //encryption starts
    defer db.Close()

    Encryption_Initialization();

    IDe:= new(big.Int) 
    IDe=Encryption_process(ID1) 
    HGBe:= new(big.Int) 
    HGBe=Encryption_process(HGB1) 
    SAO2e:= new(big.Int) 
    SAO2e=Encryption_process(SAO21) 
    SVO2e:= new(big.Int) 
    SVO2e=Encryption_process(SVO21) 
    PAO2e:= new(big.Int) 
    PAO2e=Encryption_process(PAO21) 
    PVO2e:= new(big.Int) 
    PVO2e=Encryption_process(PVO21) 
    COe:= new(big.Int) 
    COe=Encryption_process(CO1) 
    CAO2e:= new(big.Int) 
    CAO2e=Encryption_process(cao2s) 

    DIFe:= new(big.Int) 
    DIFe=Encryption_process(difs) 

    tene:= new(big.Int) 
    tene=Encryption_process("10") 
    

    //DO2 Calculation
    DO2calce := new(big.Int) 
    DO2calce.Mul(DO2calce.Mul(CAO2e,COe),tene)
    //fmt.Print("DO2 :")
    //fmt.Println(DO2calce.Mod(DO2calce,pkey[b1]))

    //VO2 Calc
    VO2calce := new(big.Int) 
    VO2calce.Mul(VO2calce.Mul(COe,DIFe),tene)
    //fmt.Print("VO2 :")
    //fmt.Println(VO2calce.Mod(VO2calce,pkey[b1]))

    cao2_r1 := "11";
    cvo2_r1 := "8";
    co_r1 := "3";
    
    hgb_r2 := "11";

    cao2_re:= new(big.Int) 
    cao2_re=Encryption_process(cao2_r1)

    cvo2_re:= new(big.Int) 
    cvo2_re=Encryption_process(cvo2_r1)

    co_r1e:= new(big.Int) 
    co_r1e=Encryption_process(co_r1)

    
    hgb_re:= new(big.Int) 
    hgb_re=Encryption_process(hgb_r2)

    hgb_range = append(hgb_range,hgb_re);

    DO2ra := new(big.Int) 
    DO2ra.Mul(DO2ra.Mul(cao2_re,co_r1e),tene)

    do2_range = append(do2_range,DO2ra);

    VO2ra := new(big.Int) 
    VO2ra.Mul(VO2ra.Mul(co_r1e,cvo2_re),tene)

    vo2_range = append(vo2_range,VO2ra);
    b1++

    //insertion into gorti table
    insert, err := db1.Prepare("INSERT INTO gortis ( ID , HGB , SAO2 , SVO2 , PAO2 , PVO2 , CO , DO2 , VO2 ) VALUES ( ? , ? , ? , ? , ? , ? , ? , ? , ? )")
        if err != nil {
            panic(err.Error())
        }
        
        insert.Exec(IDe.String() , HGBe.String() , SAO2e.String() , SVO2e.String() , PAO2e.String() , PVO2e.String() , COe.String() , DO2calce.String() , VO2calce.String())
        defer insert.Close()
        
    defer db1.Close()
    fmt.Println("Insertion into gorti table is successfull")





    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }
    
    b1=0
}

func Analysis() {

    db2, err2 := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/")
    if err2 != nil {
    fmt.Println(err2.Error())
    } else {
    fmt.Println("Database connection established...")
    }

    _,err2 = db2.Exec("USE ards2")
    if err2 != nil {
    fmt.Println(err2.Error())
    } else {
    fmt.Println("DB selected successfully..")
    }

    db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/")
    if err != nil {
    fmt.Println(err.Error())
    } else {
    fmt.Println("Database connection established...")
    }

    _,err = db.Exec("USE ards2")
    if err != nil {
    fmt.Println(err.Error())
    } else {
    fmt.Println("DB selected successfully..")
    }

    fmt.Println("Fetching from database...")
    rows, err := db.Query("select SNO , ID , HGB , SAO2 , SVO2 , PAO2 , PVO2 , CO ,DO2 , VO2 from gortis")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    for rows.Next() {
        err := rows.Scan(&SNO1, &ID1, &HGB1, &SAO21, &SVO21, &PAO21 , &PVO21 , &CO1 , &DO21 ,&VO21)
        if err != nil {
            log.Fatal(err)
        }


        var temp =0;
        var temp1 =0;
        var temp2 =0;
        var ards="No"

        bi1 := new(big.Int)
        bi1, ok := bi1.SetString( DO21 , 10)
        if !ok {
        fmt.Println("SetString: error")
        return
        }
        bi2 := new(big.Int)
        bi2, ok = bi2.SetString( VO21 , 10)
        if !ok {
        fmt.Println("SetString: error")
        return
        }
        bi3 := new(big.Int)
        bi3, ok = bi3.SetString( HGB1 , 10)
        if !ok {
        fmt.Println("SetString: error")
        return
        }

var c1 =0;
       var c2 =0;
       var c3 =0;

       c1=do2_range[b1].Cmp(bi1)
       fmt.Println(c1)

       c2=vo2_range[b1].Cmp(bi2)
       fmt.Println(c1)

       c3=hgb_range[b1].Cmp(bi3)
       fmt.Println(c1)

       if (c1==0 || c1==1) {
        temp=1
       }
       if (c2==0 || c2==1) {
        temp1=1
       }
       if (c3==0 || c3==1) {
        temp2=1
       }
     
       if ( (temp==1 && temp2==1) || (temp==1 && temp1==1) || (temp1==1 || temp2==1 )) {
                ards="Yes";
        }
       fmt.Println(ards)


        b1++

            //insertion into eanalysis table
    insert, err := db2.Prepare("INSERT INTO eanalysis ( ID , HGB , SAO2 , SVO2 , PAO2 , PVO2 , CO , DO2 , VO2, ARDS ) VALUES ( ? , ? , ? , ? , ? , ? , ? , ? , ? , ? )")
        if err != nil {
            panic(err.Error())
        }
        
        insert.Exec(ID1 , HGB1 , SAO21 , SVO21 , PAO21 , PVO21 , CO1 , DO21 , VO21 , ards)
        defer insert.Close()
        
    defer db2.Close()
    fmt.Println("Insertion into eanalysis table is successfull")
    b1=0



    }

}

func Decryption_process(id1 string) (n1 *big.Int) {
        
    id1b := new(big.Int)
    id1b, ok := id1b.SetString( id1 , 10)
    if !ok {
        fmt.Println("SetString: error")
        return
    }
    n1 = new(big.Int) 
    n1.Mod(id1b,pkey[b1]);
    return 
}

func Decryption() {
    
    db3, err3 := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/")
    if err3 != nil {
    fmt.Println(err3.Error())
    } else {
    fmt.Println("Database connection established...")
    }

    _,err3 = db3.Exec("USE ards2")
    if err3 != nil {
    fmt.Println(err3.Error())
    } else {
    fmt.Println("DB selected successfully..")
    }

   db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/")
    if err != nil {
    fmt.Println(err.Error())
    } else {
    fmt.Println("Database connection established...")
    }

    _,err = db.Exec("USE ards1")
    if err != nil {
    fmt.Println(err.Error())
    } else {
    fmt.Println("DB selected successfully..")
    }

    fmt.Println("Fetching from database...")
    rows, err := db.Query("select SNO , ID , HGB , SAO2 , SVO2 , PAO2 , PVO2 , CO ,DO2 , VO2 , ARDS from eanalysis")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    for rows.Next() {
        err := rows.Scan(&SNO1, &ID1, &HGB1, &SAO21, &SVO21, &PAO21 , &PVO21 , &CO1 , &DO21 ,&VO21, &ards)
        if err != nil {
            log.Fatal(err)
        }
        log.SetFlags(0)
        
        IDd:= new(big.Int) 
        IDd =Decryption_process(ID1)
        
        HGBd:= new(big.Int)
        HGBd =Decryption_process(HGB1)
        
        SAO2d:= new(big.Int) 
        SAO2d =Decryption_process(SAO21)
        
        SVO2d:= new(big.Int) 
        SVO2d =Decryption_process(SVO21)
        
        PAO2d:= new(big.Int) 
        PAO2d =Decryption_process(PAO21)
        
        PVO2d:= new(big.Int) 
        PVO2d =Decryption_process(PVO21)
        
        COd:= new(big.Int) 
        COd =Decryption_process(CO1)
        
        DO2d:= new(big.Int) 
        DO2d =Decryption_process(DO21)
        
        VO2d:= new(big.Int) 
        VO2d =Decryption_process(VO21)

        var hgbcheck =hgb_range[b1].String();
        fmt.Println(Decryption_process(hgbcheck));


        
    fmt.Println("---Decryption values---")
    fmt.Println("SNO :" + SNO1)
    fmt.Print("ID :")
    fmt.Println(IDd)
    fmt.Print("HGB :")
    fmt.Println(HGBd) 
    fmt.Print("SAO2 :")
    fmt.Println(SAO2d)
    fmt.Print("SVO2 :" )
    fmt.Println(SVO2d)
    fmt.Print("PAO2 :")
    fmt.Println(PAO2d)
    fmt.Print("PVO2 :")
    fmt.Println(PVO2d)
    fmt.Print("CO :")
    fmt.Println(COd)
    fmt.Print("DO2 :")
    fmt.Println(DO2d)
    fmt.Print("VO2 :")
    fmt.Println(VO2d)
    fmt.Print("ARDS :")
    fmt.Println(ards)
    fmt.Println("-----------------------")
    fmt.Println()

                //insertion into eanalysis table
    insert, err := db3.Prepare("INSERT INTO doutput ( ID , HGB , SAO2 , SVO2 , PAO2 , PVO2 , CO , DO2 , VO2, ARDS ) VALUES ( ? , ? , ? , ? , ? , ? , ? , ? , ? , ? )")
        if err != nil {
            panic(err.Error())
        }
        
        insert.Exec(IDd.String() , HGBd.String() , SAO2d.String() , SVO2d.String() , PAO2d.String() , PVO2d.String() , COd.String() , DO2d.String() , VO2d.String() , ards)
        defer insert.Close()
        
    defer db3.Close()
    fmt.Println("Insertion into doutput table is successfull")

    b1++
}
defer db.Close()
}


func main() {

    
    rand.Seed(time.Now().UnixNano())
    GetData()
    Encryption()
    Analysis()
    Decryption()
    
    
}
Defined on line 677
