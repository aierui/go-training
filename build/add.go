package main

func myFunction(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	myFunction(66, 77)
}

//

// cmd
// go tool compile -S -N -l add.go
/**
"".myFunction STEXT nosplit size=49 args=0x20 locals=0x0
        0x0000 00000 (add.go:3) TEXT    "".myFunction(SB), NOSPLIT|ABIInternal, $0-32
        0x0000 00000 (add.go:3) FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (add.go:3) FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (add.go:3) FUNCDATA        $2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (add.go:3) PCDATA  $0, $0
        0x0000 00000 (add.go:3) PCDATA  $1, $0
        0x0000 00000 (add.go:3) MOVQ    $0, "".~r2+24(SP) 		// 初始化第一个返回值
        0x0009 00009 (add.go:3) MOVQ    $0, "".~r3+32(SP) 		// 初始化第二个返回值
        0x0012 00018 (add.go:4) MOVQ    "".a+8(SP), AX 			// AX = 66
        0x0017 00023 (add.go:4) ADDQ    "".b+16(SP), AX 		// AX = AX + 77 = 143
        0x001c 00028 (add.go:4) MOVQ    AX, "".~r2+24(SP) 		// (24)SP = AX = 143
        0x0021 00033 (add.go:4) MOVQ    "".a+8(SP), AX 			// AX = 66
        0x0026 00038 (add.go:4) SUBQ    "".b+16(SP), AX  		// AX = AX - 77 = -11
        0x002b 00043 (add.go:4) MOVQ    AX, "".~r3+32(SP) 		// (32)SP = AX = -11
        0x0030 00048 (add.go:4) RET
        0x0000 48 c7 44 24 18 00 00 00 00 48 c7 44 24 20 00 00  H.D$.....H.D$ ..
        0x0010 00 00 48 8b 44 24 08 48 03 44 24 10 48 89 44 24  ..H.D$.H.D$.H.D$
        0x0020 18 48 8b 44 24 08 48 2b 44 24 10 48 89 44 24 20  .H.D$.H+D$.H.D$
        0x0030 c3                                               .
"".main STEXT size=68 args=0x0 locals=0x28
        0x0000 00000 (add.go:7) TEXT    "".main(SB), ABIInternal, $40-0
        0x0000 00000 (add.go:7) MOVQ    (TLS), CX
        0x0009 00009 (add.go:7) CMPQ    SP, 16(CX)
        0x000d 00013 (add.go:7) JLS     61
        0x000f 00015 (add.go:7) SUBQ    $40, SP
        0x0013 00019 (add.go:7) MOVQ    BP, 32(SP)
        0x0018 00024 (add.go:7) LEAQ    32(SP), BP
        0x001d 00029 (add.go:7) FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x001d 00029 (add.go:7) FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x001d 00029 (add.go:7) FUNCDATA        $2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x001d 00029 (add.go:8) PCDATA  $0, $0
        0x001d 00029 (add.go:8) PCDATA  $1, $0
        0x001d 00029 (add.go:8) MOVQ    $66, (SP)
        0x0025 00037 (add.go:8) MOVQ    $77, 8(SP)
        0x002e 00046 (add.go:8) CALL    "".myFunction(SB)
        0x0033 00051 (add.go:9) MOVQ    32(SP), BP
        0x0038 00056 (add.go:9) ADDQ    $40, SP
        0x003c 00060 (add.go:9) RET
        0x003d 00061 (add.go:9) NOP
        0x003d 00061 (add.go:7) PCDATA  $1, $-1
        0x003d 00061 (add.go:7) PCDATA  $0, $-1
        0x003d 00061 (add.go:7) CALL    runtime.morestack_noctxt(SB)
        0x0042 00066 (add.go:7) JMP     0
        0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 2e 48  eH..%....H;a.v.H
        0x0010 83 ec 28 48 89 6c 24 20 48 8d 6c 24 20 48 c7 04  ..(H.l$ H.l$ H..
        0x0020 24 42 00 00 00 48 c7 44 24 08 4d 00 00 00 e8 00  $B...H.D$.M.....
        0x0030 00 00 00 48 8b 6c 24 20 48 83 c4 28 c3 e8 00 00  ...H.l$ H..(....
        0x0040 00 00 eb bc                                      ....
        rel 5+4 t=16 TLS+0
        rel 47+4 t=8 "".myFunction+0
        rel 62+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFINFO dupok size=0
        0x0000 6d 61 69 6e                                      main
go.loc."".myFunction SDWARFLOC size=0
go.info."".myFunction SDWARFINFO size=90
        0x0000 03 22 22 2e 6d 79 46 75 6e 63 74 69 6f 6e 00 00  ."".myFunction..
        0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01  ................
        0x0020 9c 00 00 00 00 01 0f 61 00 00 03 00 00 00 00 01  .......a........
        0x0030 9c 0f 62 00 00 03 00 00 00 00 02 91 08 0f 7e 72  ..b...........~r
        0x0040 32 00 01 03 00 00 00 00 02 91 10 0f 7e 72 33 00  2...........~r3.
        0x0050 01 03 00 00 00 00 02 91 18 00                    ..........
        rel 15+8 t=1 "".myFunction+0
        rel 23+8 t=1 "".myFunction+49
        rel 33+4 t=29 gofile../Users/jinrong6/develop/sinawww/go-training/Golang_Puzzlers/src/puzzlers/build/add.go+0
        rel 43+4 t=28 go.info.int+0
        rel 54+4 t=28 go.info.int+0
        rel 68+4 t=28 go.info.int+0
        rel 82+4 t=28 go.info.int+0
go.range."".myFunction SDWARFRANGE size=0
go.isstmt."".myFunction SDWARFMISC size=0
        0x0000 04 09 01 09 02 05 01 1a 00                       .........
go.loc."".main SDWARFLOC size=0
go.info."".main SDWARFINFO size=33
        0x0000 03 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
        0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
        0x0020 00                                               .
        rel 9+8 t=1 "".main+0
        rel 17+8 t=1 "".main+68
        rel 27+4 t=29 gofile../Users/jinrong6/develop/sinawww/go-training/Golang_Puzzlers/src/puzzlers/build/add.go+0
go.range."".main SDWARFRANGE size=0
go.isstmt."".main SDWARFMISC size=0
        0x0000 04 0f 04 0e 03 08 01 09 02 16 00                 ...........
""..inittask SNOPTRDATA size=24
        0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0010 00 00 00 00 00 00 00 00                          ........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
        0x0000 01 00 00 00 00 00 00 00                          ........
*/