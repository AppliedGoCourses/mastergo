// In C, local variables do not survive their function,
// even if the function returns a pointer to that variable.
//
// Run this as
//
//     cc functionsandpointers.c && ./a.out
//
// (or "gcc" if you have gcc installed rather than cc)
//
// The compiler should issue a warning about returning the address of stack memory,
// and the value x should print as "6" instead of the expected value "1".

#include <stdio.h>

int *f() {
	int a;
	a = 1;
	return &a;
}

int g(int i) {
	return i++;
}

int main(int argc, char **argv) {
	int *x = f(); // x points to a location on the call stack
	g(5); // write new things to the call stack
	printf("x points to %p and has the value %d\n", x, *x); // Oops
}