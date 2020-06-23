package pkg


var Id = 1

/*
>go tool compile -S assembly1.go
go.cuinfo.packagename. SDWARFINFO dupok size=0
        0x0000 70 6b 67                                         pkg
"".Id SNOPTRDATA size=8
        0x0000 df 74 01 00 00 00 00 00                          .t......

 */