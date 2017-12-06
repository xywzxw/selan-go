server
{
    listen {{.port}};
    server_name {{.cname}};
    root {{.path}};
}