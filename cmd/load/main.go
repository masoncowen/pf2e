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
    party []pf.Character
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

        char, err := pbc.Build.Convert()
        check(err)
        fmt.Println(char)
        // jsonChar, err := json.Marshal(char)
        // check(err)
        // fmt.Printf("%s\n", jsonChar)
        p.party = append(p.party, char)
    }
    fmt.Println(p)
    jsonParty, err := json.Marshal(p)
    check(err)
    fmt.Printf("%s\n", jsonParty)
}
