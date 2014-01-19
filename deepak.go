package main;

import "fmt"
import "time"
import "math/rand"




func main() {

    starts     := [4]string{"Experiential truth ", "The physical world ", "Non-judgment ", "Quantum physics "}
    middles    := [4]string{"nurtures an ", "projects onto ", "imparts reality to ", "constructs with "}
    qualifiers := [4]string{"abundance of ", "the barrier of ", "self-righteous ", "potential "}
    finishes   := [4]string{"marvel.", "choices.", "creativity.", "actions."}

    rand.Seed(time.Now().UnixNano())

    fmt.Print(starts[rand.Intn(4)])
    fmt.Print(middles[rand.Intn(4)])
    fmt.Print(qualifiers[rand.Intn(4)])
    fmt.Println(finishes[rand.Intn(4)])

}