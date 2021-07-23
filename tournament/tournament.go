// package tournament implements a routine that tallys the 
// result of games played in a tournament
package tournament

import "strings"
import "io"
import "io/ioutil"
import "fmt"
import "sort"

type GameRecord struct {
    name string
    matchesPlayed int
    won int
    draw int
    loss int
    points int
}

// Tally takes input from r and writes output to w.
func Tally(r io.Reader, w io.Writer) error {
    tally := map[string]*GameRecord{}
    gRec := []*GameRecord{}
    f, err := ioutil.ReadAll(r)
    n := 0
    for _, rec := range strings.Split(string(f),"\n") {
        x := strings.Split(rec,";")
        if len(x) == 3 {
            if tally[x[0]] == nil {
                gRec = append(gRec, &GameRecord{name: x[0], matchesPlayed:0, won:0, draw:0, loss:0, points:0})
                tally[x[0]] = gRec[n]
                n++
            }
            if tally[x[1]] == nil {
                gRec = append(gRec, &GameRecord{name: x[1], matchesPlayed:0, won:0, draw:0, loss:0, points:0})
                tally[x[1]] = gRec[n]
                n++
            }
            tally[x[0]].matchesPlayed += 1
            tally[x[1]].matchesPlayed += 1
            if x[2] == "win" {
                tally[x[0]].won += 1
                tally[x[1]].loss += 1
            } else if x[2] == "loss" {
                tally[x[0]].loss += 1
                tally[x[1]].won += 1
            } else if x[2] == "draw" {
                tally[x[0]].draw += 1
                tally[x[1]].draw += 1
            } else {
                return fmt.Errorf("Tally: unknown result")
            }
        } else if len(x) > 1 {
            return fmt.Errorf("Tally: incorrect record format")
        }
    }
    for k, _ := range tally {
        tally[k].points = 3 * tally[k].won + tally[k].draw
    }
    sort.Slice(gRec, func(i, j int) bool {
        if gRec[i].points == gRec[j].points {
            return gRec[i].name < gRec[j].name
        }
        return gRec[i].points > gRec[j].points
    })
    w.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))
    for _, rec := range gRec {
        s := fmt.Sprintf("%-30s | %2d | %2d | %2d | %2d | %2d\n", rec.name, rec.matchesPlayed, rec.won, rec.draw, rec.loss, rec.points)
        w.Write([]byte(s))
    }

    return err
}
