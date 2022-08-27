_loop_while:                            ## @loop_while
## %bb.0:
	pushq	%rbp
	movq	%rsp, %rbp
	movl	$1, %eax
	cmpq	%rsi, %rdi
	jge	LBB0_3
## %bb.1:
	movl	$1, %eax
LBB0_2:                                 ## =>This Inner Loop Header: Depth=1
	leaq	(%rsi,%rdi), %rcx
	imulq	%rcx, %rax
	addq	$1, %rdi
	cmpq	%rdi, %rsi
	jne	LBB0_2
LBB0_3:
	popq	%rbp
	retq
                                        ## -- End function
.subsections_via_symbols
