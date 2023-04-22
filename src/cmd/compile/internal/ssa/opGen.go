// Code generated from _gen/*Ops.go using 'go generate'; DO NOT EDIT.

package ssa

import (
	"cmd/internal/obj"
	"cmd/internal/obj/x86"
)

const (
	BlockInvalid BlockKind = iota

	BlockAMD64EQ
	BlockAMD64NE
	BlockAMD64LT
	BlockAMD64LE
	BlockAMD64GT
	BlockAMD64GE
	BlockAMD64OS
	BlockAMD64OC
	BlockAMD64ULT
	BlockAMD64ULE
	BlockAMD64UGT
	BlockAMD64UGE
	BlockAMD64EQF
	BlockAMD64NEF
	BlockAMD64ORD
	BlockAMD64NAN
	BlockAMD64JUMPTABLE

	BlockPlain
	BlockIf
	BlockDefer
	BlockRet
	BlockRetJmp
	BlockExit
	BlockJumpTable
	BlockFirst
)

var blockString = [...]string{
	BlockInvalid: "BlockInvalid",

	BlockAMD64EQ:        "EQ",
	BlockAMD64NE:        "NE",
	BlockAMD64LT:        "LT",
	BlockAMD64LE:        "LE",
	BlockAMD64GT:        "GT",
	BlockAMD64GE:        "GE",
	BlockAMD64OS:        "OS",
	BlockAMD64OC:        "OC",
	BlockAMD64ULT:       "ULT",
	BlockAMD64ULE:       "ULE",
	BlockAMD64UGT:       "UGT",
	BlockAMD64UGE:       "UGE",
	BlockAMD64EQF:       "EQF",
	BlockAMD64NEF:       "NEF",
	BlockAMD64ORD:       "ORD",
	BlockAMD64NAN:       "NAN",
	BlockAMD64JUMPTABLE: "JUMPTABLE",

	BlockPlain:     "Plain",
	BlockIf:        "If",
	BlockDefer:     "Defer",
	BlockRet:       "Ret",
	BlockRetJmp:    "RetJmp",
	BlockExit:      "Exit",
	BlockJumpTable: "JumpTable",
	BlockFirst:     "First",
}

func (k BlockKind) String() string { return blockString[k] }
func (k BlockKind) AuxIntType() string {
	switch k {
	}
	return ""
}

