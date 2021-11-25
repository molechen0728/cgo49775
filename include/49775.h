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

#ifdef __cplusplus
}
#endif