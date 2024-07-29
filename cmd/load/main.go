package main

import (
	"encoding/json"
	"fmt"
    "path/filepath"
	pb "internal/pathbuilder"
	pf "internal/pathfinder"
	"os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}
 
type encounter struct {}
type npc struct {}

type basicCampaign struct {
    Party []pf.Character
    npcs []npc
    encounters []encounter
}

func main() {
    pf2eDir := os.Getenv("PF2E_DIR")
    if pf2eDir == "" {
        homeDir, err := os.UserHomeDir()
        if err != nil {
            panic(err)
        }
        pf2eDir = filepath.Join(homeDir, ".pf2e")
    }
    buildDir := filepath.Join(pf2eDir, "pathbuilder")
    builds, err := os.ReadDir(buildDir)
    if err != nil {
        if os.IsNotExist(err) {
            os.MkdirAll(buildDir, os.ModePerm)
            fmt.Println("Created directory for pathbuilder builds, please grab json files off pathbuilder and store in .pf2e/pathbuilder")
        } else {
            panic(err)
        }
    }
    var c = basicCampaign{}
    for _, file := range builds {
        jsonData, err := os.ReadFile(filepath.Join(buildDir, file.Name()))
        check(err)

        pbc := pb.PathbuilderJSON{}
        err2 := json.Unmarshal([]byte(jsonData), &pbc)
        check(err2)

        char, err3 := pbc.Build.Convert()
        check(err3)
        c.Party = append(c.Party, char)
    }
    jsonParty, err := json.Marshal(c)
    check(err)
    campaignsDir := filepath.Join(pf2eDir, "campaigns")
    errWriteFile := os.WriteFile(filepath.Join(campaignsDir, "foo.json"), jsonParty, 0644)
    if errWriteFile != nil {
        if os.IsNotExist(errWriteFile) {
            os.MkdirAll(campaignsDir, os.ModePerm)
            errWriteFile = os.WriteFile(filepath.Join(campaignsDir, "foo.json"), jsonParty, 0644)
        }
        if errWriteFile != nil {
            panic(errWriteFile)
        }
    }
}
