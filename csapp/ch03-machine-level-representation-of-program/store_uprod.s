_store_uprod:                           ## @store_uprod
## %bb.0:
	pushq	%rbp
	movq	%rsp, %rbp
	movq	%rdx, %rax
	mulq	%rsi
	movq	%rdx, 8(%rdi)
	movq	%rax, (%rdi)
	popq	%rbp
	retq
                                        ## -- End function
.subsections_via_symbols
