//
//  Solution to the SwiftStack PartyCo company Party problem.
//  See README for the problem description.
//
package main

import "fmt"

type Person struct {
    name string
    boss string
    score float64
}

//
//  Canned PartyCo employee data.
//
var party_co = []*Person{
    &Person{
        name: "Al Buquerque", 
        boss: "", 
        score: 2.0,
    },
    &Person{
        name: "Ferb Jinglemore",
        boss: "Al Buquerque",
        score: 12.1,
    },
    &Person{
        name: "Click N. Clack",
        boss: "Al Buquerque",
        score: 34.3,
    },
    &Person{
        name: "Carl Balgruuf",
        boss: "Ferb Jinglemore",
        score: -0.4,
    },
    &Person{
        name: "Moe Shroom",
        boss: "Carl Balgruuf",
        score: 44.91,
    },
    &Person{
        name: "Jerky McGetsDrunkAndPeesInYourFridge",
        boss: "Carl Balgruuf",
        score: -9999.99,
    },
    &Person{
        name: "Howard M. Burgers",
        boss: "Click N. Clack",
        score: 14.4,
    },
    &Person{
        name: "Soren de Kiester",
        boss: "Click N. Clack",
        score: 25,
    },
}

func validateList(invitees []*Person) bool {
    // Validate an invite list.
    for _, in1 := range invitees {
        for _, in2 := range invitees {
            if in2.name == in1.boss {
                return false
            }
        }
    }
    return true
}

func allowedLists(party_co []*Person, c chan []*Person) {
    // Generate all allowed guest lists for parties of two and greater.
    defer close(c)
    for n := 1; n <= len(party_co); n++ {
        n_combos := make(chan []*Person)
        go combinations(party_co, n, n_combos)
        for guest_list := range n_combos {
            if validateList(guest_list) {
                c <- guest_list
            }
        }
    }
}

func score(invitees []*Person) float64 {
    // Party potential for a guest list.
    score := 0.0
    for _, person := range invitees {
        score += person.score
    }
    return score
}

func hasBossInList(list []*Person, name string) bool {
    // Returns true is some Person in list has name as boss.
    for _, p := range list {
        if p.boss == name {
            return true
        }
    }
    return false
}

func main() {
    // We're lookging for the invite list with the highest party potential that
    // meets the criteria, whether or not the CEO is invited; we'll also
    // look for the 'best' party with the CEO as an attendee.
    wildest := 0.0
    var party []*Person
    wildestCEO := 0.0
    var partyCEO []*Person

    // allowedLists() evalutes potential guest lists to see if they
    // meet the problem criteria and communicates them back with a
    // channel.
    c := make(chan []*Person)
    go allowedLists(party_co, c)
    for list := range(c) {
        score := score(list)
        if score > wildest {
            wildest = score
            party = list
        }
        if hasBossInList(list, "") {
            if score > wildestCEO {
                wildestCEO = score
                partyCEO = list
            }
        }
    }
    fmt.Println("Wildest party potential: ", wildest)
    for _, person := range party {
        fmt.Println(person.name, person.score)
    }
    fmt.Println("\nWildest party potential (CEO invited): ", wildestCEO)
    for _, person := range partyCEO {
        fmt.Println(person.name, person.score)
    }
}

func combinations(iterable []*Person, r int, c chan []*Person) {
    //
    //  This was cribbed from the Go playground and modified a bit.
    //  (Use {}*Person as iterable and communicate results in a channel.)
    //
    defer close(c)
    pool := iterable
    n := len(pool)

    if r > n {
        return
    }

    indices := make([]int, r)
    for i := range indices {
        indices[i] = i
    }

    result := make([]*Person, r)
    for i, el := range indices {
        result[i] = pool[el]
    }

    c <- result

    for {
        i := r - 1
        for ; i >= 0 && indices[i] == i+n-r; i -= 1 {
        }

        if i < 0 {
            return
        }

        indices[i] += 1
        for j := i + 1; j < r; j += 1 {
            indices[j] = indices[j-1] + 1
        }

        new_result := make([]*Person, r)
        copy(new_result, result)
        result = new_result
        for ; i < len(indices); i += 1 {
            result[i] = pool[indices[i]]
        }
        c <- result
    }
}
