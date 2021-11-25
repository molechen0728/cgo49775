#ifdef __cplusplus
extern "C"
{
#endif

    typedef struct
    {
        char *data;
        unsigned int len;
    } SomeData;

    SomeData new_data();

    void delete_data(SomeData *data);
    
    float *new_float_arrary();
    
    void delete_float_arrary(float *arr);
    
#ifdef __cplusplus
}
#endif