const (
	OpInvalid Op = iota

	OpAMD64ADDSS
	OpAMD64ADDSD
	OpAMD64SUBSS
	OpAMD64SUBSD
	OpAMD64MULSS
	OpAMD64MULSD
	OpAMD64DIVSS
	OpAMD64DIVSD
	OpAMD64MOVSSload
	OpAMD64MOVSDload
	OpAMD64MOVSSconst
	OpAMD64MOVSDconst
	OpAMD64MOVSSloadidx1
	OpAMD64MOVSSloadidx4
	OpAMD64MOVSDloadidx1
	OpAMD64MOVSDloadidx8
	OpAMD64MOVSSstore
	OpAMD64MOVSDstore
	OpAMD64MOVSSstoreidx1
	OpAMD64MOVSSstoreidx4
	OpAMD64MOVSDstoreidx1
	OpAMD64MOVSDstoreidx8
	OpAMD64ADDSSload
	OpAMD64ADDSDload
	OpAMD64SUBSSload
	OpAMD64SUBSDload
	OpAMD64MULSSload
	OpAMD64MULSDload
	OpAMD64DIVSSload
	OpAMD64DIVSDload
	OpAMD64ADDSSloadidx1
	OpAMD64ADDSSloadidx4
	OpAMD64ADDSDloadidx1
	OpAMD64ADDSDloadidx8
	OpAMD64SUBSSloadidx1
	OpAMD64SUBSSloadidx4
	OpAMD64SUBSDloadidx1
	OpAMD64SUBSDloadidx8
	OpAMD64MULSSloadidx1
	OpAMD64MULSSloadidx4
	OpAMD64MULSDloadidx1
	OpAMD64MULSDloadidx8
	OpAMD64DIVSSloadidx1
	OpAMD64DIVSSloadidx4
	OpAMD64DIVSDloadidx1
	OpAMD64DIVSDloadidx8
	OpAMD64ADDQ
	OpAMD64ADDL
	OpAMD64ADDQconst
	OpAMD64ADDLconst
	OpAMD64ADDQconstmodify
	OpAMD64ADDLconstmodify
	OpAMD64SUBQ
	OpAMD64SUBL
	OpAMD64SUBQconst
	OpAMD64SUBLconst
	OpAMD64MULQ
	OpAMD64MULL
	OpAMD64MULQconst
	OpAMD64MULLconst
	OpAMD64MULLU
	OpAMD64MULQU
	OpAMD64HMULQ
	OpAMD64HMULL
	OpAMD64HMULQU
	OpAMD64HMULLU
	OpAMD64AVGQU
	OpAMD64DIVQ
	OpAMD64DIVL
	OpAMD64DIVW
	OpAMD64DIVQU
	OpAMD64DIVLU
	OpAMD64DIVWU
	OpAMD64NEGLflags
	OpAMD64ADDQcarry
	OpAMD64ADCQ
	OpAMD64ADDQconstcarry
	OpAMD64ADCQconst
	OpAMD64SUBQborrow
	OpAMD64SBBQ
	OpAMD64SUBQconstborrow
	OpAMD64SBBQconst
	OpAMD64MULQU2
	OpAMD64DIVQU2
	OpAMD64ANDQ
	OpAMD64ANDL
	OpAMD64ANDQconst
	OpAMD64ANDLconst
	OpAMD64ANDQconstmodify
	OpAMD64ANDLconstmodify
	OpAMD64ORQ
	OpAMD64ORL
	OpAMD64ORQconst
	OpAMD64ORLconst
	OpAMD64ORQconstmodify
	OpAMD64ORLconstmodify
	OpAMD64XORQ
	OpAMD64XORL
	OpAMD64XORQconst
	OpAMD64XORLconst
	OpAMD64XORQconstmodify
	OpAMD64XORLconstmodify
	OpAMD64CMPQ
	OpAMD64CMPL
	OpAMD64CMPW
	OpAMD64CMPB
	OpAMD64CMPQconst
	OpAMD64CMPLconst
	OpAMD64CMPWconst
	OpAMD64CMPBconst
	OpAMD64CMPQload
	OpAMD64CMPLload
	OpAMD64CMPWload
	OpAMD64CMPBload
	OpAMD64CMPQconstload
	OpAMD64CMPLconstload
	OpAMD64CMPWconstload
	OpAMD64CMPBconstload
	OpAMD64CMPQloadidx8
	OpAMD64CMPQloadidx1
	OpAMD64CMPLloadidx4
	OpAMD64CMPLloadidx1
	OpAMD64CMPWloadidx2
	OpAMD64CMPWloadidx1
	OpAMD64CMPBloadidx1
	OpAMD64CMPQconstloadidx8
	OpAMD64CMPQconstloadidx1
	OpAMD64CMPLconstloadidx4
	OpAMD64CMPLconstloadidx1
	OpAMD64CMPWconstloadidx2
	OpAMD64CMPWconstloadidx1
	OpAMD64CMPBconstloadidx1
	OpAMD64UCOMISS
	OpAMD64UCOMISD
	OpAMD64BTL
	OpAMD64BTQ
	OpAMD64BTCL
	OpAMD64BTCQ
	OpAMD64BTRL
	OpAMD64BTRQ
	OpAMD64BTSL
	OpAMD64BTSQ
	OpAMD64BTLconst
	OpAMD64BTQconst
	OpAMD64BTCLconst
	OpAMD64BTCQconst
	OpAMD64BTRLconst
	OpAMD64BTRQconst
	OpAMD64BTSLconst
	OpAMD64BTSQconst
	OpAMD64TESTQ
	OpAMD64TESTL
	OpAMD64TESTW
	OpAMD64TESTB
	OpAMD64TESTQconst
	OpAMD64TESTLconst
	OpAMD64TESTWconst
	OpAMD64TESTBconst
	OpAMD64SHLQ
	OpAMD64SHLL
	OpAMD64SHLQconst
	OpAMD64SHLLconst
	OpAMD64SHRQ
	OpAMD64SHRL
	OpAMD64SHRW
	OpAMD64SHRB
	OpAMD64SHRQconst
	OpAMD64SHRLconst
	OpAMD64SHRWconst
	OpAMD64SHRBconst
	OpAMD64SARQ
	OpAMD64SARL
	OpAMD64SARW
	OpAMD64SARB
	OpAMD64SARQconst
	OpAMD64SARLconst
	OpAMD64SARWconst
	OpAMD64SARBconst
	OpAMD64SHRDQ
	OpAMD64SHLDQ
	OpAMD64ROLQ
	OpAMD64ROLL
	OpAMD64ROLW
	OpAMD64ROLB
	OpAMD64RORQ
	OpAMD64RORL
	OpAMD64RORW
	OpAMD64RORB
	OpAMD64ROLQconst
	OpAMD64ROLLconst
	OpAMD64ROLWconst
	OpAMD64ROLBconst
	OpAMD64ADDLload
	OpAMD64ADDQload
	OpAMD64SUBQload
	OpAMD64SUBLload
	OpAMD64ANDLload
	OpAMD64ANDQload
	OpAMD64ORQload
	OpAMD64ORLload
	OpAMD64XORQload
	OpAMD64XORLload
	OpAMD64ADDLloadidx1
	OpAMD64ADDLloadidx4
	OpAMD64ADDLloadidx8
	OpAMD64ADDQloadidx1
	OpAMD64ADDQloadidx8
	OpAMD64SUBLloadidx1
	OpAMD64SUBLloadidx4
	OpAMD64SUBLloadidx8
	OpAMD64SUBQloadidx1
	OpAMD64SUBQloadidx8
	OpAMD64ANDLloadidx1
	OpAMD64ANDLloadidx4
	OpAMD64ANDLloadidx8
	OpAMD64ANDQloadidx1
	OpAMD64ANDQloadidx8
	OpAMD64ORLloadidx1
	OpAMD64ORLloadidx4
	OpAMD64ORLloadidx8
	OpAMD64ORQloadidx1
	OpAMD64ORQloadidx8
	OpAMD64XORLloadidx1
	OpAMD64XORLloadidx4
	OpAMD64XORLloadidx8
	OpAMD64XORQloadidx1
	OpAMD64XORQloadidx8
	OpAMD64ADDQmodify
	OpAMD64SUBQmodify
	OpAMD64ANDQmodify
	OpAMD64ORQmodify
	OpAMD64XORQmodify
	OpAMD64ADDLmodify
	OpAMD64SUBLmodify
	OpAMD64ANDLmodify
	OpAMD64ORLmodify
	OpAMD64XORLmodify
	OpAMD64ADDQmodifyidx1
	OpAMD64ADDQmodifyidx8
	OpAMD64SUBQmodifyidx1
	OpAMD64SUBQmodifyidx8
	OpAMD64ANDQmodifyidx1
	OpAMD64ANDQmodifyidx8
	OpAMD64ORQmodifyidx1
	OpAMD64ORQmodifyidx8
	OpAMD64XORQmodifyidx1
	OpAMD64XORQmodifyidx8
	OpAMD64ADDLmodifyidx1
	OpAMD64ADDLmodifyidx4
	OpAMD64ADDLmodifyidx8
	OpAMD64SUBLmodifyidx1
	OpAMD64SUBLmodifyidx4
	OpAMD64SUBLmodifyidx8
	OpAMD64ANDLmodifyidx1
	OpAMD64ANDLmodifyidx4
	OpAMD64ANDLmodifyidx8
	OpAMD64ORLmodifyidx1
	OpAMD64ORLmodifyidx4
	OpAMD64ORLmodifyidx8
	OpAMD64XORLmodifyidx1
	OpAMD64XORLmodifyidx4
	OpAMD64XORLmodifyidx8
	OpAMD64ADDQconstmodifyidx1
	OpAMD64ADDQconstmodifyidx8
	OpAMD64ANDQconstmodifyidx1
	OpAMD64ANDQconstmodifyidx8
	OpAMD64ORQconstmodifyidx1
	OpAMD64ORQconstmodifyidx8
	OpAMD64XORQconstmodifyidx1
	OpAMD64XORQconstmodifyidx8
	OpAMD64ADDLconstmodifyidx1
	OpAMD64ADDLconstmodifyidx4
	OpAMD64ADDLconstmodifyidx8
	OpAMD64ANDLconstmodifyidx1
	OpAMD64ANDLconstmodifyidx4
	OpAMD64ANDLconstmodifyidx8
	OpAMD64ORLconstmodifyidx1
	OpAMD64ORLconstmodifyidx4
	OpAMD64ORLconstmodifyidx8
	OpAMD64XORLconstmodifyidx1
	OpAMD64XORLconstmodifyidx4
	OpAMD64XORLconstmodifyidx8
	OpAMD64NEGQ
	OpAMD64NEGL
	OpAMD64NOTQ
	OpAMD64NOTL
	OpAMD64BSFQ
	OpAMD64BSFL
	OpAMD64BSRQ
	OpAMD64BSRL
	OpAMD64CMOVQEQ
	OpAMD64CMOVQNE
	OpAMD64CMOVQLT
	OpAMD64CMOVQGT
	OpAMD64CMOVQLE
	OpAMD64CMOVQGE
	OpAMD64CMOVQLS
	OpAMD64CMOVQHI
	OpAMD64CMOVQCC
	OpAMD64CMOVQCS
	OpAMD64CMOVLEQ
	OpAMD64CMOVLNE
	OpAMD64CMOVLLT
	OpAMD64CMOVLGT
	OpAMD64CMOVLLE
	OpAMD64CMOVLGE
	OpAMD64CMOVLLS
	OpAMD64CMOVLHI
	OpAMD64CMOVLCC
	OpAMD64CMOVLCS
	OpAMD64CMOVWEQ
	OpAMD64CMOVWNE
	OpAMD64CMOVWLT
	OpAMD64CMOVWGT
	OpAMD64CMOVWLE
	OpAMD64CMOVWGE
	OpAMD64CMOVWLS
	OpAMD64CMOVWHI
	OpAMD64CMOVWCC
	OpAMD64CMOVWCS
	OpAMD64CMOVQEQF
	OpAMD64CMOVQNEF
	OpAMD64CMOVQGTF
	OpAMD64CMOVQGEF
	OpAMD64CMOVLEQF
	OpAMD64CMOVLNEF
	OpAMD64CMOVLGTF
	OpAMD64CMOVLGEF
	OpAMD64CMOVWEQF
	OpAMD64CMOVWNEF
	OpAMD64CMOVWGTF
	OpAMD64CMOVWGEF
	OpAMD64BSWAPQ
	OpAMD64BSWAPL
	OpAMD64POPCNTQ
	OpAMD64POPCNTL
	OpAMD64SQRTSD
	OpAMD64SQRTSS
	OpAMD64ROUNDSD
	OpAMD64VFMADD231SD
	OpAMD64SBBQcarrymask
	OpAMD64SBBLcarrymask
	OpAMD64SETEQ
	OpAMD64SETNE
	OpAMD64SETL
	OpAMD64SETLE
	OpAMD64SETG
	OpAMD64SETGE
	OpAMD64SETB
	OpAMD64SETBE
	OpAMD64SETA
	OpAMD64SETAE
	OpAMD64SETO
	OpAMD64SETEQstore
	OpAMD64SETNEstore
	OpAMD64SETLstore
	OpAMD64SETLEstore
	OpAMD64SETGstore
	OpAMD64SETGEstore
	OpAMD64SETBstore
	OpAMD64SETBEstore
	OpAMD64SETAstore
	OpAMD64SETAEstore
	OpAMD64SETEQF
	OpAMD64SETNEF
	OpAMD64SETORD
	OpAMD64SETNAN
	OpAMD64SETGF
	OpAMD64SETGEF
	OpAMD64MOVBQSX
	OpAMD64MOVBQZX
	OpAMD64MOVWQSX
	OpAMD64MOVWQZX
	OpAMD64MOVLQSX
	OpAMD64MOVLQZX
	OpAMD64MOVLconst
	OpAMD64MOVQconst
	OpAMD64CVTTSD2SL
	OpAMD64CVTTSD2SQ
	OpAMD64CVTTSS2SL
	OpAMD64CVTTSS2SQ
	OpAMD64CVTSL2SS
	OpAMD64CVTSL2SD
	OpAMD64CVTSQ2SS
	OpAMD64CVTSQ2SD
	OpAMD64CVTSD2SS
	OpAMD64CVTSS2SD
	OpAMD64MOVQi2f
	OpAMD64MOVQf2i
	OpAMD64MOVLi2f
	OpAMD64MOVLf2i
	OpAMD64PXOR
	OpAMD64LEAQ
	OpAMD64LEAL
	OpAMD64LEAW
	OpAMD64LEAQ1
	OpAMD64LEAL1
	OpAMD64LEAW1
	OpAMD64LEAQ2
	OpAMD64LEAL2
	OpAMD64LEAW2
	OpAMD64LEAQ4
	OpAMD64LEAL4
	OpAMD64LEAW4
	OpAMD64LEAQ8
	OpAMD64LEAL8
	OpAMD64LEAW8
	OpAMD64MOVBload
	OpAMD64MOVBQSXload
	OpAMD64MOVWload
	OpAMD64MOVWQSXload
	OpAMD64MOVLload
	OpAMD64MOVLQSXload
	OpAMD64MOVQload
	OpAMD64MOVBstore
	OpAMD64MOVWstore
	OpAMD64MOVLstore
	OpAMD64MOVQstore
	OpAMD64MOVOload
	OpAMD64MOVOstore
	OpAMD64MOVBloadidx1
	OpAMD64MOVWloadidx1
	OpAMD64MOVWloadidx2
	OpAMD64MOVLloadidx1
	OpAMD64MOVLloadidx4
	OpAMD64MOVLloadidx8
	OpAMD64MOVQloadidx1
	OpAMD64MOVQloadidx8
	OpAMD64MOVBstoreidx1
	OpAMD64MOVWstoreidx1
	OpAMD64MOVWstoreidx2
	OpAMD64MOVLstoreidx1
	OpAMD64MOVLstoreidx4
	OpAMD64MOVLstoreidx8
	OpAMD64MOVQstoreidx1
	OpAMD64MOVQstoreidx8
	OpAMD64MOVBstoreconst
	OpAMD64MOVWstoreconst
	OpAMD64MOVLstoreconst
	OpAMD64MOVQstoreconst
	OpAMD64MOVOstoreconst
	OpAMD64MOVBstoreconstidx1
	OpAMD64MOVWstoreconstidx1
	OpAMD64MOVWstoreconstidx2
	OpAMD64MOVLstoreconstidx1
	OpAMD64MOVLstoreconstidx4
	OpAMD64MOVQstoreconstidx1
	OpAMD64MOVQstoreconstidx8
	OpAMD64DUFFZERO
	OpAMD64REPSTOSQ
	OpAMD64CALLstatic
	OpAMD64CALLtail
	OpAMD64CALLclosure
	OpAMD64CALLinter
	OpAMD64DUFFCOPY
	OpAMD64REPMOVSQ
	OpAMD64InvertFlags
	OpAMD64LoweredGetG
	OpAMD64LoweredGetClosurePtr
	OpAMD64LoweredGetCallerPC
	OpAMD64LoweredGetCallerSP
	OpAMD64LoweredNilCheck
	OpAMD64LoweredWB
	OpAMD64LoweredHasCPUFeature
	OpAMD64LoweredPanicBoundsA
	OpAMD64LoweredPanicBoundsB
	OpAMD64LoweredPanicBoundsC
	OpAMD64FlagEQ
	OpAMD64FlagLT_ULT
	OpAMD64FlagLT_UGT
	OpAMD64FlagGT_UGT
	OpAMD64FlagGT_ULT
	OpAMD64MOVBatomicload
	OpAMD64MOVLatomicload
	OpAMD64MOVQatomicload
	OpAMD64XCHGB
	OpAMD64XCHGL
	OpAMD64XCHGQ
	OpAMD64XADDLlock
	OpAMD64XADDQlock
	OpAMD64AddTupleFirst32
	OpAMD64AddTupleFirst64
	OpAMD64CMPXCHGLlock
	OpAMD64CMPXCHGQlock
	OpAMD64ANDBlock
	OpAMD64ANDLlock
	OpAMD64ORBlock
	OpAMD64ORLlock
	OpAMD64PrefetchT0
	OpAMD64PrefetchNTA
	OpAMD64ANDNQ
	OpAMD64ANDNL
	OpAMD64BLSIQ
	OpAMD64BLSIL
	OpAMD64BLSMSKQ
	OpAMD64BLSMSKL
	OpAMD64BLSRQ
	OpAMD64BLSRL
	OpAMD64TZCNTQ
	OpAMD64TZCNTL
	OpAMD64LZCNTQ
	OpAMD64LZCNTL
	OpAMD64MOVBEWstore
	OpAMD64MOVBELload
	OpAMD64MOVBELstore
	OpAMD64MOVBEQload
	OpAMD64MOVBEQstore
	OpAMD64MOVBELloadidx1
	OpAMD64MOVBELloadidx4
	OpAMD64MOVBELloadidx8
	OpAMD64MOVBEQloadidx1
	OpAMD64MOVBEQloadidx8
	OpAMD64MOVBEWstoreidx1
	OpAMD64MOVBEWstoreidx2
	OpAMD64MOVBELstoreidx1
	OpAMD64MOVBELstoreidx4
	OpAMD64MOVBELstoreidx8
	OpAMD64MOVBEQstoreidx1
	OpAMD64MOVBEQstoreidx8
	OpAMD64SARXQ
	OpAMD64SARXL
	OpAMD64SHLXQ
	OpAMD64SHLXL
	OpAMD64SHRXQ
	OpAMD64SHRXL
	OpAMD64SARXLload
	OpAMD64SARXQload
	OpAMD64SHLXLload
	OpAMD64SHLXQload
	OpAMD64SHRXLload
	OpAMD64SHRXQload
	OpAMD64SARXLloadidx1
	OpAMD64SARXLloadidx4
	OpAMD64SARXLloadidx8
	OpAMD64SARXQloadidx1
	OpAMD64SARXQloadidx8
	OpAMD64SHLXLloadidx1
	OpAMD64SHLXLloadidx4
	OpAMD64SHLXLloadidx8
	OpAMD64SHLXQloadidx1
	OpAMD64SHLXQloadidx8
	OpAMD64SHRXLloadidx1
	OpAMD64SHRXLloadidx4
	OpAMD64SHRXLloadidx8
	OpAMD64SHRXQloadidx1
	OpAMD64SHRXQloadidx8

	OpAdd8
	OpAdd16
	OpAdd32
	OpAdd64
	OpAddPtr
	OpAdd32F
	OpAdd64F
	OpSub8
	OpSub16
	OpSub32
	OpSub64
	OpSubPtr
	OpSub32F
	OpSub64F
	OpMul8
	OpMul16
	OpMul32
	OpMul64
	OpMul32F
	OpMul64F
	OpDiv32F
	OpDiv64F
	OpHmul32
	OpHmul32u
	OpHmul64
	OpHmul64u
	OpMul32uhilo
	OpMul64uhilo
	OpMul32uover
	OpMul64uover
	OpAvg32u
	OpAvg64u
	OpDiv8
	OpDiv8u
	OpDiv16
	OpDiv16u
	OpDiv32
	OpDiv32u
	OpDiv64
	OpDiv64u
	OpDiv128u
	OpMod8
	OpMod8u
	OpMod16
	OpMod16u
	OpMod32
	OpMod32u
	OpMod64
	OpMod64u
	OpAnd8
	OpAnd16
	OpAnd32
	OpAnd64
	OpOr8
	OpOr16
	OpOr32
	OpOr64
	OpXor8
	OpXor16
	OpXor32
	OpXor64
	OpLsh8x8
	OpLsh8x16
	OpLsh8x32
	OpLsh8x64
	OpLsh16x8
	OpLsh16x16
	OpLsh16x32
	OpLsh16x64
	OpLsh32x8
	OpLsh32x16
	OpLsh32x32
	OpLsh32x64
	OpLsh64x8
	OpLsh64x16
	OpLsh64x32
	OpLsh64x64
	OpRsh8x8
	OpRsh8x16
	OpRsh8x32
	OpRsh8x64
	OpRsh16x8
	OpRsh16x16
	OpRsh16x32
	OpRsh16x64
	OpRsh32x8
	OpRsh32x16
	OpRsh32x32
	OpRsh32x64
	OpRsh64x8
	OpRsh64x16
	OpRsh64x32
	OpRsh64x64
	OpRsh8Ux8
	OpRsh8Ux16
	OpRsh8Ux32
	OpRsh8Ux64
	OpRsh16Ux8
	OpRsh16Ux16
	OpRsh16Ux32
	OpRsh16Ux64
	OpRsh32Ux8
	OpRsh32Ux16
	OpRsh32Ux32
	OpRsh32Ux64
	OpRsh64Ux8
	OpRsh64Ux16
	OpRsh64Ux32
	OpRsh64Ux64
	OpEq8
	OpEq16
	OpEq32
	OpEq64
	OpEqPtr
	OpEqInter
	OpEqSlice
	OpEq32F
	OpEq64F
	OpNeq8
	OpNeq16
	OpNeq32
	OpNeq64
	OpNeqPtr
	OpNeqInter
	OpNeqSlice
	OpNeq32F
	OpNeq64F
	OpLess8
	OpLess8U
	OpLess16
	OpLess16U
	OpLess32
	OpLess32U
	OpLess64
	OpLess64U
	OpLess32F
	OpLess64F
	OpLeq8
	OpLeq8U
	OpLeq16
	OpLeq16U
	OpLeq32
	OpLeq32U
	OpLeq64
	OpLeq64U
	OpLeq32F
	OpLeq64F
	OpCondSelect
	OpAndB
	OpOrB
	OpEqB
	OpNeqB
	OpNot
	OpNeg8
	OpNeg16
	OpNeg32
	OpNeg64
	OpNeg32F
	OpNeg64F
	OpCom8
	OpCom16
	OpCom32
	OpCom64
	OpCtz8
	OpCtz16
	OpCtz32
	OpCtz64
	OpCtz8NonZero
	OpCtz16NonZero
	OpCtz32NonZero
	OpCtz64NonZero
	OpBitLen8
	OpBitLen16
	OpBitLen32
	OpBitLen64
	OpBswap16
	OpBswap32
	OpBswap64
	OpBitRev8
	OpBitRev16
	OpBitRev32
	OpBitRev64
	OpPopCount8
	OpPopCount16
	OpPopCount32
	OpPopCount64
	OpRotateLeft64
	OpRotateLeft32
	OpRotateLeft16
	OpRotateLeft8
	OpSqrt
	OpSqrt32
	OpFloor
	OpCeil
	OpTrunc
	OpRound
	OpRoundToEven
	OpAbs
	OpCopysign
	OpFMA
	OpPhi
	OpCopy
	OpConvert
	OpConstBool
	OpConstString
	OpConstNil
	OpConst8
	OpConst16
	OpConst32
	OpConst64
	OpConst32F
	OpConst64F
	OpConstInterface
	OpConstSlice
	OpInitMem
	OpArg
	OpArgIntReg
	OpArgFloatReg
	OpAddr
	OpLocalAddr
	OpSP
	OpSB
	OpSPanchored
	OpLoad
	OpDereference
	OpStore
	OpMove
	OpZero
	OpStoreWB
	OpMoveWB
	OpZeroWB
	OpWBend
	OpWB
	OpHasCPUFeature
	OpPanicBounds
	OpPanicExtend
	OpClosureCall
	OpStaticCall
	OpInterCall
	OpTailCall
	OpClosureLECall
	OpStaticLECall
	OpInterLECall
	OpTailLECall
	OpSignExt8to16
	OpSignExt8to32
	OpSignExt8to64
	OpSignExt16to32
	OpSignExt16to64
	OpSignExt32to64
	OpZeroExt8to16
	OpZeroExt8to32
	OpZeroExt8to64
	OpZeroExt16to32
	OpZeroExt16to64
	OpZeroExt32to64
	OpTrunc16to8
	OpTrunc32to8
	OpTrunc32to16
	OpTrunc64to8
	OpTrunc64to16
	OpTrunc64to32
	OpCvt32to32F
	OpCvt32to64F
	OpCvt64to32F
	OpCvt64to64F
	OpCvt32Fto32
	OpCvt32Fto64
	OpCvt64Fto32
	OpCvt64Fto64
	OpCvt32Fto64F
	OpCvt64Fto32F
	OpCvtBoolToUint8
	OpRound32F
	OpRound64F
	OpIsNonNil
	OpIsInBounds
	OpIsSliceInBounds
	OpNilCheck
	OpGetG
	OpGetClosurePtr
	OpGetCallerPC
	OpGetCallerSP
	OpPtrIndex
	OpOffPtr
	OpSliceMake
	OpSlicePtr
	OpSliceLen
	OpSliceCap
	OpSlicePtrUnchecked
	OpComplexMake
	OpComplexReal
	OpComplexImag
	OpStringMake
	OpStringPtr
	OpStringLen
	OpIMake
	OpITab
	OpIData
	OpStructMake0
	OpStructMake1
	OpStructMake2
	OpStructMake3
	OpStructMake4
	OpStructSelect
	OpArrayMake0
	OpArrayMake1
	OpArraySelect
	OpStoreReg
	OpLoadReg
	OpFwdRef
	OpUnknown
	OpVarDef
	OpVarLive
	OpKeepAlive
	OpInlMark
	OpInt64Make
	OpInt64Hi
	OpInt64Lo
	OpAdd32carry
	OpAdd32withcarry
	OpSub32carry
	OpSub32withcarry
	OpAdd64carry
	OpSub64borrow
	OpSignmask
	OpZeromask
	OpSlicemask
	OpSpectreIndex
	OpSpectreSliceIndex
	OpCvt32Uto32F
	OpCvt32Uto64F
	OpCvt32Fto32U
	OpCvt64Fto32U
	OpCvt64Uto32F
	OpCvt64Uto64F
	OpCvt32Fto64U
	OpCvt64Fto64U
	OpSelect0
	OpSelect1
	OpSelectN
	OpSelectNAddr
	OpMakeResult
	OpAtomicLoad8
	OpAtomicLoad32
	OpAtomicLoad64
	OpAtomicLoadPtr
	OpAtomicLoadAcq32
	OpAtomicLoadAcq64
	OpAtomicStore8
	OpAtomicStore32
	OpAtomicStore64
	OpAtomicStorePtrNoWB
	OpAtomicStoreRel32
	OpAtomicStoreRel64
	OpAtomicExchange32
	OpAtomicExchange64
	OpAtomicAdd32
	OpAtomicAdd64
	OpAtomicCompareAndSwap32
	OpAtomicCompareAndSwap64
	OpAtomicCompareAndSwapRel32
	OpAtomicAnd8
	OpAtomicAnd32
	OpAtomicOr8
	OpAtomicOr32
	OpAtomicAdd32Variant
	OpAtomicAdd64Variant
	OpAtomicExchange32Variant
	OpAtomicExchange64Variant
	OpAtomicCompareAndSwap32Variant
	OpAtomicCompareAndSwap64Variant
	OpAtomicAnd8Variant
	OpAtomicAnd32Variant
	OpAtomicOr8Variant
	OpAtomicOr32Variant
	OpPubBarrier
	OpClobber
	OpClobberReg
	OpPrefetchCache
	OpPrefetchCacheStreamed
)

