package error_handling

import "log"

func CheckErr(err error) {
        if err != nil {
                log.Fatal(err)
        }
}
