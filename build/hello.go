package hello


// cmd
// GOOS=linux GOARCH=amd64 go tool compile -S hello.go
func hello(a int) int {
	c := a + 2
	return c
}

/**
// output

"".hello STEXT nosplit size=15 args=0x10 locals=0x0
        0x0000 00000 (hello.go:6)       TEXT    "".hello(SB), NOSPLIT|ABIInternal, $0-16
        0x0000 00000 (hello.go:6)       FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (hello.go:6)       FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (hello.go:6)       FUNCDATA        $2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (hello.go:7)       PCDATA  $0, $0
        0x0000 00000 (hello.go:7)       PCDATA  $1, $0
        0x0000 00000 (hello.go:7)       MOVQ    "".a+8(SP), AX
        0x0005 00005 (hello.go:7)       ADDQ    $2, AX
        0x0009 00009 (hello.go:8)       MOVQ    AX, "".~r1+16(SP)
        0x000e 00014 (hello.go:8)       RET
        0x0000 48 8b 44 24 08 48 83 c0 02 48 89 44 24 10 c3     H.D$.H...H.D$..
go.cuinfo.packagename. SDWARFINFO dupok size=0
        0x0000 68 65 6c 6c 6f                                   hello
go.loc."".hello SDWARFLOC size=102
        0x0000 ff ff ff ff ff ff ff ff 00 00 00 00 00 00 00 00  ................
        0x0010 09 00 00 00 00 00 00 00 0f 00 00 00 00 00 00 00  ................
        0x0020 01 00 50 00 00 00 00 00 00 00 00 00 00 00 00 00  ..P.............
        0x0030 00 00 00 ff ff ff ff ff ff ff ff 00 00 00 00 00  ................
        0x0040 00 00 00 00 00 00 00 00 00 00 00 0f 00 00 00 00  ................
        0x0050 00 00 00 01 00 9c 00 00 00 00 00 00 00 00 00 00  ................
        0x0060 00 00 00 00 00 00                                ......
        rel 8+8 t=1 "".hello+0
        rel 59+8 t=1 "".hello+0
go.info."".hello SDWARFINFO size=71
        0x0000 03 22 22 2e 68 65 6c 6c 6f 00 00 00 00 00 00 00  ."".hello.......
        0x0010 00 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00  ................
        0x0020 01 0b 63 00 07 00 00 00 00 00 00 00 00 10 61 00  ..c...........a.
        0x0030 00 06 00 00 00 00 00 00 00 00 0f 7e 72 31 00 01  ...........~r1..
        0x0040 06 00 00 00 00 00 00                             .......
        rel 10+8 t=1 "".hello+0
        rel 18+8 t=1 "".hello+15
        rel 28+4 t=29 gofile../Users/jinrong6/develop/sinawww/go-training/Golang_Puzzlers/src/puzzlers/build/hello.go+0
        rel 37+4 t=28 go.info.int+0
        rel 41+4 t=28 go.loc."".hello+0
        rel 50+4 t=28 go.info.int+0
        rel 54+4 t=28 go.loc."".hello+51
        rel 65+4 t=28 go.info.int+0
go.range."".hello SDWARFRANGE size=0
go.isstmt."".hello SDWARFMISC size=0
        0x0000 04 05 01 04 02 05 01 01 00                       .........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
        0x0000 01 00 00 00 00 00 00 00                          ........


 */