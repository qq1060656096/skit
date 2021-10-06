
#include <stdio.h>
#include <stdlib.h>
#include <sys/time.h>

//使用select实现精确到1微秒(0.000001秒)的sleep
void sleep_us(unsigned int nusecs)
{
    struct timeval	tval;

    tval.tv_sec = nusecs / 1000000;
    tval.tv_usec = nusecs % 1000000;
    select(0, NULL, NULL, NULL, &tval);

}

int Add(int a, int b) {
    while(1) {
        sleep_us(1000000);
    }
	return a + b;
}