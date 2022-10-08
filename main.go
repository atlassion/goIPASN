package main

import (
	"bufio"
    "fmt"
    "os"
    "os/user"
    "strings"
    "github.com/zmap/go-iptree/iptree"
    "io/ioutil"
    "encoding/json"
)


func main() {
	usr, _ := user.Current()
    uhdir := usr.HomeDir
	IPASN,_ := os.Open(uhdir+"/asn/IPASN.DAT")
	defer IPASN.Close()

	t := iptree.New()

	scanner := bufio.NewScanner(IPASN)
	
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), ";") == false {
			words := strings.Fields(scanner.Text())
			t.AddByString(words[0], words[1]+" "+words[0] )
		}
	}
	 
	AsnJson, _ := os.Open(uhdir+"/asn/asn.json")
	defer AsnJson.Close()

   	byteAsnJson, _ := ioutil.ReadAll(AsnJson)

    sc := bufio.NewScanner(os.Stdin)

    buf := make([]byte, 0, 64*1024)
    sc.Buffer(buf, 1024*1024)

    for sc.Scan() {

	    asn_a, _, _ := t.GetByString(sc.Text())
	    asn_b :=  strings.Split(fmt.Sprintf("%v",asn_a)," ")[0]

        var resAsnJson map[string]interface{}
	   	json.Unmarshal([]byte(byteAsnJson), &resAsnJson)
        
	   	fmt.Println(sc.Text(),asn_a,resAsnJson[asn_b])
	    
}
}
