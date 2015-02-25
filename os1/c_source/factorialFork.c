#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <unistd.h>

int main(void)
{
    int fact = 1;
    int i;
    pid_t x;


    for (i = 5; i > 0; i--) {
        fact = fact * i;
        printf("Fact -> %d\n", fact);

        x = fork();

        if (x > 0) {
            exit(0);
        }
    }
}
