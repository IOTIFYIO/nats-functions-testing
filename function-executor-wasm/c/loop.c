// add.c
float loop (int iterations)
{
        int total, i ;
        for ( i = 0; i < iterations; i ++ ) {
                total += i ;
        }
        return (((float)total)/((float)iterations)) ;
}
