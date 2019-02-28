package goluckies

import (
        "math/rand"
        "time"
)

func RandInts(min int, max int, perms int) []int {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        p := r.Perm(max - min + 1)
        for _, v := range p {
                if v == 0 {
                        v +=1
                }
        }
        return p[:perms]
}
