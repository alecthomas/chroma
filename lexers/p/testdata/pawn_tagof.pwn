stock Print({_, Float, bool}:arg, arg_tag=tagof(arg))
{
    switch(arg_tag)
    {
        case (tagof(Float:)):
            PrintFloat(Float:arg);
        default:
            PrintInt(_:arg);
    }
}
