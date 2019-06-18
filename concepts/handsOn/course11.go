package main

import (
        "encoding/json"
        "fmt"
        "log"
        "errors"
)

type person struct {
        First string
        Last string
        Sayings []string
}

type sqrtError struct {
        lat string
        long string
        err error
}

type customErr struct {
        info string
}

func (ce customErr) Error() string {
        return fmt.Sprintf("This is my custom Error: %v\n", ce.info)
}

func testIt(e error) {
        fmt.Println("This ran, ", e, "\n")
}

func(se sqrtError) Error() string {
        return fmt.Sprintf("math error: %v, %v, %v\n", se.lat, se.long, se.err)
}

func sqrt(f float64) (float64, error) {
        if f < 0 {
                e := errors.New("you can't do that man...!")
                return 0, sqrtError{
                                lat: "lat",
                                long: "long",
                                err: e,
                        }
        }
        return 42, nil
}

func main() {
        p1 := person{
                First: "Joe",
                Last: "Barrows",
                Sayings: []string{"one more please","do you bring enough for everyone?","I'm not picky"},
        }

        bs, err := toJSON(p1)
        if err != nil {
                log.Panic("This is my error: ", err)
        }
        fmt.Println(string(bs))


        c1 := customErr{
                info: "Need more coffee",
        }

        testIt(c1)

        _, err = sqrt(-10)
        if err != nil {
                log.Println("This is my error: ", err)
        }

}

func toJSON(a interface{}) ([]byte, error) {
        bs, err := json.Marshal(a)
        if err != nil {
                return []byte{}, fmt.Errorf("Problem with toJSON...%v\n", err)
        }
        return bs, nil
}
