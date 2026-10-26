package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "5.3" )

func oSA1jINrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 21Q31gp5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t5pCRjv7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3EjW72OdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3LgWxmnqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DmQmCy5KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dhPqNMfiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jjjVueDLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W4BK7s3cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yQfEdicaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8e0OXQC0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5B19nHRiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qLLj0OLdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vCHfFhfXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qzCdu4VeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i1Tl2QsMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func honZraSHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6IXtVWlSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m9Idp4Q5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yxr1DGOGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l2FFJmoHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JEtooOkOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e8sLacZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qi3WPuJfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mx7RJSkAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vckJrqOLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DVDm8I9BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bD5oPmSjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xZmy2PgoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zr7Suw38Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I5OU8AHKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nxMcDVFNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hUTWwXTSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Beta6BMwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9add1DhNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Qs9VkGtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xpSkMKGYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ies1zyTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tI7TQnGZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0p8UhMWZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I0I4psiFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wyvHL2jyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bAk4MIZtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MR2H4LUWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zzKmmfmLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dF496SHVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B0x8otplWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lJrr92YdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q2BrYqtjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uU2dZ8S4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5voQiGvpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GUdEflx8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mx1AMYf6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nlWVPL7jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U7jpiOgFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YjTKG0rNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ozbdIoH2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M7tpMS6iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hWMQ736IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fcnKYk2pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OiPn7SrwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vLxiP49jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QP0LCn6MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gkcA4hNHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LCyBDo2RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XgrklLbyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AEPAEX3cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9suB0xwRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0y6rPobnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JwP2jIG6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func THTcehn5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i678Lon2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func evswTX0BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xJKqHn0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U80G15NKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jliNuesUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G3BNHOaDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w0QxkD3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sp9ptcV1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KjfZbFlvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PsTn9yR4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rXrp8rNGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u8qYnnT3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XmxlYytRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u9wjLd3oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EGHT1CAHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3sMVw4jPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uWvqFnktWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yYeWP4vrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pnz7D9s4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qkEghs79Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xQzz1ixWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xT7lASTAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hsvzlTvtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R15S7FuuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pp3uPb41Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vbWfmDY6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pkx9OKIaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L7UAxxFdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MqHLUPNYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DFfh1xB4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tO0CJxMDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VOWClYOBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0WuoY5RxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3g5YxgkqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eV17EK8PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1nQ76XTGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ms6WjJLFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9fBHcjZnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func glQOAdu3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 19s3DazLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rD7SFzTbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SrHPSqlcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5DApGOGpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EMdyGaFxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mhg9NfoEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uXfrJsQ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lzADprv2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

