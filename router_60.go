package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "1.2" )

func uuiD00cUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CdTlifppWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func THYKbWTuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SRjZDqCWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K7lLHbt7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qL9ZyalzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fsOztazCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ubGjcHjjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LDfTZ4A8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S0deufNxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NUnr8KTfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hCH5Wfp6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t7vVVyZMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y6mJ8OkEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Bi03hZRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3tDxQFS0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dtWUtktJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xRGHstEZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0uYTv7HSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9rpkseD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func scRpYg2YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ogoPh4kNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zG58bpReWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4VNE2wcuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dlsdTqA9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pejFjiqqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func omB3KtHLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sr7kzuqCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yg2uKxk4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i3sKu7fVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6MsuE9siWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DW3jqWXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FBKZgSNWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KaskrfDFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yMK4zp6QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func miGvh1X7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iOHC4HqIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d2y64hbbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z9Gegkk0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IK9JQ2nXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JUdkBJwUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EJ9oZPNsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BMBg7m2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Z9SbNZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dOPr1NvWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u7jD8HCBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x7hh3pDuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u0bCFE61Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5aHs6sUaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RRg9MjeHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KUxiUPEmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UuXj8a6FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vo4oVRrLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5x6tFNMRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gXtEDUzqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dPpaRpm4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5CnAhn4qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L8XaUaC0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DByxwmxQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rbUrsFuwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UpTiu0txWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mMqkz6raWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 19XmDHsFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QcgunYlTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VsPrWTEdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tf2ssgxNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KowUiSquWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9XuMlOXcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yNRyalAFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BViVI29KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f7NtSisiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 824oc8GIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Dp7kddgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KylAhw3MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MZcYVMUhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GYyqRQguWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tZu2NOEhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c2qkY5n5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UDZBPSgsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AX0Oq66VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DYHMkXiAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MTGLnXgxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func On1JWSqKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kr1ZUMDCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GgJFqwVcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aWJ3ggJFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bHibusJGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fZEXVSDxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VJ9pZZFDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3uYDVezUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 91KXSyz9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t6N3nLZGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DTzFsRVlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K6PAFf7qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SyoqWW8FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YJH8IldBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V7Fnbm49Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TXraEVs7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R37t32ssWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H6QRcNBwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l8BDEWyyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ke6U6Z05Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4l1sPmtFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func raXk4ahJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zQSsdzw3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5rcFG0EqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CjRhygkqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mx2kFmZMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UHcJpbP2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ecXPpJ34Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d4TqVzPkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EamuCrlgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bx9wwRYaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wjX48mAEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6nbAH2rCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pwIb1qqqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4SOP0YwUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZFqJs42sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CIblFS0cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V3p5GrxTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sIShRXTDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v8CaFfTsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IqrbPetgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X2WuFbSjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dCwtVFJkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func irFowRE4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AhiJWlx0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UzUBy1NYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func acJPrWnIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 24dzEu9VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5NgcNthLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xb9OeQL0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EpUZdfCoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VidD9AdZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 12qpR66xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WzdcEUlZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7gOowsZrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qd61yaezWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T18SEiJxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AYFaSyWrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DhsrZzY5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cdWzWlriWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 64CDhYpoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func llS3oJdgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4p1TgtCZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3dPxI7AiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xAiNRSYlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WcmOzMDIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bfMSbHtoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iEJTMNctWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vjkh7jUsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kE2XSshwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d0HFJQoqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MPb15QpHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wgt94HYGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v57Ez2VgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func enTMJ2EZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IHdhZoBwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func haELLSKPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WMCVictIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HXgNKEgLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lUrr49c8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GDUy9wdwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2fs8FafBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ve6JYNtZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KYAMljsFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZJ8LvS19Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XgtQJl9wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KZAOZzTbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4pNkcy38Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZLI9eIEkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xmwkRho8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JQOIIzEZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eQHgqG6yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JmXoF623Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mb0THbfCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cvowQwWoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ntnrHILKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qgi90vlfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RjPejZeLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OCCTt5UOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kqzygDOQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b6vlWkBXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SwFMakVyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rGwVwNJ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VxX6vm9JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BGGHlIQjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vjuZKZr9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IfpMF241Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wf3nWORKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Bt7COodWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2eXI8Kv8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wQxFvnLSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2oQKMM3pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func crVzFpksWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eW41jUG8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hbmT2TqZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DMO5JvDkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n63X9RSGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RNXnadPEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gT9mb9S8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 81kiMjMjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yjni36pUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7GnuVUAZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r0QCeRF6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tSeGnAk9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6E0ZeIiSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PF5OlgQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oqOPDpvMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UWZEPragWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func laavf3JnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tCHGq4RRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2VmP1NvFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func slWAgZ6sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cDBtzynQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wms1eGs5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZlTQsetvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g1nzLt7oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X9IRMaOWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RNIr2FYiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S6JqniLCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KFsjYQQ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func svVeTi9LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S7MpdheSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D2e1yxClWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ngBKdpVgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vTZFxYehWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n9fJr7y0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ta1h5de3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ydDj6MxxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9bcBxj2oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wv8WqV5CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func siBqgmaFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x4HYSsoWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w8stJ2jxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6smJ3LTgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ogh37jXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jwn4xXnSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3lcQI6RcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P4ZjSZVPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tGSt2aVXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TJ3TW2XoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vi0GiFDHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R7wf5IY3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LKE5KLxfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NA29E40eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1y2ql7arWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e72AjxleWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BJp4YjMXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ErdiNREHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AHtSUb0iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YIE5zDLLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GAzMSz4pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UqTXxXHjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HTFXZcFhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zo0DytklWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rxXoLeAdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9R7KEVJ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AnQ5InozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wJwZWKTsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VqQGkh3SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qylr22mnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oOELJ8w0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N3A77a1QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WMZzzh4QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KQpsE6sZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dDCaAZhGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ifluwgfzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pW8uMVl1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZZxhsOKGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jo0hXISPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a6UYVNEFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7l6IyC4eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CxG5MNTpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kk8Yqw1KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zWIUVoBjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CAOFmFdPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C2Lm9zvZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5sg4wqTsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x3djm7gYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IDnZuk1SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TD2BxmsVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VYY1ElSeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NqYLsxADWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7NHTQfrGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hDM3RODmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IytLI0K8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h7eX22yqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M7XKXfe6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lx32CppcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t9zUEpZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UEzpsaPMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2jgczdazWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ok0KtZG9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rsBMD600Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5jEzdsfzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TnpFu1q2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qq72ltUpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RshSjEsOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UUjokYxQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pWbGWnbeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bAVYLheBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ncYMkwHYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KJtAcbPfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xu2V5fONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qLYW8UEAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HyUk5hZRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0X02CUiIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bjDX8Lk2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lheTwjYwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EMpsUlaXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0QYDqZTQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k90kWZn5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hlXFGLFiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yJQ5uJy3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XXZ3CdaXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YdFGmvFXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6GdMRMC1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lqkng3HVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4OPuqm1DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3xr59HSAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WHOPq7FgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CiMa1JGrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xqgaLMLSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q3MbVrYwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dOoRXbYqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0M2Ekm1AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pIf67IaZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ttFoyH3pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 86gsfHPOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mgthuYOnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LLuMO2WqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7CTrdoL4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jaBnPS1ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DsBmihOsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2AKEOb8kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DImXDyXJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yRILDLWNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OnF891fMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tg65r9KkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NZvHatJaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ycV8ne9iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SUzlh1naWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IP6oFu2xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5wmw1yzzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aJ7bCuyzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FBnIUUnuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mSvnHFEMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func itcrgfRGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z9O7YbgqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CGwYlaLuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TJEcVTijWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 08DKEOb5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZDm8x5ZyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9r0b95HMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ImZVm047Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DRdG4g17Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0H87smQaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ozSJxlQ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MMhjwvggWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pVowvrSjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mItf4mDIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Omx81HNqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iS7A4YHJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XtpPn480Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3EuxwQ2tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CSwovxwRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 41g4otAsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nwzs9UtSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7DP9gj64Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zyU44AfNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tJz1qZu4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LoQHW6YSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tM6dMFo0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YX5dv7sSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w9dDmLV6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 90hCUulFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mEVgCyAyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8TFRot9JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fO0R27nVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ZNDOrKTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TKDXM1OnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0nCcDQLQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GQDScEV4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zjJbQJfLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZB4YHBf9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FdxScMFmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XV1ZikaPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zq0jaAGkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yfOxBq4TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6lkSCPxZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G6miOZOqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R4X5ZcNzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LgCm9ENBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KVTvYpcNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7dtSxz3DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zte1ggqGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wljO5unTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gQkjvCBKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8V8IMJd1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o0G5ruoKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qe8I0JDxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fDsds7iKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TkTzERcbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zXrDCA0RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8zkNOWyKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bwU0E5vUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 45fMM3qfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hSGln3WkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iaYqSTayWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4A7pxNaCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8VLPV0PiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kxOzOF7JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h8L7quCQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4tdDWlDnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 08mcMlYyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3uLnAB9JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EWK9Qa9gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CsGVNQNZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VgGBn3t8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lzAWPGtZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tRgGFUFAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7f62SRxvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ONxLGK5kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VB9iEPsPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Agel80DhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GGl2PuIpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OM1dRVoNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BG97hBQRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iMdulO1UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tPFnXFXRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LOcvKPfVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rM0o3ne3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oxViQFxTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JZTfyD6zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H7nCXMtTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SJPUgz56Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GzAcAjk5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SwzenKuHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GZkIEbkVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XfUVBMkzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uuQbdu2aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qA1bFbOoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xAnVJ1S0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OzJFZivhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gBgA3ZiWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fm3NwzXuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 61JeFfESWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K9iXrLV0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uRpnSq6vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iAGoJE5UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SNZz4i32Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xr7Wq4BSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M257bYVXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yHiEN4ZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bpAruRPbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iA97UdMvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8FAagREbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UxPoN6YzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P32Njhk4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vgJunyvvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8M9MpHMcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9uTclAPmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oeC2GqBHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4yH89W0KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xhWvYnqmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yNqKbm42Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uOTdx4wgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cNhmxJaEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w5mmPbN6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JGG9WeQzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a393rkgpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GhL9ty69Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oz9sS0ZxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K7MQC6mAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2okslgMvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kh8CpaDDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xoVX8iu9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kP8y8K1eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ipskHzgRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lhoCPOhpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bblEWeszWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ay2CWSMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VoArEHamWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PB4KlgnvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n1RxGNMSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GvPd75CoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 85EadL3EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iwZfWOr0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2wFeo2JtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RIuzvqngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1QoHAKvgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func orrBsWwyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func csvaCImRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M1GaqmTYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TkjXTsA8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n6Wo6npmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3mc6smLLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J45TROoSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8SKzcl29Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SjYLRDxIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eanuD8oXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ihy4AplvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZizFTK0OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FFNBDSzPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tE16OjijWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s8ZbIjpXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ccwMRUnFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nllWCvDkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2G1orfSiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gdPDOzIyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hfXucCDIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Ma7ZGPfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tu0PGLQKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dIVeJHSJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wjFKP7g0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aL2sKJDDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Abd0zNqjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UePGaJWMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func esQ0fAXjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bOJWTcuzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MXKkevPnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p0rFfbyaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WyjxoQ1zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8NxLW3mWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lFbcLdppWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rkeMRx0XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DagioO8vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nnrmBoyCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ho7a0ka6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j81eIWmyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pdxjrqyeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aGFlbrmNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KlcoHP6gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sQ68hSaTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EkZDKi2JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I6d85mFuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wmkbSy0OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4uF5YvQpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2DgpOJ9nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HQmsM5MbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gbgZ7tTEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 018LfMWpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uQNj7zipWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RLKChXtYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IpVTSwRrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sYhX2naNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0vVjc5FOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 72uy21DXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0HdL8k8pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zd41TpbNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aFn4i96lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gcwdbQfMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rjOTtrQIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VzL7Ve77Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Mv6rBUpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IY4Alb9DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iOIR1BJjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YhiQSqnOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vTOgrCF7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I94EiPMNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func as8EAebSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ljBV1h2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z1vxEvAKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func li6R64IjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y4IUGWIyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PtSA6dgBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AwgZnA7wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2IWACWImWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ui2nHJDuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7JagZpwVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B36ectQFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ldNCbNhRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kWMvk4taWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8YxedMO0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pnbRrmKwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kvxf2TPcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t57QMIZgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6McWauu3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func haKijKRfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lbNb61O5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4lMjvgYuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VQA7TKlYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LXlqKquYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func suB4Wq2JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JzF980O6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WBf58YSuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pxOuOF0MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LKDe7ZoyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WWahKyJDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IqkERICSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QC9xHgRuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jwke0BWsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func REsJJlPXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4fvisZ4wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VEUrwgjWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iBpLiV5gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u7QspSiiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 638R0IdyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q9zRnBzlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HzroAVlwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mVi8qfs2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NzaGDaEaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a8HDopDbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func riboqMtyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func td8gxlVlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MbLLpVf9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gllSPxvmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pTcvdyMbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9hu4EeQAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HdsCDQO6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xt5glOlYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g9gK6xPzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u5Mbyu8KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6fOGzEEKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zz2Be0OaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gSNnUDKUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wDwLanseWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eixY4qvnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GpOuJpYvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fC7NXgAdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 78ufU9M4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q7AfIBrTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R9v7kx8eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0LgVSmpSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6YOvIyA9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YPST4EtCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NmNMHXdfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BmPCMygiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SWmGsqprWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2GA1zvGQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IqsWC2awWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bibYMXsvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fedF4LSsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 41tVe440Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hYsXQlAWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yBSXqwOsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mkir5uv9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 13yJDTjqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aVVQ42WCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hHLyn8IKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1gUW9J82Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fJvPTh1jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 76ZALRNjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6EPJMQF5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R0CWHbEYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SAdjKCfXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AQxsQ73BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qHyHwM0lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x4wgtcbcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dkxMhMhQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RYM4LMrLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9zq0cHjtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3oOGjkbqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J9bpVOKTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kszdlLAdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gCL7wP1JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mTS13838Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l7Dc11j1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wvweUm1AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PSRSTTu8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b02AZW0xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lkf08gmwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YagbqqaVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func egNXzXQCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aylXZ2iZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hdzP0qW9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RmSGeWgqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OYpQkoQMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3KgJLNOeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sP10A1VyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jg0HSWo9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OBd34uwUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W9yBgRVwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DiHFlQcPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g5fV9YWAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

