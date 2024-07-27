package main

import (
	"encoding/json"
	"fmt"
	pb "internal/pathbuilder"
	pf "internal/pathfinder"
	"os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}
 
type party struct {
    Party []pf.Character
}

func main() {
    var p = party{}
    files := [4]string{"campaign/characters/Potato.json",
    "campaign/characters/Sonja.json",
    "campaign/characters/Gord.json",
    "campaign/characters/Otic.json"}

    for _, file := range files {
        jsonData, err := os.ReadFile(file)
        check(err)

        pbc := pb.PathbuilderJSON{}
        err2 := json.Unmarshal([]byte(jsonData), &pbc)
        check(err2)

        char, err3 := pbc.Build.Convert()
        check(err3)
        // fmt.Println(char)
        // jsonChar, err4 := json.Marshal(char)
        // check(err4)
        // fmt.Printf("%s\n", jsonChar)
        p.Party = append(p.Party, char)
    }
    fmt.Println(p)
    jsonParty, err := json.Marshal(p)
    check(err)
    fmt.Printf("%s\n", jsonParty)
}
