	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 12, 0	sdk_version 12, 3
	.globl	_rfact                          ## -- Begin function rfact
	.p2align	4, 0x90
_rfact:                                 ## @rfact
	.cfi_startproc
## %bb.0:
	movl	$1, %eax
	cmpq	$2, %rdi
	jl	LBB0_2
## %bb.1:
	pushq	%rbp
	.cfi_def_cfa_offset 16
	.cfi_offset %rbp, -16
	movq	%rsp, %rbp
	.cfi_def_cfa_register %rbp
	pushq	%rbx
	pushq	%rax
	.cfi_offset %rbx, -24
	movq	%rdi, %rbx
	addq	$-1, %rdi
	callq	_rfact
	imulq	%rbx, %rax
	addq	$8, %rsp
	popq	%rbx
	popq	%rbp
LBB0_2:
	retq
	.cfi_endproc
                                        ## -- End function
.subsections_via_symbols
