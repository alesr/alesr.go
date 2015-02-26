
#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <unistd.h>
#include <pthread.h>

void* print_xs (void* unused) {
    int valorXs = 100;
    while (valorXs != 0) {
    fputc ('x', stderr);
    valorXs --;
  }
  return NULL;
}

void* print_os (void* unused) {
    int valorOs = 100;
    while (valorOs != 0) {
    fputc ('0', stderr);
    valorOs --;
  }
  return NULL;
}


int main () {

  pthread_t thread_xs;
  pthread_t thread_os;

  pthread_create (&thread_xs, NULL, &print_xs, NULL);
  pthread_create (&thread_os, NULL, &print_os, NULL);


  pthread_join(thread_xs, NULL);
  pthread_join(thread_os, NULL);

  return(0);
  }