var opcodeTable = [...]opInfo{
	{name: "OpInvalid"},

	{
		name:         "ADDSS",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		asm:          x86.AADDSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{1, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:         "ADDSD",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		asm:          x86.AADDSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{1, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:         "SUBSS",
		argLen:       2,
		resultInArg0: true,
		asm:          x86.ASUBSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{1, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:         "SUBSD",
		argLen:       2,
		resultInArg0: true,
		asm:          x86.ASUBSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{1, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:         "MULSS",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		asm:          x86.AMULSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{1, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:         "MULSD",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		asm:          x86.AMULSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{1, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:         "DIVSS",
		argLen:       2,
		resultInArg0: true,
		asm:          x86.ADIVSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{1, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:         "DIVSD",
		argLen:       2,
		resultInArg0: true,
		asm:          x86.ADIVSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{1, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:           "MOVSSload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:           "MOVSDload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:              "MOVSSconst",
		auxType:           auxFloat32,
		argLen:            0,
		rematerializeable: true,
		asm:               x86.AMOVSS,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:              "MOVSDconst",
		auxType:           auxFloat64,
		argLen:            0,
		rematerializeable: true,
		asm:               x86.AMOVSD,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:      "MOVSSloadidx1",
		auxType:   auxSymOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       x86.AMOVSS,
		scale:     1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:      "MOVSSloadidx4",
		auxType:   auxSymOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       x86.AMOVSS,
		scale:     4,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:      "MOVSDloadidx1",
		auxType:   auxSymOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       x86.AMOVSD,
		scale:     1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:      "MOVSDloadidx8",
		auxType:   auxSymOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       x86.AMOVSD,
		scale:     8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:           "MOVSSstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVSS,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{0, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
			},
		},
	},
	{
		name:           "MOVSDstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVSD,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{0, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
			},
		},
	},
	{
		name:      "MOVSSstoreidx1",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVSS,
		scale:     1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{0, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
			},
		},
	},
	{
		name:      "MOVSSstoreidx4",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVSS,
		scale:     4,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{0, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
			},
		},
	},
	{
		name:      "MOVSDstoreidx1",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVSD,
		scale:     1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{0, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
			},
		},
	},
	{
		name:      "MOVSDstoreidx8",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVSD,
		scale:     8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{0, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
			},
		},
	},
	{
		name:           "ADDSSload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AADDSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:           "ADDSDload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AADDSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:           "SUBSSload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.ASUBSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:           "SUBSDload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.ASUBSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:           "MULSSload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AMULSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:           "MULSDload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AMULSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:           "DIVSSload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.ADIVSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:           "DIVSDload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.ADIVSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:         "ADDSSloadidx1",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		symEffect:    SymRead,
		asm:          x86.AADDSS,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{2, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:         "ADDSSloadidx4",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		symEffect:    SymRead,
		asm:          x86.AADDSS,
		scale:        4,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{2, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:         "ADDSDloadidx1",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		symEffect:    SymRead,
		asm:          x86.AADDSD,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{2, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:         "ADDSDloadidx8",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		symEffect:    SymRead,
		asm:          x86.AADDSD,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{2, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:         "SUBSSloadidx1",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		symEffect:    SymRead,
		asm:          x86.ASUBSS,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{2, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:         "SUBSSloadidx4",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		symEffect:    SymRead,
		asm:          x86.ASUBSS,
		scale:        4,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{2, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:         "SUBSDloadidx1",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		symEffect:    SymRead,
		asm:          x86.ASUBSD,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{2, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:         "SUBSDloadidx8",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		symEffect:    SymRead,
		asm:          x86.ASUBSD,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{2, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:         "MULSSloadidx1",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		symEffect:    SymRead,
		asm:          x86.AMULSS,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{2, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:         "MULSSloadidx4",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		symEffect:    SymRead,
		asm:          x86.AMULSS,
		scale:        4,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{2, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:         "MULSDloadidx1",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		symEffect:    SymRead,
		asm:          x86.AMULSD,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{2, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:         "MULSDloadidx8",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		symEffect:    SymRead,
		asm:          x86.AMULSD,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{2, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:         "DIVSSloadidx1",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		symEffect:    SymRead,
		asm:          x86.ADIVSS,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{2, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:         "DIVSSloadidx4",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		symEffect:    SymRead,
		asm:          x86.ADIVSS,
		scale:        4,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{2, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:         "DIVSDloadidx1",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		symEffect:    SymRead,
		asm:          x86.ADIVSD,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{2, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:         "DIVSDloadidx8",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		symEffect:    SymRead,
		asm:          x86.ADIVSD,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{2, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:         "ADDQ",
		argLen:       2,
		commutative:  true,
		clobberFlags: true,
		asm:          x86.AADDQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ADDL",
		argLen:       2,
		commutative:  true,
		clobberFlags: true,
		asm:          x86.AADDL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ADDQconst",
		auxType:      auxInt32,
		argLen:       1,
		clobberFlags: true,
		asm:          x86.AADDQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ADDLconst",
		auxType:      auxInt32,
		argLen:       1,
		clobberFlags: true,
		asm:          x86.AADDL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "ADDQconstmodify",
		auxType:        auxSymValAndOff,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AADDQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "ADDLconstmodify",
		auxType:        auxSymValAndOff,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AADDL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "SUBQ",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASUBQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SUBL",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASUBL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SUBQconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASUBQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SUBLconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASUBL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "MULQ",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AIMULQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "MULL",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AIMULL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "MULQconst",
		auxType:      auxInt32,
		argLen:       1,
		clobberFlags: true,
		asm:          x86.AIMUL3Q,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "MULLconst",
		auxType:      auxInt32,
		argLen:       1,
		clobberFlags: true,
		asm:          x86.AIMUL3L,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "MULLU",
		argLen:       2,
		commutative:  true,
		clobberFlags: true,
		asm:          x86.AMULL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},     // AX
				{1, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			clobbers: 4, // DX
			outputs: []outputInfo{
				{1, 0},
				{0, 1}, // AX
			},
		},
	},
	{
		name:         "MULQU",
		argLen:       2,
		commutative:  true,
		clobberFlags: true,
		asm:          x86.AMULQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},     // AX
				{1, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			clobbers: 4, // DX
			outputs: []outputInfo{
				{1, 0},
				{0, 1}, // AX
			},
		},
	},
	{
		name:         "HMULQ",
		argLen:       2,
		clobberFlags: true,
		asm:          x86.AIMULQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},     // AX
				{1, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			clobbers: 1, // AX
			outputs: []outputInfo{
				{0, 4}, // DX
			},
		},
	},
	{
		name:         "HMULL",
		argLen:       2,
		clobberFlags: true,
		asm:          x86.AIMULL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},     // AX
				{1, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			clobbers: 1, // AX
			outputs: []outputInfo{
				{0, 4}, // DX
			},
		},
	},
	{
		name:         "HMULQU",
		argLen:       2,
		clobberFlags: true,
		asm:          x86.AMULQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},     // AX
				{1, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			clobbers: 1, // AX
			outputs: []outputInfo{
				{0, 4}, // DX
			},
		},
	},
	{
		name:         "HMULLU",
		argLen:       2,
		clobberFlags: true,
		asm:          x86.AMULL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},     // AX
				{1, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			clobbers: 1, // AX
			outputs: []outputInfo{
				{0, 4}, // DX
			},
		},
	},
	{
		name:         "AVGQU",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "DIVQ",
		auxType:      auxBool,
		argLen:       2,
		clobberFlags: true,
		asm:          x86.AIDIVQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},     // AX
				{1, 49147}, // AX CX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 1}, // AX
				{1, 4}, // DX
			},
		},
	},
	{
		name:         "DIVL",
		auxType:      auxBool,
		argLen:       2,
		clobberFlags: true,
		asm:          x86.AIDIVL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},     // AX
				{1, 49147}, // AX CX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 1}, // AX
				{1, 4}, // DX
			},
		},
	},
	{
		name:         "DIVW",
		auxType:      auxBool,
		argLen:       2,
		clobberFlags: true,
		asm:          x86.AIDIVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},     // AX
				{1, 49147}, // AX CX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 1}, // AX
				{1, 4}, // DX
			},
		},
	},
	{
		name:         "DIVQU",
		argLen:       2,
		clobberFlags: true,
		asm:          x86.ADIVQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},     // AX
				{1, 49147}, // AX CX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 1}, // AX
				{1, 4}, // DX
			},
		},
	},
	{
		name:         "DIVLU",
		argLen:       2,
		clobberFlags: true,
		asm:          x86.ADIVL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},     // AX
				{1, 49147}, // AX CX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 1}, // AX
				{1, 4}, // DX
			},
		},
	},
	{
		name:         "DIVWU",
		argLen:       2,
		clobberFlags: true,
		asm:          x86.ADIVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},     // AX
				{1, 49147}, // AX CX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 1}, // AX
				{1, 4}, // DX
			},
		},
	},
	{
		name:         "NEGLflags",
		argLen:       1,
		resultInArg0: true,
		asm:          x86.ANEGL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ADDQcarry",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		asm:          x86.AADDQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ADCQ",
		argLen:       3,
		commutative:  true,
		resultInArg0: true,
		asm:          x86.AADCQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ADDQconstcarry",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		asm:          x86.AADDQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ADCQconst",
		auxType:      auxInt32,
		argLen:       2,
		resultInArg0: true,
		asm:          x86.AADCQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SUBQborrow",
		argLen:       2,
		resultInArg0: true,
		asm:          x86.ASUBQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SBBQ",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ASBBQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SUBQconstborrow",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		asm:          x86.ASUBQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SBBQconst",
		auxType:      auxInt32,
		argLen:       2,
		resultInArg0: true,
		asm:          x86.ASBBQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "MULQU2",
		argLen:       2,
		commutative:  true,
		clobberFlags: true,
		asm:          x86.AMULQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},     // AX
				{1, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 4}, // DX
				{1, 1}, // AX
			},
		},
	},
	{
		name:         "DIVQU2",
		argLen:       3,
		clobberFlags: true,
		asm:          x86.ADIVQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4},     // DX
				{1, 1},     // AX
				{2, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 1}, // AX
				{1, 4}, // DX
			},
		},
	},
	{
		name:         "ANDQ",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AANDQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ANDL",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AANDL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ANDQconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AANDQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ANDLconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AANDL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "ANDQconstmodify",
		auxType:        auxSymValAndOff,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AANDQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "ANDLconstmodify",
		auxType:        auxSymValAndOff,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AANDL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ORQ",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AORQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ORL",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AORL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ORQconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AORQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ORLconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AORL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "ORQconstmodify",
		auxType:        auxSymValAndOff,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AORQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "ORLconstmodify",
		auxType:        auxSymValAndOff,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AORL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "XORQ",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AXORQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "XORL",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AXORL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "XORQconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AXORQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "XORLconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AXORL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "XORQconstmodify",
		auxType:        auxSymValAndOff,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AXORQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "XORLconstmodify",
		auxType:        auxSymValAndOff,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AXORL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:   "CMPQ",
		argLen: 2,
		asm:    x86.ACMPQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "CMPL",
		argLen: 2,
		asm:    x86.ACMPL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "CMPW",
		argLen: 2,
		asm:    x86.ACMPW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "CMPB",
		argLen: 2,
		asm:    x86.ACMPB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:    "CMPQconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     x86.ACMPQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:    "CMPLconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     x86.ACMPL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:    "CMPWconst",
		auxType: auxInt16,
		argLen:  1,
		asm:     x86.ACMPW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:    "CMPBconst",
		auxType: auxInt8,
		argLen:  1,
		asm:     x86.ACMPB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "CMPQload",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ACMPQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "CMPLload",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ACMPL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "CMPWload",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ACMPW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "CMPBload",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ACMPB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "CMPQconstload",
		auxType:        auxSymValAndOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ACMPQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "CMPLconstload",
		auxType:        auxSymValAndOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ACMPL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "CMPWconstload",
		auxType:        auxSymValAndOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ACMPW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "CMPBconstload",
		auxType:        auxSymValAndOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ACMPB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:      "CMPQloadidx8",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymRead,
		asm:       x86.ACMPQ,
		scale:     8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:        "CMPQloadidx1",
		auxType:     auxSymOff,
		argLen:      4,
		commutative: true,
		symEffect:   SymRead,
		asm:         x86.ACMPQ,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:      "CMPLloadidx4",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymRead,
		asm:       x86.ACMPL,
		scale:     4,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:        "CMPLloadidx1",
		auxType:     auxSymOff,
		argLen:      4,
		commutative: true,
		symEffect:   SymRead,
		asm:         x86.ACMPL,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:      "CMPWloadidx2",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymRead,
		asm:       x86.ACMPW,
		scale:     2,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:        "CMPWloadidx1",
		auxType:     auxSymOff,
		argLen:      4,
		commutative: true,
		symEffect:   SymRead,
		asm:         x86.ACMPW,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:        "CMPBloadidx1",
		auxType:     auxSymOff,
		argLen:      4,
		commutative: true,
		symEffect:   SymRead,
		asm:         x86.ACMPB,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:      "CMPQconstloadidx8",
		auxType:   auxSymValAndOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       x86.ACMPQ,
		scale:     8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:        "CMPQconstloadidx1",
		auxType:     auxSymValAndOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymRead,
		asm:         x86.ACMPQ,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:      "CMPLconstloadidx4",
		auxType:   auxSymValAndOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       x86.ACMPL,
		scale:     4,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:        "CMPLconstloadidx1",
		auxType:     auxSymValAndOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymRead,
		asm:         x86.ACMPL,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:      "CMPWconstloadidx2",
		auxType:   auxSymValAndOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       x86.ACMPW,
		scale:     2,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:        "CMPWconstloadidx1",
		auxType:     auxSymValAndOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymRead,
		asm:         x86.ACMPW,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:        "CMPBconstloadidx1",
		auxType:     auxSymValAndOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymRead,
		asm:         x86.ACMPB,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:   "UCOMISS",
		argLen: 2,
		asm:    x86.AUCOMISS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{1, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:   "UCOMISD",
		argLen: 2,
		asm:    x86.AUCOMISD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{1, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:   "BTL",
		argLen: 2,
		asm:    x86.ABTL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "BTQ",
		argLen: 2,
		asm:    x86.ABTQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "BTCL",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ABTCL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "BTCQ",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ABTCQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "BTRL",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ABTRL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "BTRQ",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ABTRQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "BTSL",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ABTSL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "BTSQ",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ABTSQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:    "BTLconst",
		auxType: auxInt8,
		argLen:  1,
		asm:     x86.ABTL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:    "BTQconst",
		auxType: auxInt8,
		argLen:  1,
		asm:     x86.ABTQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "BTCLconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ABTCL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "BTCQconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ABTCQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "BTRLconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ABTRL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "BTRQconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ABTRQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "BTSLconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ABTSL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "BTSQconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ABTSQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:        "TESTQ",
		argLen:      2,
		commutative: true,
		asm:         x86.ATESTQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:        "TESTL",
		argLen:      2,
		commutative: true,
		asm:         x86.ATESTL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:        "TESTW",
		argLen:      2,
		commutative: true,
		asm:         x86.ATESTW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:        "TESTB",
		argLen:      2,
		commutative: true,
		asm:         x86.ATESTB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:    "TESTQconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     x86.ATESTQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:    "TESTLconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     x86.ATESTL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:    "TESTWconst",
		auxType: auxInt16,
		argLen:  1,
		asm:     x86.ATESTW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:    "TESTBconst",
		auxType: auxInt8,
		argLen:  1,
		asm:     x86.ATESTB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SHLQ",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHLQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SHLL",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHLL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SHLQconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHLQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SHLLconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHLL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SHRQ",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHRQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SHRL",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHRL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SHRW",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHRW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SHRB",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHRB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SHRQconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHRQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SHRLconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHRL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SHRWconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHRW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SHRBconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHRB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SARQ",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASARQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SARL",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASARL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SARW",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASARW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SARB",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASARB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SARQconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASARQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SARLconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASARL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SARWconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASARW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SARBconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASARB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SHRDQ",
		argLen:       3,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHRQ,
		reg: regInfo{
			inputs: []inputInfo{
				{2, 2},     // CX
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SHLDQ",
		argLen:       3,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHLQ,
		reg: regInfo{
			inputs: []inputInfo{
				{2, 2},     // CX
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ROLQ",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AROLQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ROLL",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AROLL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ROLW",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AROLW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ROLB",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AROLB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "RORQ",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ARORQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "RORL",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ARORL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "RORW",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ARORW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "RORB",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ARORB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ROLQconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AROLQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ROLLconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AROLL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ROLWconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AROLW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ROLBconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AROLB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "ADDLload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AADDL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "ADDQload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AADDQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "SUBQload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.ASUBQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "SUBLload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.ASUBL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "ANDLload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AANDL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "ANDQload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AANDQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "ORQload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AORQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "ORLload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AORL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "XORQload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AXORQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "XORLload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AXORL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ADDLloadidx1",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		clobberFlags: true,
		symEffect:    SymRead,
		asm:          x86.AADDL,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ADDLloadidx4",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		clobberFlags: true,
		symEffect:    SymRead,
		asm:          x86.AADDL,
		scale:        4,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ADDLloadidx8",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		clobberFlags: true,
		symEffect:    SymRead,
		asm:          x86.AADDL,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ADDQloadidx1",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		clobberFlags: true,
		symEffect:    SymRead,
		asm:          x86.AADDQ,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ADDQloadidx8",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		clobberFlags: true,
		symEffect:    SymRead,
		asm:          x86.AADDQ,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SUBLloadidx1",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		clobberFlags: true,
		symEffect:    SymRead,
		asm:          x86.ASUBL,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SUBLloadidx4",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		clobberFlags: true,
		symEffect:    SymRead,
		asm:          x86.ASUBL,
		scale:        4,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SUBLloadidx8",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		clobberFlags: true,
		symEffect:    SymRead,
		asm:          x86.ASUBL,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SUBQloadidx1",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		clobberFlags: true,
		symEffect:    SymRead,
		asm:          x86.ASUBQ,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SUBQloadidx8",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		clobberFlags: true,
		symEffect:    SymRead,
		asm:          x86.ASUBQ,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ANDLloadidx1",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		clobberFlags: true,
		symEffect:    SymRead,
		asm:          x86.AANDL,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ANDLloadidx4",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		clobberFlags: true,
		symEffect:    SymRead,
		asm:          x86.AANDL,
		scale:        4,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ANDLloadidx8",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		clobberFlags: true,
		symEffect:    SymRead,
		asm:          x86.AANDL,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ANDQloadidx1",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		clobberFlags: true,
		symEffect:    SymRead,
		asm:          x86.AANDQ,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ANDQloadidx8",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		clobberFlags: true,
		symEffect:    SymRead,
		asm:          x86.AANDQ,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ORLloadidx1",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		clobberFlags: true,
		symEffect:    SymRead,
		asm:          x86.AORL,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ORLloadidx4",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		clobberFlags: true,
		symEffect:    SymRead,
		asm:          x86.AORL,
		scale:        4,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ORLloadidx8",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		clobberFlags: true,
		symEffect:    SymRead,
		asm:          x86.AORL,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ORQloadidx1",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		clobberFlags: true,
		symEffect:    SymRead,
		asm:          x86.AORQ,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ORQloadidx8",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		clobberFlags: true,
		symEffect:    SymRead,
		asm:          x86.AORQ,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "XORLloadidx1",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		clobberFlags: true,
		symEffect:    SymRead,
		asm:          x86.AXORL,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "XORLloadidx4",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		clobberFlags: true,
		symEffect:    SymRead,
		asm:          x86.AXORL,
		scale:        4,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "XORLloadidx8",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		clobberFlags: true,
		symEffect:    SymRead,
		asm:          x86.AXORL,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "XORQloadidx1",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		clobberFlags: true,
		symEffect:    SymRead,
		asm:          x86.AXORQ,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "XORQloadidx8",
		auxType:      auxSymOff,
		argLen:       4,
		resultInArg0: true,
		clobberFlags: true,
		symEffect:    SymRead,
		asm:          x86.AXORQ,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "ADDQmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AADDQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "SUBQmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.ASUBQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "ANDQmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AANDQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "ORQmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AORQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "XORQmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AXORQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "ADDLmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AADDL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "SUBLmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.ASUBL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "ANDLmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AANDL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "ORLmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AORL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "XORLmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AXORL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ADDQmodifyidx1",
		auxType:      auxSymOff,
		argLen:       4,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AADDQ,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ADDQmodifyidx8",
		auxType:      auxSymOff,
		argLen:       4,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AADDQ,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "SUBQmodifyidx1",
		auxType:      auxSymOff,
		argLen:       4,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.ASUBQ,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "SUBQmodifyidx8",
		auxType:      auxSymOff,
		argLen:       4,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.ASUBQ,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ANDQmodifyidx1",
		auxType:      auxSymOff,
		argLen:       4,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AANDQ,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ANDQmodifyidx8",
		auxType:      auxSymOff,
		argLen:       4,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AANDQ,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ORQmodifyidx1",
		auxType:      auxSymOff,
		argLen:       4,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AORQ,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ORQmodifyidx8",
		auxType:      auxSymOff,
		argLen:       4,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AORQ,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "XORQmodifyidx1",
		auxType:      auxSymOff,
		argLen:       4,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AXORQ,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "XORQmodifyidx8",
		auxType:      auxSymOff,
		argLen:       4,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AXORQ,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ADDLmodifyidx1",
		auxType:      auxSymOff,
		argLen:       4,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AADDL,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ADDLmodifyidx4",
		auxType:      auxSymOff,
		argLen:       4,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AADDL,
		scale:        4,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ADDLmodifyidx8",
		auxType:      auxSymOff,
		argLen:       4,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AADDL,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "SUBLmodifyidx1",
		auxType:      auxSymOff,
		argLen:       4,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.ASUBL,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "SUBLmodifyidx4",
		auxType:      auxSymOff,
		argLen:       4,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.ASUBL,
		scale:        4,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "SUBLmodifyidx8",
		auxType:      auxSymOff,
		argLen:       4,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.ASUBL,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ANDLmodifyidx1",
		auxType:      auxSymOff,
		argLen:       4,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AANDL,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ANDLmodifyidx4",
		auxType:      auxSymOff,
		argLen:       4,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AANDL,
		scale:        4,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ANDLmodifyidx8",
		auxType:      auxSymOff,
		argLen:       4,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AANDL,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ORLmodifyidx1",
		auxType:      auxSymOff,
		argLen:       4,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AORL,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ORLmodifyidx4",
		auxType:      auxSymOff,
		argLen:       4,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AORL,
		scale:        4,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ORLmodifyidx8",
		auxType:      auxSymOff,
		argLen:       4,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AORL,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "XORLmodifyidx1",
		auxType:      auxSymOff,
		argLen:       4,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AXORL,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "XORLmodifyidx4",
		auxType:      auxSymOff,
		argLen:       4,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AXORL,
		scale:        4,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "XORLmodifyidx8",
		auxType:      auxSymOff,
		argLen:       4,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AXORL,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ADDQconstmodifyidx1",
		auxType:      auxSymValAndOff,
		argLen:       3,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AADDQ,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ADDQconstmodifyidx8",
		auxType:      auxSymValAndOff,
		argLen:       3,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AADDQ,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ANDQconstmodifyidx1",
		auxType:      auxSymValAndOff,
		argLen:       3,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AANDQ,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ANDQconstmodifyidx8",
		auxType:      auxSymValAndOff,
		argLen:       3,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AANDQ,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ORQconstmodifyidx1",
		auxType:      auxSymValAndOff,
		argLen:       3,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AORQ,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ORQconstmodifyidx8",
		auxType:      auxSymValAndOff,
		argLen:       3,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AORQ,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "XORQconstmodifyidx1",
		auxType:      auxSymValAndOff,
		argLen:       3,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AXORQ,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "XORQconstmodifyidx8",
		auxType:      auxSymValAndOff,
		argLen:       3,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AXORQ,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ADDLconstmodifyidx1",
		auxType:      auxSymValAndOff,
		argLen:       3,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AADDL,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ADDLconstmodifyidx4",
		auxType:      auxSymValAndOff,
		argLen:       3,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AADDL,
		scale:        4,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ADDLconstmodifyidx8",
		auxType:      auxSymValAndOff,
		argLen:       3,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AADDL,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ANDLconstmodifyidx1",
		auxType:      auxSymValAndOff,
		argLen:       3,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AANDL,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ANDLconstmodifyidx4",
		auxType:      auxSymValAndOff,
		argLen:       3,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AANDL,
		scale:        4,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ANDLconstmodifyidx8",
		auxType:      auxSymValAndOff,
		argLen:       3,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AANDL,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ORLconstmodifyidx1",
		auxType:      auxSymValAndOff,
		argLen:       3,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AORL,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ORLconstmodifyidx4",
		auxType:      auxSymValAndOff,
		argLen:       3,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AORL,
		scale:        4,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ORLconstmodifyidx8",
		auxType:      auxSymValAndOff,
		argLen:       3,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AORL,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "XORLconstmodifyidx1",
		auxType:      auxSymValAndOff,
		argLen:       3,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AXORL,
		scale:        1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "XORLconstmodifyidx4",
		auxType:      auxSymValAndOff,
		argLen:       3,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AXORL,
		scale:        4,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "XORLconstmodifyidx8",
		auxType:      auxSymValAndOff,
		argLen:       3,
		clobberFlags: true,
		symEffect:    SymRead | SymWrite,
		asm:          x86.AXORL,
		scale:        8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "NEGQ",
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ANEGQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "NEGL",
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ANEGL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "NOTQ",
		argLen:       1,
		resultInArg0: true,
		asm:          x86.ANOTQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "NOTL",
		argLen:       1,
		resultInArg0: true,
		asm:          x86.ANOTL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "BSFQ",
		argLen: 1,
		asm:    x86.ABSFQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "BSFL",
		argLen:       1,
		clobberFlags: true,
		asm:          x86.ABSFL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "BSRQ",
		argLen: 1,
		asm:    x86.ABSRQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "BSRL",
		argLen:       1,
		clobberFlags: true,
		asm:          x86.ABSRL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVQEQ",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVQEQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVQNE",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVQNE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVQLT",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVQLT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVQGT",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVQGT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVQLE",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVQLE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVQGE",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVQGE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVQLS",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVQLS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVQHI",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVQHI,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVQCC",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVQCC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVQCS",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVQCS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVLEQ",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVLEQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVLNE",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVLNE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVLLT",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVLLT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVLGT",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVLGT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVLLE",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVLLE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVLGE",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVLGE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVLLS",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVLLS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVLHI",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVLHI,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVLCC",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVLCC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVLCS",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVLCS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVWEQ",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVWEQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVWNE",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVWNE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVWLT",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVWLT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVWGT",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVWGT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVWLE",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVWLE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVWGE",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVWGE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVWLS",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVWLS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVWHI",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVWHI,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVWCC",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVWCC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVWCS",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVWCS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVQEQF",
		argLen:       3,
		resultInArg0: true,
		needIntTemp:  true,
		asm:          x86.ACMOVQNE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVQNEF",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVQNE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVQGTF",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVQHI,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVQGEF",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVQCC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVLEQF",
		argLen:       3,
		resultInArg0: true,
		needIntTemp:  true,
		asm:          x86.ACMOVLNE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVLNEF",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVLNE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVLGTF",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVLHI,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVLGEF",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVLCC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVWEQF",
		argLen:       3,
		resultInArg0: true,
		needIntTemp:  true,
		asm:          x86.ACMOVWNE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVWNEF",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVWNE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVWGTF",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVWHI,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "CMOVWGEF",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVWCC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "BSWAPQ",
		argLen:       1,
		resultInArg0: true,
		asm:          x86.ABSWAPQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "BSWAPL",
		argLen:       1,
		resultInArg0: true,
		asm:          x86.ABSWAPL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "POPCNTQ",
		argLen:       1,
		clobberFlags: true,
		asm:          x86.APOPCNTQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "POPCNTL",
		argLen:       1,
		clobberFlags: true,
		asm:          x86.APOPCNTL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "SQRTSD",
		argLen: 1,
		asm:    x86.ASQRTSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:   "SQRTSS",
		argLen: 1,
		asm:    x86.ASQRTSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:    "ROUNDSD",
		auxType: auxInt8,
		argLen:  1,
		asm:     x86.AROUNDSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:         "VFMADD231SD",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.AVFMADD231SD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{1, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{2, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:   "SBBQcarrymask",
		argLen: 1,
		asm:    x86.ASBBQ,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "SBBLcarrymask",
		argLen: 1,
		asm:    x86.ASBBL,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "SETEQ",
		argLen: 1,
		asm:    x86.ASETEQ,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "SETNE",
		argLen: 1,
		asm:    x86.ASETNE,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "SETL",
		argLen: 1,
		asm:    x86.ASETLT,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "SETLE",
		argLen: 1,
		asm:    x86.ASETLE,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "SETG",
		argLen: 1,
		asm:    x86.ASETGT,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "SETGE",
		argLen: 1,
		asm:    x86.ASETGE,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "SETB",
		argLen: 1,
		asm:    x86.ASETCS,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "SETBE",
		argLen: 1,
		asm:    x86.ASETLS,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "SETA",
		argLen: 1,
		asm:    x86.ASETHI,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "SETAE",
		argLen: 1,
		asm:    x86.ASETCC,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "SETO",
		argLen: 1,
		asm:    x86.ASETOS,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "SETEQstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.ASETEQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "SETNEstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.ASETNE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "SETLstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.ASETLT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "SETLEstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.ASETLE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "SETGstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.ASETGT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "SETGEstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.ASETGE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "SETBstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.ASETCS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "SETBEstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.ASETLS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "SETAstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.ASETHI,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "SETAEstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.ASETCC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "SETEQF",
		argLen:       1,
		clobberFlags: true,
		needIntTemp:  true,
		asm:          x86.ASETEQ,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "SETNEF",
		argLen:       1,
		clobberFlags: true,
		needIntTemp:  true,
		asm:          x86.ASETNE,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "SETORD",
		argLen: 1,
		asm:    x86.ASETPC,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "SETNAN",
		argLen: 1,
		asm:    x86.ASETPS,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "SETGF",
		argLen: 1,
		asm:    x86.ASETHI,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "SETGEF",
		argLen: 1,
		asm:    x86.ASETCC,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "MOVBQSX",
		argLen: 1,
		asm:    x86.AMOVBQSX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "MOVBQZX",
		argLen: 1,
		asm:    x86.AMOVBLZX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "MOVWQSX",
		argLen: 1,
		asm:    x86.AMOVWQSX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "MOVWQZX",
		argLen: 1,
		asm:    x86.AMOVWLZX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "MOVLQSX",
		argLen: 1,
		asm:    x86.AMOVLQSX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "MOVLQZX",
		argLen: 1,
		asm:    x86.AMOVL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:              "MOVLconst",
		auxType:           auxInt32,
		argLen:            0,
		rematerializeable: true,
		asm:               x86.AMOVL,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:              "MOVQconst",
		auxType:           auxInt64,
		argLen:            0,
		rematerializeable: true,
		asm:               x86.AMOVQ,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "CVTTSD2SL",
		argLen: 1,
		asm:    x86.ACVTTSD2SL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "CVTTSD2SQ",
		argLen: 1,
		asm:    x86.ACVTTSD2SQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "CVTTSS2SL",
		argLen: 1,
		asm:    x86.ACVTTSS2SL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "CVTTSS2SQ",
		argLen: 1,
		asm:    x86.ACVTTSS2SQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "CVTSL2SS",
		argLen: 1,
		asm:    x86.ACVTSL2SS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:   "CVTSL2SD",
		argLen: 1,
		asm:    x86.ACVTSL2SD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:   "CVTSQ2SS",
		argLen: 1,
		asm:    x86.ACVTSQ2SS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:   "CVTSQ2SD",
		argLen: 1,
		asm:    x86.ACVTSQ2SD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:   "CVTSD2SS",
		argLen: 1,
		asm:    x86.ACVTSD2SS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:   "CVTSS2SD",
		argLen: 1,
		asm:    x86.ACVTSS2SD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:   "MOVQi2f",
		argLen: 1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:   "MOVQf2i",
		argLen: 1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "MOVLi2f",
		argLen: 1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:   "MOVLf2i",
		argLen: 1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "PXOR",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		asm:          x86.APXOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{1, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:              "LEAQ",
		auxType:           auxSymOff,
		argLen:            1,
		rematerializeable: true,
		symEffect:         SymAddr,
		asm:               x86.ALEAQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:              "LEAL",
		auxType:           auxSymOff,
		argLen:            1,
		rematerializeable: true,
		symEffect:         SymAddr,
		asm:               x86.ALEAL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:              "LEAW",
		auxType:           auxSymOff,
		argLen:            1,
		rematerializeable: true,
		symEffect:         SymAddr,
		asm:               x86.ALEAW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:        "LEAQ1",
		auxType:     auxSymOff,
		argLen:      2,
		commutative: true,
		symEffect:   SymAddr,
		asm:         x86.ALEAQ,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:        "LEAL1",
		auxType:     auxSymOff,
		argLen:      2,
		commutative: true,
		symEffect:   SymAddr,
		asm:         x86.ALEAL,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:        "LEAW1",
		auxType:     auxSymOff,
		argLen:      2,
		commutative: true,
		symEffect:   SymAddr,
		asm:         x86.ALEAW,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:      "LEAQ2",
		auxType:   auxSymOff,
		argLen:    2,
		symEffect: SymAddr,
		asm:       x86.ALEAQ,
		scale:     2,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:      "LEAL2",
		auxType:   auxSymOff,
		argLen:    2,
		symEffect: SymAddr,
		asm:       x86.ALEAL,
		scale:     2,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:      "LEAW2",
		auxType:   auxSymOff,
		argLen:    2,
		symEffect: SymAddr,
		asm:       x86.ALEAW,
		scale:     2,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:      "LEAQ4",
		auxType:   auxSymOff,
		argLen:    2,
		symEffect: SymAddr,
		asm:       x86.ALEAQ,
		scale:     4,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:      "LEAL4",
		auxType:   auxSymOff,
		argLen:    2,
		symEffect: SymAddr,
		asm:       x86.ALEAL,
		scale:     4,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:      "LEAW4",
		auxType:   auxSymOff,
		argLen:    2,
		symEffect: SymAddr,
		asm:       x86.ALEAW,
		scale:     4,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:      "LEAQ8",
		auxType:   auxSymOff,
		argLen:    2,
		symEffect: SymAddr,
		asm:       x86.ALEAQ,
		scale:     8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:      "LEAL8",
		auxType:   auxSymOff,
		argLen:    2,
		symEffect: SymAddr,
		asm:       x86.ALEAL,
		scale:     8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:      "LEAW8",
		auxType:   auxSymOff,
		argLen:    2,
		symEffect: SymAddr,
		asm:       x86.ALEAW,
		scale:     8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "MOVBload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVBLZX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "MOVBQSXload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVBQSX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "MOVWload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVWLZX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "MOVWQSXload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVWQSX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "MOVLload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "MOVLQSXload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVLQSX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "MOVQload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "MOVBstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "MOVWstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "MOVLstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "MOVQstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "MOVOload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVUPS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
			},
			outputs: []outputInfo{
				{0, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			},
		},
	},
	{
		name:           "MOVOstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVUPS,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2147418112}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
				{0, 4295016447}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15 SB
			},
		},
	},
	{
		name:        "MOVBloadidx1",
		auxType:     auxSymOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymRead,
		asm:         x86.AMOVBLZX,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:        "MOVWloadidx1",
		auxType:     auxSymOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymRead,
		asm:         x86.AMOVWLZX,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:      "MOVWloadidx2",
		auxType:   auxSymOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       x86.AMOVWLZX,
		scale:     2,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:        "MOVLloadidx1",
		auxType:     auxSymOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymRead,
		asm:         x86.AMOVL,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:      "MOVLloadidx4",
		auxType:   auxSymOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       x86.AMOVL,
		scale:     4,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:      "MOVLloadidx8",
		auxType:   auxSymOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       x86.AMOVL,
		scale:     8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:        "MOVQloadidx1",
		auxType:     auxSymOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymRead,
		asm:         x86.AMOVQ,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:      "MOVQloadidx8",
		auxType:   auxSymOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       x86.AMOVQ,
		scale:     8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:        "MOVBstoreidx1",
		auxType:     auxSymOff,
		argLen:      4,
		commutative: true,
		symEffect:   SymWrite,
		asm:         x86.AMOVB,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:        "MOVWstoreidx1",
		auxType:     auxSymOff,
		argLen:      4,
		commutative: true,
		symEffect:   SymWrite,
		asm:         x86.AMOVW,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:      "MOVWstoreidx2",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVW,
		scale:     2,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:        "MOVLstoreidx1",
		auxType:     auxSymOff,
		argLen:      4,
		commutative: true,
		symEffect:   SymWrite,
		asm:         x86.AMOVL,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:      "MOVLstoreidx4",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVL,
		scale:     4,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:      "MOVLstoreidx8",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVL,
		scale:     8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:        "MOVQstoreidx1",
		auxType:     auxSymOff,
		argLen:      4,
		commutative: true,
		symEffect:   SymWrite,
		asm:         x86.AMOVQ,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:      "MOVQstoreidx8",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVQ,
		scale:     8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "MOVBstoreconst",
		auxType:        auxSymValAndOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "MOVWstoreconst",
		auxType:        auxSymValAndOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "MOVLstoreconst",
		auxType:        auxSymValAndOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "MOVQstoreconst",
		auxType:        auxSymValAndOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "MOVOstoreconst",
		auxType:        auxSymValAndOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVUPS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:        "MOVBstoreconstidx1",
		auxType:     auxSymValAndOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymWrite,
		asm:         x86.AMOVB,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:        "MOVWstoreconstidx1",
		auxType:     auxSymValAndOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymWrite,
		asm:         x86.AMOVW,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:      "MOVWstoreconstidx2",
		auxType:   auxSymValAndOff,
		argLen:    3,
		symEffect: SymWrite,
		asm:       x86.AMOVW,
		scale:     2,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:        "MOVLstoreconstidx1",
		auxType:     auxSymValAndOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymWrite,
		asm:         x86.AMOVL,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:      "MOVLstoreconstidx4",
		auxType:   auxSymValAndOff,
		argLen:    3,
		symEffect: SymWrite,
		asm:       x86.AMOVL,
		scale:     4,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:        "MOVQstoreconstidx1",
		auxType:     auxSymValAndOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymWrite,
		asm:         x86.AMOVQ,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:      "MOVQstoreconstidx8",
		auxType:   auxSymValAndOff,
		argLen:    3,
		symEffect: SymWrite,
		asm:       x86.AMOVQ,
		scale:     8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "DUFFZERO",
		auxType:        auxInt64,
		argLen:         2,
		faultOnNilArg0: true,
		unsafePoint:    true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 128}, // DI
			},
			clobbers: 128, // DI
		},
	},
	{
		name:           "REPSTOSQ",
		argLen:         4,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 128}, // DI
				{1, 2},   // CX
				{2, 1},   // AX
			},
			clobbers: 130, // CX DI
		},
	},
	{
		name:         "CALLstatic",
		auxType:      auxCallOff,
		argLen:       -1,
		clobberFlags: true,
		call:         true,
		reg: regInfo{
			clobbers: 2147483631, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 g R15 X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
		},
	},
	{
		name:         "CALLtail",
		auxType:      auxCallOff,
		argLen:       -1,
		clobberFlags: true,
		call:         true,
		tailCall:     true,
		reg: regInfo{
			clobbers: 2147483631, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 g R15 X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
		},
	},
	{
		name:         "CALLclosure",
		auxType:      auxCallOff,
		argLen:       -1,
		clobberFlags: true,
		call:         true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 4},     // DX
				{0, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			clobbers: 2147483631, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 g R15 X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
		},
	},
	{
		name:         "CALLinter",
		auxType:      auxCallOff,
		argLen:       -1,
		clobberFlags: true,
		call:         true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			clobbers: 2147483631, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 g R15 X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
		},
	},
	{
		name:           "DUFFCOPY",
		auxType:        auxInt64,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		faultOnNilArg1: true,
		unsafePoint:    true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 128}, // DI
				{1, 64},  // SI
			},
			clobbers: 65728, // SI DI X0
		},
	},
	{
		name:           "REPMOVSQ",
		argLen:         4,
		faultOnNilArg0: true,
		faultOnNilArg1: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 128}, // DI
				{1, 64},  // SI
				{2, 2},   // CX
			},
			clobbers: 194, // CX SI DI
		},
	},
	{
		name:   "InvertFlags",
		argLen: 1,
		reg:    regInfo{},
	},
	{
		name:   "LoweredGetG",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:      "LoweredGetClosurePtr",
		argLen:    0,
		zeroWidth: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 4}, // DX
			},
		},
	},
	{
		name:              "LoweredGetCallerPC",
		argLen:            0,
		rematerializeable: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:              "LoweredGetCallerSP",
		argLen:            1,
		rematerializeable: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "LoweredNilCheck",
		argLen:         2,
		clobberFlags:   true,
		nilCheck:       true,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49151}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "LoweredWB",
		auxType:      auxInt64,
		argLen:       1,
		clobberFlags: true,
		reg: regInfo{
			clobbers: 2147418112, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14
			outputs: []outputInfo{
				{0, 2048}, // R11
			},
		},
	},
	{
		name:              "LoweredHasCPUFeature",
		auxType:           auxSym,
		argLen:            0,
		rematerializeable: true,
		symEffect:         SymNone,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:    "LoweredPanicBoundsA",
		auxType: auxInt64,
		argLen:  3,
		call:    true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4}, // DX
				{1, 8}, // BX
			},
		},
	},
	{
		name:    "LoweredPanicBoundsB",
		auxType: auxInt64,
		argLen:  3,
		call:    true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2}, // CX
				{1, 4}, // DX
			},
		},
	},
	{
		name:    "LoweredPanicBoundsC",
		auxType: auxInt64,
		argLen:  3,
		call:    true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1}, // AX
				{1, 2}, // CX
			},
		},
	},
	{
		name:   "FlagEQ",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:   "FlagLT_ULT",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:   "FlagLT_UGT",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:   "FlagGT_UGT",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:   "FlagGT_ULT",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:           "MOVBatomicload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "MOVLatomicload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "MOVQatomicload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "XCHGB",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		faultOnNilArg1: true,
		hasSideEffects: true,
		symEffect:      SymRdWr,
		asm:            x86.AXCHGB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "XCHGL",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		faultOnNilArg1: true,
		hasSideEffects: true,
		symEffect:      SymRdWr,
		asm:            x86.AXCHGL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "XCHGQ",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		faultOnNilArg1: true,
		hasSideEffects: true,
		symEffect:      SymRdWr,
		asm:            x86.AXCHGQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "XADDLlock",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		hasSideEffects: true,
		symEffect:      SymRdWr,
		asm:            x86.AXADDL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "XADDQlock",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		hasSideEffects: true,
		symEffect:      SymRdWr,
		asm:            x86.AXADDQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "AddTupleFirst32",
		argLen: 2,
		reg:    regInfo{},
	},
	{
		name:   "AddTupleFirst64",
		argLen: 2,
		reg:    regInfo{},
	},
	{
		name:           "CMPXCHGLlock",
		auxType:        auxSymOff,
		argLen:         4,
		clobberFlags:   true,
		faultOnNilArg0: true,
		hasSideEffects: true,
		symEffect:      SymRdWr,
		asm:            x86.ACMPXCHGL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 1},     // AX
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			clobbers: 1, // AX
			outputs: []outputInfo{
				{1, 0},
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "CMPXCHGQlock",
		auxType:        auxSymOff,
		argLen:         4,
		clobberFlags:   true,
		faultOnNilArg0: true,
		hasSideEffects: true,
		symEffect:      SymRdWr,
		asm:            x86.ACMPXCHGQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 1},     // AX
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			clobbers: 1, // AX
			outputs: []outputInfo{
				{1, 0},
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "ANDBlock",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		hasSideEffects: true,
		symEffect:      SymRdWr,
		asm:            x86.AANDB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "ANDLlock",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		hasSideEffects: true,
		symEffect:      SymRdWr,
		asm:            x86.AANDL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "ORBlock",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		hasSideEffects: true,
		symEffect:      SymRdWr,
		asm:            x86.AORB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "ORLlock",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		hasSideEffects: true,
		symEffect:      SymRdWr,
		asm:            x86.AORL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "PrefetchT0",
		argLen:         2,
		hasSideEffects: true,
		asm:            x86.APREFETCHT0,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "PrefetchNTA",
		argLen:         2,
		hasSideEffects: true,
		asm:            x86.APREFETCHNTA,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:         "ANDNQ",
		argLen:       2,
		clobberFlags: true,
		asm:          x86.AANDNQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "ANDNL",
		argLen:       2,
		clobberFlags: true,
		asm:          x86.AANDNL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "BLSIQ",
		argLen:       1,
		clobberFlags: true,
		asm:          x86.ABLSIQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "BLSIL",
		argLen:       1,
		clobberFlags: true,
		asm:          x86.ABLSIL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "BLSMSKQ",
		argLen:       1,
		clobberFlags: true,
		asm:          x86.ABLSMSKQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "BLSMSKL",
		argLen:       1,
		clobberFlags: true,
		asm:          x86.ABLSMSKL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "BLSRQ",
		argLen: 1,
		asm:    x86.ABLSRQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "BLSRL",
		argLen: 1,
		asm:    x86.ABLSRL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "TZCNTQ",
		argLen:       1,
		clobberFlags: true,
		asm:          x86.ATZCNTQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "TZCNTL",
		argLen:       1,
		clobberFlags: true,
		asm:          x86.ATZCNTL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "LZCNTQ",
		argLen:       1,
		clobberFlags: true,
		asm:          x86.ALZCNTQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:         "LZCNTL",
		argLen:       1,
		clobberFlags: true,
		asm:          x86.ALZCNTL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "MOVBEWstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVBEW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "MOVBELload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVBEL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "MOVBELstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVBEL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:           "MOVBEQload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVBEQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "MOVBEQstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVBEQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:        "MOVBELloadidx1",
		auxType:     auxSymOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymRead,
		asm:         x86.AMOVBEL,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:      "MOVBELloadidx4",
		auxType:   auxSymOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       x86.AMOVBEL,
		scale:     4,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:      "MOVBELloadidx8",
		auxType:   auxSymOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       x86.AMOVBEL,
		scale:     8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:        "MOVBEQloadidx1",
		auxType:     auxSymOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymRead,
		asm:         x86.AMOVBEQ,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:      "MOVBEQloadidx8",
		auxType:   auxSymOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       x86.AMOVBEQ,
		scale:     8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:        "MOVBEWstoreidx1",
		auxType:     auxSymOff,
		argLen:      4,
		commutative: true,
		symEffect:   SymWrite,
		asm:         x86.AMOVBEW,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:      "MOVBEWstoreidx2",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVBEW,
		scale:     2,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:        "MOVBELstoreidx1",
		auxType:     auxSymOff,
		argLen:      4,
		commutative: true,
		symEffect:   SymWrite,
		asm:         x86.AMOVBEL,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:      "MOVBELstoreidx4",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVBEL,
		scale:     4,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:      "MOVBELstoreidx8",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVBEL,
		scale:     8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:        "MOVBEQstoreidx1",
		auxType:     auxSymOff,
		argLen:      4,
		commutative: true,
		symEffect:   SymWrite,
		asm:         x86.AMOVBEQ,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:      "MOVBEQstoreidx8",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVBEQ,
		scale:     8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{2, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
		},
	},
	{
		name:   "SARXQ",
		argLen: 2,
		asm:    x86.ASARXQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "SARXL",
		argLen: 2,
		asm:    x86.ASARXL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "SHLXQ",
		argLen: 2,
		asm:    x86.ASHLXQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "SHLXL",
		argLen: 2,
		asm:    x86.ASHLXL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "SHRXQ",
		argLen: 2,
		asm:    x86.ASHRXQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:   "SHRXL",
		argLen: 2,
		asm:    x86.ASHRXL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "SARXLload",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ASARXL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "SARXQload",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ASARXQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "SHLXLload",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ASHLXL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "SHLXQload",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ASHLXQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "SHRXLload",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ASHRXL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "SHRXQload",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ASHRXQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "SARXLloadidx1",
		auxType:        auxSymOff,
		argLen:         4,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ASARXL,
		scale:          1,
		reg: regInfo{
			inputs: []inputInfo{
				{2, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "SARXLloadidx4",
		auxType:        auxSymOff,
		argLen:         4,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ASARXL,
		scale:          4,
		reg: regInfo{
			inputs: []inputInfo{
				{2, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "SARXLloadidx8",
		auxType:        auxSymOff,
		argLen:         4,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ASARXL,
		scale:          8,
		reg: regInfo{
			inputs: []inputInfo{
				{2, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "SARXQloadidx1",
		auxType:        auxSymOff,
		argLen:         4,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ASARXQ,
		scale:          1,
		reg: regInfo{
			inputs: []inputInfo{
				{2, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "SARXQloadidx8",
		auxType:        auxSymOff,
		argLen:         4,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ASARXQ,
		scale:          8,
		reg: regInfo{
			inputs: []inputInfo{
				{2, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "SHLXLloadidx1",
		auxType:        auxSymOff,
		argLen:         4,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ASHLXL,
		scale:          1,
		reg: regInfo{
			inputs: []inputInfo{
				{2, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "SHLXLloadidx4",
		auxType:        auxSymOff,
		argLen:         4,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ASHLXL,
		scale:          4,
		reg: regInfo{
			inputs: []inputInfo{
				{2, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "SHLXLloadidx8",
		auxType:        auxSymOff,
		argLen:         4,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ASHLXL,
		scale:          8,
		reg: regInfo{
			inputs: []inputInfo{
				{2, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "SHLXQloadidx1",
		auxType:        auxSymOff,
		argLen:         4,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ASHLXQ,
		scale:          1,
		reg: regInfo{
			inputs: []inputInfo{
				{2, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "SHLXQloadidx8",
		auxType:        auxSymOff,
		argLen:         4,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ASHLXQ,
		scale:          8,
		reg: regInfo{
			inputs: []inputInfo{
				{2, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "SHRXLloadidx1",
		auxType:        auxSymOff,
		argLen:         4,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ASHRXL,
		scale:          1,
		reg: regInfo{
			inputs: []inputInfo{
				{2, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "SHRXLloadidx4",
		auxType:        auxSymOff,
		argLen:         4,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ASHRXL,
		scale:          4,
		reg: regInfo{
			inputs: []inputInfo{
				{2, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "SHRXLloadidx8",
		auxType:        auxSymOff,
		argLen:         4,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ASHRXL,
		scale:          8,
		reg: regInfo{
			inputs: []inputInfo{
				{2, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "SHRXQloadidx1",
		auxType:        auxSymOff,
		argLen:         4,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ASHRXQ,
		scale:          1,
		reg: regInfo{
			inputs: []inputInfo{
				{2, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},
	{
		name:           "SHRXQloadidx8",
		auxType:        auxSymOff,
		argLen:         4,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ASHRXQ,
		scale:          8,
		reg: regInfo{
			inputs: []inputInfo{
				{2, 49135},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
				{1, 49151},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 g R15 SB
			},
			outputs: []outputInfo{
				{0, 49135}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R15
			},
		},
	},

	{
		name:        "Add8",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Add16",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Add32",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Add64",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:    "AddPtr",
		argLen:  2,
		generic: true,
	},
	{
		name:        "Add32F",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Add64F",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:    "Sub8",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Sub16",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Sub32",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Sub64",
		argLen:  2,
		generic: true,
	},
	{
		name:    "SubPtr",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Sub32F",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Sub64F",
		argLen:  2,
		generic: true,
	},
	{
		name:        "Mul8",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Mul16",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Mul32",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Mul64",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Mul32F",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Mul64F",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:    "Div32F",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Div64F",
		argLen:  2,
		generic: true,
	},
	{
		name:        "Hmul32",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Hmul32u",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Hmul64",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Hmul64u",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Mul32uhilo",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Mul64uhilo",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Mul32uover",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Mul64uover",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:    "Avg32u",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Avg64u",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Div8",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Div8u",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Div16",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Div16u",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Div32",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Div32u",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Div64",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Div64u",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Div128u",
		argLen:  3,
		generic: true,
	},
	{
		name:    "Mod8",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Mod8u",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Mod16",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Mod16u",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Mod32",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Mod32u",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Mod64",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Mod64u",
		argLen:  2,
		generic: true,
	},
	{
		name:        "And8",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "And16",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "And32",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "And64",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Or8",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Or16",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Or32",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Or64",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Xor8",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Xor16",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Xor32",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Xor64",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:    "Lsh8x8",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh8x16",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh8x32",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh8x64",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh16x8",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh16x16",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh16x32",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh16x64",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh32x8",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh32x16",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh32x32",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh32x64",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh64x8",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh64x16",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh64x32",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh64x64",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh8x8",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh8x16",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh8x32",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh8x64",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh16x8",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh16x16",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh16x32",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh16x64",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh32x8",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh32x16",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh32x32",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh32x64",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh64x8",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh64x16",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh64x32",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh64x64",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh8Ux8",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh8Ux16",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh8Ux32",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh8Ux64",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh16Ux8",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh16Ux16",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh16Ux32",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh16Ux64",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh32Ux8",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh32Ux16",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh32Ux32",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh32Ux64",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh64Ux8",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh64Ux16",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh64Ux32",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh64Ux64",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:        "Eq8",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Eq16",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Eq32",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Eq64",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "EqPtr",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:    "EqInter",
		argLen:  2,
		generic: true,
	},
	{
		name:    "EqSlice",
		argLen:  2,
		generic: true,
	},
	{
		name:        "Eq32F",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Eq64F",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Neq8",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Neq16",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Neq32",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Neq64",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "NeqPtr",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:    "NeqInter",
		argLen:  2,
		generic: true,
	},
	{
		name:    "NeqSlice",
		argLen:  2,
		generic: true,
	},
	{
		name:        "Neq32F",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Neq64F",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:    "Less8",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Less8U",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Less16",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Less16U",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Less32",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Less32U",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Less64",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Less64U",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Less32F",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Less64F",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Leq8",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Leq8U",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Leq16",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Leq16U",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Leq32",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Leq32U",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Leq64",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Leq64U",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Leq32F",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Leq64F",
		argLen:  2,
		generic: true,
	},
	{
		name:    "CondSelect",
		argLen:  3,
		generic: true,
	},
	{
		name:        "AndB",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "OrB",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "EqB",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "NeqB",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:    "Not",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Neg8",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Neg16",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Neg32",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Neg64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Neg32F",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Neg64F",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Com8",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Com16",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Com32",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Com64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Ctz8",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Ctz16",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Ctz32",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Ctz64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Ctz8NonZero",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Ctz16NonZero",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Ctz32NonZero",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Ctz64NonZero",
		argLen:  1,
		generic: true,
	},
	{
		name:    "BitLen8",
		argLen:  1,
		generic: true,
	},
	{
		name:    "BitLen16",
		argLen:  1,
		generic: true,
	},
	{
		name:    "BitLen32",
		argLen:  1,
		generic: true,
	},
	{
		name:    "BitLen64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Bswap16",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Bswap32",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Bswap64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "BitRev8",
		argLen:  1,
		generic: true,
	},
	{
		name:    "BitRev16",
		argLen:  1,
		generic: true,
	},
	{
		name:    "BitRev32",
		argLen:  1,
		generic: true,
	},
	{
		name:    "BitRev64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "PopCount8",
		argLen:  1,
		generic: true,
	},
	{
		name:    "PopCount16",
		argLen:  1,
		generic: true,
	},
	{
		name:    "PopCount32",
		argLen:  1,
		generic: true,
	},
	{
		name:    "PopCount64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "RotateLeft64",
		argLen:  2,
		generic: true,
	},
	{
		name:    "RotateLeft32",
		argLen:  2,
		generic: true,
	},
	{
		name:    "RotateLeft16",
		argLen:  2,
		generic: true,
	},
	{
		name:    "RotateLeft8",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Sqrt",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Sqrt32",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Floor",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Ceil",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Trunc",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Round",
		argLen:  1,
		generic: true,
	},
	{
		name:    "RoundToEven",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Abs",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Copysign",
		argLen:  2,
		generic: true,
	},
	{
		name:    "FMA",
		argLen:  3,
		generic: true,
	},
	{
		name:      "Phi",
		argLen:    -1,
		zeroWidth: true,
		generic:   true,
	},
	{
		name:    "Copy",
		argLen:  1,
		generic: true,
	},
	{
		name:         "Convert",
		argLen:       2,
		resultInArg0: true,
		zeroWidth:    true,
		generic:      true,
	},
	{
		name:    "ConstBool",
		auxType: auxBool,
		argLen:  0,
		generic: true,
	},
	{
		name:    "ConstString",
		auxType: auxString,
		argLen:  0,
		generic: true,
	},
	{
		name:    "ConstNil",
		argLen:  0,
		generic: true,
	},
	{
		name:    "Const8",
		auxType: auxInt8,
		argLen:  0,
		generic: true,
	},
	{
		name:    "Const16",
		auxType: auxInt16,
		argLen:  0,
		generic: true,
	},
	{
		name:    "Const32",
		auxType: auxInt32,
		argLen:  0,
		generic: true,
	},
	{
		name:    "Const64",
		auxType: auxInt64,
		argLen:  0,
		generic: true,
	},
	{
		name:    "Const32F",
		auxType: auxFloat32,
		argLen:  0,
		generic: true,
	},
	{
		name:    "Const64F",
		auxType: auxFloat64,
		argLen:  0,
		generic: true,
	},
	{
		name:    "ConstInterface",
		argLen:  0,
		generic: true,
	},
	{
		name:    "ConstSlice",
		argLen:  0,
		generic: true,
	},
	{
		name:      "InitMem",
		argLen:    0,
		zeroWidth: true,
		generic:   true,
	},
	{
		name:      "Arg",
		auxType:   auxSymOff,
		argLen:    0,
		zeroWidth: true,
		symEffect: SymRead,
		generic:   true,
	},
	{
		name:      "ArgIntReg",
		auxType:   auxNameOffsetInt8,
		argLen:    0,
		zeroWidth: true,
		generic:   true,
	},
	{
		name:      "ArgFloatReg",
		auxType:   auxNameOffsetInt8,
		argLen:    0,
		zeroWidth: true,
		generic:   true,
	},
	{
		name:      "Addr",
		auxType:   auxSym,
		argLen:    1,
		symEffect: SymAddr,
		generic:   true,
	},
	{
		name:      "LocalAddr",
		auxType:   auxSym,
		argLen:    2,
		symEffect: SymAddr,
		generic:   true,
	},
	{
		name:      "SP",
		argLen:    0,
		zeroWidth: true,
		generic:   true,
	},
	{
		name:      "SB",
		argLen:    0,
		zeroWidth: true,
		generic:   true,
	},
	{
		name:      "SPanchored",
		argLen:    2,
		zeroWidth: true,
		generic:   true,
	},
	{
		name:    "Load",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Dereference",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Store",
		auxType: auxTyp,
		argLen:  3,
		generic: true,
	},
	{
		name:    "Move",
		auxType: auxTypSize,
		argLen:  3,
		generic: true,
	},
	{
		name:    "Zero",
		auxType: auxTypSize,
		argLen:  2,
		generic: true,
	},
	{
		name:    "StoreWB",
		auxType: auxTyp,
		argLen:  3,
		generic: true,
	},
	{
		name:    "MoveWB",
		auxType: auxTypSize,
		argLen:  3,
		generic: true,
	},
	{
		name:    "ZeroWB",
		auxType: auxTypSize,
		argLen:  2,
		generic: true,
	},
	{
		name:    "WBend",
		argLen:  1,
		generic: true,
	},
	{
		name:    "WB",
		auxType: auxInt64,
		argLen:  1,
		generic: true,
	},
	{
		name:      "HasCPUFeature",
		auxType:   auxSym,
		argLen:    0,
		symEffect: SymNone,
		generic:   true,
	},
	{
		name:    "PanicBounds",
		auxType: auxInt64,
		argLen:  3,
		call:    true,
		generic: true,
	},
	{
		name:    "PanicExtend",
		auxType: auxInt64,
		argLen:  4,
		call:    true,
		generic: true,
	},
	{
		name:    "ClosureCall",
		auxType: auxCallOff,
		argLen:  -1,
		call:    true,
		generic: true,
	},
	{
		name:    "StaticCall",
		auxType: auxCallOff,
		argLen:  -1,
		call:    true,
		generic: true,
	},
	{
		name:    "InterCall",
		auxType: auxCallOff,
		argLen:  -1,
		call:    true,
		generic: true,
	},
	{
		name:    "TailCall",
		auxType: auxCallOff,
		argLen:  -1,
		call:    true,
		generic: true,
	},
	{
		name:    "ClosureLECall",
		auxType: auxCallOff,
		argLen:  -1,
		call:    true,
		generic: true,
	},
	{
		name:    "StaticLECall",
		auxType: auxCallOff,
		argLen:  -1,
		call:    true,
		generic: true,
	},
	{
		name:    "InterLECall",
		auxType: auxCallOff,
		argLen:  -1,
		call:    true,
		generic: true,
	},
	{
		name:    "TailLECall",
		auxType: auxCallOff,
		argLen:  -1,
		call:    true,
		generic: true,
	},
	{
		name:    "SignExt8to16",
		argLen:  1,
		generic: true,
	},
	{
		name:    "SignExt8to32",
		argLen:  1,
		generic: true,
	},
	{
		name:    "SignExt8to64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "SignExt16to32",
		argLen:  1,
		generic: true,
	},
	{
		name:    "SignExt16to64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "SignExt32to64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "ZeroExt8to16",
		argLen:  1,
		generic: true,
	},
	{
		name:    "ZeroExt8to32",
		argLen:  1,
		generic: true,
	},
	{
		name:    "ZeroExt8to64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "ZeroExt16to32",
		argLen:  1,
		generic: true,
	},
	{
		name:    "ZeroExt16to64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "ZeroExt32to64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Trunc16to8",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Trunc32to8",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Trunc32to16",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Trunc64to8",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Trunc64to16",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Trunc64to32",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt32to32F",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt32to64F",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt64to32F",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt64to64F",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt32Fto32",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt32Fto64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt64Fto32",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt64Fto64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt32Fto64F",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt64Fto32F",
		argLen:  1,
		generic: true,
	},
	{
		name:    "CvtBoolToUint8",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Round32F",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Round64F",
		argLen:  1,
		generic: true,
	},
	{
		name:    "IsNonNil",
		argLen:  1,
		generic: true,
	},
	{
		name:    "IsInBounds",
		argLen:  2,
		generic: true,
	},
	{
		name:    "IsSliceInBounds",
		argLen:  2,
		generic: true,
	},
	{
		name:    "NilCheck",
		argLen:  2,
		generic: true,
	},
	{
		name:      "GetG",
		argLen:    1,
		zeroWidth: true,
		generic:   true,
	},
	{
		name:    "GetClosurePtr",
		argLen:  0,
		generic: true,
	},
	{
		name:    "GetCallerPC",
		argLen:  0,
		generic: true,
	},
	{
		name:    "GetCallerSP",
		argLen:  1,
		generic: true,
	},
	{
		name:    "PtrIndex",
		argLen:  2,
		generic: true,
	},
	{
		name:    "OffPtr",
		auxType: auxInt64,
		argLen:  1,
		generic: true,
	},
	{
		name:    "SliceMake",
		argLen:  3,
		generic: true,
	},
	{
		name:    "SlicePtr",
		argLen:  1,
		generic: true,
	},
	{
		name:    "SliceLen",
		argLen:  1,
		generic: true,
	},
	{
		name:    "SliceCap",
		argLen:  1,
		generic: true,
	},
	{
		name:    "SlicePtrUnchecked",
		argLen:  1,
		generic: true,
	},
	{
		name:    "ComplexMake",
		argLen:  2,
		generic: true,
	},
	{
		name:    "ComplexReal",
		argLen:  1,
		generic: true,
	},
	{
		name:    "ComplexImag",
		argLen:  1,
		generic: true,
	},
	{
		name:    "StringMake",
		argLen:  2,
		generic: true,
	},
	{
		name:    "StringPtr",
		argLen:  1,
		generic: true,
	},
	{
		name:    "StringLen",
		argLen:  1,
		generic: true,
	},
	{
		name:    "IMake",
		argLen:  2,
		generic: true,
	},
	{
		name:    "ITab",
		argLen:  1,
		generic: true,
	},
	{
		name:    "IData",
		argLen:  1,
		generic: true,
	},
	{
		name:    "StructMake0",
		argLen:  0,
		generic: true,
	},
	{
		name:    "StructMake1",
		argLen:  1,
		generic: true,
	},
	{
		name:    "StructMake2",
		argLen:  2,
		generic: true,
	},
	{
		name:    "StructMake3",
		argLen:  3,
		generic: true,
	},
	{
		name:    "StructMake4",
		argLen:  4,
		generic: true,
	},
	{
		name:    "StructSelect",
		auxType: auxInt64,
		argLen:  1,
		generic: true,
	},
	{
		name:    "ArrayMake0",
		argLen:  0,
		generic: true,
	},
	{
		name:    "ArrayMake1",
		argLen:  1,
		generic: true,
	},
	{
		name:    "ArraySelect",
		auxType: auxInt64,
		argLen:  1,
		generic: true,
	},
	{
		name:    "StoreReg",
		argLen:  1,
		generic: true,
	},
	{
		name:    "LoadReg",
		argLen:  1,
		generic: true,
	},
	{
		name:      "FwdRef",
		auxType:   auxSym,
		argLen:    0,
		symEffect: SymNone,
		generic:   true,
	},
	{
		name:    "Unknown",
		argLen:  0,
		generic: true,
	},
	{
		name:      "VarDef",
		auxType:   auxSym,
		argLen:    1,
		zeroWidth: true,
		symEffect: SymNone,
		generic:   true,
	},
	{
		name:      "VarLive",
		auxType:   auxSym,
		argLen:    1,
		zeroWidth: true,
		symEffect: SymRead,
		generic:   true,
	},
	{
		name:      "KeepAlive",
		argLen:    2,
		zeroWidth: true,
		generic:   true,
	},
	{
		name:    "InlMark",
		auxType: auxInt32,
		argLen:  1,
		generic: true,
	},
	{
		name:    "Int64Make",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Int64Hi",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Int64Lo",
		argLen:  1,
		generic: true,
	},
	{
		name:        "Add32carry",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Add32withcarry",
		argLen:      3,
		commutative: true,
		generic:     true,
	},
	{
		name:    "Sub32carry",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Sub32withcarry",
		argLen:  3,
		generic: true,
	},
	{
		name:        "Add64carry",
		argLen:      3,
		commutative: true,
		generic:     true,
	},
	{
		name:    "Sub64borrow",
		argLen:  3,
		generic: true,
	},
	{
		name:    "Signmask",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Zeromask",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Slicemask",
		argLen:  1,
		generic: true,
	},
	{
		name:    "SpectreIndex",
		argLen:  2,
		generic: true,
	},
	{
		name:    "SpectreSliceIndex",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Cvt32Uto32F",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt32Uto64F",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt32Fto32U",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt64Fto32U",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt64Uto32F",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt64Uto64F",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt32Fto64U",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt64Fto64U",
		argLen:  1,
		generic: true,
	},
	{
		name:      "Select0",
		argLen:    1,
		zeroWidth: true,
		generic:   true,
	},
	{
		name:      "Select1",
		argLen:    1,
		zeroWidth: true,
		generic:   true,
	},
	{
		name:    "SelectN",
		auxType: auxInt64,
		argLen:  1,
		generic: true,
	},
	{
		name:    "SelectNAddr",
		auxType: auxInt64,
		argLen:  1,
		generic: true,
	},
	{
		name:    "MakeResult",
		argLen:  -1,
		generic: true,
	},
	{
		name:    "AtomicLoad8",
		argLen:  2,
		generic: true,
	},
	{
		name:    "AtomicLoad32",
		argLen:  2,
		generic: true,
	},
	{
		name:    "AtomicLoad64",
		argLen:  2,
		generic: true,
	},
	{
		name:    "AtomicLoadPtr",
		argLen:  2,
		generic: true,
	},
	{
		name:    "AtomicLoadAcq32",
		argLen:  2,
		generic: true,
	},
	{
		name:    "AtomicLoadAcq64",
		argLen:  2,
		generic: true,
	},
	{
		name:           "AtomicStore8",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicStore32",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicStore64",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicStorePtrNoWB",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicStoreRel32",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicStoreRel64",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicExchange32",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicExchange64",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicAdd32",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicAdd64",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicCompareAndSwap32",
		argLen:         4,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicCompareAndSwap64",
		argLen:         4,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicCompareAndSwapRel32",
		argLen:         4,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicAnd8",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicAnd32",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicOr8",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicOr32",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicAdd32Variant",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicAdd64Variant",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicExchange32Variant",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicExchange64Variant",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicCompareAndSwap32Variant",
		argLen:         4,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicCompareAndSwap64Variant",
		argLen:         4,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicAnd8Variant",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicAnd32Variant",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicOr8Variant",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicOr32Variant",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "PubBarrier",
		argLen:         1,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:      "Clobber",
		auxType:   auxSymOff,
		argLen:    0,
		symEffect: SymNone,
		generic:   true,
	},
	{
		name:    "ClobberReg",
		argLen:  0,
		generic: true,
	},
	{
		name:           "PrefetchCache",
		argLen:         2,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "PrefetchCacheStreamed",
		argLen:         2,
		hasSideEffects: true,
		generic:        true,
	},
}

func (o Op) Asm() obj.As          { return opcodeTable[o].asm }
func (o Op) Scale() int16         { return int16(opcodeTable[o].scale) }
func (o Op) String() string       { return opcodeTable[o].name }
func (o Op) SymEffect() SymEffect { return opcodeTable[o].symEffect }
func (o Op) IsCall() bool         { return opcodeTable[o].call }
func (o Op) IsTailCall() bool     { return opcodeTable[o].tailCall }
func (o Op) HasSideEffects() bool { return opcodeTable[o].hasSideEffects }
func (o Op) UnsafePoint() bool    { return opcodeTable[o].unsafePoint }
func (o Op) ResultInArg0() bool   { return opcodeTable[o].resultInArg0 }

var registersAMD64 = [...]Register{
	{0, x86.REG_AX, 0, "AX"},
	{1, x86.REG_CX, 1, "CX"},
	{2, x86.REG_DX, 2, "DX"},
	{3, x86.REG_BX, 3, "BX"},
	{4, x86.REGSP, -1, "SP"},
	{5, x86.REG_BP, 4, "BP"},
	{6, x86.REG_SI, 5, "SI"},
	{7, x86.REG_DI, 6, "DI"},
	{8, x86.REG_R8, 7, "R8"},
	{9, x86.REG_R9, 8, "R9"},
	{10, x86.REG_R10, 9, "R10"},
	{11, x86.REG_R11, 10, "R11"},
	{12, x86.REG_R12, 11, "R12"},
	{13, x86.REG_R13, 12, "R13"},
	{14, x86.REGG, -1, "g"},
	{15, x86.REG_R15, 13, "R15"},
	{16, x86.REG_X0, -1, "X0"},
	{17, x86.REG_X1, -1, "X1"},
	{18, x86.REG_X2, -1, "X2"},
	{19, x86.REG_X3, -1, "X3"},
	{20, x86.REG_X4, -1, "X4"},
	{21, x86.REG_X5, -1, "X5"},
	{22, x86.REG_X6, -1, "X6"},
	{23, x86.REG_X7, -1, "X7"},
	{24, x86.REG_X8, -1, "X8"},
	{25, x86.REG_X9, -1, "X9"},
	{26, x86.REG_X10, -1, "X10"},
	{27, x86.REG_X11, -1, "X11"},
	{28, x86.REG_X12, -1, "X12"},
	{29, x86.REG_X13, -1, "X13"},
	{30, x86.REG_X14, -1, "X14"},
	{31, x86.REG_X15, -1, "X15"},
	{32, 0, -1, "SB"},
}
var paramIntRegAMD64 = []int8{0, 3, 1, 7, 6, 8, 9, 10, 11}
var paramFloatRegAMD64 = []int8{16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
var gpRegMaskAMD64 = regMask(49135)
var fpRegMaskAMD64 = regMask(2147418112)
var specialRegMaskAMD64 = regMask(2147483648)
var framepointerRegAMD64 = int8(5)
var linkRegAMD64 = int8(-1)
