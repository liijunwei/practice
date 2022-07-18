	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 12, 0	sdk_version 12, 3
	.globl	_gets                           ## -- Begin function gets
	.p2align	4, 0x90
_gets:                                  ## @gets
	.cfi_startproc
## %bb.0:
	pushq	%rbp
	.cfi_def_cfa_offset 16
	.cfi_offset %rbp, -16
	movq	%rsp, %rbp
	.cfi_def_cfa_register %rbp
	pushq	%r14
	pushq	%rbx
	.cfi_offset %rbx, -32
	.cfi_offset %r14, -24
	movq	%rdi, %r14
	xorl	%ebx, %ebx
	.p2align	4, 0x90
LBB0_1:                                 ## =>This Inner Loop Header: Depth=1
	callq	_getchar
	cmpl	$-1, %eax
	je	LBB0_4
## %bb.2:                               ##   in Loop: Header=BB0_1 Depth=1
	cmpl	$10, %eax
	je	LBB0_6
## %bb.3:                               ##   in Loop: Header=BB0_1 Depth=1
	movb	%al, (%r14,%rbx)
	addq	$1, %rbx
	jmp	LBB0_1
LBB0_4:
	testq	%rbx, %rbx
	je	LBB0_5
LBB0_6:
	movb	$0, (%r14,%rbx)
LBB0_7:
	movq	%r14, %rax
	popq	%rbx
	popq	%r14
	popq	%rbp
	retq
LBB0_5:
	xorl	%r14d, %r14d
	jmp	LBB0_7
	.cfi_endproc
                                        ## -- End function
	.globl	_echo                           ## -- Begin function echo
	.p2align	4, 0x90
_echo:                                  ## @echo
	.cfi_startproc
## %bb.0:
	pushq	%rbp
	.cfi_def_cfa_offset 16
	.cfi_offset %rbp, -16
	movq	%rsp, %rbp
	.cfi_def_cfa_register %rbp
	pushq	%rbx
	subq	$24, %rsp ## allocate 24 bytes on stack
	.cfi_offset %rbx, -24
	movq	___stack_chk_guard@GOTPCREL(%rip), %rax
	movq	(%rax), %rax
	movq	%rax, -16(%rbp)
	leaq	-24(%rbp), %rbx
	movq	%rbx, %rdi
	callq	_gets
	movq	%rbx, %rdi
	callq	_puts
	movq	___stack_chk_guard@GOTPCREL(%rip), %rax
	movq	(%rax), %rax
	cmpq	-16(%rbp), %rax
	jne	LBB1_2
## %bb.1:
	addq	$24, %rsp ## deallocate 24 bytes on stack
	popq	%rbx
	popq	%rbp
	retq
LBB1_2:
	callq	___stack_chk_fail
	.cfi_endproc
                                        ## -- End function
.subsections_via_symbols
