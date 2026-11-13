package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "3.7" )

func ko7CnU5ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PnIqFv8NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M7LevtYeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func whQYEcydWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7FT0TBnYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SCXzjpaDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jecRyXq0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZMVB4dFEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G1g4n2pNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DpmS6alhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DQlj2zWoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5SAiluQ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XqVANDcLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hx0HPzk3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rVkASBMfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jJRJ8yMxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OdeliihuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lvn8xlL7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JJk0mQrsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rCMNEjAzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Od8zs0WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Mi6e9LPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DTeYR7mcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eQgEEXfwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d5qDlT8XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w0jJPMEWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k75R0t45Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 56yaOWP5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vm0hSPDPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1qYivtsbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AsnRJxXTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8kG4BLF9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nBYmwMMwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oNIQkdIPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WEB3jE2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func chLgjZatWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CxenFOfcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qCN3vud8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TDUyuhEMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d4gmdY4oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jus43lbDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1z07FECSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T9HG0mAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4EaBIGA7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1hDcj1GAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DfxQcRW3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qb1Q22ZqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func txGo0T2LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZJx01DSBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3uB5SKQCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gEEQcNVRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VFkaet9WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3BEfhCBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EbsfPmoJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A9UdcKtMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wOftMHzqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZYIY4AaIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d12fhW91Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fmjHi1jvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xAXrbP5CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p2xQILFzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LNgnD1flWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FrSbq1lgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wjLScrHlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ewVHRGZYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S0FsUNsbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SaFyXoDtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mf8LTCRnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eexff6uTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9gfy7Ig7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oIySeiY3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5XrXDcZWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v8o2ciRhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2yZ8d1oBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jduAmccNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nSWPybwIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sXP7xhY0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hqRkYYLxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S86c0olqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ybIyO1kWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vCNjRKzlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wvItVlJJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dtVz4kdpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pr7xwszGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xOfeiDCBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gk1Upo6XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LGzsBzP0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bTrmNXrYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qh7qkBNiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nFCprIbRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C83WyusPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Mqku1dyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 83kf9k5VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eRQ2uRNSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dxu5jVe5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CFV3yOsXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JSTNI7WFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eJOxFlbnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ra113CwBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xgr1Ik2rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0jP1fvtZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cVUPfDfGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v9ZONc91Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3SVnTng4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8dQswd2pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2DOVM6VkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GLhTIo2tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O7EUryAdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wg7OJeyCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NmQyj2fXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w28LzNhZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AljJGIX1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SfEvV6arWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bmrww8S3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LM7qzFp8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FghKUYasWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Suaq1VWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zZYPNA1nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R95DmKOTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ydv6xng4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ikAkxivTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2f0hMCfJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I3I6aTS6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wJpe2BLsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0YCtulNvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mobOQkmzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ijbuopTsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 14MLPhDNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1jUrovtjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iFwvgCb8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PUSOWs1XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0VM57ZMXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pFN8mEUMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bxTXuaDKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wHitLalbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lbjq2SzcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4iLmpAT2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c1tCd1a5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CY5GXRj1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bbti2IZ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ebxM9OngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9zTq4eK0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eRlCc3FwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vodWSGWeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yXqTHoe9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DW0KDB6QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZsNLwY91Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U08tvXjyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wx93ihd6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lwzgc7NuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ki9herUoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ni6UQhHLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7GiaaAe3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9DhDVbinWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2TirUbxEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1EuFHSu0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ek09fr5LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8VTr6Qe5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vBNVcPmVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iYFH8251Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oHQIGnYNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xfv3rpjjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S8NIpcrXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ctJ4uIauWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xGEDZyhhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8PWGZShUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qx8qSinHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L5c6e2t6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iFOou9WJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZWmGIAISWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oYJh0XMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y9OWrfeMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XlwlOmqyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zhoPIaXbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bLYV0UHDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YpaeS5QiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jPFxFbWuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0kVfIE4qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 43JUzpGZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bzBrNqHRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MnEINcGEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 25n6wtocWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yX5yoWmQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0GWpcPYGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iJcQh5VMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mjZnZL2eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 34XrUjabWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yaakghBwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XDfwYx58Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FNQRBdTdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hUDCfYlQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yCN7pXlhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MfQmYgX4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VBVnoWNGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UiK8ZLbTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HWmW6mzjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RK7mYLdtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bGovvF0VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n1scchtXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7aL11xxKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 16RR0foTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EiXUT1qGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C8c3P5GwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nnLMV09nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RMkEzENeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GBQ6Iq6hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OlvpOv6DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xnkWPOACWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d7IDOYIrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b8OvdyS3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UPTQTTWvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Swe2pwAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4NnRVFOjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lrHqi40DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9RAfw5oQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FrUvSe0pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GzuwRJaLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G7fecHslWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yC9acaXFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HtryHCLAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wG8Z5dONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DWSU84JsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kirQokTZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PbeGpcXhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NeoRv8dHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AfCMCtljWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3lDTWZjvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gj4Iy2yaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8YXhElkHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rj6UxaPwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TeBIbU9vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4PUQaCc6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UJJO2VLlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MA0mGP23Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2LpgYrmMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MfKdDr77Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sTMoQPEFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func myTjUJ5nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bpV5lIW3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Knpou6XJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 44i7UVNWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LJsjHgTwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BMmQZ0Q6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hQrUYS4WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MfodI3P2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wdHg5DVKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AU2y4ZGXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EgbF3v2YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func clv2bXP0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0FsDkJTFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o1GtW4PyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func klwPO2drWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t38ETli0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ptmLA0L5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p8xrHVJbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tPuAyW7qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 01y41C03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LjejQSQ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DE65bMsXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jVlAP7HdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LS7PvcJkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OOItQvj0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yn9yyriQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JqPNhyRaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KE1IlailWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NNPRdbtRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gUyMzhIgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4tAYmuhhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0DFiD2lxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sqfA3lRjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fdz5TImXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CgeQgWKRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6DEkVykDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oo5ZXpHJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IkJAG49hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YwEyfQK8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FVJxtD87Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9a49QeJmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KJ6w1LHMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3JBJfqLzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lMwy8VJhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SrgqTsPeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IamVgPkrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MhNITddfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HVl3dkveWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ylZN9qCGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q8vsYlWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ltraK2g9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 67GuvhqqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nZ6ajO46Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dZFQHNESWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RGewOgO6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rrEGjguoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TZVWUouEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eeLLxqPDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DRV1JZMTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LrEdX8S3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bHv92DHnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R0BUHw6VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2YFybQcVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qw1ZyFtxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sA74DC3pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k1WECsbbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2OSnfqj7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vykojqSmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XVKdry63Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KnYOWDC9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hXEGFQAmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G6uUeOYcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1dJJGEX3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NLVB9vNMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zLEkdxPOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dVuHGjTtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vc1VWENCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nzRe1WAAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tSgwayJ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jIo995KPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SieYtQ8tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HXuSth9xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VCGl8GkfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QTkiEXtaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JQOawY82Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

